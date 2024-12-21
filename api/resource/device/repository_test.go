package device_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"

	"northpole-shop/api/resource/device"
	mockDB "northpole-shop/db/mock/db"
	testUtil "northpole-shop/util/test"
)

func TestRepository_List(t *testing.T) {
	t.Parallel()

	db, mock, err := mockDB.NewMockDB()
	testUtil.NoError(t, err)

	repo := device.NewRepository(db)

	mockRows := sqlmock.NewRows([]string{"id", "name", "serial_number", "device_type_id", "location_id"}).
		AddRow(uuid.New(), "Device1", "Serial1", uuid.NullUUID{}, uuid.NullUUID{}).
		AddRow(uuid.New(), "Device2", "Serial2", uuid.NullUUID{}, uuid.NullUUID{})

	mock.ExpectQuery("^SELECT (.+) FROM \"devices\"").WillReturnRows(mockRows)

	devices, err := repo.List(0, 0, uuid.Nil)
	testUtil.NoError(t, err)
	testUtil.Equal(t, len(devices), 2)
}

