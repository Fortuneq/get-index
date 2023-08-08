package usecase

import (
	"context"
	"getProject/internal/entity"
)

type ParseIndexInputDTO struct {
	ID int `json:"id"`
}

type GetIndexInteractor interface {
	Execute(ctx context.Context, p ParseIndexInputDTO) ([]entity.Data, error)
}

type getIndexDataUseCase struct {
	indexRepo IndexRepository
}

func (u *getIndexDataUseCase) Execute(ctx context.Context, p ParseIndexInputDTO) ([]entity.Data, error) {
	data, err := u.indexRepo.GetByIndex(ctx, p.ID)
	switch err {
	case ErrNotFound:
		break
	case nil:
		return []entity.Data{}, entity.NewError("bad input id ", entity.ErrCodeBadInput)
	default:
		return []entity.Data{}, entity.NewError(err.Error(), entity.ErrCodeInternal)
	}

	return data, nil
}

var _ GetIndexInteractor = (*getIndexDataUseCase)(nil)

func NewCreateUserUseCase(indexRepo IndexRepository) *getIndexDataUseCase {
	return &getIndexDataUseCase{indexRepo}
}
