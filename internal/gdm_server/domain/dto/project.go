package dto

type Project struct {
	ID            int    `json:"id"`
	Link          string `json:"link"`
	ContainerName string `json:"containerName"`
	UpToDate      bool   `json:"upToDate"`
	Running       bool   `json:"running"`
	Deleted       bool   `json:"deleted"`
}

type CreateProject struct {
	Link          string `json:"link"`
	ContainerName string `json:"containerName"`
	UpToDate      bool   `json:"upToDate,omitempty"`
	Running       bool   `json:"running,omitempty"`
	Deleted       bool   `json:"deleted,omitempty"`
}

type Projects []Project
