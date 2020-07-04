package runner

import (
	"os/exec"
)

type RCmd struct {
	c      *exec.Cmd
	output []byte
	done bool
	err error
}

func (c *RCmd) Wait() (string, error){
	for ! c.CheckDone() {

	}
	return c.Collect()
}

func (c *RCmd) CheckDone() bool {
	return c.done
}

func (c *RCmd) Collect() (string, error) {
	return string(c.output), c.err
}

func (c *RCmd) start() {
	c.output, c.err = c.c.CombinedOutput()
	c.done = true
}