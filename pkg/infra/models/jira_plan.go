package models

// PlanPageScheme represents a cursor-paginated page of plans in Jira.
type PlanPageScheme struct {
	NextPageCursor string            `json:"nextPageCursor,omitempty"` // The cursor for the next page.
	Last           bool              `json:"last,omitempty"`           // Indicates if this is the last page.
	Size           int               `json:"size,omitempty"`           // The number of items in this page.
	Total          int               `json:"total,omitempty"`          // The total number of items.
	Values         []*PlanListScheme `json:"values,omitempty"`         // The plans in this page.
}

// PlanTeamPageScheme represents a cursor-paginated page of plan teams in Jira.
type PlanTeamPageScheme struct {
	NextPageCursor string                `json:"nextPageCursor,omitempty"` // The cursor for the next page.
	Last           bool                  `json:"last,omitempty"`           // Indicates if this is the last page.
	Size           int                   `json:"size,omitempty"`           // The number of items in this page.
	Total          int                   `json:"total,omitempty"`          // The total number of items.
	Values         []*PlanTeamListScheme `json:"values,omitempty"`         // The plan teams in this page.
}

// PlanListScheme represents a plan list item in Jira.
type PlanListScheme struct {
	ID           string                   `json:"id"`                     // The ID of the plan.
	Name         string                   `json:"name"`                   // The name of the plan.
	Status       string                   `json:"status"`                 // The status of the plan.
	ScenarioID   string                   `json:"scenarioId"`             // The scenario ID of the plan.
	IssueSources []*PlanIssueSourceScheme `json:"issueSources,omitempty"` // The issue sources of the plan.
}

// PlanScheme represents the full detail of a plan in Jira.
type PlanScheme struct {
	ID                   int                              `json:"id"`                             // The ID of the plan.
	Name                 string                           `json:"name"`                           // The name of the plan.
	Status               string                           `json:"status"`                         // The status of the plan.
	LeadAccountID        string                           `json:"leadAccountId,omitempty"`         // The account ID of the plan lead.
	LastSaved            string                           `json:"lastSaved,omitempty"`             // The last saved timestamp of the plan.
	IssueSources         []*PlanIssueSourceScheme         `json:"issueSources,omitempty"`          // The issue sources of the plan.
	ExclusionRules       *PlanExclusionRulesScheme        `json:"exclusionRules,omitempty"`        // The exclusion rules of the plan.
	Scheduling           *PlanSchedulingScheme            `json:"scheduling,omitempty"`            // The scheduling configuration of the plan.
	CrossProjectReleases []*PlanCrossProjectReleaseScheme `json:"crossProjectReleases,omitempty"`  // The cross-project releases of the plan.
	CustomFields         []*PlanCustomFieldScheme         `json:"customFields,omitempty"`          // The custom fields of the plan.
	Permissions          []*PlanPermissionScheme          `json:"permissions,omitempty"`            // The permissions of the plan.
}

// PlanIssueSourceScheme represents an issue source of a plan in Jira.
type PlanIssueSourceScheme struct {
	Type  string `json:"type"`  // The type of the issue source.
	Value int    `json:"value"` // The value of the issue source.
}

// PlanExclusionRulesScheme represents the exclusion rules of a plan in Jira.
type PlanExclusionRulesScheme struct {
	NumberOfDaysToShowCompletedIssues *int  `json:"numberOfDaysToShowCompletedIssues,omitempty"` // The number of days to show completed issues. nil = 30-day default, 0 = exclude all completed.
	IssueIDs                         []int `json:"issueIds,omitempty"`                  // The excluded issue IDs.
	IssueTypeIDs                     []int `json:"issueTypeIds,omitempty"`              // The excluded issue type IDs.
	WorkStatusIDs                    []int `json:"workStatusIds,omitempty"`             // The excluded work status IDs.
	WorkStatusCategoryIDs            []int `json:"workStatusCategoryIds,omitempty"`     // The excluded work status category IDs.
	ReleaseIDs                       []int `json:"releaseIds,omitempty"`                // The excluded release IDs.
}

// PlanSchedulingScheme represents the scheduling configuration of a plan in Jira.
type PlanSchedulingScheme struct {
	Dependencies  string              `json:"dependencies"`          // The dependency scheduling mode.
	Estimation    string              `json:"estimation"`            // The estimation mode.
	InferredDates string              `json:"inferredDates"`         // The inferred dates mode.
	StartDate     *PlanDateFieldScheme `json:"startDate,omitempty"`  // The start date field configuration.
	EndDate       *PlanDateFieldScheme `json:"endDate,omitempty"`    // The end date field configuration.
}

