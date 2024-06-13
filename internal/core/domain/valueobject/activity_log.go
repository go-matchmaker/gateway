package valueobject

import (
	"github.com/google/uuid"
	"time"
)

type ActivityLog struct {
	UserID    uuid.UUID `json:"user_id"`
	IP        string    `json:"ip"`
	LastLogin time.Time `json:"last_login"`
}
