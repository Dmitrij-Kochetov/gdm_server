package project

import "github.com/google/uuid"

type Projects []Project

type Project struct {
	ID            uuid.UUID
	Link          string
	ContainerName string
	UpToDate      bool
	Running       bool
	Deleted       bool
}
