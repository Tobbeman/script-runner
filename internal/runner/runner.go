package runner

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Runner struct {
	root string
}

func New(root string) *Runner {
	return &Runner{
		root: root,
	}
}

func (r *Runner) List() ([]string, error) {
	var files []string
	cleanDir := filepath.Clean(r.root)
	err := filepath.Walk(cleanDir, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() {
			files = append(files, strings.Replace(path, cleanDir, "", 1)[1:])
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (r *Runner) Run(file string, args []string) (string, error) {
	cmd, err := r.RunAsync(file, args)
	if err != nil {
		return "", err
	}
	return cmd.Wait()
}

func (r *Runner) RunAsync(file string, args []string) (*RCmd, error) {
	if !insideRoot(r.root, file) {
		return nil, &Error{
			"File outside root",
			nil,
		}
	}

	p := path.Join(r.root, file)

	executor, err := exec.LookPath(p)
	if err != nil {
		return nil, convertError(err).Err
	}

	c := RCmd{
		&exec.Cmd{
			Path: executor,
			Args: args,
		},
		[]byte{},
		false,
		nil,
		time.Time{},
		time.Time{},
		sync.RWMutex{},
	}

	go c.start()
	return &c, nil
}

//==============

func insideRoot(root, file string) bool {
	if strings.Contains(file, "..") {
		return false
	}
	return true
}
