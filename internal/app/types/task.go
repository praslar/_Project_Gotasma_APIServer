package types

import "time"

type (
	Task struct {
		Label            string    `json:"label,omitempty" bson:"label" validate:"required"`
		TaskID           string    `json:"task_id,omitempty" bson:"task_id" validate:"required"`
		ProjectID        string    `json:"project_id,omitempty" bson:"project_id"`
		Type             string    `json:"type,omitempty" bson:"type" validate:"required"`
		Parent           string    `json:"parent,omitempty" bson:"parent"`
		CreaterID        string    `json:"creater_id,omitempty" bson:"creater_id"`
		Effort           int       `json:"effort,omitempty" bson:"effort" validate:"required"`
		Start            int       `json:"start,omitempty" bson:"start" validate:"gt=1262304000000" validate:"required"`
		Duration         int       `json:"duration,omitempty" bson:"duration" validate:"required"`
		EstimateDuration int       `json:"estimate_duration,omitempty" bson:"estimate_duration" validate:"required"`
		End              int       `json:"end,omitempty" bson:"end"  validate:"gt=1262304000000" validate:"required"`
		Parents          []string  `json:"parents,omitempty" bson:"parents" `
		Children         []string  `json:"children,omitempty" bson:"children" `
		AllChildren      []string  `json:"all_children,omitempty" bson:"all_children"`
		UpdateAt         time.Time `json:"updated_at,omitempty" bson:"updated_at" validate:"required"`
		CreatedAt        time.Time `json:"created_at,omitempty" bson:"created_at"`
	}

	TaskDetailInfo struct {
		Label     string    `json:"label,omitempty" bson:"label" validate:"required"`
		TaskID    string    `json:"task_id,omitempty" bson:"task_id" validate:"required"`
		ProjectID string    `json:"project_id,omitempty" bson:"project_id"`
		Type      string    `json:"type,omitempty" bson:"type" validate:"required"`
		Effort    int       `json:"effort,omitempty" bson:"effort" validate:"required"`
		Start     time.Time `json:"start,omitempty" bson:"start" validate:"gt=1262304000000" validate:"required"`
		Duration  int       `json:"duration,omitempty" bson:"duration" validate:"required"`
		End       time.Time `json:"end,omitempty" bson:"end"  validate:"gt=1262304000000" validate:"required"`
		UpdateAt  time.Time `json:"updated_at,omitempty" bson:"updated_at" validate:"required"`
		CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`
	}

	TaskInfo struct {
		TaskID string `json:"task_id,omitempty" bson:"task_id" validate:"required"`
		Label  string `json:"label,omitempty" bson:"label" validate:"required"`
		Start  int    `json:"start,omitempty" bson:"start" validate:"gt=1262304000000" validate:"required"`
		End    int    `json:"end,omitempty" bson:"end"  validate:"gt=1262304000000" validate:"required"`
	}
)
