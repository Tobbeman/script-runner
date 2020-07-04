package runner

import (
	"os/exec"
	"path"
	"strings"
)

type Runner struct{
	root string
}

func New(root string) *Runner {
	return &Runner{
		root: root,
	}
}

func insideRoot(root, file string) bool {
	if strings.Contains(file, ".."){
		return false
	}
	return true
}

func(r *Runner) Run(file string, args []string) (string, error) {
	cmd, err := r.RunAsync(file, args)
	if err != nil {
		return "", err
	}
	return cmd.Wait()
}

func (r *Runner) RunAsync(file string, args []string) (*RCmd, error) {
	if ! insideRoot(r.root, file) {
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
		&exec.Cmd {
			Path:   executor,
			Args:   args,
		},
		[]byte{},
		false,
		nil,
	}

	go c.start()
	return &c, nil
}
