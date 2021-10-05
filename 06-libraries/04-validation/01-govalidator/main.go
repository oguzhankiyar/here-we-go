package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"regexp"
	"time"
)

func main() {
	Sample("IsEmail", IsEmail)
	Sample("IsAlpha", IsAlpha)
	Sample("IsNumeric", IsNumeric)
	Sample("IsAlphanumeric", IsAlphanumeric)
	Sample("IsFloat", IsFloat)
	Sample("IsPositive", IsPositive)
	Sample("IsNegative", IsNegative)
	Sample("IsTime", IsTime)
	Sample("IsUnixTime", IsUnixTime)
	Sample("IsJSON", IsJSON)
	Sample("IsIP", IsIP)
	Sample("IsRequestURI", IsRequestURI)
	Sample("IsDataURI", IsDataURI)
	Sample("IsIn", IsIn)
	Sample("InRange", InRange)
	Sample("IsNull", IsNull)
	Sample("IsNotNull", IsNotNull)
	Sample("HasLowerCase", HasLowerCase)
	Sample("HasUpperCase", HasUpperCase)
	Sample("HasWhitespace", HasWhitespace)
	Sample("HasWhitespaceOnly", HasWhitespaceOnly)
	Sample("ValidateStruct", ValidateStruct)
	Sample("ValidateArray", ValidateArray)
	Sample("ValidateMap", ValidateMap)
	Sample("TagMap", TagMap)
}

func IsEmail() {
	fn := func(str string) {
		valid := govalidator.IsEmail(str)
		fmt.Printf("%q > %v\n", str, valid)
	}

	fn("admin@golang.org")
	fn("asd.com")
}

func IsAlpha() {
	fn := func(str string) {
		valid := govalidator.IsAlpha(str)
		fmt.Printf("%q > %v\n", str, valid)
	}

	fn("123")
	fn("12.33")
	fn("gopher")
	fn("***")
	fn("")
}

func IsNumeric() {
	fn := func(str string) {
		valid := govalidator.IsNumeric(str)
		fmt.Printf("%q > %v\n", str, valid)
	}

	fn("123")
	fn("12.33")
	fn("gopher")
	fn("***")
	fn("")
}

func IsAlphanumeric() {
	fn := func(str string) {
		valid := govalidator.IsAlphanumeric(str)
		fmt.Printf("%q > %v\n", str, valid)
	}

	fn("123")
	fn("12.33")
	fn("gopher")
	fn("gopher!")
	fn("***")
	fn("")
}

func IsFloat() {
	fn := func(str string) {
		valid := govalidator.IsFloat(str)
		fmt.Printf("%q > %v\n", str, valid)
	}

	fn("123")
	fn("12.33")
	fn("gopher")
	fn("***")
	fn("")
}

func IsPositive() {
	fn := func(in float64) {
		valid := govalidator.IsPositive(in)
		fmt.Printf("%v > %v\n", in, valid)
	}

	fn(10)
	fn(12.33)
	fn(0)
	fn(-1)
	fn(-125)
}

func IsNegative() {
	fn := func(in float64) {
		valid := govalidator.IsNegative(in)
		fmt.Printf("%v > %v\n", in, valid)
	}

	fn(10)
	fn(12.33)
	fn(0)
	fn(-1)
	fn(-125)
}

func IsTime() {
	fn := func(str string, format string) {
		valid := govalidator.IsTime(str, format)
		fmt.Printf("%q > %v\n", str, valid)
	}

	fn("2021-01-02T15:04:05Z", time.RFC3339)
	fn("2021-01-02", time.RFC3339)
	fn("", "HH:mm")
}

func IsUnixTime() {
	fn := func(str string) {
		valid := govalidator.IsUnixTime(str)
		fmt.Printf("%q > %v\n", str, valid)
	}

	fn("123")
	fn("-123")
	fn("")
}

func IsJSON() {
	fn := func(str string) {
		valid := govalidator.IsJSON(str)
		fmt.Printf("%q > %v\n", str, valid)
	}

	fn("{}")
	fn("[1]")
	fn("{id: 1}")
	fn(`{"id": 1}`)
	fn("123")
	fn("")
}

func IsIP() {
	fn := func(str string) {
		valid := govalidator.IsIP(str)
		fmt.Printf("%q > %v\n", str, valid)
	}

	fn("123")
	fn("127.0.0.1")
	fn("10.80.16.12")
	fn("2001:0db8:85a3:0000:0000:8a2e:0370:7334")
	fn("")
}

func IsRequestURI() {
	fn := func(str string) {
		valid := govalidator.IsRequestURI(str)
		fmt.Printf("%q > %v\n", str, valid)
	}

	fn("123")
	fn("127.0.0.1")
	fn("gopher")
	fn("localhost")
	fn("google.com")
	fn("http://")
	fn("https://localhost")
	fn("tcp://test.com")
	fn("ftp://10.80.16.12")
	fn("")
}

func IsDataURI() {
	fn := func(str string) {
		valid := govalidator.IsDataURI(str)
		fmt.Printf("%q > %v\n", str, valid)
	}

	fn("123")
	fn("http://")
	fn("https://localhost")
	fn("data:text/plain;base64,SGVsbG8sIEdvcGhlciE=")
}

