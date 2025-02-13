package option

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestApply(t *testing.T) {
	u := &User{}
	Apply[User](u, WithName("Tom"), WithAge(18))
	assert.Equal(t, u, &User{name: "Tom", age: 18})
}
func TestApplyErr(t *testing.T) {
	u := &User{}
	err := ApplyErr[User](u, WithNameErr("Tom"), WithAgeErr(18))
	require.NoError(t, err)
	assert.Equal(t, u, &User{name: "Tom", age: 18})
	err = ApplyErr[User](u, WithNameErr(""), WithAgeErr(18))
	assert.Equal(t, err, errors.New("name empty"))
}

func WithName(s string) Option[User] {
	return func(u *User) {
		u.name = s
	}
}

func WithAge(i int) Option[User] {
	return func(u *User) {
		u.age = i
	}
}

func WithAgeErr(i int) OptionErr[User] {
	return func(u *User) error {
		if i < 18 {
			return errors.New("age error")
		}
		u.age = i
		return nil
	}
}

func WithNameErr(s string) OptionErr[User] {
	return func(u *User) error {
		if s == "" {
			return errors.New("name empty")
		}
		u.name = s
		return nil
	}
}

type User struct {
	name string
	age  int
}
