package dto

type Project struct {
	ID            int    `json:"id"`
	Link          string `json:"link"`
	ContainerName string `json:"containerName"`
	Description   string `json:"description"`
	UpToDate      bool   `json:"upToDate"`
	Deleted       bool   `json:"deleted"`
}

type CreateProject struct {
	Link          string `json:"link"`
	ContainerName string `json:"containerName"`
	Description   string `json:"description,omitempty"`
	UpToDate      bool   `json:"upToDate,omitempty"`
	Deleted       bool   `json:"deleted,omitempty"`
}

type Projects []Project
