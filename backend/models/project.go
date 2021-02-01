package models

import (
	"encoding/json"
)

type Project struct {
	Id     string `bson:"_id" json:"id"`
	UserId string `bson:"userId" json:"userId"`
	Name   string `bson:"name" json:"name"`
}

func (p *Project) Bytes() []byte {

	bytes, err := json.Marshal(p)

	if err != nil {
		return nil
	}
	return bytes
}

func (p *Project) String() string {

	return string(p.Bytes())
}

func (p *Project) ToDTO(tasks []*Task) *ProjectDTO {
	return &ProjectDTO{
		Id:     p.Id,
		Name:   p.Name,
		UserId: p.UserId,
		Tasks:  tasks,
	}
}

type ProjectDTO struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	UserId string  `json:"userId"`
	Tasks  []*Task `json:"tasks"`
}

type CreateProject struct {
	Name   string `json:"name"`
	UserId string `json:"userId"`
}

type UpdateProject struct {
	Id string `json:"id"`
	UserId string `json:"userId"`
	Name string `json:"name"`
}
