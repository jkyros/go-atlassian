package jira

import (
	"context"
	"encoding/json"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// JPOPlanConnector defines the methods available from the JPO (Advanced Roadmaps) Plan API.
//
// These are private/internal JIRA APIs used by Advanced Roadmaps (formerly Portfolio for Jira).
// They are not part of the official Atlassian REST API but are required for programmatic
// access to plan data until the v3 Plans API reaches full parity.
type JPOPlanConnector interface {

	// Get retrieves a JPO plan configuration.
	//
	// GET /rest/jpo/1.0/plans/{planID}/
	Get(ctx context.Context, planID int) (*models.JPOPlanScheme, *models.ResponseScheme, error)

	// GetDetail retrieves the full plan configuration including scenarios.
	//
	// POST /rest/jpo/1.0/plans/detail
	GetDetail(ctx context.Context, planID int) (*models.JPOPlanDetailResponseScheme, json.RawMessage, *models.ResponseScheme, error)

	// GetBacklog retrieves the full backlog for a plan scenario.
	// Issues are returned with numeric JIRA IDs (not human-readable keys) that
	// require resolution via an IDResolver.
	//
	// POST /rest/jpo/1.0/backlog
	GetBacklog(ctx context.Context, planID, scenarioID int, filter *models.JPOBacklogFilterScheme) (*models.JPOBacklogResponseScheme, json.RawMessage, *models.ResponseScheme, error)

	// GetTeams retrieves team data for a plan.
	// Returns raw JSON since the response shape varies by JIRA configuration.
	//
	// GET /rest/jpo/1.0/teams/plan/{planID}
	GetTeams(ctx context.Context, planID int) (json.RawMessage, *models.ResponseScheme, error)

	// List retrieves all available JPO plans.
	//
	// GET /rest/jpo/1.0/plans/list
	List(ctx context.Context) ([]*models.JPOPlanListItemScheme, *models.ResponseScheme, error)

	// GetViewPreferences retrieves the per-user view preferences for a plan view.
	// This is a private/internal API — not part of the official Atlassian REST API.
	// The viewID can be discovered via browser DevTools on the plan page.
	//
	// GET /rest/jpo/1.0/views/{viewID}/preferences
	GetViewPreferences(ctx context.Context, viewID int) (*models.JPOViewPreferencesResponseScheme, *models.ResponseScheme, error)
}
