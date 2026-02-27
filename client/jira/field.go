package jira

import (
	"context"

	"github.com/surajrajput1024/go-atlassian-cloud/constants"
	"github.com/surajrajput1024/go-atlassian-cloud/types"
)

func (j *Client) GetFields() ([]types.FieldResponse, error) {
	var out []types.FieldResponse
	if err := j.doJSON(context.Background(), "GET", j.path(constants.JiraPathField), nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}
