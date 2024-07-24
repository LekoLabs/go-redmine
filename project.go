package redmine

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type projectRequest struct {
	Project Project `json:"project"`
}

type projectResult struct {
	Project Project `json:"project"`
}

type projectsResult struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	Id                 int               `json:"id"`
	Parent             IdName            `json:"parent,omitempty"`
	Name               string            `json:"name"`
	Identifier         string            `json:"identifier"`
	Description        string            `json:"description,omitempty"`
	CreatedOn          string            `json:"created_on,omitempty"`
	UpdatedOn          string            `json:"updated_on,omitempty"`
	IsPublic           bool              `json:"is_public,omitempty"`
	ParentID           int               `json:"parent_id,omitempty"`
	InheritMembers     bool              `json:"inherit_members,omitempty"`
	TrackerIDs         []int             `json:"tracker_ids,omitempty"`
	EnabledModuleNames []string          `json:"enabled_module_names,omitempty"`
	CustomFields       []*CustomField    `json:"custom_fields,omitempty"`
	CustomFieldValues  map[string]string `json:"custom_field_values,omitempty"`
}

func (c *Client) Project(id string) (*Project, error) {
	res, err := c.Get(c.endpoint + "/projects/" + id + ".json?key=" + c.apikey)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r projectResult
	if res.StatusCode != 200 {
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}
	return &r.Project, nil
}

func (c *Client) Projects() ([]Project, error) {
	url := fmt.Sprintf("%s/projects.json?key=%s%s", c.endpoint, c.apikey, c.getPaginationClause())
	return c.fetchProjects(url)
}

func (c *Client) ProjectsByFilter(f map[string]string) ([]Project, error) {
	filter := mapToQueryString(f)
	url := fmt.Sprintf("%s/projects.json?%s&key=%s%s", c.endpoint, filter, c.apikey, c.getPaginationClause())
	return c.fetchProjects(url)
}

func (c *Client) fetchProjects(url string) ([]Project, error) {
	res, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r projectsResult
	if res.StatusCode != 200 {
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}
	return r.Projects, nil
}

func (c *Client) CreateProject(project Project) (*Project, error) {
	var ir projectRequest
	ir.Project = project
	s, err := json.Marshal(ir)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.endpoint+"/projects.json?key="+c.apikey, strings.NewReader(string(s)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r projectRequest
	if res.StatusCode != 201 {
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}
	return &r.Project, nil
}

func (c *Client) UpdateProject(project Project) error {
	var ir projectRequest
	ir.Project = project
	s, err := json.Marshal(ir)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", c.endpoint+"/projects/"+strconv.Itoa(project.Id)+".json?key="+c.apikey, strings.NewReader(string(s)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return errors.New("not found")
	}
	if res.StatusCode != 200 {
		decoder := json.NewDecoder(res.Body)
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	}
	if err != nil {
		return err
	}
	return err
}

func (c *Client) DeleteProject(id string) error {
	req, err := http.NewRequest("DELETE", c.endpoint+"/projects/"+id+".json?key="+c.apikey, strings.NewReader(""))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return errors.New("not found")
	}

	decoder := json.NewDecoder(res.Body)
	if res.StatusCode != 204 {
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	}
	return err
}
