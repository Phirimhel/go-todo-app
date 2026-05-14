package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/Phirimhel/go-todo-app/internal/core/errors"
)

type User struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func NewUserUnitiliazied(FullName string, PhoneNumber *string) User {
	return NewUser(
		UnitiliaziedID,
		UnitiliaziedVersion,
		FullName,
		PhoneNumber,
	)
}

func NewUser(ID, Version int, FullName string, PhoneNumber *string) User {
	return User{
		ID:          UnitiliaziedID,
		Version:     UnitiliaziedVersion,
		FullName:    FullName,
		PhoneNumber: PhoneNumber,
	}
}

func (u *User) Validation() error {

	fullNameLength := len([]rune(u.FullName))
	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf(
			"invalid FullName length: %d, %w",
			fullNameLength,
			core_errors.ErrInvalidArgument,
		)
	}

	if u.PhoneNumber != nil {
		phoneLength := len([]rune(*u.PhoneNumber))
		if phoneLength < 8 || phoneLength > 15 {
			return fmt.Errorf(
				"invalid PhoneNumber length: %d, %w",
				phoneLength,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	re := regexp.MustCompile(`^\+[0-9]+$`)
	phoneNumber := *u.PhoneNumber
	if re.MatchString(phoneNumber) {
		return fmt.Errorf(
			"invalid PhoneNumber format: %v, %w",
			phoneNumber,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
