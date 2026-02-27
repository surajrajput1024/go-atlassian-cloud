package jira

import (
	"context"

	"github.com/surajrajput1024/go-atlassian-cloud/constants"
	"github.com/surajrajput1024/go-atlassian-cloud/types"
)

func (j *Client) GetStatuses() ([]types.StatusResponse, error) {
	var out []types.StatusResponse
	if err := j.doJSON(context.Background(), "GET", j.path(constants.JiraPathStatus), nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}
