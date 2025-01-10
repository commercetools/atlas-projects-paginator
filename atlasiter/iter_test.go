package atlasiter

import (
	"context"
	"testing"

	"go.mongodb.org/atlas/mongodbatlas"
)

type mockProjectsService struct {
	projects mongodbatlas.Projects
}

func (m *mockProjectsService) GetAllProjects(ctx context.Context, opts *mongodbatlas.ListOptions) (*mongodbatlas.Projects, *mongodbatlas.Response, error) {
	return &m.projects, nil, nil
}

func TestAllProjects(t *testing.T) {
	t.Run("Should yield all projects given", func(t *testing.T) {
		// Given a mockProjectsService with two projects
		projectsService := &mockProjectsService{mongodbatlas.Projects{
			Results: []*mongodbatlas.Project{
				{
					ID:   "1",
					Name: "Project 1",
				},
				{
					ID:   "2",
					Name: "Project 2",
				},
			},
		}}
		ctx := context.Background()
		projects := []*mongodbatlas.Project{}

		// When we iterate over all projects
		for project, err := range AllProjects(ctx, projectsService) {
			projects = append(projects, project)
			if err != nil {
				t.Errorf("got %v, want nil", err)
			}
		}

		// Then we should get both projects
		if len(projects) != 2 {
			t.Errorf("got %d, want 2", len(projects))
		}
	})

	t.Run("Should yield an error if GetAllProjects fails", func(t *testing.T) {
		// Given a mockProjectsService that fails
		projectsService := &mockProjectsService{}
		ctx := context.Background()
		projects := []*mongodbatlas.Project{}

		// When we iterate over all AllProjects
		for project, err := range AllProjects(ctx, projectsService) {
			projects = append(projects, project)
			if err == nil {
				t.Errorf("got nil, want an error")
			}
		}

		// Then we should get an Error
		if len(projects) != 0 {
			t.Errorf("got %d, want 0", len(projects))
		}
	})

	t.Run("Should stop iterating if yield returns false", func(t *testing.T) {
		// Given a mockProjectsService with two TestAllProjects
		projectsService := &mockProjectsService{mongodbatlas.Projects{
			Results: []*mongodbatlas.Project{
				{
					ID:   "1",
					Name: "Project 1",
				},
				{
					ID:   "2",
					Name: "Project 2",
				},
			},
		}}
		ctx := context.Background()
		projects := []*mongodbatlas.Project{}

		// When we iterate over all projects and stop after the first one
		for project, err := range AllProjects(ctx, projectsService) {
			projects = append(projects, project)
			if err != nil {
				t.Errorf("got %v, want nil", err)
			}
			break
		}

		// Then we should get only the first TestAllProjects
		if len(projects) != 1 {
			t.Errorf("got %d, want 1", len(projects))
		}
	})
}
