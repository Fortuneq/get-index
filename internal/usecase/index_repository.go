package usecase

import (
	"context"
	"getProject/internal/entity"
)

type IndexRepository interface {
	GetByIndex(ctx context.Context, id int) ([]entity.Data, error)
}
