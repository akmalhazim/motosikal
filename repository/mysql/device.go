package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/akmalhazim/motosikal/models"
)

type deviceRepo struct {
	db *sql.DB
}

func (repo *deviceRepo) List(ctx context.Context) ([]*models.Device, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT `id`, `name`, `last_ping` FROM `devices`")
	if err != nil {
		return nil, err
	}

	devices := make([]*models.Device, 0)
	for rows.Next() {
		device := models.Device{}

		err = rows.Scan(&device.ID, &device.Name, &device.LastPing)
		if err != nil {
			return nil, err
		}

		devices = append(devices, &device)
	}

	return devices, nil
}

func (repo *deviceRepo) Save(ctx context.Context, device models.Device) error {
	return errors.New("unimplemented")
}

func NewDeviceRepo(db *sql.DB) *deviceRepo {
	return &deviceRepo{db}
}
