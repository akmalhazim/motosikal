package mysql

import (
	"context"
	"errors"

	"github.com/akmalhazim/motosikal/models"
)

type deviceRepo struct{}

func (repo *deviceRepo) Save(ctx context.Context, device models.Device) error {
	return errors.New("unimplemented")
}
