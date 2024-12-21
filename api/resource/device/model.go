package device

import (

	"github.com/google/uuid"
)

type Location struct {
	ID            uuid.UUID `gorm:"primarykey"`
	Name          string
	CreatedTime   int64
}

type LocationDTO struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
}

func (b *Location) ToDto() *LocationDTO {
	return &LocationDTO{
		ID:            b.ID.String(),
		Name:          b.Name,
	}
}

type DeviceType struct {
	ID            uuid.UUID `gorm:"primarykey"`
	Name          string
	CreatedTime   int64
}

type DeviceTypeDTO struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
}

func (b *DeviceType) ToDto() *DeviceTypeDTO {
	return &DeviceTypeDTO{
		ID:            b.ID.String(),
		Name:          b.Name,
	}
}

type Device struct {
	ID            uuid.UUID `gorm:"primarykey"`
	Name          string
	SerialNumber  string
	DeviceTypeID  uuid.UUID
	DeviceType    DeviceType
	LocationID    uuid.UUID
	Location      Location
	CreatedTime   int64
}
type Devices []*Device

type DeviceDTO struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	SerialNumber  string `json:"serial_number"`
	CreatedTime   int64  `json:"created_time"`
	DeviceType    DeviceTypeDTO `json:"device_type,omitempty"`
	Location      LocationDTO `json:"location,omitempty"`
}

func (b *Device) ToDto() *DeviceDTO {
	return &DeviceDTO{
		ID:            b.ID.String(),
		Name:          b.Name,
		SerialNumber:  b.SerialNumber,
		CreatedTime:   b.CreatedTime,
		DeviceType:    *b.DeviceType.ToDto(),
		Location:      *b.Location.ToDto(),
	}
}

type DeviceForm struct {
	Name          string `json:"name" form:"required,max=255"`
	SerialNumber  string `json:"serial_number" form:"required,max=255"`
	LocationId    string `json:"location_id" form:"required,uuid"`
	DeviceTypeId  string `json:"device_type_id" form:"required,uuid"`
}


func (bs Devices) ToDto() []*DeviceDTO {
	dtos := make([]*DeviceDTO, len(bs))
	for i, v := range bs {
		dtos[i] = v.ToDto()
	}

	return dtos
}

func (f *DeviceForm) ToModel() *Device {
	deviceTypeID, _ := uuid.Parse(f.DeviceTypeId)
	locationID, _ := uuid.Parse(f.LocationId)
	return &Device{
		Name:          f.Name,
		SerialNumber:  f.SerialNumber,
		DeviceTypeID:  deviceTypeID,
		LocationID:    locationID,
	}
}
