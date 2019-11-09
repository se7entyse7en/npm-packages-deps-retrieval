package store

import "context"

type Record struct {
	ID           string
	Name         string
	Version      string
	Dependencies map[string]string
}

type Store interface {
	Save(context.Context, *Record) error
	Get(context.Context, string) (*Record, error)
	All(context.Context) (map[string]*Record, error)

	Open(context.Context) error
	Close(context.Context) error
}

type MemoryStore struct {
	memory map[string]*Record
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{memory: make(map[string]*Record)}
}

func (ms *MemoryStore) Save(ctx context.Context, r *Record) error {
	ms.memory[r.ID] = r
	return nil
}

func (ms *MemoryStore) Get(ctx context.Context, id string) (*Record, error) {
	return ms.memory[id], nil
}

func (ms *MemoryStore) All(ctx context.Context) (map[string]*Record, error) {
	return ms.memory, nil
}

func (ms *MemoryStore) Open(context.Context) error {
	return nil
}

func (ms *MemoryStore) Close(context.Context) error {
	return nil
}
