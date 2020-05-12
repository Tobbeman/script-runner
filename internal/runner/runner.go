package runner

import (
	"bytes"
	"os/exec"
	"path"
	"strings"
)

type runner struct{
	root string
}

func New(root string) *runner {
	return &runner{
		root: root,
	}
}

func insideRoot(root, file string) bool {
	if strings.Contains(file, ".."){
		return false
	}
	return true
}

func(r *runner) Run(file string, args []string) (string, error) {
	if ! insideRoot(r.root, file) {
		return "", &Error{
			"File outside root",
			nil,
		}
	}

	p := path.Join(r.root, file)

	executor, err := exec.LookPath(p)
	if err != nil {
		return "", convertError(err).Err
	}

	execCmd := exec.Command(executor, args...)

	res, err := execCmd.Output()
	if err != nil {
		return "", convertError(err).Err
	}
	return string(res), nil
}

func (r *runner) RunAsync(file string, args []string) (*rCmd, error) {
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

	var output bytes.Buffer

	c := rCmd{
		&exec.Cmd {
			Path:   executor,
			Args:   args,
		},
		&output,
		false,
		nil,
	}

	go c.start()
	return &c, nil
}
