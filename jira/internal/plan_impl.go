package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewPlanService creates a new instance of PlanService.
func NewPlanService(client service.Connector, version string) (*PlanService, error) {

	if version == "" {
		return nil, fmt.Errorf("jira: %w", model.ErrNoVersionProvided)
	}

	return &PlanService{
		internalClient: &internalPlanImpl{c: client, version: version},
	}, nil
}

// PlanService provides methods for interacting with the Jira Plans API (Advanced Roadmaps).
type PlanService struct {
	internalClient jira.PlanConnector
}

// Gets returns a paginated list of plans.
//
// GET /rest/api/{2-3}/plans/plan
func (p *PlanService) Gets(ctx context.Context, opts *model.PlanGetsOptions) (*model.PlanPageScheme, *model.ResponseScheme, error) {
	return p.internalClient.Gets(ctx, opts)
}

// Get returns a plan by its ID.
//
// GET /rest/api/{2-3}/plans/plan/{planId}
func (p *PlanService) Get(ctx context.Context, planID int) (*model.PlanScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, planID)
}

// Create creates a new plan.
//
// POST /rest/api/{2-3}/plans/plan
func (p *PlanService) Create(ctx context.Context, payload *model.PlanCreateScheme) (int64, *model.ResponseScheme, error) {
	return p.internalClient.Create(ctx, payload)
}

// Update updates a plan using JSON Patch operations.
//
// PUT /rest/api/{2-3}/plans/plan/{planId}
func (p *PlanService) Update(ctx context.Context, planID int, payload interface{}) (*model.ResponseScheme, error) {
	return p.internalClient.Update(ctx, planID, payload)
}

// Archive archives a plan.
//
// PUT /rest/api/{2-3}/plans/plan/{planId}/archive
func (p *PlanService) Archive(ctx context.Context, planID int) (*model.ResponseScheme, error) {
	return p.internalClient.Archive(ctx, planID)
}

// Trash moves a plan to the trash.
//
// PUT /rest/api/{2-3}/plans/plan/{planId}/trash
func (p *PlanService) Trash(ctx context.Context, planID int) (*model.ResponseScheme, error) {
	return p.internalClient.Trash(ctx, planID)
}

// Duplicate duplicates a plan.
//
// POST /rest/api/{2-3}/plans/plan/{planId}/duplicate
func (p *PlanService) Duplicate(ctx context.Context, planID int, payload *model.PlanDuplicateScheme) (int64, *model.ResponseScheme, error) {
	return p.internalClient.Duplicate(ctx, planID, payload)
}

// GetTeams returns a paginated list of teams in a plan.
//
// GET /rest/api/{2-3}/plans/plan/{planId}/team
func (p *PlanService) GetTeams(ctx context.Context, planID int, opts *model.PlanTeamGetsOptions) (*model.PlanTeamPageScheme, *model.ResponseScheme, error) {
	return p.internalClient.GetTeams(ctx, planID, opts)
}

// GetAtlassianTeam returns an Atlassian team in a plan.
//
// GET /rest/api/{2-3}/plans/plan/{planId}/team/atlassian/{atlassianTeamId}
func (p *PlanService) GetAtlassianTeam(ctx context.Context, planID int, atlassianTeamID string) (*model.PlanAtlassianTeamScheme, *model.ResponseScheme, error) {
	return p.internalClient.GetAtlassianTeam(ctx, planID, atlassianTeamID)
}

// AddAtlassianTeam adds an Atlassian team to a plan.
//
// POST /rest/api/{2-3}/plans/plan/{planId}/team/atlassian
func (p *PlanService) AddAtlassianTeam(ctx context.Context, planID int, payload *model.PlanAtlassianTeamCreateScheme) (*model.ResponseScheme, error) {
	return p.internalClient.AddAtlassianTeam(ctx, planID, payload)
}

// UpdateAtlassianTeam updates an Atlassian team in a plan using JSON Patch operations.
//
// PUT /rest/api/{2-3}/plans/plan/{planId}/team/atlassian/{atlassianTeamId}
func (p *PlanService) UpdateAtlassianTeam(ctx context.Context, planID int, atlassianTeamID string, payload interface{}) (*model.ResponseScheme, error) {
	return p.internalClient.UpdateAtlassianTeam(ctx, planID, atlassianTeamID, payload)
}

// RemoveAtlassianTeam removes an Atlassian team from a plan.
//
// DELETE /rest/api/{2-3}/plans/plan/{planId}/team/atlassian/{atlassianTeamId}
func (p *PlanService) RemoveAtlassianTeam(ctx context.Context, planID int, atlassianTeamID string) (*model.ResponseScheme, error) {
	return p.internalClient.RemoveAtlassianTeam(ctx, planID, atlassianTeamID)
}

// GetPlanOnlyTeam returns a plan-only team.
//
// GET /rest/api/{2-3}/plans/plan/{planId}/team/planonly/{planOnlyTeamId}
func (p *PlanService) GetPlanOnlyTeam(ctx context.Context, planID int, planOnlyTeamID int) (*model.PlanOnlyTeamScheme, *model.ResponseScheme, error) {
	return p.internalClient.GetPlanOnlyTeam(ctx, planID, planOnlyTeamID)
}

