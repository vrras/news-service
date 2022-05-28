package news

import (
	"time"
)

type News struct {
	ID        int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Title     string    `gorm:"type:varchar(255)" json:"title"`
	Content   string    `json:"content"`
	Topic     string    `gorm:"type:varchar(255)" json:"topic"`
	Tags      string    `gorm:"type:text" json:"tags"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FilterSpec struct {
	Topic  string
	Status string
}
