package git

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	log "github.com/Sirupsen/logrus"
)

var (
	//ErrBadURI bad git uri
	ErrBadURI = errors.New("bad uri")
	//ErrBadName bad repo name
	ErrBadName = errors.New("bad repo name")
	//ErrBadWorkspace invalid workspace dir
	ErrBadWorkspace = errors.New("bad workspace")
	//ErrBadCmd bad git subcommand
	ErrBadCmd = errors.New("bad git subcommand")
	//ErrRemoteAdd git remote add failed
	ErrRemoteAdd = errors.New("git remote add failed")
	//ErrBadRef invalid ref
	ErrBadRef = errors.New("bad ref")
	//ErrBadFile invalid filepath
	ErrBadFile = errors.New("bad file")
	//ErrBadCommit invalid commit
	ErrBadCommit = errors.New("bad commit")
)

const (
	//GitSshWrapper GIT_SSH
	GitSshWrapper = "git_ssh_wrapper"
	//GitSshWrapperScript GIT_SSH script
	GitSshWrapperScript = `#!/bin/sh

ssh -i %s $1 $2`
)

//Client git client for executing git commands
type Client struct {
	URI    string
	Branch string
	Path   string
}

//NewClient create new client
func NewClient(path, uri, branch string) (*Client, error) {
	if len(uri) == 0 {
		return nil, ErrBadURI
	}

	client := &Client{
		URI:    uri,
		Branch: branch,
	}

	if err := client.initRepo(path); err != nil {
		return nil, err
	}
	return client, nil
}

//String impl
func (client *Client) String() string {
	return fmt.Sprintf("uri: %s, branch: %s", client.URI, client.Branch)
}

//initRepo init empty repo
func (client *Client) initRepo(path string) error {
	if len(path) == 0 {
		path = "/var/lib/harbor/workspace/"
	}
	if !filepath.IsAbs(path) {
		log.Errorln("bad workspace", path)
		return ErrBadWorkspace
	}
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}
	client.Path = path

	wrapperpath := filepath.Join(path, GitSshWrapper)
	wrapperScript := fmt.Sprintf(GitSshWrapperScript, os.Getenv("PK_PATH"))
	err := ioutil.WriteFile(wrapperpath, []byte(wrapperScript), 0755)
	if err != nil {
		return err
	}

	return nil
}

func gitCmd(path string, args ...string) (*exec.Cmd, error) {
	if len(path) == 0 {
		log.Errorln("bad workspace", path)
		return nil, ErrBadWorkspace
	}
	if len(args) < 1 {
		return nil, ErrBadCmd
	}

	cmd := exec.Command("git", args...)
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), "GIT_SSH="+filepath.Join(path, GitSshWrapper))
	trace(cmd)
	return cmd, nil
}

//git init
func (client *Client) Init() error {
	cmd, err := gitCmd(client.Path, "init")
	if err != nil {
		return err
	}
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

//git remote add
func (client *Client) RemoteAdd() error {
	cmd, err := gitCmd(client.Path, "remote", "add", "origin", client.URI)
	if err != nil {
		return err
	}
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

//pull the update info from remote branch
func (client *Client) Pull() error {
	cmd, err := gitCmd(client.Path, "pull", "origin", client.Branch)
	if err != nil {
		return err
	}
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

//reset local content
func (client *Client) Reset() error {
	cmd, err := gitCmd(client.Path, "reset", "--hard", "HEAD")
	if err != nil {
		return err
	}
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// Trace writes each command to standard error (preceded by a ‘$ ’) before it
// is executed. Used for debugging your build.
func trace(cmd *exec.Cmd) {
	log.Infoln("$", cmd.Path, strings.Join(cmd.Args, " "))
}
