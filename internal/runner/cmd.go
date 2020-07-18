package runner

import (
	"os/exec"
	"time"
)

type RCmd struct {
	cmd       *exec.Cmd
	output    []byte
	done      bool
	err       error
	StartTime time.Time
	EndTime   time.Time
}

func (c *RCmd) Wait() (string, error) {
	for !c.CheckDone() {

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
	c.StartTime = time.Now()
	c.output, c.err = c.cmd.CombinedOutput()
	c.EndTime = time.Now()
	c.done = true
}
