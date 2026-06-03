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

func (u *User) ValidateUser() error {

	fullNameLength := len([]rune(u.FullName))
	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf(
			"invalid FullName length: %d, %w",
			fullNameLength,
			core_errors.ErrInvalidArgument,
		)
	}

	if u.PhoneNumber != nil {

		phoneNumber := *u.PhoneNumber
		phoneLength := len([]rune(phoneNumber))
		if phoneLength < 8 || phoneLength > 15 {
			return fmt.Errorf(
				"invalid PhoneNumber length: %d, %w",
				phoneLength,
				core_errors.ErrInvalidArgument,
			)
		}

		re := regexp.MustCompile(`^\+[0-9]+$`)
		if !re.MatchString(phoneNumber) {
			return fmt.Errorf(
				"invalid PhoneNumber format: %v, %w",
				phoneNumber,
				core_errors.ErrInvalidArgument,
			)
		}

	}

	return nil
}

func (u *User) ApplyPatch(patch UserPatch) error {

	if err := patch.ValidatePatch(); err != nil {
		return fmt.Errorf("[user/domain]: failed validate user patch: %w", err)
	}

	tmp := *u
	if patch.FullName.Set {
		tmp.FullName = *patch.FullName.Value
	}
	if patch.PhoneNumber.Set {
		tmp.PhoneNumber = patch.PhoneNumber.Value
	}
	if err := tmp.ValidateUser(); err != nil {
		return fmt.Errorf("[user/domain]: failed validate user: %w", err)
	}
	*u = tmp

	return nil
}

type UserPatch struct {
	FullName    Nullable[string]
	PhoneNumber Nullable[string]
}

func (p *UserPatch) ValidatePatch() error {
	if p.FullName.Set && p.FullName.Value == nil {
		return fmt.Errorf(
			"FullNamme can't be patched to null %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
