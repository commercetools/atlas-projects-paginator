package atlasiter

import (
	"context"

	"go.mongodb.org/atlas/mongodbatlas"
)

const nextPage = "next"

type projectPaginator struct {
	projectsService ProjectsService
	page            int
	hasNext         bool
}

// newProjectPaginator creates a new ProjectIterator.
func newProjectPaginator(projectsService ProjectsService) projectsIterator {
	return &projectPaginator{projectsService, 0, true}
}

type projectsIterator interface {
	value(ctx context.Context) (*mongodbatlas.Projects, error)
	next() bool
}

// next returns true if there is at least one more project page.
func (p *projectPaginator) next() bool {
	return p.hasNext
}

// value fetches the next page of Atlas Project Results.
func (p *projectPaginator) value(ctx context.Context) (*mongodbatlas.Projects, error) {
	projects, err := p.getNextPageOfProjects(ctx)
	if err != nil {
		return nil, err
	}

	p.setNextPage(projects)

	return projects, err
}

func (p *projectPaginator) getNextPageOfProjects(ctx context.Context) (projects *mongodbatlas.Projects, err error) {
	if p.hasNext {
		projects, _, err = p.projectsService.GetAllProjects(ctx, &mongodbatlas.ListOptions{
			PageNum: p.page,
		})
		if err != nil {
			return nil, err
		}
	}

	return projects, nil
}

func (p *projectPaginator) setNextPage(projects *mongodbatlas.Projects) {
	p.hasNext = false

	for i := range projects.Links {
		if projects.Links[i].Rel == nextPage {
			p.hasNext = true
			p.page += 1
		}
	}
}
