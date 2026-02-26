package jira

import jiraclient "github.com/surajsinghrajput/go-atlassian-cloud/client/jira"

type Client = jiraclient.Client
type ProjectSearchParams = jiraclient.ProjectSearchParams

var New = jiraclient.New
