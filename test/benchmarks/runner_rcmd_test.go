package benchmarks

import (
	"os/exec"
	"sync"
	"testing"
)

type wCmdChan struct {
	c    *exec.Cmd
	err  error
	done chan bool
}

func (w *wCmdChan) start() {
	go func() {
		w.err = w.c.Run()
		w.done <- true
	}()
}

func (w *wCmdChan) check() bool {
	return <-w.done
}

type wCmdMutext struct {
	c    *exec.Cmd
	err  error
	done bool
	sync.RWMutex
}

func (w *wCmdMutext) start() {
	w.Lock()
	w.err = w.c.Run()
	w.done = true
	w.Unlock()
}

func (w *wCmdMutext) check() bool {
	w.RLock()
	defer w.RUnlock()
	return w.done
}

func BenchmarkChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd := exec.Command("ls", "-la")
		w := wCmdChan{
			c:    cmd,
			done: make(chan bool),
		}
		w.start()
		for w.check() != true {
		}
	}
}

func BenchmarkMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd := exec.Command("ls", "-la")
		w := wCmdMutext{
			c:    cmd,
			done: false,
		}
		w.start()
		for w.check() != true {
		}
	}
}
