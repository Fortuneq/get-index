package postgres

import (
	"context"
	"getProject/internal/entity"
	"getProject/internal/usecase"
	"github.com/jmoiron/sqlx"
)

type IndexRepository struct {
	db *sqlx.DB
}

func (r *IndexRepository) GetByIndex(ctx context.Context, id int) ([]entity.Data, error) {

	//Абстрактный sql ,  с которого получаем данные

	//q := "SELECT id, name FROM user WHERE id = :id"
	//
	//stmt, err := r.db.PrepareNamedContext(ctx, q)
	//if err != nil {
	//	return entity.Data{}, fmt.Errorf("prepare statement: %w", err)
	//}
	//defer stmt.Close()
	//
	//var u entity.Data
	//err = stmt.GetContext(ctx, &u, map[string]any{"name": name})
	//if err != nil {
	//	return []entity.Data{}, fmt.Errorf("get context: %w", err)
	//}

	return []entity.Data{}, nil
}

var _ usecase.IndexRepository = (*IndexRepository)(nil)

func NewIndexRepository(db *sqlx.DB) *IndexRepository {
	return &IndexRepository{db}
}
