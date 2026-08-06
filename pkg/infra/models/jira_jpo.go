package models

// JPOPlanScheme represents a JPO (Advanced Roadmaps) plan configuration.
type JPOPlanScheme struct {
	ID                 int                    `json:"id"`
	Title              string                 `json:"title"`
	PlanningUnit       string                 `json:"planningUnit,omitempty"`
	IssueSources       []*JPOIssueSourceScheme `json:"issueSources,omitempty"`
	RankAgainstStories bool                   `json:"rankAgainstStories,omitempty"`
	CreatedTimestamp   int64                  `json:"createdTimestamp,omitempty"`
}

// JPOIssueSourceScheme represents an issue source in a JPO plan.
type JPOIssueSourceScheme struct {
	ID    int    `json:"id"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

// JPOPlanDetailRequestScheme is the request body for POST /rest/jpo/1.0/plans/detail.
type JPOPlanDetailRequestScheme struct {
	PlanID int `json:"planId"`
}

// JPOPlanDetailResponseScheme is the response from POST /rest/jpo/1.0/plans/detail.
type JPOPlanDetailResponseScheme struct {
	ID           int                     `json:"id"`
	Title        string                  `json:"title"`
	IssueSources []*JPOIssueSourceScheme `json:"issueSources,omitempty"`
	Scenarios    []*JPOScenarioScheme    `json:"scenarios,omitempty"`
}

// JPOScenarioScheme represents a scenario (sandbox) within a JPO plan.
type JPOScenarioScheme struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// JPOBacklogRequestScheme is the request body for POST /rest/jpo/1.0/backlog.
type JPOBacklogRequestScheme struct {
	PlanID     int                      `json:"planId"`
	ScenarioID int                      `json:"scenarioId"`
	Filter     *JPOBacklogFilterScheme  `json:"filter"`
}

// JPOBacklogFilterScheme controls which issues the backlog endpoint returns.
type JPOBacklogFilterScheme struct {
	IncludeCompleted            bool  `json:"includeCompleted"`
	IncludeCompletedSince       int64 `json:"includeCompletedSince"`
	IncludeIssueLinks           bool  `json:"includeIssueLinks"`
	PerformDependencyCompletion bool  `json:"performDependencyCompletion"`
}

// JPOBacklogResponseScheme is the response from POST /rest/jpo/1.0/backlog.
type JPOBacklogResponseScheme struct {
	Issues []*JPOBacklogIssueScheme `json:"issues"`
	More   bool                     `json:"more"`
}

// JPOBacklogIssueScheme represents an issue in the JPO backlog response.
// IssueKey is the numeric issue number (e.g. 544), not the full key (e.g. OCPSTRAT-544).
// Type, Status, Project, and Parent are numeric JIRA IDs that need resolution via IDResolver.
type JPOBacklogIssueScheme struct {
	ID             string                     `json:"id"`
	IssueKey       int                        `json:"issueKey"`
	IssueSources   []int                      `json:"issueSources,omitempty"`
	JiraValues     *JPOBacklogJiraValuesScheme `json:"jiraValues"`
	ScenarioValues map[string]interface{}     `json:"scenarioValues,omitempty"`
}

// JPOBacklogJiraValuesScheme holds the JIRA field values for a backlog issue.
// Numeric fields (Type, Project, Components) are JIRA internal IDs.
type JPOBacklogJiraValuesScheme struct {
	Type          int                    `json:"type"`
	Project       int                    `json:"project"`
	Status        string                 `json:"status"`
	Summary       string                 `json:"summary"`
	LexoRank      string                 `json:"lexoRank,omitempty"`
	Labels        []string               `json:"labels,omitempty"`
	Excluded      bool                   `json:"excluded,omitempty"`
	Components    []int                  `json:"components,omitempty"`
	FixVersions   []string               `json:"fixVersions,omitempty"`
	BaselineStart *int64                 `json:"baselineStart,omitempty"`
	BaselineEnd   *int64                 `json:"baselineEnd,omitempty"`
	DueDate       *int64                 `json:"dueDate,omitempty"`
	Assignee      string                 `json:"assignee,omitempty"`
	Reporter      string                 `json:"reporter,omitempty"`
	Parent        *string                `json:"parent,omitempty"`
	StoryPoints   *float64               `json:"storyPoints,omitempty"`
	Priority      string                 `json:"priority,omitempty"`
	Team          *string                `json:"team,omitempty"`
	CustomFields  map[string]interface{} `json:"customFields,omitempty"`
}

// JPOPlanListItemScheme represents a plan from GET /rest/jpo/1.0/plans/list.
type JPOPlanListItemScheme struct {
	ID                   int    `json:"id"`
	Title                string `json:"title"`
	ReadOnly             bool   `json:"readOnly,omitempty"`
	ProgramID            *int   `json:"programId,omitempty"`
	PortfolioPlanVersion int    `json:"portfolioPlanVersion,omitempty"`
}

// JPOViewPreferencesResponseScheme is the top-level response from
// GET /rest/jpo/1.0/views/{viewID}/preferences.
type JPOViewPreferencesResponseScheme struct {
	Value *JPOViewPreferencesValueScheme `json:"value"`
}

// JPOViewPreferencesValueScheme holds the versioned preferences blob.
type JPOViewPreferencesValueScheme struct {
	Version     int                        `json:"version"`
	Preferences *JPOViewPreferencesScheme  `json:"preferences"`
}

// JPOViewPreferencesScheme contains the per-user view preference sections.
type JPOViewPreferencesScheme struct {
	Filters       map[string]*JPOViewFilterScheme       `json:"filtersV1,omitempty"`
	FilterOptions *JPOViewFilterOptionsScheme            `json:"filterOptionsV1,omitempty"`
	Sorting       *JPOViewVisualisationsScheme           `json:"visualisationsV2,omitempty"`
	FieldColumns  *JPOViewFieldColumnsScheme             `json:"listFieldColumnsV0,omitempty"`
	ViewMode      *JPOViewModeScheme                     `json:"viewModeV0,omitempty"`
}

// JPOViewFilterScheme is a single named filter entry.
type JPOViewFilterScheme struct {
	ID    string      `json:"id,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// JPOViewHierarchyRangeScheme is the value of HIERARCHY_FILTER_ID or HIERARCHY_RANGE_FILTER_ID.
type JPOViewHierarchyRangeScheme struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// JPOViewFilterOptionsScheme holds filter behavior flags.
type JPOViewFilterOptionsScheme struct {
	IncludeAncestors   bool     `json:"includeAncestors"`
	IncludeDescendants bool     `json:"includeDescendants"`
	VisibleFilterIDs   []string `json:"orderedVisibleFilterIds,omitempty"`
}

// JPOViewVisualisationsScheme holds sorting and grouping settings.
type JPOViewVisualisationsScheme struct {
	Sorting  *JPOViewSortingScheme `json:"sorting,omitempty"`
	Grouping string                `json:"grouping,omitempty"`
}

// JPOViewSortingScheme holds sort field and direction.
type JPOViewSortingScheme struct {
	Field     string `json:"field"`
	Type      string `json:"type,omitempty"`
	Direction string `json:"direction"`
}

// JPOViewFieldColumnsScheme holds visible field columns.
type JPOViewFieldColumnsScheme struct {
	Columns []*JPOViewColumnScheme `json:"columns,omitempty"`
}

// JPOViewColumnScheme is a single visible column.
type JPOViewColumnScheme struct {
	ID        string `json:"id"`
	IsVisible bool   `json:"isVisible"`
}

// JPOViewModeScheme holds the view display mode (list, timeline, etc).
type JPOViewModeScheme struct {
	Mode string `json:"mode"`
}
