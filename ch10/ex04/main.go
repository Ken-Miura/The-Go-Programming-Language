// Copyright 2017 Ken Miura
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: " + os.Args[0] + " 'package name'")
		fmt.Println("usage: " + os.Args[0] + " fmt")
		os.Exit(1)
	}

	importPath, err := findImportPathByPackageName(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dependentPackages, err := findDependencies(importPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, pkg := range dependentPackages {
		fmt.Println(pkg)
	}
}

func findImportPathByPackageName(pkg string) (string, error) {
	cmd := exec.Command("go", "list", "-json", pkg)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("failed to connect with %s's stdout: %v", cmd.Args[0], err)
	}
	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("failed to start command (%s): %v", cmd.Args[0], err)
	}
	var packageInfo struct {
		ImportPath string
	}
	if err := json.NewDecoder(stdout).Decode(&packageInfo); err != nil {
		return "", fmt.Errorf("failed to decode as json: %v", err)
	}
	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("failed to wait for command (%s): %v", cmd.Args[0], err)
	}
	return packageInfo.ImportPath, nil
}

func findDependencies(importPath string) ([]string, error) {
	cmd := exec.Command("go", "list", "-f",
		`{ "Name":"{{.Name}}", "Deps":["{{ join .Deps "\", \"" }}"]},`, "...")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to connect with %s's stdout: %v", cmd.Args[0], err)
	}
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start command (%s): %v", cmd.Args[0], err)
	}
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, stdout)
	if err != nil {
		return nil, err
	}
	tmp := buf.String()
	i := strings.LastIndex(tmp, ",")
	tmp = "[" + tmp[:i] + "]"

	type packageInfo struct {
		Name string
		Deps []string
	}
	var allPackages []packageInfo
	if err := json.NewDecoder(strings.NewReader(tmp)).Decode(&allPackages); err != nil {
		return nil, fmt.Errorf("failed to decode as json: %v", err)
	}
	if err := cmd.Wait(); err != nil {
		return nil, fmt.Errorf("failed to wait for command (%s): %v", cmd.Args[0], err)
	}
	var dependencies []string
loop:
	for _, pkg := range allPackages {
		for _, dep := range pkg.Deps {
			if dep == importPath {
				dependencies = append(dependencies, pkg.Name)
				continue loop
			}
		}
	}
	return dependencies, nil
}
