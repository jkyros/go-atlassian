package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewJPOPlanService creates a new instance of JPOPlanService.
func NewJPOPlanService(client service.Connector) *JPOPlanService {
	return &JPOPlanService{
		internalClient: &internalJPOPlanServiceImpl{c: client},
	}
}

// JPOPlanService provides methods for interacting with the JPO (Advanced Roadmaps) Plan API.
type JPOPlanService struct {
	internalClient jira.JPOPlanConnector
}

// Get retrieves a JPO plan configuration.
//
// GET /rest/jpo/1.0/plans/{planID}/
func (j *JPOPlanService) Get(ctx context.Context, planID int) (*model.JPOPlanScheme, *model.ResponseScheme, error) {
	return j.internalClient.Get(ctx, planID)
}

// GetDetail retrieves the full plan configuration including scenarios.
//
// POST /rest/jpo/1.0/plans/detail
func (j *JPOPlanService) GetDetail(ctx context.Context, planID int) (*model.JPOPlanDetailResponseScheme, json.RawMessage, *model.ResponseScheme, error) {
	return j.internalClient.GetDetail(ctx, planID)
}

// GetBacklog retrieves the full backlog for a plan scenario.
//
// POST /rest/jpo/1.0/backlog
func (j *JPOPlanService) GetBacklog(ctx context.Context, planID, scenarioID int, filter *model.JPOBacklogFilterScheme) (*model.JPOBacklogResponseScheme, json.RawMessage, *model.ResponseScheme, error) {
	return j.internalClient.GetBacklog(ctx, planID, scenarioID, filter)
}

// GetTeams retrieves team data for a plan.
//
// GET /rest/jpo/1.0/teams/plan/{planID}
func (j *JPOPlanService) GetTeams(ctx context.Context, planID int) (json.RawMessage, *model.ResponseScheme, error) {
	return j.internalClient.GetTeams(ctx, planID)
}

// List retrieves all available JPO plans.
//
// GET /rest/jpo/1.0/plans/list
func (j *JPOPlanService) List(ctx context.Context) ([]*model.JPOPlanListItemScheme, *model.ResponseScheme, error) {
	return j.internalClient.List(ctx)
}

// GetViewPreferences retrieves the per-user view preferences for a plan view.
//
// GET /rest/jpo/1.0/views/{viewID}/preferences
func (j *JPOPlanService) GetViewPreferences(ctx context.Context, viewID int) (*model.JPOViewPreferencesResponseScheme, *model.ResponseScheme, error) {
	return j.internalClient.GetViewPreferences(ctx, viewID)
}

type internalJPOPlanServiceImpl struct {
	c service.Connector
}

func (i *internalJPOPlanServiceImpl) Get(ctx context.Context, planID int) (*model.JPOPlanScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/jpo/1.0/plans/%d/", planID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	plan := new(model.JPOPlanScheme)
	response, err := i.c.Call(request, plan)
	if err != nil {
		return nil, response, err
	}

	return plan, response, nil
}

func (i *internalJPOPlanServiceImpl) GetDetail(ctx context.Context, planID int) (*model.JPOPlanDetailResponseScheme, json.RawMessage, *model.ResponseScheme, error) {

	endpoint := "rest/jpo/1.0/plans/detail"

	payload := &model.JPOPlanDetailRequestScheme{PlanID: planID}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, nil, err
	}

	// Use Do instead of Call so we can capture the raw JSON alongside the typed response.
	httpResp, err := i.c.Do(request)
	if err != nil {
		return nil, nil, nil, err
	}
	defer httpResp.Body.Close()

	rawBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, nil, nil, err
	}

	response := &model.ResponseScheme{
		Response: httpResp,
		Code:     httpResp.StatusCode,
		Endpoint: httpResp.Request.URL.String(),
		Method:   httpResp.Request.Method,
	}
	response.Bytes.Write(rawBody)

	if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		return nil, json.RawMessage(rawBody), response, fmt.Errorf("jpo plans/detail returned %d", httpResp.StatusCode)
	}

	detail := new(model.JPOPlanDetailResponseScheme)
	if err := json.Unmarshal(rawBody, detail); err != nil {
		return nil, json.RawMessage(rawBody), response, err
	}

	return detail, json.RawMessage(rawBody), response, nil
}

func (i *internalJPOPlanServiceImpl) GetBacklog(ctx context.Context, planID, scenarioID int, filter *model.JPOBacklogFilterScheme) (*model.JPOBacklogResponseScheme, json.RawMessage, *model.ResponseScheme, error) {

	endpoint := "rest/jpo/1.0/backlog"

	payload := &model.JPOBacklogRequestScheme{
		PlanID:     planID,
		ScenarioID: scenarioID,
		Filter:     filter,
	}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, nil, err
	}

	httpResp, err := i.c.Do(request)
	if err != nil {
		return nil, nil, nil, err
	}
	defer httpResp.Body.Close()

	rawBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, nil, nil, err
	}

	response := &model.ResponseScheme{
		Response: httpResp,
		Code:     httpResp.StatusCode,
		Endpoint: httpResp.Request.URL.String(),
		Method:   httpResp.Request.Method,
	}
	response.Bytes.Write(rawBody)

	if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		return nil, json.RawMessage(rawBody), response, fmt.Errorf("jpo backlog returned %d", httpResp.StatusCode)
	}

	backlog := new(model.JPOBacklogResponseScheme)
	if err := json.Unmarshal(rawBody, backlog); err != nil {
		return nil, json.RawMessage(rawBody), response, err
	}

	return backlog, json.RawMessage(rawBody), response, nil
}

func (i *internalJPOPlanServiceImpl) GetTeams(ctx context.Context, planID int) (json.RawMessage, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/jpo/1.0/teams/plan/%d", planID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	httpResp, err := i.c.Do(request)
	if err != nil {
		return nil, nil, err
	}
	defer httpResp.Body.Close()

	rawBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, nil, err
	}

	response := &model.ResponseScheme{
		Response: httpResp,
		Code:     httpResp.StatusCode,
		Endpoint: httpResp.Request.URL.String(),
		Method:   httpResp.Request.Method,
	}
	response.Bytes.Write(rawBody)

	if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		return nil, response, fmt.Errorf("jpo teams returned %d", httpResp.StatusCode)
	}

	return json.RawMessage(rawBody), response, nil
}

func (i *internalJPOPlanServiceImpl) List(ctx context.Context) ([]*model.JPOPlanListItemScheme, *model.ResponseScheme, error) {

	endpoint := "rest/jpo/1.0/plans/list"

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var plans []*model.JPOPlanListItemScheme
	response, err := i.c.Call(request, &plans)
	if err != nil {
		return nil, response, err
	}

	return plans, response, nil
}

func (i *internalJPOPlanServiceImpl) GetViewPreferences(ctx context.Context, viewID int) (*model.JPOViewPreferencesResponseScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/jpo/1.0/views/%d/preferences", viewID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	prefs := new(model.JPOViewPreferencesResponseScheme)
	response, err := i.c.Call(request, prefs)
	if err != nil {
		return nil, response, err
	}

	return prefs, response, nil
}