// CreatePlanOnlyTeam creates a plan-only team.
//
// POST /rest/api/{2-3}/plans/plan/{planId}/team/planonly
func (p *PlanService) CreatePlanOnlyTeam(ctx context.Context, planID int, payload *model.PlanOnlyTeamCreateScheme) (*model.ResponseScheme, error) {
	return p.internalClient.CreatePlanOnlyTeam(ctx, planID, payload)
}

// UpdatePlanOnlyTeam updates a plan-only team using JSON Patch operations.
//
// PUT /rest/api/{2-3}/plans/plan/{planId}/team/planonly/{planOnlyTeamId}
func (p *PlanService) UpdatePlanOnlyTeam(ctx context.Context, planID int, planOnlyTeamID int, payload interface{}) (*model.ResponseScheme, error) {
	return p.internalClient.UpdatePlanOnlyTeam(ctx, planID, planOnlyTeamID, payload)
}

// DeletePlanOnlyTeam deletes a plan-only team.
//
// DELETE /rest/api/{2-3}/plans/plan/{planId}/team/planonly/{planOnlyTeamId}
func (p *PlanService) DeletePlanOnlyTeam(ctx context.Context, planID int, planOnlyTeamID int) (*model.ResponseScheme, error) {
	return p.internalClient.DeletePlanOnlyTeam(ctx, planID, planOnlyTeamID)
}

type internalPlanImpl struct {
	c       service.Connector
	version string
}

func (i *internalPlanImpl) Gets(ctx context.Context, opts *model.PlanGetsOptions) (*model.PlanPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}

	if opts != nil {

		if opts.IncludeTrashed {
			params.Add("includeTrashed", "true")
		}

		if opts.IncludeArchived {
			params.Add("includeArchived", "true")
		}

		if opts.Cursor != "" {
			params.Add("cursor", opts.Cursor)
		}

		if opts.MaxResults > 0 {
			params.Add("maxResults", strconv.Itoa(opts.MaxResults))
		}
	}

	endpoint := fmt.Sprintf("rest/api/%v/plans/plan", i.version)
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.PlanPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalPlanImpl) Get(ctx context.Context, planID int) (*model.PlanScheme, *model.ResponseScheme, error) {

	if planID == 0 {
		return nil, nil, fmt.Errorf("jira: %w", model.ErrNoPlanID)
	}

	endpoint := fmt.Sprintf("rest/api/%v/plans/plan/%v", i.version, planID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	plan := new(model.PlanScheme)
	response, err := i.c.Call(request, plan)
	if err != nil {
		return nil, response, err
	}

	return plan, response, nil
}

func (i *internalPlanImpl) Create(ctx context.Context, payload *model.PlanCreateScheme) (int64, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/plans/plan", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return 0, nil, err
	}

	var planID int64
	response, err := i.c.Call(request, &planID)
	if err != nil {
		return 0, response, err
	}

	return planID, response, nil
}

