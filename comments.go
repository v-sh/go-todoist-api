package todoist

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

const CommentsEndpoint = "comments"

type Comment struct {
	Id         string                 `json:"id"`
	TaskId     string                 `json:"task_id"`
	ProjectId  string                 `json:"project_id"`
	PostedAt   string                 `json:"posted_at"`
	Content    string                 `json:"content"`
	Attachment map[string]interface{} `json:"attachment"`
}

type Attachment struct {
	ResourceType string `json:"resource_type"`
	FileName     string `json:"file_name"`
	FileSize     int    `json:"file_size"`
	FileType     string `json:"file_type"`
	FileUrl      string `json:"file_url"`
	UploadState  string `json:"upload_state"`
}

type ImageAttachment struct {
	Attachment

	LargeThumbnail  []interface{} `json:"tn_l"`
	MediumThumbnail []interface{} `json:"tn_m"`
	SmallThumbnail  []interface{} `json:"tn_s"`
}

type AudioAttachment struct {
	Attachment

	FileDuration int `json:"file_duration"`
}

// region GetComments

type GetCommentsParams map[string]string

//goland:noinspection GoUnusedExportedFunction
func MakeGetCommentsParams() *GetCommentsParams {
	params := make(GetCommentsParams)
	return &params
}

func (p *GetCommentsParams) WithProjectId(projectId string) *GetCommentsParams {
	if projectId != "" {
		(*p)["project_id"] = projectId
	}

	return p
}

func (p *GetCommentsParams) WithTaskId(taskId string) *GetCommentsParams {
	if taskId != "" {
		(*p)["task_id"] = taskId
	}

	return p
}

func (t *Todoist) GetComments(ctx context.Context, params *GetCommentsParams) (comments []Comment, err error) {
	comments = make([]Comment, 0)
	err = t.request(ctx, http.MethodGet, CommentsEndpoint, *params, nil, &comments)

	return
}

// endregion

// region AddComment

type AddCommentParams map[string]interface{}

//goland:noinspection GoUnusedExportedFunction
func MakeAddCommentParams() *AddCommentParams {
	params := make(AddCommentParams)
	return &params
}

func (p *AddCommentParams) WithTaskId(taskId string) *AddCommentParams {
	if taskId != "" {
		(*p)["task_id"] = taskId
	}

	return p
}

func (p *AddCommentParams) WithProjectId(projectId string) *AddCommentParams {
	if projectId != "" {
		(*p)["project_id"] = projectId
	}

	return p
}

func (p *AddCommentParams) WithContent(content string) *AddCommentParams {
	if content != "" {
		(*p)["content"] = content
	}

	return p
}

func (p *AddCommentParams) WithAttachment(attachment interface{}) *AddCommentParams {
	if attachment != nil {
		(*p)["attachment"] = attachment
	}

	return p
}

func (t *Todoist) AddComment(ctx context.Context, params *AddCommentParams) (comment *Comment, err error) {
	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}

	comment = new(Comment)
	err = t.request(ctx, http.MethodPost, CommentsEndpoint, nil, bytes.NewBuffer(payload), comment)

	return
}

// endregion

// region GetComment

func (t *Todoist) GetComment(ctx context.Context, commentId string) (comment *Comment, err error) {
	comment = new(Comment)
	encodedCommentId := url.PathEscape(commentId)
	err = t.request(ctx, http.MethodGet, CommentsEndpoint+"/"+encodedCommentId, nil, nil, comment)

	return
}

// endregion

// region UpdateComment

type UpdateCommentParams map[string]interface{}

//goland:noinspection GoUnusedExportedFunction
func MakeUpdateCommentParams() *UpdateCommentParams {
	params := make(UpdateCommentParams)
	return &params
}

func (p *UpdateCommentParams) WithContent(content string) *UpdateCommentParams {
	if content != "" {
		(*p)["content"] = content
	}

	return p
}

func (t *Todoist) UpdateComment(ctx context.Context, commentId string, params *UpdateCommentParams) (err error) {
	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}

	encodedCommentId := url.PathEscape(commentId)
	return t.request(ctx, http.MethodPost, CommentsEndpoint+"/"+encodedCommentId, nil, bytes.NewBuffer(payload), nil)
}

// endregion

// region DeleteComment

func (t *Todoist) DeleteComment(ctx context.Context, commentId string) (err error) {
	encodedCommentId := url.PathEscape(commentId)
	return t.request(ctx, http.MethodDelete, CommentsEndpoint+"/"+encodedCommentId, nil, nil, nil)
}

// endregion
