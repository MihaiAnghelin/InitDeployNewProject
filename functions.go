package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func readRepoURL() string {
	for {
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		input = strings.TrimSuffix(input, "\n")
		// Check if the input is a valid URL
		if !strings.HasPrefix(input, "https://") {
			err = fmt.Errorf("invalid URL. Please try again")
			fmt.Println(err)
			continue
		}

		// Check if the input is a valid gitlab repository
		if !strings.Contains(input, "gitlab.com") {
			err = fmt.Errorf("invalid gitlab repository. Please try again")
			fmt.Println(err)
			continue
		}

		if err == nil {
			return input
		}
	}
}

func checkDotnetVersion() string {
	cmd, err := exec.Command("dotnet", "--version").Output()

	if err != nil {
		fmt.Println("Please install dotnet on your system")
		os.Exit(1)
	}

	output := strings.TrimSuffix(string(cmd), "\n")

	// Get the major version of dotnet and the minor version
	// Example: 5.0
	// Major version: 5
	// Minor version: 0
	majorVersion := output[:strings.Index(output, ".")]
	minorVersion := "0"

	output = majorVersion + "." + minorVersion

	return output
}

func checkNodeVersion() string {
	cmd, err := exec.Command("node", "--version").Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(cmd)
}

func printPath() {
	cmd, err := exec.Command("echo", "$PATH").Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(cmd))
}

func checkGit() string {
	cmd, err := exec.Command("git", "--version").Output()

	if err != nil {
		fmt.Println("Please install git on your system")
		os.Exit(1)
	}

	return string(cmd)
}

func selectPath() string {
	for {
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			_ = fmt.Errorf("invalid path. Please try again")
			continue
		}

		input = strings.TrimSuffix(input, "\n")

		if err == nil {
			return input
		}
	}
}

func initDotnetProject(pathStr string, dotnetVersion string) {
	createPathIfNotExist(pathStr)
	projectName := pathStr[strings.LastIndex(pathStr, "/")+1:]

	{ // Create a new dotnet solution
		cmd := exec.Command("dotnet", "new", "sln")
		cmd.Dir = pathStr
		err := cmd.Run()

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	{ // Create a new dotnet project
		cmd := exec.Command("dotnet", "new", "webapi", "--language", "C#", "--name", projectName)
		cmd.Dir = pathStr
		err := cmd.Run()

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	{ // Add the project to the solution
		cmd := exec.Command("dotnet", "sln", "add", projectName+"/"+projectName+".csproj")
		cmd.Dir = pathStr
		err := cmd.Run()

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	// Create a new dockerfile for the project
	dockerfilePath := pathStr + "/" + projectName + "/Dockerfile"

	createFile(
		dockerfilePath,
		DotnetDockerfile,
		Replacer{
			Old: "{{DOTNET_VERSION}}",
			New: dotnetVersion,
		},
		Replacer{
			Old: "{{PROJECT_NAME}}",
			New: projectName,
		},
	)

	// Create a new dockerignore for the project
	dockerignorePath := pathStr + "/" + projectName + "/.dockerignore"

	createFile(
		dockerignorePath,
		DotnetDockerignore,
	)

	// Create a new gitignore for the project
	gitignorePath := pathStr + "/.gitignore"

	createFile(
		gitignorePath,
		Gitignore,
	)

	// Create a new gitlab-ci.yml for the project
	gitlabCIPath := pathStr + "/.gitlab-ci.yml"

	createFile(
		gitlabCIPath,
		GitlabCI,
	)

}
