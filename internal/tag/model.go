package tag

import (
	"time"
)

type Tag struct {
	ID        int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Tag       string    `gorm:"type:varchar(255)" json:"tag"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
