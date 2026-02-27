package jira

import "github.com/surajrajput1024/go-atlassian-cloud/types"

var (
	MockPermissionSchemeListResponse = types.PermissionSchemeListResponse{
		PermissionSchemes: []types.PermissionSchemeResponse{{ID: 1, Name: "Default"}},
	}
	MockPermissionGrantsResponse = types.PermissionGrantsResponse{
		Permissions: []types.PermissionGrant{{ID: 100, Permission: "BROWSE_PROJECTS"}},
	}
	MockPermissionGrant = types.PermissionGrant{
		ID:         100,
		Permission: "BROWSE_PROJECTS",
		Holder:     types.PermissionHolder{Type: "group", Parameter: "jira-users"},
	}
	MockProjectResponse       = types.ProjectResponse{ID: "10000", Key: "PROJ", Name: "Test"}
	MockProjectSearchResponse = types.ProjectSearchResponse{
		Total: 1,
		Values: []types.ProjectResponse{{Key: "P1"}},
	}
	MockPermissionSchemeResponse = types.PermissionSchemeResponse{
		ID: 10000, Name: "Example scheme", Description: "desc",
	}
	MockProjectRolesMapResponse = types.ProjectRolesMapResponse{
		"Developers": "https://site.atlassian.net/rest/api/3/project/PROJ/role/10000",
	}
	MockProjectRoleResponse = types.ProjectRoleResponse{
		ID: 10000, Name: "Developers", Description: "Dev role",
		Actors: []types.ProjectRoleActor{{ID: 1, Type: "atlassian-group-role-actor"}},
	}
	MockProjectRoleResponseWithActors = types.ProjectRoleResponse{
		ID: 10000, Name: "Developers",
		Actors: []types.ProjectRoleActor{{ID: 1, Type: "atlassian-group-role-actor"}},
	}
	MockGroupResponse = types.GroupResponse{
		GroupID: "group-uuid", Name: "jira-users", Self: "https://site/group?groupId=group-uuid",
	}
	MockGroupResponseCreated = types.GroupResponse{
		GroupID: "new-group-uuid", Name: "power-users",
	}
	MockWorkflowSchemeProjectAssociationsResponse = types.WorkflowSchemeProjectAssociationsResponse{
		Values: []types.WorkflowSchemeProjectAssociation{{
			ProjectIDs:     []string{"10000"},
			WorkflowScheme: types.WorkflowSchemeRef{ID: 10032, Name: "Example scheme"},
		}},
	}
)
