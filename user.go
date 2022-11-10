package main

import (
	"database/sql"
	"time"
)

type DBManager struct {
	db *sql.DB
}

func NewDBManager(db *sql.DB) *DBManager {
	return &DBManager{
		db: db,
	}
}

type User struct {
	Id        int
	Name      string
	Lastname  string
	CreatedAt time.Time
}

func (b *DBManager) Create(user *User) (*User, error) {
	query := `
		insert into users(
		                  name,
		                  lastname
		) values ($1,$2)
			returning id,name,lastname,created_at
`
	row := b.db.QueryRow(
		query,
		user.Name,
		user.Lastname,
	)

	var result User
	err := row.Scan(
		&result.Id,
		&result.Name,
		&result.Lastname,
		&result.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (b *DBManager) Get(id int) (*User, error) {
	var result User

	query := `
		select 
			id,
			name,
			lastname,
			created_at
		from users
		where id = $1
`
	row := b.db.QueryRow(query, id)
	err := row.Scan(
		&result.Id,
		&result.Name,
		&result.Lastname,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (b *DBManager) GetAll() ([]*User, error) {
	var users []*User

	query := `
		select 
			id,
			name,
			lastname,
			created_at
		from users
`

	rows, err := b.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Lastname,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)

	}
	return users, nil
}

func (b *DBManager) Update(user *User) (*User, error) {
	query := `
		update users set
		                 name=$1,
		                 lastname=$2
		where id=$3
		returning id,name,lastname,created_at
`

	row := b.db.QueryRow(
		query,
		user.Name,
		user.Lastname,
		user.Id,
	)

	var result User
	err := row.Scan(
		&result.Id,
		&result.Name,
		&result.Lastname,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (b *DBManager) Delete(id int) error {
	query := "delete from users where id=$1"

	_, err := b.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
