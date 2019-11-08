package worker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/queue"
	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/store"
)

type Worker struct {
	q queue.Queue
	s store.Store
}

func NewWorker(q queue.Queue, s store.Store) *Worker {
	return &Worker{q: q, s: s}
}

type dependencies struct {
	Dependencies map[string]string `json:"dependencies,omitempty"`
}

func (d *dependencies) normalize() {
	for k, v := range d.Dependencies {
		if strings.HasPrefix(v, "~") || strings.HasPrefix(v, "^") {
			v = v[1:]
		}

		d.Dependencies[k] = v
	}
}

func (w *Worker) Start() error {
	for {
		len, err := w.q.Len()
		if err != nil {
			return nil
		}

		if len <= 0 {
			break
		}

		fmt.Printf("queue length: %d\n", len)
		id, err := w.q.Remove()
		if err != nil {
			return err
		}

		splitted := strings.Split(id, "@")
		packageName, packageVersion := splitted[0], splitted[1]
		deps, err := w.getDependencies(packageName, packageVersion)
		if err != nil {
			return err
		}

		if err := w.storeDependencies(id, packageName, packageVersion, deps); err != nil {
			return err
		}

		if err := w.enqueueDependencies(deps); err != nil {
			return err
		}
	}

	dbContent, err := w.s.All()
	if err != nil {
		return err
	}

	fmt.Printf("store: %v\n", dbContent)

	return nil
}

func (w *Worker) getDependencies(packageName, packageVersion string) (*dependencies, error) {
	fmt.Printf("getting dependencies for package `%s`, version `%s`\n", packageName, packageVersion)
	endpoint := fmt.Sprintf("https://registry.npmjs.org/%s/%s",
		packageName, packageVersion)
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	deps := &dependencies{}
	if err := json.NewDecoder(resp.Body).Decode(deps); err != nil {
		return nil, err
	}

	deps.normalize()
	return deps, nil
}

func (w *Worker) storeDependencies(id, packageName, packageVersion string, deps *dependencies) error {
	record := &store.Record{
		ID:           id,
		Name:         packageName,
		Version:      packageVersion,
		Dependencies: deps.Dependencies,
	}
	fmt.Printf("storing record: %v\n", record)
	return w.s.Save(record)
}

func (w *Worker) enqueueDependencies(deps *dependencies) error {
	for name, version := range deps.Dependencies {
		id := fmt.Sprintf("%s@%s", name, version)
		fmt.Printf("enqueuing package: %s\n", id)
		if err := w.q.Add(id); err != nil {
			return err
		}
	}

	return nil
}
