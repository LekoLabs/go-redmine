package redmine

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type membershipsResult struct {
	Memberships []Membership `json:"memberships"`
}

type membershipResult struct {
	Membership Membership `json:"membership"`
}

type membershipRequest struct {
	Membership Membership `json:"membership"`
}

type Membership struct {
	Id      int      `json:"id"`
	Project IdName   `json:"project"`
	User    IdName   `json:"user"`
	Roles   []IdName `json:"roles"`
	Groups  []IdName `json:"groups"`
}

func (c *Client) Memberships(projectId int) ([]Membership, error) {
	res, err := c.Get(c.endpoint + "/projects/" + strconv.Itoa(projectId) + "/memberships.json?key=" + c.apikey + c.getPaginationClause())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r membershipsResult
	if res.StatusCode == 404 {
		return nil, errors.New("not Found")
	}
	if res.StatusCode != 200 {
		err = errorFromResp(decoder, res.StatusCode)
	} else {
		err = decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}
	return r.Memberships, nil
}

func (c *Client) Membership(id int) (*Membership, error) {
	res, err := c.Get(c.endpoint + "/memberships/" + strconv.Itoa(id) + ".json?key=" + c.apikey)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r membershipResult
	if res.StatusCode == 404 {
		return nil, errors.New("not Found")
	}
	if res.StatusCode != 200 {
		err = errorFromResp(decoder, res.StatusCode)
	} else {
		err = decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}
	return &r.Membership, nil
}

func (c *Client) CreateMembership(membership Membership, userName ...string) (*Membership, error) {
	var ir membershipRequest
	ir.Membership = membership
	s, err := json.Marshal(ir)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.endpoint+"/memberships.json?key="+c.apikey, strings.NewReader(string(s)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if len(userName) > 0 {
		req.Header.Set("X-Redmine-Switch-User", userName[0])	
	}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r membershipRequest
	if res.StatusCode != 201 {
		err = errorFromResp(decoder, res.StatusCode)
	} else {
		err = decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}
	return &r.Membership, nil
}

func (c *Client) UpdateMembership(membership Membership, userName ...string) error {
	var ir membershipRequest
	ir.Membership = membership
	s, err := json.Marshal(ir)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", c.endpoint+"/memberships/"+strconv.Itoa(membership.Id)+".json?key="+c.apikey, strings.NewReader(string(s)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if len(userName) > 0 {
		req.Header.Set("X-Redmine-Switch-User", userName[0])	
	}
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return errors.New("not Found")
	}
	if res.StatusCode != 200 {
		decoder := json.NewDecoder(res.Body)
		err = errorFromResp(decoder, res.StatusCode)
	}
	if err != nil {
		return err
	}
	return err
}

func (c *Client) DeleteMembership(id int, userName ...string) error {
	req, err := http.NewRequest("DELETE", c.endpoint+"/memberships/"+strconv.Itoa(id)+".json?key="+c.apikey, strings.NewReader(""))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if len(userName) > 0 {
		req.Header.Set("X-Redmine-Switch-User", userName[0])	
	}
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return errors.New("not Found")
	}

	decoder := json.NewDecoder(res.Body)
	if res.StatusCode != 200 {
		err = errorFromResp(decoder, res.StatusCode)
	}
	return err
}
