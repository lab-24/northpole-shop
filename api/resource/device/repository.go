package device

import (
	"time"

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

func (d *Device) BeforeCreate(tx *gorm.DB) (err error) {
    d.CreatedTime = time.Now().Unix()
    return nil
}

func (r *Repository) Create(device *Device) (*Device, error) {
	if err := r.db.Create(device).Error; err != nil {
		return nil, err
	}

	return device, nil
}

func (r *Repository) Read(id uuid.UUID) (*Device, error) {
	device := &Device{}
	if err := r.db.Preload("Location").Preload("DeviceType").Where("id = ?", id).First(&device).Error; err != nil {
		return nil, err
	}

	return device, nil
}

func (r *Repository) Update(device *Device) (int64, error) {
	result := r.db.Model(&Device{}).
		Select("Name", "SerialNumber", "DeviceTypeID", "LocationID").
		Where("id = ?", device.ID).
		Updates(device)

	return result.RowsAffected, result.Error
}

func (r *Repository) Delete(id uuid.UUID) (int64, error) {
	result := r.db.Where("id = ?", id).Delete(&Device{})

	return result.RowsAffected, result.Error
}
