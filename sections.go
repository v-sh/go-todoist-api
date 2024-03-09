package todoist

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

const SectionsEndpoint = "sections"

type Section struct {
	Id        string `json:"id"`
	ProjectId string `json:"project_id"`
	Order     int    `json:"order"`
	Name      string `json:"name"`
}

// region GetSections

type GetSectionsParams map[string]string

//goland:noinspection GoUnusedExportedFunction
func MakeGetSectionsParams() *GetSectionsParams {
	params := make(GetSectionsParams)
	return &params
}

func (p *GetSectionsParams) WithProjectId(projectId string) *GetSectionsParams {
	if projectId != "" {
		(*p)["project_id"] = projectId
	}

	return p
}

func (t *Todoist) GetSections(ctx context.Context, params *GetSectionsParams) (sections []Section, err error) {
	sections = make([]Section, 0)
	err = t.request(ctx, http.MethodGet, SectionsEndpoint, *params, nil, &sections)

	return
}

// endregion

// region AddSection

type AddSectionParams map[string]interface{}

//goland:noinspection GoUnusedExportedFunction
func MakeAddSectionParams() *AddSectionParams {
	params := make(AddSectionParams)
	return &params
}

func (p *AddSectionParams) WithName(name string) *AddSectionParams {
	if name != "" {
		(*p)["name"] = name
	}

	return p
}

func (p *AddSectionParams) WithProjectId(projectId string) *AddSectionParams {
	if projectId != "" {
		(*p)["project_id"] = projectId
	}

	return p
}

func (p *AddSectionParams) WithOrder(order int) *AddSectionParams {
	if order != 0 {
		(*p)["order"] = order
	}

	return p
}

func (t *Todoist) AddSection(ctx context.Context, params *AddSectionParams) (section *Section, err error) {
	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}

	section = new(Section)
	err = t.request(ctx, http.MethodPost, SectionsEndpoint, nil, bytes.NewBuffer(payload), section)

	return
}

// endregion

// region GetSection

func (t *Todoist) GetSection(ctx context.Context, sectionId string) (section *Section, err error) {
	section = new(Section)
	encodedSectionId := url.PathEscape(sectionId)
	err = t.request(ctx, http.MethodGet, SectionsEndpoint+"/"+encodedSectionId, nil, nil, section)
	return
}

// endregion

// region UpdateSection

type UpdateSectionParams map[string]interface{}

//goland:noinspection GoUnusedExportedFunction
func MakeUpdateSectionParams() *UpdateSectionParams {
	params := make(UpdateSectionParams)
	return &params
}

func (p *UpdateSectionParams) WithName(name string) *UpdateSectionParams {
	if name != "" {
		(*p)["name"] = name
	}

	return p
}

func (t *Todoist) UpdateSection(ctx context.Context, sectionId string, params *UpdateSectionParams) (err error) {
	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}
	encodedSectionId := url.PathEscape(sectionId)
	return t.request(ctx, http.MethodPost, SectionsEndpoint+"/"+encodedSectionId, nil, bytes.NewBuffer(payload), nil)
}

// endregion

// region DeleteSection

func (t *Todoist) DeleteSection(ctx context.Context, sectionId string) (err error) {
	encodedSectionId := url.PathEscape(sectionId)
	return t.request(ctx, http.MethodDelete, SectionsEndpoint+"/"+encodedSectionId, nil, nil, nil)
}

// endregion
