package storage

import (
	"errors"
	"sync"
)

type Task struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type MemoryStore struct {
	blockFlag     sync.RWMutex
	autoIncrement int64
	tasks         map[int64]*Task
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		tasks: make(map[int64]*Task),
	}
}

func (s *MemoryStore) Create(title string) *Task {
	s.blockFlag.Lock()
	defer s.blockFlag.Unlock()
	s.autoIncrement++
	t := &Task{ID: s.autoIncrement, Title: title, Done: false}
	s.tasks[t.ID] = t
	return t
}

func (s *MemoryStore) Delete(id int64) *Task {
	s.blockFlag.Lock()
	defer s.blockFlag.Unlock()
	t := s.tasks[id]
	delete(s.tasks, id)
	return t
}

func (s *MemoryStore) Get(id int64) (*Task, error) {
	s.blockFlag.RLock()
	defer s.blockFlag.RUnlock()
	t, ok := s.tasks[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return t, nil
}

func (s *MemoryStore) List() []*Task {
	s.blockFlag.RLock()
	defer s.blockFlag.RUnlock()
	out := make([]*Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		out = append(out, t)
	}
	return out
}
