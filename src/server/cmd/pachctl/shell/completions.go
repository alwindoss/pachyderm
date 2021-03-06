package shell

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/pachyderm/pachyderm/src/client"
	"github.com/pachyderm/pachyderm/src/client/pfs"
	"github.com/pachyderm/pachyderm/src/client/pps"
	"github.com/pachyderm/pachyderm/src/server/pkg/cmdutil"
	"github.com/pachyderm/pachyderm/src/server/pkg/errutil"
	"github.com/pachyderm/pachyderm/src/server/pkg/pretty"
	pps_pretty "github.com/pachyderm/pachyderm/src/server/pps/pretty"

	prompt "github.com/c-bata/go-prompt"
	"github.com/docker/go-units"
)

// RepoCompletion completes repo parameters of the form <repo>
func RepoCompletion(_, text string, maxCompletions int64) []prompt.Suggest {
	c, err := client.NewOnUserMachine("user-completion")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	ris, err := c.ListRepo()
	if err != nil {
		log.Fatal(err)
	}
	var result []prompt.Suggest
	for _, ri := range ris {
		result = append(result, prompt.Suggest{
			Text:        ri.Repo.Name,
			Description: fmt.Sprintf("%s (%s)", ri.Description, units.BytesSize(float64(ri.SizeBytes))),
		})
	}
	return result
}

// BranchCompletion completes branch parameters of the form <repo>@<branch>
func BranchCompletion(flag, text string, maxCompletions int64) []prompt.Suggest {
	c, err := client.NewOnUserMachine("user-completion")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	partialFile := cmdutil.ParsePartialFile(text)
	var result []prompt.Suggest
	if partialFile.Commit.Repo.Name == "" || // nothing typed yet
		(len(partialFile.Commit.ID) == 0 && text[len(text)-1] != '@') { // partial repo typed
		return RepoCompletion(flag, text, maxCompletions)
	} else if partialFile.Commit.ID == "" || // repo@ typed, no commit/branch yet
		(len(partialFile.Path) == 0 && text[len(text)-1] != ':') { // partial commit/branch typed
		bis, err := c.ListBranch(partialFile.Commit.Repo.Name)
		if err != nil {
			log.Fatal(err)
		}
		for _, bi := range bis {
			head := "-"
			if bi.Head != nil {
				head = bi.Head.ID
			}
			result = append(result, prompt.Suggest{
				Text:        fmt.Sprintf("%s@%s:", partialFile.Commit.Repo.Name, bi.Branch.Name),
				Description: fmt.Sprintf("(%s)", head),
			})
		}
	}
	return result
}

// FileCompletion completes file parameters of the form <repo>@<branch>:/file
func FileCompletion(flag, text string, maxCompletions int64) []prompt.Suggest {
	c, err := client.NewOnUserMachine("user-completion")
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()
	partialFile := cmdutil.ParsePartialFile(text)
	var result []prompt.Suggest
	if partialFile.Commit.Repo.Name == "" || // nothing typed yet
		(len(partialFile.Commit.ID) == 0 && text[len(text)-1] != '@') { // partial repo typed
		return RepoCompletion(flag, text, maxCompletions)
	} else if partialFile.Commit.ID == "" || // repo@ typed, no commit/branch yet
		(len(partialFile.Path) == 0 && text[len(text)-1] != ':') { // partial commit/branch typed
		return BranchCompletion(flag, text, maxCompletions)
	} else { // typing the file
		if err := c.GlobFileF(partialFile.Commit.Repo.Name, partialFile.Commit.ID, partialFile.Path+"*", func(fi *pfs.FileInfo) error {
			if maxCompletions > 0 {
				maxCompletions--
			} else {
				return errutil.ErrBreak
			}
			result = append(result, prompt.Suggest{
				Text: fmt.Sprintf("%s@%s:%s", partialFile.Commit.Repo.Name, partialFile.Commit.ID, fi.File.Path),
			})
			return nil
		}); err != nil {
			log.Fatal(err)
		}
	}
	return result
}

// FilesystemCompletion completes file parameters from the local filesystem (not from pfs).
func FilesystemCompletion(_, text string, maxCompletions int64) []prompt.Suggest {
	dir := filepath.Dir(text)
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	var result []prompt.Suggest
	for _, fi := range fis {
		result = append(result, prompt.Suggest{
			Text: filepath.Join(dir, fi.Name()),
		})
	}
	return result
}

// PipelineCompletion completes pipeline parameters of the form <pipeline>
func PipelineCompletion(_, text string, maxCompletions int64) []prompt.Suggest {
	c, err := client.NewOnUserMachine("user-completion")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	pis, err := c.ListPipeline()
	if err != nil {
		log.Fatal(err)
	}
	var result []prompt.Suggest
	for _, pi := range pis {
		result = append(result, prompt.Suggest{
			Text:        pi.Pipeline.Name,
			Description: pi.Description,
		})
	}
	return result
}

func jobDesc(ji *pps.JobInfo) string {
	statusString := ""
	if ji.Finished == nil {
		statusString = fmt.Sprintf("%s for %s", pps_pretty.JobState(ji.State), pretty.Since(ji.Started))
	} else {
		statusString = fmt.Sprintf("%s %s", pps_pretty.JobState(ji.State), pretty.Ago(ji.Finished))
	}
	return fmt.Sprintf("%s: %s - %s", ji.Pipeline.Name, pps_pretty.Progress(ji), statusString)
}

// JobCompletion completes job parameters of the form <job>
func JobCompletion(_, text string, maxCompletions int64) []prompt.Suggest {
	c, err := client.NewOnUserMachine("user-completion")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	var result []prompt.Suggest
	if err := c.ListJobF("", nil, nil, 0, false, func(ji *pps.JobInfo) error {
		if maxCompletions > 0 {
			maxCompletions--
		} else {
			return errutil.ErrBreak
		}
		result = append(result, prompt.Suggest{
			Text:        ji.Job.ID,
			Description: jobDesc(ji),
		})
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return result
}
