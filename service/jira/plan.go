package jira

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type PlanConnector interface {

	// Gets returns a paginated list of plans.
	//
	// GET /rest/api/{2-3}/plans/plan
	Gets(ctx context.Context, opts *model.PlanGetsOptions) (*model.PlanPageScheme, *model.ResponseScheme, error)

	// Get returns a plan by its ID.
	//
	// GET /rest/api/{2-3}/plans/plan/{planId}
	Get(ctx context.Context, planID int) (*model.PlanScheme, *model.ResponseScheme, error)

	// Create creates a new plan.
	//
	// POST /rest/api/{2-3}/plans/plan
	Create(ctx context.Context, payload *model.PlanCreateScheme) (int64, *model.ResponseScheme, error)

	// Update updates a plan using JSON Patch operations.
	//
	// PUT /rest/api/{2-3}/plans/plan/{planId}
	Update(ctx context.Context, planID int, payload interface{}) (*model.ResponseScheme, error)

	// Archive archives a plan.
	//
	// PUT /rest/api/{2-3}/plans/plan/{planId}/archive
	Archive(ctx context.Context, planID int) (*model.ResponseScheme, error)

	// Trash moves a plan to the trash.
	//
	// PUT /rest/api/{2-3}/plans/plan/{planId}/trash
	Trash(ctx context.Context, planID int) (*model.ResponseScheme, error)

	// Duplicate duplicates a plan.
	//
	// POST /rest/api/{2-3}/plans/plan/{planId}/duplicate
	Duplicate(ctx context.Context, planID int, payload *model.PlanDuplicateScheme) (int64, *model.ResponseScheme, error)

	// GetTeams returns a paginated list of teams in a plan.
	//
	// GET /rest/api/{2-3}/plans/plan/{planId}/team
	GetTeams(ctx context.Context, planID int, opts *model.PlanTeamGetsOptions) (*model.PlanTeamPageScheme, *model.ResponseScheme, error)

	// GetAtlassianTeam returns an Atlassian team in a plan.
	//
	// GET /rest/api/{2-3}/plans/plan/{planId}/team/atlassian/{atlassianTeamId}
	GetAtlassianTeam(ctx context.Context, planID int, atlassianTeamID string) (*model.PlanAtlassianTeamScheme, *model.ResponseScheme, error)

	// AddAtlassianTeam adds an Atlassian team to a plan.
	//
	// POST /rest/api/{2-3}/plans/plan/{planId}/team/atlassian
	AddAtlassianTeam(ctx context.Context, planID int, payload *model.PlanAtlassianTeamCreateScheme) (*model.ResponseScheme, error)

	// UpdateAtlassianTeam updates an Atlassian team in a plan using JSON Patch operations.
	//
	// PUT /rest/api/{2-3}/plans/plan/{planId}/team/atlassian/{atlassianTeamId}
	UpdateAtlassianTeam(ctx context.Context, planID int, atlassianTeamID string, payload interface{}) (*model.ResponseScheme, error)

	// RemoveAtlassianTeam removes an Atlassian team from a plan.
	//
	// DELETE /rest/api/{2-3}/plans/plan/{planId}/team/atlassian/{atlassianTeamId}
	RemoveAtlassianTeam(ctx context.Context, planID int, atlassianTeamID string) (*model.ResponseScheme, error)

	// GetPlanOnlyTeam returns a plan-only team.
	//
	// GET /rest/api/{2-3}/plans/plan/{planId}/team/planonly/{planOnlyTeamId}
	GetPlanOnlyTeam(ctx context.Context, planID int, planOnlyTeamID int) (*model.PlanOnlyTeamScheme, *model.ResponseScheme, error)

	// CreatePlanOnlyTeam creates a plan-only team.
	//
	// POST /rest/api/{2-3}/plans/plan/{planId}/team/planonly
	CreatePlanOnlyTeam(ctx context.Context, planID int, payload *model.PlanOnlyTeamCreateScheme) (*model.ResponseScheme, error)

	// UpdatePlanOnlyTeam updates a plan-only team using JSON Patch operations.
	//
	// PUT /rest/api/{2-3}/plans/plan/{planId}/team/planonly/{planOnlyTeamId}
	UpdatePlanOnlyTeam(ctx context.Context, planID int, planOnlyTeamID int, payload interface{}) (*model.ResponseScheme, error)

	// DeletePlanOnlyTeam deletes a plan-only team.
	//
	// DELETE /rest/api/{2-3}/plans/plan/{planId}/team/planonly/{planOnlyTeamId}
	DeletePlanOnlyTeam(ctx context.Context, planID int, planOnlyTeamID int) (*model.ResponseScheme, error)
}
