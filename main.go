package main

import (
	"os"
	"path"
)

/**
this function panic if an error is encountered
*/
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// path to your project
	projectPath := "/Users/nawfal/projects"

	// project name
	projectName := "project-name"

	// append more directories or subdirectories
	strDir := []string{}
	strDir = append(strDir, "cmd")
	strDir = append(strDir, "internal")
	strDir = append(strDir, "pkg")
	strDir = append(strDir, "vendor")
	strDir = append(strDir, "api")
	strDir = append(strDir, "web")
	strDir = append(strDir, "config")
	strDir = append(strDir, "init")
	strDir = append(strDir, "script")
	strDir = append(strDir, "scripts")
	strDir = append(strDir, "build")
	strDir = append(strDir, "deployment")
	strDir = append(strDir, "test")

	// iterates through directories
	for i := 0; i < len(strDir); i++ {
		pathDir := path.Join(projectPath, projectName, strDir[i])
		err := os.MkdirAll(pathDir, os.ModePerm)
		check(err)
	}

}
