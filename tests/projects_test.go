package redmine_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	lkredmine "github.com/LekoLabs/go-redmine"
)

func createSingleProject(t *testing.T, project_id string, client_redmine *lkredmine.Client) {
	preparedProject := lkredmine.ProjectToCreate{
		Name:        "Test Project X",
		Identifier:  project_id,
		Description: "This is the test project X.",
		IsPublic:    false,
	}
	project, err := client_redmine.CreateProject(preparedProject)
	assert.Nil(t, err, "Found the following error when trying to create the project:", err)
	// fmt.Printf("Succeeded in creating the following project: %v", project)
	_ = project
}

func getSingleProject(t *testing.T, project_id string, client_redmine *lkredmine.Client) {
	project, err := client_redmine.Project(project_id)
	assert.Nilf(t, err, "Found the following error when trying to fetch project %s: %v", project_id, err)
	assert.NotEmpty(t, project, "Project supposed to be not empty.")
	// fmt.Printf("Project found: %v", project)
	_ = project
}

func deleteSingleProject(t *testing.T, project_id string, client_redmine *lkredmine.Client) {
	loadEnvsFromDotenv(t)
	err := client_redmine.DeleteProject(project_id)
	assert.Nil(t, err, "Found the following error when trying to delete the project:", err)
	// fmt.Printf("Succeeded in deleting the following project: %s", project_id)
}

func Test_CreateDeleteSingleProject(t *testing.T) {
	loadEnvsFromDotenv(t)
	stat, host, apik := getRedmineHostApikFromEnvs()
	if stat {
		if doesRedmineHostExist(host) {
			clientRedmine := lkredmine.NewClient(host, apik)
			project_id := "test-project-x"
			createSingleProject(t, project_id, clientRedmine)
			getSingleProject(t, project_id, clientRedmine)
			deleteSingleProject(t, project_id, clientRedmine)
			return
		}
		t.Log("Environment variables found but redmine host not detected; test not run.")
		return
	}
	t.Log("Environment variables not found; test not run.")
}

func Test_GetListProjects(t *testing.T) {
	loadEnvsFromDotenv(t)
	stat, host, apik := getRedmineHostApikFromEnvs()
	if stat {
		if doesRedmineHostExist(host) {
			clientRedmine := lkredmine.NewClient(host, apik)
			projects, err := clientRedmine.Projects()
			assert.Nil(t, err, "Found the following error when trying to get projects:", err)
			// fmt.Printf("Found the following projects: %v", projects)
			_ = projects
			return
		}
		t.Log("Environment variables found but redmine host not detected; test not run.")
		return
	}
	t.Log("Environment variables not found; test not run.")
}
