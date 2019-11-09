package api

import (
	"context"
	fmt "fmt"

	"github.com/se7entyse7en/npm-packages-deps-retrieval/internal/store"
)

type ApiServer struct {
	s store.Store
}

func (s *ApiServer) GetDependencies(ctx context.Context, r *DependenciesRequest) (*DependenciesResponse, error) {
	name, version := s.parsePackage(r)
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
	packageID := fmt.Sprintf("%s@%s", packageName, packageVersion)
	record, err := s.s.Get(ctx, packageID)
	if err != nil {
		return nil, err
	}

	if record == nil {
		return nil, fmt.Errorf("record `%s` not found", packageID)
	}

	var deps []*Dependency
	for name, version := range record.Dependencies {
		depsSubTree, err := s.buildDependenciesTree(ctx, name, version)
		if err != nil {
			return nil, err
		}

		deps = append(deps, depsSubTree)
	}

	return &Dependency{
		Name:         packageName,
		Version:      packageVersion,
		Dependencies: deps,
	}, nil
}

func NewApiServer(s store.Store) *ApiServer {
	return &ApiServer{s: s}
}
