package model_repository

import "time"

type Category struct {
	Id               int64     `json:"id"`
	Name             string    `json:"name"`
	Version          int16     `json:"version"`
	IndexedAt        time.Time `json:"indexedAt"`
	CreatedBy        string    `json:"createdBy"`
	CreationDate     string    `json:"creationDate"`
	ModifiedBy       string    `json:"modifiedBy"`
	LastModifiedDate string    `json:"lastModifiedDate"`
}
