package main

import (
	"fmt"
	"strconv"
)

func main() {
	Sample("Atoi", Atoi)
	Sample("Itoa", Itoa)
	Sample("Parse", Parse)
	Sample("Format", Format)
	Sample("Append", Append)
	Sample("Quote", Quote)
	Sample("Other", Other)
}

func Atoi() {
	r, err := strconv.Atoi("65")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", r)
	}
}

func Itoa() {
	r := strconv.Itoa(65)
	fmt.Printf("%#v\n", r)
}

func Parse() {
	resultInt, errInt := strconv.ParseInt("65", 10, 32)
	if errInt != nil {
		fmt.Println(errInt)
	} else {
		fmt.Printf("%#v\n", resultInt)
	}

	resultUint, errUint := strconv.ParseUint("65", 10, 64)
	if errUint != nil {
		fmt.Println(errUint)
	} else {
		fmt.Printf("%#v\n", resultUint)
	}

	resultBool, errBool := strconv.ParseBool("True")
	if errBool != nil {
		fmt.Println(errBool)
	} else {
		fmt.Printf("%#v\n", resultBool)
	}

	resultFloat, errFloat := strconv.ParseFloat("7.2", 64)
	if errFloat != nil {
		fmt.Println(errFloat)
	} else {
		fmt.Printf("%#v\n", resultFloat)
	}

	resultComplex, errComplex := strconv.ParseComplex("1+4i", 64)
	if errComplex != nil {
		fmt.Println(errComplex)
	} else {
		fmt.Printf("%#v\n", resultComplex)
	}
}

func Format() {
	fmt.Printf("%v -> %#v\n", true, strconv.FormatBool(true))
	fmt.Printf("%v -> %#v\n", 7.20, strconv.FormatFloat(7.20, 'E', -1, 32))
	fmt.Printf("%v -> %#v\n", complex(1, 4), strconv.FormatComplex(complex(1, 4), 'E', -1, 64))
	fmt.Printf("%v -> %#v\n", -100, strconv.FormatInt(-100, 10))
	fmt.Printf("%v -> %#v\n", 100, strconv.FormatUint(100, 10))
}

func Append() {
	b10 := []byte("int (base 10):")
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))

	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))

	b := []byte("quote:")
	b = strconv.AppendQuote(b, `"Fran & Freddie's Diner"`)
	fmt.Println(string(b))

	bqr := []byte("rune:")
	bqr = strconv.AppendQuoteRune(bqr, '☺')
	fmt.Println(string(bqr))
}

func Quote() {
	var s string
	var err error

	s = strconv.Quote("Hey, Gopher!")
	fmt.Println(s)

	s = strconv.QuoteRune('☺')
	fmt.Println(s)

	s = strconv.QuoteRuneToGraphic('☺')
	fmt.Println(s)

	s = strconv.QuoteRuneToGraphic('\u263a')
	fmt.Println(s)

	s = strconv.QuoteRuneToGraphic('\u000a')
	fmt.Println(s)

	s = strconv.QuoteRuneToGraphic('	') // tab character
	fmt.Println(s)

	s, err = strconv.Unquote("You can't unquote a string without quotes")
	fmt.Printf("%q, %v\n", s, err)

	s, err = strconv.Unquote("\"The string must be either double-quoted\"")
	fmt.Printf("%q, %v\n", s, err)

	s, err = strconv.Unquote("`or backquoted.`")
	fmt.Printf("%q, %v\n", s, err)

	s, err = strconv.Unquote("'\u263a'") // single character only allowed in single quotes
	fmt.Printf("%q, %v\n", s, err)

	s, err = strconv.Unquote("'\u2639\u2639'") // not single character
	fmt.Printf("%q, %v\n", s, err)
}

func Other() {
	fmt.Println("IsGraphic:", strconv.IsGraphic('\u263a'))
	fmt.Println("IsPrint:", strconv.IsPrint('\u263a'))
	fmt.Println("CanBackquote:", strconv.CanBackquote("`Hey, there!`"))
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}