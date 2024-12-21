package device

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) List(startTime int64, endTime int64, locationId uuid.UUID) (Devices, error) {
	devices := make([]*Device, 0)
	// filter by dates or location
    queryDB := r.db.Preload("Location").Preload("DeviceType").Session(&gorm.Session{})
	if startTime > 0 {
		queryDB = queryDB.Where("created_time > ?", startTime)
	}
	if endTime > 0 {
		queryDB = queryDB.Where("created_time < ?", endTime)
	}
	if locationId != uuid.Nil {
		queryDB = queryDB.Where("location_id = ?", locationId)
	}
	if err := queryDB.Find(&devices).Error; err != nil {
		return nil, err
	}

	return devices, nil
}

