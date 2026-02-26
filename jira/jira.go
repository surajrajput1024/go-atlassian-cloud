package jira

import jiraclient "github.com/surajrajput1024/go-atlassian-cloud/client/jira"

type Client = jiraclient.Client
type ProjectSearchParams = jiraclient.ProjectSearchParams

var New = jiraclient.New
