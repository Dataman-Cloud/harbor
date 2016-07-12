package git

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego"
	"github.com/vmware/harbor/models"
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

ssh -i %s -o "StrictHostKeyChecking no" $1 $2`
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
	fmt.Println("=================================")
	os.Chmod(wrapperpath, 0400)
	fmt.Println(wrapperpath)
	fmt.Println("=================================")
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
	return cmd.Run()
}

//git remote add
func (client *Client) RemoteAdd() error {
	cmd, err := gitCmd(client.Path, "remote", "add", "origin", client.URI)

	if err != nil {
		return err
	}
	return cmd.Run()
}

//pull the update info from remote branch
func (client *Client) Pull() error {
	cmd, err := gitCmd(client.Path, "pull", "origin", client.Branch)
	if err != nil {
		return err
	}
	return cmd.Run()
}

//reset local content
func (client *Client) Reset() error {
	cmd, err := gitCmd(client.Path, "reset", "--hard", "HEAD")
	if err != nil {
		return err
	}
	return cmd.Run()
}

// Trace writes each command to standard error (preceded by a ‘$ ’) before it
// is executed. Used for debugging your build.
func trace(cmd *exec.Cmd) {
	log.Infoln("$", cmd.Path, strings.Join(cmd.Args, " "))
}

//fetch catalog info from local files
func FetchRepoInfo(repository *models.Repository) {
	dir := path.Join("/go/bin/harborCatalog/library", repository.Name)
	beego.Info(fmt.Sprintf("%s:%s", "the dir is ", dir))
	repository.Category = readFile(path.Join(dir, "category"))
	repository.Description = readFile(path.Join(dir, "description"))
	repository.DockerCompose = readFile(path.Join(dir, "docker_compose.yml"))
	repository.Readme = readFile(path.Join(dir, "README.md"))
	repository.Catalog = readFile(path.Join(dir, "catalog.yml"))
	repository.MarathonConfig = readFile(path.Join(dir, "marathon_config.yml"))
	repository.Icon = fmt.Sprintf("%s/%s/%s.%s", "/api/v3/repositories/icons", repository.Name, repository.Name, "png")
	log.Println(fmt.Sprintf("%s:%s", "the icon path is", repository.Icon))
}

func readFile(path string) string {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(fmt.Sprintf("%s:%s", "file not exists:", path))
		return ""
	}
	return string(contents)
}