// PlanDateFieldScheme represents a date field configuration in a plan in Jira.
type PlanDateFieldScheme struct {
	Type               string `json:"type"`                          // The type of the date field.
	DateCustomFieldID  *int64 `json:"dateCustomFieldId,omitempty"`   // The custom field ID for the date field.
}

// PlanCrossProjectReleaseScheme represents a cross-project release in a plan in Jira.
type PlanCrossProjectReleaseScheme struct {
	Name       string `json:"name,omitempty"`       // The name of the cross-project release.
	ReleaseIDs []int  `json:"releaseIds,omitempty"` // The release IDs in the cross-project release.
}

// PlanCustomFieldScheme represents a custom field configuration in a plan in Jira.
type PlanCustomFieldScheme struct {
	CustomFieldID int  `json:"customFieldId"`    // The ID of the custom field.
	Filter        bool `json:"filter,omitempty"` // Indicates if the custom field is used as a filter.
}

// PlanPermissionScheme represents a permission entry in a plan in Jira.
type PlanPermissionScheme struct {
	Type   string                      `json:"type"`             // The type of the permission.
	Holder *PlanPermissionHolderScheme `json:"holder,omitempty"` // The holder of the permission.
}

// PlanPermissionHolderScheme represents a permission holder in a plan in Jira.
type PlanPermissionHolderScheme struct {
	Type  string `json:"type"`  // The type of the permission holder.
	Value string `json:"value"` // The value of the permission holder.
}

// PlanTeamListScheme represents a plan team list item in Jira.
type PlanTeamListScheme struct {
	ID   string `json:"id"`             // The ID of the team.
	Name string `json:"name,omitempty"` // The name of the team.
	Type string `json:"type"`           // The type of the team.
}

// PlanAtlassianTeamScheme represents an Atlassian team in a plan in Jira.
type PlanAtlassianTeamScheme struct {
	ID            string   `json:"id"`                        // The ID of the Atlassian team.
	PlanningStyle string   `json:"planningStyle"`             // The planning style of the team.
	Capacity      *float64 `json:"capacity,omitempty"`        // The capacity of the team.
	IssueSourceID *int     `json:"issueSourceId,omitempty"`   // The issue source ID of the team.
	SprintLength  *int     `json:"sprintLength,omitempty"`    // The sprint length of the team.
}

// PlanOnlyTeamScheme represents a plan-only team in a plan in Jira.
type PlanOnlyTeamScheme struct {
	ID               int      `json:"id"`                          // The ID of the plan-only team.
	Name             string   `json:"name"`                        // The name of the plan-only team.
	PlanningStyle    string   `json:"planningStyle"`               // The planning style of the team.
	Capacity         *float64 `json:"capacity,omitempty"`          // The capacity of the team.
	IssueSourceID    *int     `json:"issueSourceId,omitempty"`     // The issue source ID of the team.
	SprintLength     *int     `json:"sprintLength,omitempty"`      // The sprint length of the team.
	MemberAccountIDs []string `json:"memberAccountIds,omitempty"` // The account IDs of team members.
}

// PlanCreateScheme represents the payload to create a plan in Jira.
type PlanCreateScheme struct {
	Name                 string                                `json:"name"`                            // The name of the plan.
	IssueSources         []*PlanIssueSourceCreateScheme        `json:"issueSources"`                    // The issue sources of the plan.
	Scheduling           *PlanSchedulingCreateScheme           `json:"scheduling"`                      // The scheduling configuration of the plan.
	LeadAccountID        string                                `json:"leadAccountId,omitempty"`         // The account ID of the plan lead.
	ExclusionRules       *PlanExclusionRulesCreateScheme       `json:"exclusionRules,omitempty"`        // The exclusion rules of the plan.
	Permissions          []*PlanPermissionCreateScheme         `json:"permissions,omitempty"`           // The permissions of the plan.
	CrossProjectReleases []*PlanCrossProjectReleaseCreateScheme `json:"crossProjectReleases,omitempty"` // The cross-project releases of the plan.
	CustomFields         []*PlanCustomFieldCreateScheme        `json:"customFields,omitempty"`          // The custom fields of the plan.
}

// PlanIssueSourceCreateScheme represents the payload for an issue source when creating a plan in Jira.
type PlanIssueSourceCreateScheme struct {
	Type  string `json:"type"`  // The type of the issue source.
	Value int    `json:"value"` // The value of the issue source.
}

// PlanSchedulingCreateScheme represents the payload for scheduling configuration when creating a plan in Jira.
type PlanSchedulingCreateScheme struct {
	Estimation    string                    `json:"estimation"`              // The estimation mode.
	Dependencies  string                    `json:"dependencies,omitempty"`  // The dependency scheduling mode.
	InferredDates string                    `json:"inferredDates,omitempty"` // The inferred dates mode.
	StartDate     *PlanDateFieldCreateScheme `json:"startDate,omitempty"`    // The start date field configuration.
	EndDate       *PlanDateFieldCreateScheme `json:"endDate,omitempty"`      // The end date field configuration.
}

