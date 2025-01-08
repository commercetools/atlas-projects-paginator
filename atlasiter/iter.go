package atlasiter

import (
	"context"
	"iter"

	"go.mongodb.org/atlas/mongodbatlas"
)

// AllProjects returns an iterator that will yield all projects in the Atlas organization.
func AllProjects(ctx context.Context, client *mongodbatlas.Client) iter.Seq2[*mongodbatlas.Project, error] {
	return func(yield func(*mongodbatlas.Project, error) bool) {
		for projects, err := range allProjectPages(ctx, client) {
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

func allProjectPages(ctx context.Context, client *mongodbatlas.Client) iter.Seq2[*mongodbatlas.Projects, error] {
	projectIterator := newProjectPaginator(client)

	return func(yield func(*mongodbatlas.Projects, error) bool) {
		for projectIterator.next() {
			projects, err := projectIterator.value(ctx)

			if !yield(projects, err) {
				return
			}
		}
	}
}