func IsIn() {
	fn := func(str string, params ...string) {
		valid := govalidator.IsIn(str, params...)
		fmt.Printf("%q > %v\n", str, valid)
	}

	fn("a", "a", "b", "c")
	fn("", "a", "b", "c")
	fn("gopher")
}

func InRange() {
	fn := func(value interface{}, left interface{}, right interface{}) {
		valid := govalidator.InRange(value, left, right)
		fmt.Printf("%v > %v\n", value, valid)
	}

	fn(5, 0, 10)
	fn(100, 0, 10)
}

func IsNull() {
	fn := func(str string) {
		valid := govalidator.IsNull(str)
		fmt.Printf("%q > %v\n", str, valid)
	}

	fn("null")
	fn("gopher")
	fn("nil")
	fn("")
}

func IsNotNull() {
	fn := func(str string) {
		valid := govalidator.IsNotNull(str)
		fmt.Printf("%q > %v\n", str, valid)
	}

	fn("null")
	fn("gopher")
	fn("nil")
	fn("")
}

func HasLowerCase() {
	fn := func(str string) {
		valid := govalidator.HasLowerCase(str)
		fmt.Printf("%q > %v\n", str, valid)
	}

	fn("123")
	fn("hey")
	fn("GOPHER")
	fn("Go")
	fn("")
}

func HasUpperCase() {
	fn := func(str string) {
		valid := govalidator.HasUpperCase(str)
		fmt.Printf("%q > %v\n", str, valid)
	}

	fn("123")
	fn("hey")
	fn("GOPHER")
	fn("Go")
	fn("")
}

func HasWhitespace() {
	fn := func(str string) {
		valid := govalidator.HasWhitespace(str)
		fmt.Printf("%q > %v\n", str, valid)
	}

	fn("gopher")
	fn("h e y")
	fn("  ")
	fn("")
}

func HasWhitespaceOnly() {
	fn := func(str string) {
		valid := govalidator.HasWhitespaceOnly(str)
		fmt.Printf("%q > %v\n", str, valid)
	}

	fn("gopher")
	fn("h e y")
	fn("  ")
	fn("")
}

func ValidateStruct() {
	type Person struct {
		Name  string `valid:"required"`
		Email string `valid:"email"`
	}

	person := Person{
		Name: "Gopher",
		Email: "gopher@golang.org",
	}

	valid, err := govalidator.ValidateStruct(person)

	fmt.Println("valid:", valid)
	fmt.Println("error:", err)
}

func ValidateArray() {
	type Person struct {
		Name  string `valid:"required"`
		Email string `valid:"email"`
	}

	person1 := Person{
		Name: "Gopher 1",
		Email: "gopher1@golang.org",
	}

	person2 := Person{
		Name: "Gopher 2",
		Email: "gopher2@golang.org",
	}

	personList := []interface{}{person1, person2}

	valid := govalidator.ValidateArray(personList, func(item interface{}, _ int) bool {
		p := item.(Person)
		return govalidator.IsNotNull(p.Name)
	})

	fmt.Println("valid:", valid)
}

func ValidateMap() {
	var mapTemplate = map[string]interface{} {
		"name": "required,alpha",
		"family": "required,alpha",
		"email": "required,email",
		"cell-phone": "numeric",
		"address": map[string]interface{} {
			"line1": "required,alphanum",
			"line2": "alphanum",
			"postal-code": "numeric",
		},
	}

	var inputMap = map[string]interface{} {
		"name": "Bob",
		"family": "Smith",
		"email": "foo@bar.baz",
		"address": map[string]interface{} {
			"line1": "",
			"line2": "",
			"postal-code": "",
		},
	}

	valid, err := govalidator.ValidateMap(inputMap, mapTemplate)

	fmt.Println("valid:", valid)
	fmt.Println("error:", err)
}

func TagMap() {
	type Post struct {
		Title    string `valid:"alphanum,required"`
		Message  string `valid:"duck,ascii,not-same"`
		Message2 string `valid:"animal(dog),not-same"`
		AuthorIP string `valid:"ipv4"`
		Date     string `valid:"-"`
	}

	post := &Post{
		Title:   "MyExamplePost",
		Message: "duck",
		Message2: "dog",
		AuthorIP: "123.234.54.3",
	}

	govalidator.TagMap["duck"] = govalidator.Validator(func(str string) bool {
		return str == "duck"
	})

	govalidator.ParamTagMap["animal"] = govalidator.ParamValidator(func(str string, params ...string) bool {
		species := params[0]
		return str == species
	})

	govalidator.CustomTypeTagMap.Set("not-same", func(i interface{}, context interface{}) bool {
		switch v := context.(type) {
		case Post:
			return v.Message != v.Message2
		}
		return true
	})

	govalidator.ParamTagRegexMap["animal"] = regexp.MustCompile("^animal\\((\\w+)\\)$")

	valid, err := govalidator.ValidateStruct(post)

	fmt.Println("valid:", valid)
	fmt.Println("error:", err)
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}