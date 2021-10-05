package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	Sample("Simple", Simple)
	Sample("Struct", Struct)
	Sample("Custom", Custom)
}

func Simple() {
	gofakeit.Seed(0)

	fmt.Println("Name:", gofakeit.Name())
	fmt.Println("Email:", gofakeit.Email())
	fmt.Println("Phone:", gofakeit.Phone())
	fmt.Println("BS:", gofakeit.BS())
	fmt.Println("BeerName:", gofakeit.BeerName())
	fmt.Println("Color:", gofakeit.Color())
	fmt.Println("Company:", gofakeit.Company())
	fmt.Println("CreditCardNumber:", gofakeit.CreditCardNumber(&gofakeit.CreditCardOptions{}))
	fmt.Println("HackerPhrase:", gofakeit.HackerPhrase())
	fmt.Println("JobTitle:", gofakeit.JobTitle())
	fmt.Println("CurrencyShort:", gofakeit.CurrencyShort())
}

func Struct() {
	type Foo struct {
		Bar      string
		Int      int
		Pointer  *int
		Name     string    `fake:"{firstname}"`         // Any available function all lowercase
		Sentence string    `fake:"{sentence:3}"`        // Can call with parameters
		RandStr  string    `fake:"{randomstring:[hello,world]}"`
		Number   string    `fake:"{number:1,10}"`       // Comma separated for multiple values
		Regex    string    `fake:"{regex:[abcdef]{5}}"` // Generate string from regex
		Skip     *string   `fake:"skip"`                // Set to "skip" to not generate data for
		Created  time.Time								// Can take in a fake tag as well as a format tag
		CreatedFormat  time.Time `fake:"{year}-{month}-{day}" format:"2006-01-02"`
	}

	type FooBar struct {
		Bars    []string `fake:"{name}"`              // Array of random size (1-10) with fake function applied
		Foos    []Foo    `fakesize:"3"`               // Array of size specified with faked struct
		FooBars []Foo    `fake:"{name}" fakesize:"3"` // Array of size 3 with fake function applied
	}

	var f Foo
	gofakeit.Struct(&f)
	fmt.Println(f.Bar)
	fmt.Println(f.Int)
	fmt.Println(*f.Pointer)
	fmt.Println(f.Name)
	fmt.Println(f.Sentence)
	fmt.Println(f.RandStr)
	fmt.Println(f.Number)
	fmt.Println(f.Regex)
	fmt.Println(f.Skip)
	fmt.Println(f.Created.String())

	var fb FooBar
	gofakeit.Struct(&fb)
	fmt.Println(fb.Bars)
	fmt.Println(fb.Foos)
}

func Custom() {
	gofakeit.AddFuncLookup("friendname", gofakeit.Info{
		Category:    "custom",
		Description: "Random friend name",
		Example:     "bill",
		Output:      "string",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			return gofakeit.RandomString([]string{"bill", "bob", "sally"}), nil
		},
	})

	gofakeit.AddFuncLookup("jumbleword", gofakeit.Info{
		Category:    "jumbleword",
		Description: "Take a word and jumple it up",
		Example:     "loredlowlh",
		Output:      "string",
		Params: []gofakeit.Param{
			{Field: "word", Type: "int", Description: "Word you want to jumble"},
		},
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			word, err := info.GetString(m, "word")
			if err != nil {
				return nil, err
			}

			split := strings.Split(word, "")
			gofakeit.ShuffleStrings(split)
			return strings.Join(split, ""), nil
		},
	})

	type Foo struct {
		FriendName string `fake:"{friendname}"`
		JumbleWord string `fake:"{jumbleword:helloworld}"`
	}

	var f Foo
	gofakeit.Struct(&f)
	fmt.Printf("%s", f.FriendName) // bill
	fmt.Printf("%s", f.JumbleWord) // loredlowlh
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
