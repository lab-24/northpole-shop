package query

import (
	"github.com/google/uuid"
)

type QueryDeviceList struct {
	StartTime   int64  `in:"query=start_time;default=0"`
	EndTime     int64     `in:"query=end_time;default=0"`
	LocationId  uuid.UUID     `in:"query=location_id;omitempty"`
}
