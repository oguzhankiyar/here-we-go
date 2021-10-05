package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func main() {
	Sample("ValidateStruct", ValidateStruct)
	Sample("ValidateVariable", ValidateVariable)
}

func ValidateStruct() {
	type Address struct {
		Street string `validate:"required"`
		City   string `validate:"required"`
		Planet string `validate:"required"`
		Phone  string `validate:"required"`
	}

	type User struct {
		FirstName      string     `validate:"required"`
		LastName       string     `validate:"required"`
		Age            uint8      `validate:"gte=0,lte=130"`
		Email          string     `validate:"required,email"`
		FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
		Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
	}

	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		Addresses:      []*Address{address},
	}

	validate := validator.New()

	err := validate.Struct(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("Namespace:", err.Namespace())
			fmt.Println("Field:", err.Field())
			fmt.Println("StructNamespace:", err.StructNamespace())
			fmt.Println("StructField:", err.StructField())
			fmt.Println("Tag:", err.Tag())
			fmt.Println("ActualTag:", err.ActualTag())
			fmt.Println("Kind:", err.Kind())
			fmt.Println("Type:", err.Type())
			fmt.Println("Value:", err.Value())
			fmt.Println("Param:", err.Param())
			fmt.Println()
		}
	}
}

func ValidateVariable() {
	v := validator.New()

	fn := func(str string) {
		err := v.Var(str, "required,email")
		if err != nil {
			fmt.Println(str, ">", "error:", err)
		} else {
			fmt.Println(str, ">", "valid")
		}
	}

	fn("gopher@golang.com")
	fn("golang.com")
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}