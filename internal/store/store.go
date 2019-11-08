package store

type Record struct {
	ID           string
	Name         string
	Version      string
	Dependencies map[string]string
}

type Store interface {
	Save(*Record) error
	Get(string) (*Record, error)
	All() (map[string]*Record, error)
}

type MemoryStore struct {
	memory map[string]*Record
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{memory: make(map[string]*Record)}
}

func (ms *MemoryStore) Save(r *Record) error {
	ms.memory[r.ID] = r
	return nil
}

func (ms *MemoryStore) Get(id string) (*Record, error) {
	return ms.memory[id], nil
}

func (ms *MemoryStore) All() (map[string]*Record, error) {
	return ms.memory, nil
}
