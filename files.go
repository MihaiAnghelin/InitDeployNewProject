package main

import (
	"fmt"
	"os"
	"strings"
)

type FileType int

const (
	DotnetDockerfile FileType = iota
	DotnetDockerignore
	GitlabCI
	Gitignore
)

func getFileTemplate(fileType FileType) string {
	var templatePath string
	switch fileType {
	case DotnetDockerfile:
		templatePath = "templates/dotnet.Dockerfile"
	case DotnetDockerignore:
		templatePath = "templates/dotnet.dockerignore"
	case GitlabCI:
		templatePath = "templates/backend.gitlab-ci.yml"
	case Gitignore:
		templatePath = "templates/dotnet.gitignore"
	default:
		fmt.Println("Invalid file type")
		os.Exit(1)
	}

	return templatePath
}

type Replacer struct {
	Old string
	New string
}

func createPathIfNotExist(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// Create the path if it doesn't exist
		err = os.MkdirAll(path, 0755)
		if err != nil {
			fmt.Println("Error while creating the path")
			os.Exit(1)
		}
	}
}

func createFile(path string, fileType FileType, replacers ...Replacer) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error while creating file")
		os.Exit(1)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error while closing file")
			os.Exit(1)
		}
	}(file)

	templatePath := getFileTemplate(fileType)

	// Read the template
	template, err := os.ReadFile(templatePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	templateString := string(template)

	// Replace the placeholders in the template
	for _, elem := range replacers {
		templateString = strings.ReplaceAll(templateString, elem.Old, elem.New)
	}

	// Write to the file from the template
	_, err = file.WriteString(templateString)
	if err != nil {
		fmt.Println("Error while writing to file")
		os.Exit(1)
	}
}
