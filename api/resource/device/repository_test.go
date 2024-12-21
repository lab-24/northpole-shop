package device_test

import (
	"testing"
	"time"

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

func TestRepository_Create(t *testing.T) {
	t.Parallel()

	db, mock, err := mockDB.NewMockDB()
	testUtil.NoError(t, err)

	repo := device.NewRepository(db)

	id := uuid.New()
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO \"devices\" ").
		WithArgs(id, "Name", "SerialNumber", uuid.NullUUID{}.UUID.String(), uuid.NullUUID{}.UUID.String(), time.Now().Unix()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	device := &device.Device{ID: id, Name: "Name", SerialNumber: "SerialNumber"}
	_, err = repo.Create(device)
	testUtil.NoError(t, err)
}

func TestRepository_Read(t *testing.T) {
	t.Parallel()

	db, mock, err := mockDB.NewMockDB()
	testUtil.NoError(t, err)

	repo := device.NewRepository(db)

	id := uuid.New()
	mockRows := sqlmock.NewRows([]string{"id", "name", "serial_number"}).
		AddRow(id, "Device1", "Serial1")

	mock.ExpectQuery("^SELECT (.+) FROM \"devices\" WHERE (.+)").
		WithArgs(id, 1).
		WillReturnRows(mockRows)

	device, err := repo.Read(id)
	testUtil.NoError(t, err)
	testUtil.Equal(t, "Device1", device.Name)
}

func TestRepository_Update(t *testing.T) {
	t.Parallel()

	db, mock, err := mockDB.NewMockDB()
	testUtil.NoError(t, err)

	repo := device.NewRepository(db)

	id := uuid.New()
	_ = sqlmock.NewRows([]string{"id", "name", "serial_number"}).
		AddRow(id, "Device1", "Serial1")

	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE \"devices\" SET").
		WithArgs("Name", "SerialNumber", uuid.NullUUID{}.UUID.String(), uuid.NullUUID{}.UUID.String(), id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	device := &device.Device{ID: id, Name: "Name", SerialNumber: "SerialNumber"}
	rows, err := repo.Update(device)
	testUtil.NoError(t, err)
	testUtil.Equal(t, 1, rows)
}

func TestRepository_Delete(t *testing.T) {
	t.Parallel()

	db, mock, err := mockDB.NewMockDB()
	testUtil.NoError(t, err)

	repo := device.NewRepository(db)

	id := uuid.New()
	_ = sqlmock.NewRows([]string{"id", "name", "serial_number"}).
		AddRow(id, "Device1", "Serial1")

	mock.ExpectBegin()
	mock.ExpectExec("^DELETE FROM \"devices\" WHERE (.+)").
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	rows, err := repo.Delete(id)
	testUtil.NoError(t, err)
	testUtil.Equal(t, 1, rows)
}
