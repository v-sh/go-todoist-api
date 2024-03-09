package todoist

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"net/url"
)

const TasksEndpoint = "tasks"

type Task struct {
	Id           string   `json:"id"`
	ProjectId    string   `json:"project_id"`
	SectionId    string   `json:"section_id"`
	Content      string   `json:"content"`
	Description  string   `json:"description"`
	IsCompleted  bool     `json:"is_completed"`
	Labels       []string `json:"labels"`
	ParentId     string   `json:"parent_id"`
	Order        int      `json:"order"`
	Priority     int      `json:"priority"`
	Due          Due      `json:"due"`
	Url          string   `json:"url"`
	CommentCount int      `json:"comment_count"`
	AssigneeId   string   `json:"assignee_id"`
	AssignerId   string   `json:"assigner_id"`
}

type Due struct {
	String      string `json:"string"`
	Date        string `json:"date"`
	IsRecurring bool   `json:"is_recurring"`
	Datetime    string `json:"datetime"`
	Timezone    string `json:"timezone"`
}

// region GetTasks

type GetTasksParams map[string]string

//goland:noinspection GoUnusedExportedFunction
func MakeGetTasksParams() *GetTasksParams {
	params := make(GetTasksParams)
	return &params
}

func (p *GetTasksParams) WithProjectId(projectId string) *GetTasksParams {
	if projectId != "" {
		(*p)["project_id"] = projectId
	}

	return p
}

func (p *GetTasksParams) WithSectionId(sectionId string) *GetTasksParams {
	if sectionId != "" {
		(*p)["section_id"] = sectionId
	}

	return p
}

func (p *GetTasksParams) WithLabel(label string) *GetTasksParams {
	if label != "" {
		(*p)["label"] = label
	}

	return p
}

func (p *GetTasksParams) WithFilter(filter string) *GetTasksParams {
	if filter != "" {
		(*p)["filter"] = filter
	}

	return p
}

func (p *GetTasksParams) WithLang(lang string) *GetTasksParams {
	if lang != "" {
		(*p)["lang"] = lang
	}

	return p
}

func (p *GetTasksParams) WithIds(ids []string) *GetTasksParams {
	if ids != nil && len(ids) != 0 {
		(*p)["ids"] = strings.Join(ids, ",")
	}

	return p
}

func (t *Todoist) GetTasks(ctx context.Context, params *GetTasksParams) (tasks []Task, err error) {
	tasks = make([]Task, 0)
	err = t.request(ctx, http.MethodGet, TasksEndpoint, *params, nil, &tasks)

	return
}

// endregion

// region AddTask

type AddTaskParams map[string]interface{}

//goland:noinspection GoUnusedExportedFunction
func MakeAddTaskParams() *AddTaskParams {
	params := make(AddTaskParams)
	return &params
}

func (p *AddTaskParams) WithContent(content string) *AddTaskParams {
	if content != "" {
		(*p)["content"] = content
	}

	return p
}

func (p *AddTaskParams) WithDescription(description string) *AddTaskParams {
	if description != "" {
		(*p)["description"] = description
	}

	return p
}

func (p *AddTaskParams) WithProjectId(projectId string) *AddTaskParams {
	if projectId != "" {
		(*p)["project_id"] = projectId
	}

	return p
}

func (p *AddTaskParams) WithSectionId(sectionId string) *AddTaskParams {
	if sectionId != "" {
		(*p)["section_id"] = sectionId
	}

	return p
}

func (p *AddTaskParams) WithParentId(parentId string) *AddTaskParams {
	if parentId != "" {
		(*p)["parent_id"] = parentId
	}

	return p
}

func (p *AddTaskParams) WithOrder(order int) *AddTaskParams {
	if order != 0 {
		(*p)["order"] = order
	}

	return p
}

func (p *AddTaskParams) WithLabels(labels []string) *AddTaskParams {
	if labels != nil && len(labels) != 0 {
		(*p)["labels"] = labels
	}

	return p
}

func (p *AddTaskParams) WithPriority(priority int) *AddTaskParams {
	if priority != 0 {
		(*p)["priority"] = priority
	}

	return p
}

func (p *AddTaskParams) WithDueString(dueString string) *AddTaskParams {
	if dueString != "" {
		(*p)["due_string"] = dueString
	}

	return p
}

func (p *AddTaskParams) WithDueDate(dueDate string) *AddTaskParams {
	if dueDate != "" {
		(*p)["due_date"] = dueDate
	}

	return p
}

