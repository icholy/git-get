package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	giturls "github.com/whilp/git-urls"
)

func Path(rawurl string) (string, error) {
	u, err := giturls.Parse(rawurl)
	if err != nil {
		return "", err
	}
	p := u.Path
	p = strings.TrimPrefix(p, "/")
	// general purpose
	p = strings.TrimSuffix(p, ".git")
	// sourcehut
	if u.Host == "git.sr.ht" {
		p = strings.TrimPrefix(p, "~")
	}
	// old bitbucket/stash
	if u.Scheme == "http" || u.Scheme == "https" {
		p = strings.TrimPrefix(p, "scm/")
	}
	return filepath.Join(u.Host, filepath.FromSlash(p)), nil
}

func main() {
	// figure out if it's being invoked as a "git-get" or "git get"
	args := os.Args[1:]
	if len(os.Args) >= 2 && os.Args[1] == "get" {
		args = os.Args[2:]
	}
	// setup the command
	cmd := exec.Command("git", "clone")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	// forward all the arguments
	cmd.Args = append(cmd.Args, args...)
	// if the last argument is a remote url, append the local path
	if n := len(args); n > 0 && IsRemoteArg(args[n-1]) {
		path, err := Path(args[n-1])
		if err != nil {
			log.Fatal(err)
		}
		// figure out where to put it
		root := os.Getenv("GIT_GET_PATH")
		if root == "" {
			home, _ := os.UserHomeDir()
			root = filepath.Join(home, "src")
		}
		dir := filepath.Join(root, path)
		cmd.Args = append(cmd.Args, dir)
	}
	// run and exit
	if err := cmd.Run(); err != nil {
		code := 1
		var exit *exec.ExitError
		if errors.As(err, &exit) {
			code = exit.ExitCode()
		}
		os.Exit(code)
	}
}

func IsRemoteArg(arg string) bool {
	u, err := giturls.Parse(arg)
	if err != nil {
		return false
	}
	switch u.Scheme {
	case "http", "https", "ssh":
		return true
	}
	return false
}
