package postgresql

import (
	"context"
	"database/sql"
	"go-api-learn/domain"
	"log"

	"github.com/google/uuid"
)

type UserRepository struct {
	Conn *sql.DB
}

func NewUserRepository(conn *sql.DB) *UserRepository {
	return &UserRepository{conn}
}

func (r *UserRepository) Store(ctx context.Context, u *domain.User) (*uuid.UUID, error) {
	query := `INSERT INTO users (username, email, password, name, role) 
          VALUES ($1, $2, $3, $4, $5) RETURNING uuid`

	row := r.Conn.QueryRowContext(ctx, query, u.Username, u.Email, u.Password, u.Name, u.Role)

	newUser := &domain.User{}

	err := row.Scan(&newUser.UUID, &newUser.Username, &newUser.Email, &newUser.Name, &newUser.Role, &newUser.CreatedAt, &newUser.UpdatedAt)
	if err != nil {
		log.Println("Error inserting new user:", err)
		return nil, err
	}

	return &newUser.UUID, nil
}
