package todoist

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

const LabelsEndpoint = "labels"

type Label struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Color      string `json:"color"`
	Order      int    `json:"order"`
	IsFavorite bool   `json:"is_favorite"`
}

// region GetLabels

func (t *Todoist) GetLabels(ctx context.Context) (labels []Label, err error) {
	labels = make([]Label, 0)
	err = t.request(ctx, http.MethodGet, LabelsEndpoint, nil, nil, &labels)

	return
}

// endregion

// region AddLabel

type AddLabelParams map[string]interface{}

//goland:noinspection GoUnusedExportedFunction
func MakeAddLabelParams() *AddLabelParams {
	params := make(AddLabelParams)
	return &params
}

func (p *AddLabelParams) WithName(name string) *AddLabelParams {
	if name != "" {
		(*p)["name"] = name
	}

	return p
}

func (p *AddLabelParams) WithOrder(order int) *AddLabelParams {
	if order != 0 {
		(*p)["order"] = order
	}

	return p
}

func (p *AddLabelParams) WithColor(color string) *AddLabelParams {
	if color != "" {
		(*p)["color"] = color
	}

	return p
}

func (p *AddLabelParams) WithFavorite(favorite bool) *AddLabelParams {
	(*p)["is_favorite"] = favorite
	return p
}

func (t *Todoist) AddLabel(ctx context.Context, params *AddLabelParams) (label *Label, err error) {
	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}

	label = new(Label)
	err = t.request(ctx, http.MethodPost, LabelsEndpoint, nil, bytes.NewBuffer(payload), label)

	return
}

// endregion

// region GetLabel

func (t *Todoist) GetLabel(ctx context.Context, labelId string) (label *Label, err error) {
	label = new(Label)
	encodedLabelId := url.PathEscape(labelId)
	err = t.request(ctx, http.MethodGet, LabelsEndpoint+"/"+encodedLabelId, nil, nil, label)

	return
}

// endregion

// region UpdateLabel

type UpdateLabelParams map[string]interface{}

//goland:noinspection GoUnusedExportedFunction
func MakeUpdateLabelParams() *UpdateLabelParams {
	params := make(UpdateLabelParams)
	return &params
}

func (p *UpdateLabelParams) WithName(name string) *UpdateLabelParams {
	if name != "" {
		(*p)["name"] = name
	}

	return p
}

func (p *UpdateLabelParams) WithOrder(order int) *UpdateLabelParams {
	if order != 0 {
		(*p)["order"] = order
	}

	return p
}

func (p *UpdateLabelParams) WithColor(color string) *UpdateLabelParams {
	if color != "" {
		(*p)["color"] = color
	}

	return p
}

func (p *UpdateLabelParams) WithFavorite(favorite bool) *UpdateLabelParams {
	(*p)["is_favorite"] = favorite
	return p
}

func (t *Todoist) UpdateLabel(ctx context.Context, labelId string, params *UpdateLabelParams) (err error) {
	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}

	encodedLabelId := url.PathEscape(labelId)
	return t.request(ctx, http.MethodPost, LabelsEndpoint+"/"+encodedLabelId, nil, bytes.NewBuffer(payload), nil)
}

// endregion

// region DeleteLabel

func (t *Todoist) DeleteLabel(ctx context.Context, labelId string) (err error) {
	encodedLabelId := url.PathEscape(labelId)
	return t.request(ctx, http.MethodDelete, LabelsEndpoint+"/"+encodedLabelId, nil, nil, nil)
}

// endregion
