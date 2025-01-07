package atlasiter

import (
	"context"

	"go.mongodb.org/atlas/mongodbatlas"
)

const nextPage = "next"

type projectPaginator struct {
	client  *mongodbatlas.Client
	page    int
	hasNext bool
}

// newProjectPaginator creates a new ProjectIterator.
func newProjectPaginator(client *mongodbatlas.Client) projectsIterator {
	return &projectPaginator{client, 0, true}
}

type projectsIterator interface {
	value(ctx context.Context) (*mongodbatlas.Projects, error)
	next() bool
}

// next returns true if there is at least one more project page.
func (iter *projectPaginator) next() bool {
	return iter.hasNext
}

// value fetches the next page of Atlas Project Results.
func (iter *projectPaginator) value(ctx context.Context) (*mongodbatlas.Projects, error) {
	projects, err := iter.getNextPageOfProjects(ctx)
	if err != nil {
		return nil, err
	}

	iter.setNextPage(projects)

	return projects, err
}

func (iter *projectPaginator) getNextPageOfProjects(ctx context.Context) (projects *mongodbatlas.Projects, err error) {
	if iter.hasNext {
		projects, _, err = iter.client.Projects.GetAllProjects(ctx, &mongodbatlas.ListOptions{
			PageNum: iter.page,
		})
		if err != nil {
			return nil, err
		}
	}

	return projects, nil
}

func (iter *projectPaginator) setNextPage(projects *mongodbatlas.Projects) {
	iter.hasNext = false

	for i := range projects.Links {
		if projects.Links[i].Rel == nextPage {
			iter.hasNext = true
			iter.page += 1
		}
	}
}
