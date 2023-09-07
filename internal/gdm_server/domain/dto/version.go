package dto

type Version struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"project_id"`
	Version   string `json:"version"`
	Deleted   bool   `json:"deleted"`
}

type CreateVersion struct {
	ProjectID int    `json:"project_id"`
	Version   string `json:"version"`
	Deleted   bool   `json:"deleted"`
}

type Versions []Version
