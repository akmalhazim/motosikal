package mysql

import (
	"context"
	"errors"

	"github.com/akmalhazim/motosikal/models"
)

type surveyRepo struct{}

func (repo *surveyRepo) Save(ctx context.Context, survey models.Survey) error {
	return errors.New("unimplemented")
}
