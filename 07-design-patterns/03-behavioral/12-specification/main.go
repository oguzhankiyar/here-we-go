package main

import (
	"fmt"
)

type Specification interface {
	IsSatisfiedBy(object interface{}) bool
	And(Specification) Specification
	Or(Specification) Specification
	Not() Specification
	Relate(Specification)
}

type BaseSpecification struct {
	Specification
}

func (spec *BaseSpecification) IsSatisfiedBy(object interface{}) bool {
	return false
}

func (spec *BaseSpecification) And(other Specification) Specification {
	a := &AndSpecification{spec.Specification, other}
	a.Relate(a)
	return a
}

func (spec *BaseSpecification) Or(other Specification) Specification {
	a := &OrSpecification{spec.Specification, other}
	a.Relate(a)
	return a
}

func (spec *BaseSpecification) Not() Specification {
	a := &NotSpecification{spec.Specification}
	a.Relate(a)
	return a
}

func (spec *BaseSpecification) Relate(other Specification) {
	spec.Specification = other
}

type AndSpecification struct {
	Specification
	Other Specification
}

func (spec *AndSpecification) IsSatisfiedBy(object interface{}) bool {
	return spec.Specification.IsSatisfiedBy(object) && spec.Other.IsSatisfiedBy(object)
}

type OrSpecification struct {
	Specification
	Other Specification
}

func (spec *OrSpecification) IsSatisfiedBy(object interface{}) bool {
	return spec.Specification.IsSatisfiedBy(object) || spec.Other.IsSatisfiedBy(object)
}

type NotSpecification struct {
	Specification
}

func (spec *NotSpecification) IsSatisfiedBy(object interface{}) bool {
	return spec.Specification.IsSatisfiedBy(object)
}

type User struct {
	Name	string
	Enabled	bool
}

type UserNameSpecifiedSpecification struct {
	Specification
}

func (spec *UserNameSpecifiedSpecification) IsSatisfiedBy(object interface{}) bool {
	if user, ok := object.(User); ok {
		return len(user.Name) > 0
	}

	return false
}

func NewUserNameSpecifiedSpecification() Specification {
	a := &UserNameSpecifiedSpecification{&BaseSpecification{}}
	a.Relate(a)
	return a
}

type UserEnabledSpecification struct {
	Specification
}

func (spec *UserEnabledSpecification) IsSatisfiedBy(object interface{}) bool {
	if user, ok := object.(User); ok {
		return user.Enabled
	}

	return false
}

func NewUserEnabledSpecification() Specification {
	a := &UserEnabledSpecification{&BaseSpecification{}}
	a.Relate(a)
	return a
}

func main() {
	userNameSpecifiedSpec := NewUserNameSpecifiedSpecification()
	userEnabledSpec := NewUserEnabledSpecification()

	notValidSpec := userNameSpecifiedSpec.Not().Or(userEnabledSpec.Not())

	user := User{
		"Gopher",
		false,
	}

	result := notValidSpec.IsSatisfiedBy(user)

	fmt.Println("not valid:", result)
}