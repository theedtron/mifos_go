package models

import(
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationTemplate struct {
	ID        string `gorm:"primaryKey" json:"id"`
	NotificationType string `json:"notificationtype" binding:"required"`
	Subject  string `json:"subject" binding:"required"`
	Message  string `json:"message" binding:"required"`
	CreatedAt int64 `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (x *NotificationTemplate) FillDefaults() {
	if x.ID == "" {
		x.ID = uuid.New().String()
	}
}
