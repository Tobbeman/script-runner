package runner

import (
	"os/exec"
)

type rCmd struct {
	c      *exec.Cmd
	output []byte
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
	return string(c.output), c.err
}

func (c *rCmd) start() {
	c.output, c.err = c.c.CombinedOutput()
	c.done = true
}