func (p *AddTaskParams) WithDueDatetime(dueDatetime string) *AddTaskParams {
	if dueDatetime != "" {
		(*p)["due_datetime"] = dueDatetime
	}

	return p
}

func (p *AddTaskParams) WithDueLang(dueLang string) *AddTaskParams {
	if dueLang != "" {
		(*p)["due_lang"] = dueLang
	}

	return p
}

func (p *AddTaskParams) WithAssigneeId(assigneeId string) *AddTaskParams {
	if assigneeId != "" {
		(*p)["assignee_id"] = assigneeId
	}

	return p
}

func (t *Todoist) AddTask(ctx context.Context, params *AddTaskParams) (task *Task, err error) {
	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}

	task = new(Task)
	err = t.request(ctx, http.MethodPost, TasksEndpoint, nil, bytes.NewBuffer(payload), task)

	return
}

// endregion

// region GetTask

func (t *Todoist) GetTask(ctx context.Context, taskId string) (task *Task, err error) {
	task = new(Task)
	encodedTaskId := url.PathEscape(taskId)
	err = t.request(ctx, http.MethodGet, TasksEndpoint+"/"+encodedTaskId, nil, nil, task)

	return
}

// endregion

// region UpdateTask

type UpdateTaskParams map[string]interface{}

//goland:noinspection GoUnusedExportedFunction
func MakeUpdateTaskParams() *UpdateTaskParams {
	params := make(UpdateTaskParams)
	return &params
}

func (p *UpdateTaskParams) WithContent(content string) *UpdateTaskParams {
	if content != "" {
		(*p)["content"] = content
	}

	return p
}

func (p *UpdateTaskParams) WithDescription(description string) *UpdateTaskParams {
	if description != "" {
		(*p)["description"] = description
	}

	return p
}

func (p *UpdateTaskParams) WithLabelIds(labels []string) *UpdateTaskParams {
	if labels != nil && len(labels) != 0 {
		(*p)["label_ids"] = labels
	}

	return p
}

func (p *UpdateTaskParams) WithPriority(priority int) *UpdateTaskParams {
	if priority != 0 {
		(*p)["priority"] = priority
	}

	return p
}

func (p *UpdateTaskParams) WithDueString(dueString string) *UpdateTaskParams {
	if dueString != "" {
		(*p)["due_string"] = dueString
	}

	return p
}

func (p *UpdateTaskParams) WithDueDate(dueDate string) *UpdateTaskParams {
	if dueDate != "" {
		(*p)["due_date"] = dueDate
	}

	return p
}

func (p *UpdateTaskParams) WithDueDatetime(dueDatetime string) *UpdateTaskParams {
	if dueDatetime != "" {
		(*p)["due_datetime"] = dueDatetime
	}

	return p
}

func (p *UpdateTaskParams) WithDueLang(dueLang string) *UpdateTaskParams {
	if dueLang != "" {
		(*p)["due_lang"] = dueLang
	}

	return p
}

func (p *UpdateTaskParams) WithAssigneeId(assigneeId string) *UpdateTaskParams {
	if assigneeId != "" {
		(*p)["assignee_id"] = assigneeId
	}

	return p
}

func (t *Todoist) UpdateTask(ctx context.Context, taskId string, params *UpdateTaskParams) (err error) {
	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}
	encodedTaskId := url.PathEscape(taskId)
	return t.request(ctx, http.MethodPost, TasksEndpoint+"/"+encodedTaskId, nil, bytes.NewBuffer(payload), nil)
}

// endregion

// region CloseTask

func (t *Todoist) CloseTask(ctx context.Context, taskId string) (err error) {
	encodedTaskId := url.PathEscape(taskId)
	return t.request(ctx, http.MethodPost, TasksEndpoint+"/"+encodedTaskId+"/close", nil, nil, nil)
}

// endregion

// region ReopenTask

func (t *Todoist) ReopenTask(ctx context.Context, taskId string) (err error) {
	encodedTaskId := url.PathEscape(taskId)
	return t.request(ctx, http.MethodPost, TasksEndpoint+"/"+encodedTaskId+"/reopen", nil, nil, nil)
}

// endregion

// region DeleteTask

func (t *Todoist) DeleteTask(ctx context.Context, taskId string) (err error) {
	encodedTaskId := url.PathEscape(taskId)
	return t.request(ctx, http.MethodDelete, TasksEndpoint+"/"+encodedTaskId, nil, nil, nil)
}

// endregion
