package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Globals
const (
	CommitMsgLen = 20
)

var MainBranchOptions = []string{"main", "master"}

// Helper functions
func hasGit() bool {
	// _, err := os.Stat(".git")
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	_, err := cmd.Output()
	return err == nil
}

func isMainBranch(branchName string, mainBranches []string) bool {
	for _, mainBranch := range mainBranches {
		if branchName == mainBranch {
			return true
		}
	}
	return false
}

func getCurrentBranchName() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func prepareCommit() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Is this commit related to any projects tickets / Components / features (eg: JIRA-124, button, vpc): ")
	ticket, _ := reader.ReadString('\n')

	fmt.Print("Enter which git operation did you performed (eg: add, update, del): ")
	operation, _ := reader.ReadString('\n')

	fmt.Println("Files modified")
	if err := exec.Command("git", "status").Run(); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Print("What did you ", strings.TrimSpace(operation), ": ")
	commitMsg, _ := reader.ReadString('\n')

	switch strings.TrimSpace(operation) {
	case "add", "update", "del":
		msgLength := len(commitMsg)
		if msgLength <= CommitMsgLen {
			fmt.Printf("COMMIT MESSAGE TOO SHORT. COMMIT MESSAGE SHOULD BE AT LEAST %d CHARS LONG.\n", CommitMsgLen)
		}
	default:
		fmt.Println("Operation", strings.TrimSpace(operation), "not understood.")
		os.Exit(1)
	}

	commit(strings.TrimSpace(operation), strings.TrimSpace(ticket), strings.TrimSpace(commitMsg))
}

func commit(operation, ticket, commitMsg string) {
	currentBranchName, err := getCurrentBranchName()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if isMainBranch(currentBranchName, MainBranchOptions) {
		fmt.Println()
		fmt.Println("Not recommended to commit to", strings.Join(MainBranchOptions, " or "), "branch. Please create a new branch using \"bc branch\"")
		os.Exit(1)
	}

	exec.Command("git", "add", ".").Run()
	exec.Command("git", "commit", "-m", fmt.Sprintf("%s(%s): %s", operation, ticket, commitMsg)).Run()
}

func branch() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Is this commit related to any projects tickets / Components / features (eg: JIRA-124, button, vpc): ")
	ticket, _ := reader.ReadString('\n')

	fmt.Print("What kind of work is this (eg: bugfix, feat, demo, test): ")
	operation, _ := reader.ReadString('\n')

	fmt.Print("What work will be done in this branch: ")
	work, _ := reader.ReadString('\n')

	branchName := fmt.Sprintf("%s/%s/%s", strings.TrimSpace(operation), strings.TrimSpace(ticket), strings.TrimSpace(work))
	sanitizedBranchName := strings.ReplaceAll(branchName, " ", "-")

	exec.Command("git", "checkout", "-b", sanitizedBranchName).Run()
}

func push() {
	currentBranchName, err := getCurrentBranchName()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if isMainBranch(currentBranchName, MainBranchOptions) {
		fmt.Println("Not recommended to push to main branch. Please create a new branch using \"bc branch\"")
		os.Exit(1)
	}

	statusCmd := exec.Command("git", "status", "--porcelain")
	out, _ := statusCmd.Output()
	if len(out) != 0 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("There are still changes to be committed. Are you sure you want to push? [y/N] ")
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(response)
		if response == "y" || response == "Y" {
			exec.Command("git", "push", "--set-upstream", "origin", currentBranchName).Run()
		} else {
			fmt.Println("Please commit the changes using \"bc commit\"")
			os.Exit(1)
		}
	} else {
		exec.Command("git", "push", "--set-upstream", "origin", currentBranchName).Run()
	}
}

func openURL(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func getGitRepoURL() (string, error) {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	url := strings.TrimSpace(string(output))
	return url, nil
}

func convertGitURL(input string) string {
	// Replace ":" with "/"
	output := strings.Replace(input, ":", "/", 1)
	// Replace "git@" with "https://"
	output = strings.Replace(output, "git@", "https://", 1)
	// Remove ".git" from the end of the URL
	output = strings.TrimSuffix(output, ".git")
	return output
}

func openBrowser() {
	// Get Git repository URL
	url, err := getGitRepoURL()
	if err != nil {
		log.Fatal("Error getting Git repository URL:", err)
	}

	// Open URL using the default urlopen command
	if err := openURL(convertGitURL(url)); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if !hasGit() {
		fmt.Println("This is not a git repo. I am not needed here. Ta Ta !!!")
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments provided.")
		fmt.Println("")
		usage()
		os.Exit(1)
	}

	commandName := os.Args[1]
	switch commandName {
	case "commit":
		prepareCommit()
	case "add":
		exec.Command("git", "add", ".").Run()
	case "branch":
		branch()
	case "push":
		push()
	case "open":
		openBrowser()
	case "help":
		usage()
	case "pr":
		// Handle pull request creation
		fmt.Println("Pull request creation functionality is not implemented yet.")
		os.Exit(1)
	default:
		fmt.Println("Could not understand the command. Try running \"bc help\".")
		os.Exit(1)
	}
}

func usage() {
	fmt.Printf(`
BetterCommit (bc) is a small utility which will help you make an habit of writing better commit messages.

Usage: bc <subcommands>

Available commands:
    add         Adds all the changes to staging area
    commit      Adds and commits all the changes
    branch      Creates a new branch
    pr          Creates a new pull request
    open        Open relevant repo in browser

Example usage:

    $ bc open
`)
}
