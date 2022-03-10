package mysql

import (
	"context"
	"errors"

	"github.com/akmalhazim/motosikal/models"
)

type recordRepo struct{}

func (repo *recordRepo) ListByDeviceID(ctx context.Context, deviceID string) ([]*models.Record, error) {
	return nil, errors.New("unimplemented")
}

func (repo *recordRepo) ListBySurveyID(ctx context.Context, surveyID string) ([]*models.Record, error) {
	return nil, errors.New("unimplemented")
}

func (repo *recordRepo) Save(ctx context.Context, record models.Record) error {
	return errors.New("unimplemented")
}
