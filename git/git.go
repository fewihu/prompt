// Package git formats the information about the state of a local git repository
package git

import (
	"os/exec"
	"prompt/text"
	"regexp"
	"strconv"
	"strings"
)

const (
	retOk             = 0
	git               = "git"
	revParse          = "rev-parse"
	head              = "HEAD"
	pathSwitch        = "-C"
	verifySwitch      = "--verify"
	abbrevRefSwitch   = "--abbrev-ref"
	lsFiles           = "ls-files"
	othersSwitch      = "--others"
	excludeStdSwitch  = "--exclude-standard"
	status            = "status"
	shortSwitch       = "-s"
	noUntrackedSwitch = "-uno"
	stash             = "stash"
	list              = "list"
)

var specialCommitRefs = map[string]string{
	"REBASE":      "REBASE_HEAD",
	"CHERRY-PICK": "CHERRY_PICK_HEAD",
	"MERGE":       "MERGE_HEAD",
	"REVERT":      "REVERT_HEAD",
}

func GetStateOrBranch(pathCh <-chan *string, state chan<- text.FormattedText) {
	path := <-pathCh
	if path == nil || !inRepository(*path) {
		state <- text.Normal("")
		return
	}

	stateCh := make(chan *string)
	go getState(*path, stateCh)
	specialState := <-stateCh

	commitCh := make(chan *string)
	go getCommit(*path, commitCh)
	commit := <-commitCh

	if specialState != nil {
		state <- text.Join("  ",
			text.SpecialRef(text.Red),
			text.Join("",
				text.Red("["),
				text.BoldColor(text.Red(*specialState)),
				text.Red("]"),
				text.Black(" "),
				text.Black(*commit)))
		return
	}

	refCh := make(chan *string)
	go getRef(*path, refCh)
	branch := <-refCh

	diffCh := make(chan text.FormattedText)
	go getDiff(*path, diffCh)
	diff := <-diffCh

	if branch != nil {
		state <- text.Join(" ",
			text.Branch(text.Green),
			text.BoldColor(text.Green(*branch)),
			diff,
			text.Normal("\n"))
		return
	}
	state <- text.Normal("")
}

func inRepository(path string) bool {
	cmd := exec.Command(git, pathSwitch, path, revParse)

	if err := cmd.Start(); err != nil {
		return false
	}

	if err := cmd.Wait(); err != nil {
		return false
	}

	return cmd.ProcessState.ExitCode() == retOk
}

func getState(path string, state chan<- *string) {
	for indicator, commitRef := range specialCommitRefs {
		cmd := exec.Command(git, pathSwitch, path, revParse, verifySwitch, commitRef)
		if err := cmd.Start(); err != nil {
			continue
		}

		if err := cmd.Wait(); err != nil {
			continue
		}

		if cmd.ProcessState.ExitCode() == retOk {
			state <- &indicator
			return
		}
	}
	state <- nil
}

func getCommit(path string, commit chan<- *string) {
	c, err := exec.Command(git, pathSwitch, path, revParse, head).Output()
	if err != nil {
		commit <- nil
		return
	}
	commitHash := string(c)
	commit <- &commitHash
}

func getRef(path string, ref chan<- *string) {
	branch, err := exec.Command(git, pathSwitch, path, revParse, abbrevRefSwitch, head).Output()
	if err != nil {
		ref <- nil
		return
	}
	branchName := strings.Replace(string(branch), "\n", "", -1)
	ref <- &branchName
}

func getDiff(path string, diff chan<- text.FormattedText) {

	diffState := text.Join("", text.Green("["))

	news := make(chan int)
	modified := make(chan int)
	added := make(chan int)
	stashed := make(chan int)

	go getNumberNew(path, news)
	go getNumberWithState(path, *regexp.MustCompile("^.M.*"), modified)
	go getNumberWithState(path, *regexp.MustCompile("^A.*$"), added)
	go getNumberStashes(path, stashed)

	// if news == 0 && modified == 0 && added == 0 && stashed == 0 {
	// 	return text.Join("", diffState, text.Check(text.Green), text.Green("]"))
	// }

	formatInformation(text.BoldColor(text.Green("!")), <-news, &diffState)
	formatInformation(text.Pen(text.Green), <-modified, &diffState)
	formatInformation(text.Check(text.Green), <-added, &diffState)
	formatInformation(text.Cabinet(text.Green), <-stashed, &diffState)
	diffState = text.Join(" ", diffState, text.Green("]"))
	diff <- diffState
}

func getNumberNew(path string, ch chan<- int) {
	// git ls-files --others --exclude-standard
	out, err := exec.Command(git, pathSwitch, path, lsFiles, othersSwitch, excludeStdSwitch).Output()
	if err != nil {
		ch <- 0
		return
	}
	ch <- strings.Count(string(out), "\n")
}

func getNumberWithState(path string, reg regexp.Regexp, ch chan<- int) {
	// git status -s -uno
	out, err := exec.Command(git, pathSwitch, path, status, shortSwitch, noUntrackedSwitch).Output()
	if err != nil {
		ch <- 0
		return
	}

	var hasState int
	for _, outStr := range strings.Split(string(out), "\n") {
		hasState += len(reg.FindAllString(outStr, -1))
	}
	ch <- hasState
}

func getNumberStashes(path string, ch chan<- int) {
	out, err := exec.Command(git, pathSwitch, path, stash, list).Output()
	if err != nil {
		ch <- 0
		return
	}
	ch <- strings.Count(string(out), "\n")
}

func formatInformation(sym text.FormattedText, num int, state *text.FormattedText) {
	if num > 0 {
		*state = text.Join("", *state, text.Green(" "), sym, text.BoldColor(text.Green(strconv.Itoa(num))))
	}
}
