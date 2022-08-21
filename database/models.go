package database

import (
	"github.com/google/uuid"
)

type Link struct {
	ID  uuid.UUID `gorm:"type:uuid;primaryKey;not null;unique;default:gen_random_uuid()"`
	URL string    `form:"URL" json:"URL"`
}
