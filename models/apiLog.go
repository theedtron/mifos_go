package models

import(
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ApiLogID struct {
	ID string `uri:"id" binding:"required"`
}

type ApiLog struct {
	ID        string `gorm:"primaryKey" json:"id"`
	RequestUrl string `json:"request_url" binding:"required"`
	RequestType  string `json:"request_type" binding:"required"`
	RequestBody  string `json:"request_body" binding:"required"`
	ResponseBody  string `json:"response_body" binding:"required"`
	CreatedAt int64 `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (x *ApiLog) FillDefaults() {
	if x.ID == "" {
		x.ID = uuid.New().String()
	}
}