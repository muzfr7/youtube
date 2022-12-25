package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/muzfr7/echoblog/entities"
)

// UserRepository provides actions to manipulate users in the database.
type UserRepository interface {
	User(ID int) (entities.User, error)
	Users() ([]entities.User, error)
	Insert(user entities.User) error
	Update(user entities.User) error
	Delete(ID int) error
}

// userRepo implements the UserRepository interface.
type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo instantiates userRepo struct.
func NewUserRepo(db *sqlx.DB) UserRepository {
	return &userRepo{
		db: db,
	}
}

// User fetches a user from the database and return it.
func (r *userRepo) User(ID int) (entities.User, error) {
	var user entities.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", ID)
	return user, err
}

// Users fetches all the users from the database and return them.
func (r *userRepo) Users() ([]entities.User, error) {
	var users []entities.User
	err := r.db.Select(&users, "SELECT * FROM users")
	return users, err
}

// Insert inserts a new user into the database.
func (r *userRepo) Insert(user entities.User) error {
	_, err := r.db.NamedExec("INSERT INTO users (firstname, lastname, email, status) VALUES (:firstname, :lastname, :email, :status)", user)
	if err != nil {
		return err
	}

	return nil
}

// Update updates given user in the database.
func (r *userRepo) Update(user entities.User) error {
	_, err := r.db.NamedExec("UPDATE users SET firstname = :firstname, lastname = :lastname, email = :email, status = :status) WHERE id = :id", user)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a user from the database.
func (r *userRepo) Delete(ID int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", ID)
	if err != nil {
		return err
	}

	return nil
}
