package jira

import (
	"context"
	"net/url"

	"github.com/surajrajput1024/go-atlassian-cloud/constants"
	"github.com/surajrajput1024/go-atlassian-cloud/types"
)

func (j *Client) GetGroup(groupID, groupName string) (*types.GroupResponse, error) {
	return j.Groups.GetGroup(groupID, groupName)
}

func (j *Client) CreateGroup(req *types.GroupCreateRequest) (*types.GroupResponse, error) {
	return j.Groups.CreateGroup(req)
}

func (j *Client) DeleteGroup(groupID, groupName, swapGroup, swapGroupID string) error {
	return j.Groups.DeleteGroup(groupID, groupName, swapGroup, swapGroupID)
}

func (s *GroupService) GetGroup(groupID, groupName string) (*types.GroupResponse, error) {
	q := url.Values{}
	if groupID != "" {
		q.Set("groupId", groupID)
	}
	if groupName != "" {
		q.Set("groupname", groupName)
	}
	path := s.c.path(constants.JiraPathGroup)
	if len(q) > 0 {
		path = path + "?" + q.Encode()
	}
	var out types.GroupResponse
	if err := s.c.doJSON(context.Background(), "GET", path, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *GroupService) CreateGroup(req *types.GroupCreateRequest) (*types.GroupResponse, error) {
	var out types.GroupResponse
	if err := s.c.doJSON(context.Background(), "POST", s.c.path(constants.JiraPathGroup), req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *GroupService) DeleteGroup(groupID, groupName, swapGroup, swapGroupID string) error {
	q := url.Values{}
	if groupID != "" {
		q.Set("groupId", groupID)
	}
	if groupName != "" {
		q.Set("groupname", groupName)
	}
	if swapGroup != "" {
		q.Set("swapGroup", swapGroup)
	}
	if swapGroupID != "" {
		q.Set("swapGroupId", swapGroupID)
	}
	path := s.c.path(constants.JiraPathGroup)
	if len(q) > 0 {
		path = path + "?" + q.Encode()
	}
	return s.c.delete(context.Background(), path)
}
