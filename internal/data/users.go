package data

import (
	"context"
	"database/sql"

	"github.com/luccasbarros/the-service/internal/dto"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

func (ur *UsersRepository) GetAllUsers(ctx context.Context, limit, page uint64) ([]dto.User, error) {
	offset := (page - 1) * limit

	stmt := qb.
		Select("id, email, name, role, created_at").
		From("users").
		Limit(limit).
		Offset(offset)

	sql, args, err := stmt.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := ur.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]dto.User, 0)

	for rows.Next() {
		var user dto.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Role, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
