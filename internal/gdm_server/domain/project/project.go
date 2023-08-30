package project

import "github.com/google/uuid"

type Project struct {
	ID            uuid.UUID
	Link          string
	ContainerName string
	UpToDate      bool
	Running       bool
	Deleted       bool
}
