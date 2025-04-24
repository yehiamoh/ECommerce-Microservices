package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	pb "user-service/gen/user"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(email, password, firstName, lastName, role string) (*pb.CreateUserResponse, error) {
    tx, err := r.db.Begin()
    if err != nil {
        return nil, err
    }
    defer tx.Rollback()

    var id int
    query := `
    INSERT INTO users (email, password, first_name, last_name, role)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id
    `
    err = tx.QueryRow(query, email, password, firstName, lastName, role).Scan(&id)
    if err != nil {
        return nil, err
    }

    err = tx.Commit()
    if err != nil {
        return nil, err
    }

    var userRole pb.Role
    switch role {
    case "ADMIN":
        userRole = pb.Role_ADMIN
    case "USER":
        userRole = pb.Role_CUSTOMER
    case "SELLER":
        userRole = pb.Role_SELLER
    default:
        return nil, fmt.Errorf("invalid role: %s", role)
    }

    user := &pb.User{
        Id:        strconv.Itoa(id),
        Email:     email,
        FirstName: firstName,
        LastName:  lastName,
        Role:      userRole,
    }
    response := &pb.CreateUserResponse{
        User: user,
    }
    return response, nil
}

func (r *UserRepository) GetUser(id int) (*pb.GetUserResponse, error) {
    user := &pb.User{}
    query := `
    SELECT id, email, first_name, last_name, role
    FROM users
    WHERE id = $1
    `

    var role string
    err := r.db.QueryRow(query, id).Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &role)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }

    switch role {
    case "ADMIN":
        user.Role = pb.Role_ADMIN
    case "USER":
        user.Role = pb.Role_CUSTOMER
    case "SELLER":
        user.Role = pb.Role_SELLER
    default:
        return nil, fmt.Errorf("invalid role in database: %s", role)
    }

    return &pb.GetUserResponse{
        User: user,
    }, nil
}