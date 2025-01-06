package atlaspaginator

import (
	"context"

	"go.mongodb.org/atlas/mongodbatlas"
)

const nextPage = "next"

type projectPaginator struct {
	client  *mongodbatlas.Client
	next    int
	hasNext bool
}

// NewProjectPaginator creates a new ProjectIterator.
func NewProjectPaginator(client *mongodbatlas.Client) ProjectsIterator {
	return &projectPaginator{client, 0, true}
}

type ProjectsIterator interface {
	Value(ctx context.Context) (*mongodbatlas.Projects, error)
	Next() bool
}

// Next returns true if there is at least one more project page.
func (iter *projectPaginator) Next() bool {
	return iter.hasNext
}

// Value fetches the next page of Atlas Project Results.
func (iter *projectPaginator) Value(ctx context.Context) (*mongodbatlas.Projects, error) {
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
			PageNum: iter.next,
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
			iter.next += 1
		}
	}
}
