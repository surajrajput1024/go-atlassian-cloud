package jira

import (
	"github.com/surajsinghrajput/go-atlassian-cloud/constants"
	"github.com/surajsinghrajput/go-atlassian-cloud/types"
)

func (j *Client) GetPriorities() ([]types.PriorityResponse, error) {
	var out []types.PriorityResponse
	if err := j.getJSON(j.path(constants.JiraPathPriority), &out); err != nil {
		return nil, err
	}
	return out, nil
}
