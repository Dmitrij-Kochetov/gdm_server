package dto

type Version struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"project_id"`
	Version   string `json:"version"`
	Config    string `json:"config"`
	Deleted   bool   `json:"deleted"`
}

type CreateVersion struct {
	ProjectID int    `json:"project_id"`
	Version   string `json:"version"`
	Config    string `json:"config"`
}

type Versions []Version
