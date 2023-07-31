package data

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
}

type Role string

const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)

func (ur *Data) GetAllUsers(ctx context.Context, limit, page uint64) ([]User, error) {
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

	users := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Role, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
