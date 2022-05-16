package repository

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

const (
	TABLE_NAME string = "users"
)

var userMetadata = table.Metadata{
	Name:    TABLE_NAME,
	Columns: []string{"id", "username", "email", "firstname", "lastname", "avatar", "friend_count", "provider", "email_verified", "email_code", "password_hash", "role"},
	PartKey: []string{"id"},
}
var userTable = table.New(userMetadata)

type UserRepository interface {
	Create(context.Context, datastruct.User) (datastruct.User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, u datastruct.User) error
	UpdateVerified(ctx context.Context, email string, emailVerified bool) error
	RotateEmailCode(ctx context.Context, email string) error
	EmailTaken(ctx context.Context, email string) bool
	UsernameTaken(ctx context.Context, username string) bool
	GetById(ctx context.Context, id string) (datastruct.User, error)
	GetByEmail(ctx context.Context, email string) (datastruct.User, error)
	GetByUsername(ctx context.Context, username string) (datastruct.User, error)
}

type userRepository struct {
	sess *gocqlx.Session
}

func (r *userRepository) Create(ctx context.Context, u datastruct.User) (datastruct.User, error) {
	v := validator.New()
	err := v.Struct(u)
	if err != nil {
		return datastruct.User{}, err
	}

	stmt, names := qb.
		Insert(TABLE_NAME).
		Columns(userMetadata.Columns...).
		ToCql()

	code, err := utils.GenerateOTP(4)
	if err != nil {
		return datastruct.User{}, errors.New("Failed to generate Email Code")
	}

	u.EmailCode = code

	err = r.sess.
		Query(stmt, names).
		BindStruct(u).
		ExecRelease()
	if err != nil {
		return datastruct.User{}, err
	}

	return u, nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	stmt, names := qb.
		Delete(TABLE_NAME).
		Where(qb.Eq("id")).
		ToCql()

	err := r.sess.
		Query(stmt, names).
		BindMap((qb.M{"id": id})).
		ExecRelease()
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(ctx context.Context, u datastruct.User) error {
	b := qb.
		Update(TABLE_NAME).
		Where(qb.Eq("id"))

	if u.Username != "" {
		b.Set("username")
	}

	if u.Email != "" {
		b.Set("email")
	}

	if u.Firstname != "" {
		b.Set("firstname")
	}

	if u.Lastname != "" {
		b.Set("lastname")
	}

	if u.Avatar != "" {
		b.Set("avatar")
	}

	if u.PasswordHash != "" {
		b.Set("password_hash")
	}

	stmt, names := b.Existing().ToCql()

	err := r.sess.Query(stmt, names).
		BindMap((qb.M{
			"id":            u.Id,
			"username":      u.Username,
			"email":         u.Email,
			"firstname":     u.Firstname,
			"lastname":      u.Lastname,
			"avatar":        u.Avatar,
			"password_hash": u.PasswordHash,
		})).
		ExecRelease()
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetById(ctx context.Context, id string) (res datastruct.User, err error) {
	err = r.sess.
		Query(userTable.Get()).
		BindMap((qb.M{"id": id})).
		GetRelease(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *userRepository) RotateEmailCode(ctx context.Context, email string) error {
	code, err := utils.GenerateOTP(4)
	if err != nil {
		return errors.New("failed to generate Email Code")
	}

	stmt, names := qb.
		Update(TABLE_NAME).
		Where(qb.Eq("email")).
		Set("email_code").
		ToCql()

	err = r.sess.Query(stmt, names).
		BindMap((qb.M{
			"email":      email,
			"email_code": code,
		})).
		ExecRelease()
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) UpdateVerified(ctx context.Context, email string, emailVerified bool) error {
	stmt, names := qb.
		Update(TABLE_NAME).
		Where(qb.Eq("email")).
		Set("email_code").
		Existing().
		ToCql()

	err := r.sess.Query(stmt, names).
		BindMap((qb.M{
			"email":          email,
			"email_verified": emailVerified,
		})).
		ExecRelease()
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) EmailTaken(ctx context.Context, email string) bool {
	user := datastruct.User{}

	stmt, names := qb.
		Select(TABLE_NAME).
		Where(qb.Eq("email")).
		Limit(1).
		ToCql()

	err := r.sess.Query(stmt, names).
		BindMap((qb.M{
			"email": email,
		})).
		GetRelease(&user)
	if err != nil {
		return false
	}

	if user.Id == "" {
		return false
	}

	return true
}

func (r *userRepository) UsernameTaken(ctx context.Context, username string) bool {
	user := datastruct.User{}

	stmt, names := qb.
		Select(TABLE_NAME).
		Where(qb.Eq("username")).
		Limit(1).
		ToCql()

	err := r.sess.Query(stmt, names).
		BindMap((qb.M{
			"username": username,
		})).
		GetRelease(&user)
	if err != nil {
		return false
	}

	if user.Id == "" {
		return false
	}

	return true
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (res datastruct.User, err error) {
	stmt, names := qb.
		Select(TABLE_NAME).
		Where(qb.Eq("email")).
		ToCql()

	err = r.sess.Query(stmt, names).
		BindMap((qb.M{
			"email": email,
		})).
		GetRelease(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (res datastruct.User, err error) {
	stmt, names := qb.
		Select(TABLE_NAME).
		Where(qb.Eq("username")).
		ToCql()

	err = r.sess.Query(stmt, names).
		BindMap((qb.M{
			"username": username,
		})).
		GetRelease(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
