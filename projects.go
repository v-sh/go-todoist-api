package todoist

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

const ProjectsEndpoint = "projects"

type Project struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Color          string `json:"color"`
	ParentId       string `json:"parent_id"`
	Order          int    `json:"order"`
	CommentCount   int    `json:"comment_count"`
	IsShared       bool   `json:"is_shared"`
	IsFavorite     bool   `json:"is_favorite"`
	IsInboxProject bool   `json:"is_inbox_project"`
	IsTeamInbox    bool   `json:"is_team_inbox"`
	Url            string `json:"url"`
}

type Collaborator struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// region GetProjects

func (t *Todoist) GetProjects(ctx context.Context) (projects []Project, err error) {
	projects = make([]Project, 0)
	err = t.request(ctx, http.MethodGet, ProjectsEndpoint, nil, nil, &projects)

	return
}

// endregion

// region AddProject

type AddProjectParams map[string]interface{}

//goland:noinspection GoUnusedExportedFunction
func MakeAddProjectParams() *AddProjectParams {
	params := make(AddProjectParams)
	return &params
}

func (p *AddProjectParams) WithName(name string) *AddProjectParams {
	if name != "" {
		(*p)["name"] = name
	}

	return p
}

func (p *AddProjectParams) WithParentId(parentId string) *AddProjectParams {
	if parentId != "" {
		(*p)["parent_id"] = parentId
	}

	return p
}

func (p *AddProjectParams) WithColor(color string) *AddProjectParams {
	if color != "" {
		(*p)["color"] = color
	}

	return p
}

func (p *AddProjectParams) WithFavorite(favorite bool) *AddProjectParams {
	(*p)["is_favorite"] = favorite
	return p
}

func (t *Todoist) AddProject(ctx context.Context, params *AddProjectParams) (project *Project, err error) {
	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}

	project = new(Project)
	err = t.request(ctx, http.MethodPost, ProjectsEndpoint, nil, bytes.NewBuffer(payload), project)

	return
}

// endregion

// region GetProject

func (t *Todoist) GetProject(ctx context.Context, projectId string) (project *Project, err error) {
	project = new(Project)
	encodedProjectId := url.PathEscape(projectId)
	err = t.request(ctx, http.MethodGet, ProjectsEndpoint+"/"+encodedProjectId, nil, nil, project)
	return
}

// endregion

// region UpdateProject

type UpdateProjectParams map[string]interface{}

//goland:noinspection GoUnusedExportedFunction
func MakeUpdateProjectParams() *UpdateProjectParams {
	params := make(UpdateProjectParams)
	return &params
}

func (p *UpdateProjectParams) WithName(name string) *UpdateProjectParams {
	if name != "" {
		(*p)["name"] = name
	}

	return p
}

func (p *UpdateProjectParams) WithColor(color string) *UpdateProjectParams {
	if color != "" {
		(*p)["color"] = color
	}

	return p
}

func (p *UpdateProjectParams) WithFavorite(favorite bool) *UpdateProjectParams {
	(*p)["is_favorite"] = favorite
	return p
}

func (t *Todoist) UpdateProject(ctx context.Context, projectId string, params *UpdateProjectParams) (err error) {
	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}
	encodedProjectId := url.PathEscape(projectId)
	return t.request(ctx, http.MethodPost, ProjectsEndpoint+"/"+encodedProjectId, nil, bytes.NewBuffer(payload), nil)
}

// endregion

// region DeleteProject

func (t *Todoist) DeleteProject(ctx context.Context, projectId string) (err error) {
	encodedProjectId := url.PathEscape(projectId)
	return t.request(ctx, http.MethodDelete, ProjectsEndpoint+"/"+encodedProjectId, nil, nil, nil)
}

// endregion

// region GetCollaborators

func (t *Todoist) GetCollaborators(ctx context.Context, projectId string) (collaborators []Collaborator, err error) {
	collaborators = make([]Collaborator, 0)
	encodedProjectId := url.PathEscape(projectId)
	err = t.request(ctx, http.MethodGet, ProjectsEndpoint+"/"+encodedProjectId+"/collaborators", nil, nil, &collaborators)

	return
}

// endregion
