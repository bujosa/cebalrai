package user

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn // Database connection
}

func (ur *UserRepository) GetUserByID(ctx context.Context, userID int) (*User, error) {
	var user User
	err := ur.db.QueryRow(ctx, "SELECT id, name, email FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) GetUsers(ctx context.Context) ([]*User, error) {
	var users []*User
	rows, err := ur.db.Query(ctx, "SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