func (i *internalPlanImpl) Update(ctx context.Context, planID int, payload interface{}) (*model.ResponseScheme, error) {

	if planID == 0 {
		return nil, fmt.Errorf("jira: %w", model.ErrNoPlanID)
	}

	endpoint := fmt.Sprintf("rest/api/%v/plans/plan/%v", i.version, planID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "application/json-patch+json", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalPlanImpl) Archive(ctx context.Context, planID int) (*model.ResponseScheme, error) {

	if planID == 0 {
		return nil, fmt.Errorf("jira: %w", model.ErrNoPlanID)
	}

	endpoint := fmt.Sprintf("rest/api/%v/plans/plan/%v/archive", i.version, planID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalPlanImpl) Trash(ctx context.Context, planID int) (*model.ResponseScheme, error) {

	if planID == 0 {
		return nil, fmt.Errorf("jira: %w", model.ErrNoPlanID)
	}

	endpoint := fmt.Sprintf("rest/api/%v/plans/plan/%v/trash", i.version, planID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalPlanImpl) Duplicate(ctx context.Context, planID int, payload *model.PlanDuplicateScheme) (int64, *model.ResponseScheme, error) {

	if planID == 0 {
		return 0, nil, fmt.Errorf("jira: %w", model.ErrNoPlanID)
	}

	endpoint := fmt.Sprintf("rest/api/%v/plans/plan/%v/duplicate", i.version, planID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return 0, nil, err
	}

	var newPlanID int64
	response, err := i.c.Call(request, &newPlanID)
	if err != nil {
		return 0, response, err
	}

	return newPlanID, response, nil
}

func (i *internalPlanImpl) GetTeams(ctx context.Context, planID int, opts *model.PlanTeamGetsOptions) (*model.PlanTeamPageScheme, *model.ResponseScheme, error) {

	if planID == 0 {
		return nil, nil, fmt.Errorf("jira: %w", model.ErrNoPlanID)
	}

	params := url.Values{}

	if opts != nil {

		if opts.Cursor != "" {
			params.Add("cursor", opts.Cursor)
		}

		if opts.MaxResults > 0 {
			params.Add("maxResults", strconv.Itoa(opts.MaxResults))
		}
	}

	endpoint := fmt.Sprintf("rest/api/%v/plans/plan/%v/team", i.version, planID)
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.PlanTeamPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalPlanImpl) GetAtlassianTeam(ctx context.Context, planID int, atlassianTeamID string) (*model.PlanAtlassianTeamScheme, *model.ResponseScheme, error) {

	if planID == 0 {
		return nil, nil, fmt.Errorf("jira: %w", model.ErrNoPlanID)
	}

	if atlassianTeamID == "" {
		return nil, nil, fmt.Errorf("jira: %w", model.ErrNoAtlassianTeamID)
	}

	endpoint := fmt.Sprintf("rest/api/%v/plans/plan/%v/team/atlassian/%v", i.version, planID, atlassianTeamID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	team := new(model.PlanAtlassianTeamScheme)
	response, err := i.c.Call(request, team)
	if err != nil {
		return nil, response, err
	}

	return team, response, nil
}

func (i *internalPlanImpl) AddAtlassianTeam(ctx context.Context, planID int, payload *model.PlanAtlassianTeamCreateScheme) (*model.ResponseScheme, error) {

	if planID == 0 {
		return nil, fmt.Errorf("jira: %w", model.ErrNoPlanID)
	}

	endpoint := fmt.Sprintf("rest/api/%v/plans/plan/%v/team/atlassian", i.version, planID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalPlanImpl) UpdateAtlassianTeam(ctx context.Context, planID int, atlassianTeamID string, payload interface{}) (*model.ResponseScheme, error) {

	if planID == 0 {
		return nil, fmt.Errorf("jira: %w", model.ErrNoPlanID)
	}

	if atlassianTeamID == "" {
		return nil, fmt.Errorf("jira: %w", model.ErrNoAtlassianTeamID)
	}

	endpoint := fmt.Sprintf("rest/api/%v/plans/plan/%v/team/atlassian/%v", i.version, planID, atlassianTeamID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "application/json-patch+json", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalPlanImpl) RemoveAtlassianTeam(ctx context.Context, planID int, atlassianTeamID string) (*model.ResponseScheme, error) {

	if planID == 0 {
		return nil, fmt.Errorf("jira: %w", model.ErrNoPlanID)
	}

	if atlassianTeamID == "" {
		return nil, fmt.Errorf("jira: %w", model.ErrNoAtlassianTeamID)
	}

	endpoint := fmt.Sprintf("rest/api/%v/plans/plan/%v/team/atlassian/%v", i.version, planID, atlassianTeamID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalPlanImpl) GetPlanOnlyTeam(ctx context.Context, planID int, planOnlyTeamID int) (*model.PlanOnlyTeamScheme, *model.ResponseScheme, error) {

	if planID == 0 {
		return nil, nil, fmt.Errorf("jira: %w", model.ErrNoPlanID)
	}

	if planOnlyTeamID == 0 {
		return nil, nil, fmt.Errorf("jira: %w", model.ErrNoPlanOnlyTeamID)
	}

	endpoint := fmt.Sprintf("rest/api/%v/plans/plan/%v/team/planonly/%v", i.version, planID, planOnlyTeamID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	team := new(model.PlanOnlyTeamScheme)
	response, err := i.c.Call(request, team)
	if err != nil {
		return nil, response, err
	}

	return team, response, nil
}

func (i *internalPlanImpl) CreatePlanOnlyTeam(ctx context.Context, planID int, payload *model.PlanOnlyTeamCreateScheme) (*model.ResponseScheme, error) {

	if planID == 0 {
		return nil, fmt.Errorf("jira: %w", model.ErrNoPlanID)
	}

	endpoint := fmt.Sprintf("rest/api/%v/plans/plan/%v/team/planonly", i.version, planID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalPlanImpl) UpdatePlanOnlyTeam(ctx context.Context, planID int, planOnlyTeamID int, payload interface{}) (*model.ResponseScheme, error) {

	if planID == 0 {
		return nil, fmt.Errorf("jira: %w", model.ErrNoPlanID)
	}

	if planOnlyTeamID == 0 {
		return nil, fmt.Errorf("jira: %w", model.ErrNoPlanOnlyTeamID)
	}

	endpoint := fmt.Sprintf("rest/api/%v/plans/plan/%v/team/planonly/%v", i.version, planID, planOnlyTeamID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "application/json-patch+json", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalPlanImpl) DeletePlanOnlyTeam(ctx context.Context, planID int, planOnlyTeamID int) (*model.ResponseScheme, error) {

	if planID == 0 {
		return nil, fmt.Errorf("jira: %w", model.ErrNoPlanID)
	}

	if planOnlyTeamID == 0 {
		return nil, fmt.Errorf("jira: %w", model.ErrNoPlanOnlyTeamID)
	}

	endpoint := fmt.Sprintf("rest/api/%v/plans/plan/%v/team/planonly/%v", i.version, planID, planOnlyTeamID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
