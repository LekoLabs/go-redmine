package redmine_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	lkredmine "github.com/LekoLabs/go-redmine"
)

func Load_Dotenv(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		t.Error("Dotenv file not found")
	} else {
		t.Log("Dotenv file found; environment variables loaded")
	}
}

func createSingleProject(t *testing.T, project_id string, client_redmine *lkredmine.Client) {
	preparedProject := lkredmine.ProjectToCreate{
		Name:        "Test Project X",
		Identifier:  project_id,
		Description: "This is the test project X.",
		IsPublic:    false,
	}
	project, err := client_redmine.CreateProject(preparedProject)
	assert.Nil(t, err, "Found the following error when trying to create the project:", err)
	fmt.Printf("Succeeded in creating the following project: %v", project)
}

func deleteSingleProject(t *testing.T, project_id string, client_redmine *lkredmine.Client) {
	Load_Dotenv(t)
	err := client_redmine.DeleteProject(project_id)
	assert.Nil(t, err, "Found the following error when trying to delete the project:", err)
	fmt.Printf("Succeeded in deleting the following project: %s", project_id)
}

func Test_CreateDeleteSingleProject(t *testing.T) {
	Load_Dotenv(t)

	client_redmine := lkredmine.NewClient(
		os.Getenv("REDMINE_HOST"),
		os.Getenv("REDMINE_API_KEY"),
	)

	project_id := "test-project-x"
	createSingleProject(t, project_id, client_redmine)
	deleteSingleProject(t, project_id, client_redmine)
}

func Test_GetListProjects(t *testing.T) {
	Load_Dotenv(t)

	CLIENT_REDMINE := lkredmine.NewClient(
		os.Getenv("REDMINE_HOST"),
		os.Getenv("REDMINE_API_KEY"),
	)
	projects, err := CLIENT_REDMINE.Projects()
	assert.Nil(t, err, "Found the following error when trying to get projects:", err)
	fmt.Printf("Found the following projects: %v", projects)
}
