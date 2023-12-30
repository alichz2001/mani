package main

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

type UserServiceImpl struct {
	repo *ManiDB
}

type UserService interface {
	Register(user *User) (*User, error)
}

func NewUserService(repo *ManiDB) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (s *UserServiceImpl) Register(user *User) (*User, error) {
	ctx := context.Background()

	user.Password = hasher(user.Password)

	docMetas, err := s.repo.UsersCollection.CreateDocument(ctx, user)
	if err != nil {
		return nil, errors.New("register error")
	}

	user.ID = docMetas.Key

	return user, nil
}

func (s *UserServiceImpl) ValidateUserCredential(user *User) (*User, bool) {

	ctx := context.Background()

	q := fmt.Sprintf("FOR i IN users FILTER i.username == \"%s\" LIMIT 1 RETURN i", user.Username)
	cursor, err := s.repo.Database.Query(ctx, q, nil)
	if err != nil {
		return nil, false
	}
	defer func() { _ = cursor.Close() }()

	var readUser *User
	readUser = &User{}

	_, err = cursor.ReadDocument(ctx, readUser)
	if err != nil {
		return nil, false
	}

	if strings.Compare(hasher(user.Password), readUser.Password) == 0 {
		return readUser, true
	}

	return nil, false
}

func hasher(in string) string {
	h := sha1.New()
	h.Write([]byte(in))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
