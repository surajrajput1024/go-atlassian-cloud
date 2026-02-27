package types

import (
	"encoding/json"

	"github.com/surajrajput1024/go-atlassian-cloud/util"
)

type IssueTypeResponse struct {
	ID             string `json:"-"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Self           string `json:"self"`
	Subtask        bool   `json:"subtask"`
	HierarchyLevel int    `json:"hierarchyLevel"`
	IconURL        string `json:"iconUrl"`
}

func (r *IssueTypeResponse) UnmarshalJSON(data []byte) error {
	type issueTypeResponse IssueTypeResponse
	var aux struct {
		issueTypeResponse
		ID json.RawMessage `json:"id"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*r = IssueTypeResponse(aux.issueTypeResponse)
	id, err := util.ParseStringOrNumber(aux.ID)
	if err != nil {
		return err
	}
	r.ID = id
	return nil
}
