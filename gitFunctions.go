package main

import (
	"fmt"
	"os"
	"os/exec"
)

func gitInit(pathStr string) {
	cmd := exec.Command("git", "init", "--initial-branch=master")
	cmd.Dir = pathStr
	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func gitRemoteAdd(pathStr string, repoURL string) {
	cmd := exec.Command("git", "remote", "add", "origin", repoURL)
	cmd.Dir = pathStr
	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func gitAdd(pathStr string) {
	cmd := exec.Command("git", "add", ".")
	cmd.Dir = pathStr
	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func gitCommit(pathStr string, commitMessage string) {
	cmd := exec.Command("git", "commit", "-m", commitMessage)
	cmd.Dir = pathStr
	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func gitPush(pathStr string) {
	cmd := exec.Command("git", "push", "--set-upstream", "origin", "master")
	cmd.Dir = pathStr
	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
