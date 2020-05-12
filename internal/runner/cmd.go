package runner

import (
	"bytes"
	"os/exec"
)

type rCmd struct {
	c      *exec.Cmd
	output *bytes.Buffer
	done bool
	err error
}

func (c *rCmd) Wait() (string, error){
	for ! c.CheckFinished() {

	}
	return c.Collect()
}

func (c *rCmd) CheckFinished() bool {
	return c.done
}

func (c *rCmd) Collect() (string, error) {
	return c.output.String(), c.err
}

func (c *rCmd) start() {
	res, err := c.c.CombinedOutput()

	c.err = err
	c.output = bytes.NewBuffer(res)

	c.done = true
}