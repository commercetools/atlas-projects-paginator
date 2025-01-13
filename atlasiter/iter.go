package atlasiter

import (
	"context"
	"iter"

	"go.mongodb.org/atlas/mongodbatlas"
)

// ProjectsService is a minimal interface for the Atlas projects service.
type ProjectsService interface {
	GetAllProjects(context.Context, *mongodbatlas.ListOptions) (*mongodbatlas.Projects, *mongodbatlas.Response, error)
}

// AllProjects returns an iterator that will yield all projects in the Atlas organization.
func AllProjects(ctx context.Context, projectsService ProjectsService) iter.Seq2[*mongodbatlas.Project, error] {
	return func(yield func(*mongodbatlas.Project, error) bool) {
		for projects, err := range allProjectPages(ctx, projectsService) {
			if err != nil {
				yield(nil, err)
				return
			}

			for _, project := range projects.Results {
				if !yield(project, nil) {
					return
				}
			}
		}
	}
}

func allProjectPages(ctx context.Context, projectsService ProjectsService) iter.Seq2[*mongodbatlas.Projects, error] {
	projectIterator := newProjectPaginator(projectsService)

	return func(yield func(*mongodbatlas.Projects, error) bool) {
		for projectIterator.next() {
			projects, err := projectIterator.value(ctx)

			if !yield(projects, err) {
				return
			}
		}
	}
}
