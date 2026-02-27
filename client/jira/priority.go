package jira

import (
	"context"

	"github.com/surajrajput1024/go-atlassian-cloud/constants"
	"github.com/surajrajput1024/go-atlassian-cloud/types"
)

func (j *Client) GetPriorities() ([]types.PriorityResponse, error) {
	var out []types.PriorityResponse
	if err := j.doJSON(context.Background(), "GET", j.path(constants.JiraPathPriority), nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}
