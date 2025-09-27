package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, pass string) (*User, string)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo{
	return &userRepo{
		db: db,
	}
}

func (r userRepo) Create(user User) (*User, error){
	query := `
	INSERT INFO users (
		first_name,
		last_name,
		email,
		password,
		is_shop_owner
	)
	VALUES (
		:first_name,
		:last_name,
		:email,
		:password,
		:is_shop_owner
	)
	RETURNING id
`

	// Execute named query
	var userID int
	rows, err := r.db.NamedQuery(query, user)
	if err != nil{
		fmt.Println(err)
		return nil, err
	}
	if rows.Next() {
		rows.Scan(&userID)
	}

	user.ID = userID
	return &user, nil
}

func (r userRepo) Find(email, pass string) (*User, error) {
	var user User
	query := `
		SELECT
			id,
			first_name,
			last_name,
			email,
			password,
			is_shop_owner,
		FROM users
		WHERE email=$1 AND password=$2
		LIMIT 1
`
	err := r.db.Get(&user, query, email, pass)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // user not found
		}
		return nil, err
	}
	return &user, nil
}
