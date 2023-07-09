package data

import (
	"context"

	"github.com/luccasbarros/the-service/internal/dto"
)

func (ur *Data) GetAllUsers(ctx context.Context, limit, page uint64) ([]dto.User, error) {
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

	rows, err := ur.db.Query(ctx, sql, args...)
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
