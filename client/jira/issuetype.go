package jira

import (
	"net/url"

	"github.com/surajrajput1024/go-atlassian-cloud/constants"
	"github.com/surajrajput1024/go-atlassian-cloud/types"
)

func (j *Client) GetIssueTypes() ([]types.IssueTypeResponse, error) {
	var out []types.IssueTypeResponse
	if err := j.getJSON(j.path(constants.JiraPathIssueType), &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (j *Client) GetIssueType(id string) (*types.IssueTypeResponse, error) {
	var out types.IssueTypeResponse
	path := j.path(constants.JiraPathIssueType, url.PathEscape(id))
	if err := j.getJSON(path, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
