package data

import "database/sql"

type UsersRepository struct {
  db *sql.DB
}

func NewUserRepository(db *sql.DB) *UsersRepository {
  return &UsersRepository{
    db: db,
  }
}

func (ur *UsersRepository) GetAllUsers() ([]int, error) {
  return []int{}, nil
}