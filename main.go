package main

import "fmt"

func setupBackendProjectRepo() {
	//	Read Backend Repository URL
	fmt.Println("Backend Repository URL: ")
	backendURL := readRepoURL() /*TODO REMOVE HARDCODED VALUE */
	//backendURL := "https://gitlab.com/MihaiAnghelin/testdotnet"
	fmt.Printf("Backend Repository URL: %s\n\n", backendURL)

	//  Check dotnet version installed on the system
	dotnetVersion := checkDotnetVersion()
	fmt.Printf("Dotnet Version: %s\n\n", dotnetVersion)

	// Path to the backend project
	fmt.Println("Backend Project Path (including the project directory) : ")
	backendPath := selectPath() /*TODO REMOVE HARDCODED VALUE */
	//backendPath := "/home/mihai/Dev/GoDotnetTest"
	fmt.Printf("Backend Path: %s\n\n", backendPath)

	//	Init dotnet new project and dockerfile for the project
	initDotnetProject(backendPath, dotnetVersion)

	//	Setup git repository for the project and push to the repository
	fmt.Println("Backend - Git Init")
	gitInit(backendPath)

	fmt.Println("Backend - Git Remote Add")
	gitRemoteAdd(backendPath, backendURL)

	fmt.Println("Backend - Git Add")
	gitAdd(backendPath)

	fmt.Println("Backend - Git Commit")
	gitCommit(backendPath, "Initial commit")

	fmt.Println("Backend - Git Push")
	gitPush(backendPath)
}

func setupFrontendProjectRepo() {
	//	Read Frontend Repository URL
	fmt.Println("Frontend Repository URL: ")
	//frontendURL := readRepoURL()
	frontendURL := "https://gitlab.com/MihaiAnghelin/testnextjs"
	fmt.Printf("Frontend Repository URL: %s\n\n", frontendURL)

	//	Check node version installed on the system
	nodeVersion := checkNodeVersion()
	fmt.Printf("Node Version: %s\n\n", nodeVersion)
	//	Init next.js project and dockerfile for the project
	//	Push to frontend repository
}

func main() {
	// TODO CHECK IF ALL THE TOOLS ARE INSTALLED ON THE SYSTEM
	// TODO CHECK IF THE USER HAS ACCESS TO THE REPOSITORY

	setupBackendProjectRepo()
	//backendPath := "/home/mihai/Dev/GoDotnetTest"
	//backendURL := "https://gitlab.com/MihaiAnghelin/testdotnet"
	//

	//setupFrontendProjectRepo()

	//	Install Docker on the remote server
	//	Create new user on the remote server
	//	Setup SSH on the remote server for the new user to use with Gitlab CI/CD
	//  Setup gitlab runner on the remote server

	//	Setup Gitlab CI/CD pipeline

	//	Setup database on the remote server (MYSQL) through docker
	//	Setup nginx on the remote server as reverse-proxy through docker
	//	Setup SSL on the remote server through docker (Certbot)

}
