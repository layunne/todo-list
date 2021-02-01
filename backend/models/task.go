package models

import (
	"encoding/json"
)

type Task struct {
	Id          string `bson:"_id" json:"id"`
	ProjectId   string `bson:"projectId" json:"projectId"`
	Description string `bson:"description" json:"description"`
	Status      bool   `bson:"status" json:"status"`
	CreatedAt   int64  `bson:"createdAt" json:"createdAt"`
	FinishedAt  int64  `bson:"finishedAt" json:"finishedAt"`
	ToFinishAt  int64  `bson:"toFinishAt" json:"toFinishAt"`
}

func (t *Task) Bytes() []byte {

	bytes, err := json.Marshal(t)

	if err != nil {
		return nil
	}
	return bytes
}

func (t *Task) String() string {

	return string(t.Bytes())
}

type CreateTask struct {
	ProjectId   string `json:"projectId"`
	Description string `json:"description"`
	ToFinishAt  int64  `json:"toFinishAt"`
}

type UpdateTask struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	ToFinishAt  int64  `json:"toFinishAt"`
}
