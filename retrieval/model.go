package main

type User struct {
	ID       string `json:"_id,omitempty" validate:"-"`
	Key      string `json:"_key,omitempty" validate:"-"`
	Rev      string `json:"_rev,omitempty" validate:"-"`
	Username string `json:"username" validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required,min=5,max=20"`
}
