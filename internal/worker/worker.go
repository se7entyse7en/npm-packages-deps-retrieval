package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/queue"
	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/store"
)

type dependencies struct {
	Dependencies map[string]string `json:"dependencies,omitempty"`
}

func (d *dependencies) normalize() {
	for name, version := range d.Dependencies {
		if strings.HasPrefix(version, "~") || strings.HasPrefix(version, "^") {
			version = version[1:]
		}

		matched, err := regexp.MatchString(`^\d`, version)
		if err != nil {
			panic(err)
		}

		if !matched {
			fmt.Printf("warning, skipping package `%s` with version `%s`",
				name, version)
			delete(d.Dependencies, name)
			continue
		}

		d.Dependencies[name] = version
	}
}

type DependenciesFetcher struct {
	s store.Store
}

func NewDependenciesFetcher(s store.Store) *DependenciesFetcher {
	return &DependenciesFetcher{s: s}
}

func (f *DependenciesFetcher) Fetch(packageName, packageVersion string) (map[string]string, error) {
	fmt.Printf("getting dependencies for package `%s`, version `%s`\n", packageName, packageVersion)
	endpoint := fmt.Sprintf("https://registry.npmjs.org/%s/%s",
		packageName, packageVersion)
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 status code")
	}

	deps := &dependencies{}
	if err := json.NewDecoder(resp.Body).Decode(deps); err != nil {
		return nil, err
	}

	deps.normalize()
	return deps.Dependencies, nil
}

func (f *DependenciesFetcher) Store(ctx context.Context, id, packageName, packageVersion string, deps map[string]string) error {
	record := &store.Record{
		ID:           id,
		Name:         packageName,
		Version:      packageVersion,
		Dependencies: deps,
	}

	fmt.Printf("storing record: %v\n", record)
	return f.s.Save(ctx, record)
}

type Worker struct {
	q queue.Queue
	f *DependenciesFetcher
}

func NewWorker(q queue.Queue, s store.Store) *Worker {
	return &Worker{q: q, f: NewDependenciesFetcher(s)}
}

func (w *Worker) Start(ctx context.Context) error {
	for {
		len, err := w.q.Len(ctx)
		if err != nil {
			return nil
		}

		if len <= 0 {
			break
		}

		fmt.Printf("queue length: %d\n", len)
		id, err := w.q.Remove(ctx)
		if err != nil {
			return err
		}

		splitted := strings.Split(id, "@")
		packageName, packageVersion := splitted[0], splitted[1]
		deps, err := w.f.Fetch(packageName, packageVersion)
		if err != nil {
			return err
		}

		if err := w.f.Store(ctx, id, packageName, packageVersion, deps); err != nil {
			return err
		}

		if err := w.enqueueDependencies(ctx, deps); err != nil {
			return err
		}
	}

	dbContent, err := w.f.s.All(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("store: %v\n", dbContent)

	return nil
}

func (w *Worker) enqueueDependencies(ctx context.Context, deps map[string]string) error {
	for name, version := range deps {
		id := fmt.Sprintf("%s@%s", name, version)
		fmt.Printf("enqueuing package: %s\n", id)
		if err := w.q.Add(ctx, id); err != nil {
			return err
		}
	}

	return nil
}
