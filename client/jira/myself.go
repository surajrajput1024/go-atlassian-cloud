package jira

import (
	"github.com/surajrajput1024/go-atlassian-cloud/constants"
	"github.com/surajrajput1024/go-atlassian-cloud/types"
)

func (j *Client) GetCurrentUser() (*types.CurrentUserResponse, error) {
	var out types.CurrentUserResponse
	if err := j.getJSON(j.path(constants.JiraPathMyself), &out); err != nil {
		return nil, err
	}
	return &out, nil
}
