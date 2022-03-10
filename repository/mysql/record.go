package mysql

import (
	"context"
	"errors"

	"github.com/akmalhazim/motosikal/models"
)

type recordRepo struct{}

func (repo *recordRepo) Save(ctx context.Context, record models.Record) error {
	return errors.New("unimplemented")
}