// PlanDateFieldCreateScheme represents the payload for a date field configuration when creating a plan in Jira.
type PlanDateFieldCreateScheme struct {
	Type              string `json:"type"`                        // The type of the date field.
	DateCustomFieldID *int64 `json:"dateCustomFieldId,omitempty"` // The custom field ID for the date field.
}

// PlanExclusionRulesCreateScheme represents the payload for exclusion rules when creating a plan in Jira.
type PlanExclusionRulesCreateScheme struct {
	NumberOfDaysToShowCompletedIssues int   `json:"numberOfDaysToShowCompletedIssues"` // The number of days to show completed issues.
	IssueIDs                         []int `json:"issueIds,omitempty"`                 // The excluded issue IDs.
	IssueTypeIDs                     []int `json:"issueTypeIds,omitempty"`             // The excluded issue type IDs.
	WorkStatusIDs                    []int `json:"workStatusIds,omitempty"`            // The excluded work status IDs.
	WorkStatusCategoryIDs            []int `json:"workStatusCategoryIds,omitempty"`    // The excluded work status category IDs.
	ReleaseIDs                       []int `json:"releaseIds,omitempty"`               // The excluded release IDs.
}

// PlanPermissionCreateScheme represents the payload for a permission entry when creating a plan in Jira.
type PlanPermissionCreateScheme struct {
	Type   string                            `json:"type"`   // The type of the permission.
	Holder *PlanPermissionHolderCreateScheme `json:"holder"` // The holder of the permission.
}

// PlanPermissionHolderCreateScheme represents the payload for a permission holder when creating a plan in Jira.
type PlanPermissionHolderCreateScheme struct {
	Type  string `json:"type"`  // The type of the permission holder.
	Value string `json:"value"` // The value of the permission holder.
}

// PlanCrossProjectReleaseCreateScheme represents the payload for a cross-project release when creating a plan in Jira.
type PlanCrossProjectReleaseCreateScheme struct {
	Name       string `json:"name"`                // The name of the cross-project release.
	ReleaseIDs []int  `json:"releaseIds,omitempty"` // The release IDs in the cross-project release.
}

// PlanCustomFieldCreateScheme represents the payload for a custom field configuration when creating a plan in Jira.
type PlanCustomFieldCreateScheme struct {
	CustomFieldID int  `json:"customFieldId"`    // The ID of the custom field.
	Filter        bool `json:"filter,omitempty"` // Indicates if the custom field is used as a filter.
}

// PlanDuplicateScheme represents the payload to duplicate a plan in Jira.
type PlanDuplicateScheme struct {
	Name string `json:"name"` // The name of the duplicated plan.
}

// PlanAtlassianTeamCreateScheme represents the payload to add an Atlassian team to a plan in Jira.
type PlanAtlassianTeamCreateScheme struct {
	ID            string   `json:"id"`                        // The ID of the Atlassian team.
	PlanningStyle string   `json:"planningStyle"`             // The planning style of the team.
	Capacity      *float64 `json:"capacity,omitempty"`        // The capacity of the team.
	IssueSourceID *int     `json:"issueSourceId,omitempty"`   // The issue source ID of the team.
	SprintLength  *int     `json:"sprintLength,omitempty"`    // The sprint length of the team.
}

// PlanOnlyTeamCreateScheme represents the payload to create a plan-only team in a plan in Jira.
type PlanOnlyTeamCreateScheme struct {
	Name             string   `json:"name"`                        // The name of the plan-only team.
	PlanningStyle    string   `json:"planningStyle"`               // The planning style of the team.
	Capacity         *float64 `json:"capacity,omitempty"`          // The capacity of the team.
	IssueSourceID    *int     `json:"issueSourceId,omitempty"`     // The issue source ID of the team.
	SprintLength     *int     `json:"sprintLength,omitempty"`      // The sprint length of the team.
	MemberAccountIDs []string `json:"memberAccountIds,omitempty"` // The account IDs of team members.
}

// PlanGetsOptions represents the options for getting a list of plans in Jira.
type PlanGetsOptions struct {
	IncludeTrashed  bool   // Indicates if trashed plans should be included.
	IncludeArchived bool   // Indicates if archived plans should be included.
	Cursor          string // The cursor for pagination.
	MaxResults      int    // The maximum number of results to return.
}

// PlanTeamGetsOptions represents the options for getting a list of teams in a plan in Jira.
type PlanTeamGetsOptions struct {
	Cursor     string // The cursor for pagination.
	MaxResults int    // The maximum number of results to return.
}
