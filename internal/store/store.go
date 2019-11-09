package store

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Record struct {
	ID           string            `json:"_id,omitempty"`
	Name         string            `json:"name,omitempty"`
	Version      string            `json:"version,omitempty"`
	Dependencies map[string]string `json:"dependencies,omitempty"`
}

type Records map[string]*Record

type Store interface {
	Save(context.Context, *Record) error
	Get(context.Context, string) (*Record, error)
	All(context.Context) (map[string]*Record, error)

	Open(context.Context) error
	Close(context.Context) error
}

type MemoryStore struct {
	memory Records
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{memory: make(Records)}
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

type FileStore struct {
	MemoryStore
	path   string
	noInit bool
}

func NewFileStore(path string, noInit bool) *FileStore {
	return &FileStore{path: path, noInit: noInit}
}

func (ms *FileStore) Open(context.Context) error {
	if ms.noInit {
		ms.memory = make(Records)
		return nil
	}

	file, err := os.Open(ms.path)
	if err != nil {
		return err
	}

	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	records := &Records{}
	if err := json.Unmarshal(content, records); err != nil {
		fmt.Println(err)
		return err
	}

	ms.memory = *records
	return nil
}

func (ms *FileStore) Close(context.Context) error {
	content, err := json.MarshalIndent(ms.memory, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(ms.path, content, 0644)
}
