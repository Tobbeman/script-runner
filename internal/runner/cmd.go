package runner

import (
	"os/exec"
	"sync"
	"time"
)

type RCmd struct {
	cmd       *exec.Cmd
	output    []byte
	done      bool
	err       error
	StartTime time.Time
	EndTime   time.Time
	sync.RWMutex
}

func (c *RCmd) Wait() (string, error) {
	for !c.CheckDone() {

	}
	return c.Collect()
}

func (c *RCmd) CheckDone() bool {
	c.RLock()
	defer c.RUnlock()
	return c.done
}

func (c *RCmd) Collect() (string, error) {
	return string(c.output), c.err
}

func (c *RCmd) start() {
	c.Lock()
	c.StartTime = time.Now()
	c.output, c.err = c.cmd.CombinedOutput()
	c.EndTime = time.Now()
	c.done = true
	c.Unlock()
}
