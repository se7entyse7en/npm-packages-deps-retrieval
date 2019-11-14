package api

import (
	"context"
	fmt "fmt"
	"sync"

	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/store"
	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/worker"
)

type ApiServer struct {
	s store.Store
	f *worker.DependenciesFetcher
}

func (s *ApiServer) GetDependencies(ctx context.Context, r *DependenciesRequest) (*DependenciesResponse, error) {
	name, version := s.parsePackage(r)
	fmt.Printf("received request for `%s@%s`\n", name, version)
	depsTree, err := s.buildDependenciesTree(ctx, name, version)
	if err != nil {
		return nil, err
	}

	return &DependenciesResponse{Dependencies: depsTree}, nil
}

func (s *ApiServer) parsePackage(r *DependenciesRequest) (string, string) {
	var version string
	if version = r.GetVersion(); version == "" {
		version = "latest"
	}

	return r.GetName(), version
}

func (s *ApiServer) buildDependenciesTree(ctx context.Context, packageName, packageVersion string) (*Dependency, error) {
	deps, err := s.getDependencies(ctx, packageName, packageVersion)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	index := 0
	depsTree := make([]*Dependency, len(deps))
	errCh := make(chan error, len(deps))
	for name, version := range deps {
		wg.Add(1)

		go func(i int, name, version string) {
			defer wg.Done()

			depsSubTree, err := s.buildDependenciesTree(ctx, name, version)
			if err != nil {
				errCh <- err
				return
			}

			depsTree[i] = depsSubTree
		}(index, name, version)

		index++
	}

	wg.Wait()

	select {
	case err := <-errCh:
		return nil, err
	default:
		return &Dependency{
			Name:         packageName,
			Version:      packageVersion,
			Dependencies: depsTree,
		}, nil
	}
}

func (s *ApiServer) getDependencies(ctx context.Context, packageName, packageVersion string) (map[string]string, error) {
	packageID := fmt.Sprintf("%s@%s", packageName, packageVersion)
	fmt.Printf("looking for dependencies for `%s` in database\n", packageID)
	record, err := s.s.Get(ctx, packageID)
	if err != nil && err.Error() != "mongo: no documents in result" {
		return nil, err
	}

	if record != nil {
		return record.Dependencies, nil
	}

	fmt.Printf("cannot find dependencies for `%s`, fetching now\n", packageID)
	deps, err := s.f.Fetch(packageName, packageVersion)
	if err != nil {
		return nil, err
	}

	if err := s.f.Store(ctx, packageID, packageName, packageVersion, deps); err != nil {
		return nil, err
	}

	return deps, nil
}

func NewApiServer(s store.Store) *ApiServer {
	return &ApiServer{s: s, f: worker.NewDependenciesFetcher(s)}
}
