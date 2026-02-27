package jira

import (
	"context"

	"github.com/surajrajput1024/go-atlassian-cloud/constants"
	"github.com/surajrajput1024/go-atlassian-cloud/types"
)

func (j *Client) GetCurrentUser() (*types.CurrentUserResponse, error) {
	var out types.CurrentUserResponse
	if err := j.doJSON(context.Background(), "GET", j.path(constants.JiraPathMyself), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
