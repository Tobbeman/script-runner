package runner

import (
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

type RCmdStore struct {
	sync.RWMutex
	store map[string]*RCmd
}

type RetentionConfig struct {
	Time float64 `yaml:"time"`
	Mode string  `yaml:"mode"`
}

type RetentionMode string

const (
	Completed RetentionMode = "completed"
)

func NewStoreWithRetention(c RetentionConfig) *RCmdStore {
	s := RCmdStore{
		store: make(map[string]*RCmd),
	}
	ticker := time.NewTicker(time.Duration(c.Time) * time.Hour)
	go func() {
		for {
			select {
			case <-ticker.C:
				s.cleanup(c)
			}
		}
	}()
	return &s
}

func NewStore() *RCmdStore {
	return &RCmdStore{store: make(map[string]*RCmd)}
}

func (s *RCmdStore) Store(cmd *RCmd) string {
	s.Lock()
	defer s.Unlock()
	u := uuid.NewV4().String()
	s.store[u] = cmd
	return u
}

func (s *RCmdStore) Get(uuid string) (bool, *RCmd) {
	s.RLock()
	defer s.RUnlock()
	if cmd, found := s.store[uuid]; found {
		return true, cmd
	}
	return false, nil
}

func (s *RCmdStore) GetMap() map[string]*RCmd {
	s.RLock()
	defer s.RUnlock()
	return s.store
}

func (s *RCmdStore) cleanup(c RetentionConfig) {
	for uuid, cmd := range s.store {
		switch RetentionMode(c.Mode) {
		case Completed:
			if time.Since(cmd.EndTime).Hours() > c.Time {
				delete(s.store, uuid)
			}
		}
	}
}
