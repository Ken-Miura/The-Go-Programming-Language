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
		fmt.Println("ex: " + os.Args[0] + " fmt")
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
	pkgsString, err := getAllPackageInfoAsJson()
	if err != nil {
		return nil, err
	}
	type packageInfo struct {
		Name       string
		ImportPath string
		Deps       []string
	}
	var allPkgs []packageInfo
	if err := json.NewDecoder(strings.NewReader(pkgsString)).Decode(&allPkgs); err != nil {
		return nil, fmt.Errorf("failed to decode as json: %v", err)
	}
	var dependencies []string
loop:
	for _, pkg := range allPkgs {
		for _, dep := range pkg.Deps {
			if dep == importPath {
				dependencies = append(dependencies, pkg.Name+" (import path: "+pkg.ImportPath+")")
				continue loop
			}
		}
	}
	return dependencies, nil
}

func getAllPackageInfoAsJson() (string, error) {
	cmd := exec.Command("go", "list", "-f",
		`{ "Name":"{{.Name}}", "ImportPath":"{{.ImportPath}}", "Deps":["{{ join .Deps "\", \"" }}"]},`, "...") // ...1
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("failed to connect with %s's stdout: %v", cmd.Args[0], err)
	}
	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("failed to start command (%s): %v", cmd.Args[0], err)
	}
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, stdout)
	if err != nil {
		return "", err
	}
	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("failed to wait for command (%s): %v", cmd.Args[0], err)
	}

	// 1のフォーマットによる出力が正しいJSON配列でないので微修正 (最後にカンマ有りかつ[と]でくくられていない)
	tmp := buf.String()
	i := strings.LastIndex(tmp, ",")
	return fmt.Sprintf("[%s]", tmp[:i]), nil
}
