package jira

import (
	"github.com/surajsinghrajput/go-atlassian-cloud/constants"
	"github.com/surajsinghrajput/go-atlassian-cloud/types"
)

func (j *Client) GetFields() ([]types.FieldResponse, error) {
	var out []types.FieldResponse
	if err := j.getJSON(j.path(constants.JiraPathField), &out); err != nil {
		return nil, err
	}
	return out, nil
}
