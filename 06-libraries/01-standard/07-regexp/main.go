package main

import (
	"fmt"
	"regexp"
)

func main() {
	Sample("MatchString", MatchString)
	Sample("QuoteMeta", QuoteMeta)
	Sample("Expand", Expand)
	Sample("ExpandString", ExpandString)
	Sample("Find", Find)
	Sample("FindAll", FindAll)
	Sample("FindAllIndex", FindAllIndex)
	Sample("FindAllString", FindAllString)
	Sample("FindAllStringIndex", FindAllStringIndex)
	Sample("FindAllStringSubmatchIndex", FindAllStringSubmatchIndex)
	Sample("FindAllSubmatch", FindAllSubmatch)
	Sample("FindAllSubmatchIndex", FindAllSubmatchIndex)
	Sample("FindIndex", FindIndex)
	Sample("FindString", FindString)
	Sample("FindStringIndex", FindStringIndex)
	Sample("FindStringSubmatch", FindStringSubmatch)
	Sample("FindStringSubmatchIndex", FindStringSubmatchIndex)
	Sample("ReplaceAll", ReplaceAll)
	Sample("ReplaceAllFunc", ReplaceAllFunc)
	Sample("ReplaceAllLiteral", ReplaceAllLiteral)
	Sample("ReplaceAllLiteralString", ReplaceAllLiteralString)
	Sample("ReplaceAllString", ReplaceAllString)
	Sample("ReplaceAllStringFunc", ReplaceAllStringFunc)
	Sample("Split", Split)
	Sample("Longest", Longest)
	Sample("SubexpNames", SubexpNames)
	Sample("SubexpIndex", SubexpIndex)
	Sample("NumSubexp", NumSubexp)
}

func MatchString() {
	matched, err := regexp.MatchString("foo.*", "seafood")
	fmt.Println(matched, err) // true <nil>

	matched, err = regexp.MatchString("bar.*", "seafood")
	fmt.Println(matched, err) // true <nil>

	matched, err = regexp.MatchString("a(b", "sea(food)")
	fmt.Println(matched, err) // false missing closing ): `a(b`
}

func QuoteMeta() {
	fmt.Println(regexp.QuoteMeta(`Escaping symbols like: .+*?()|[]{}^$`)) // \.\+\*\?\(\)\|\[\]\{\}\^\$
}

func Expand() {
	content := []byte(`
		# comment line
		option1: value1
		option2: value2
	
		# another comment line
		option3: value3
	`)

	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	template := []byte("$key=$value\n")

	var result []byte

	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {
		result = pattern.Expand(result, template, content, submatches)
	}

	fmt.Print(string(result)) // option1=value1, option2=value2, option3=value3
}

func ExpandString() {
	content := `
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`

	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	template := "$key=$value\n"

	var result []byte

	for _, submatches := range pattern.FindAllStringSubmatchIndex(content, -1) {
		result = pattern.ExpandString(result, template, content, submatches)
	}

	fmt.Print(string(result)) // option1=value1, option2=value2, option3=value3
}

func Find() {
	pattern := regexp.MustCompile("test.*")
	var str []byte
	var result []byte

	str = []byte("i am tester")
	result = pattern.Find(str)
	fmt.Printf("%q -> %q\n", str, string(result)) // "tester"

	str = []byte("i am developer")
	result = pattern.Find(str)
	fmt.Printf("%q -> %q\n", str, string(result)) // ""
}

func FindAll() {
	pattern := regexp.MustCompile("foo.?")
	var str []byte
	var results [][]byte

	str = []byte("seafood fool")
	results = pattern.FindAll(str, -1)
	fmt.Printf("%q -> %q\n", str, results) // ["food" "fool"]
}

func FindAllIndex() {
	content := []byte("London")
	re := regexp.MustCompile(`o.`)

	fmt.Println(re.FindAllIndex(content, 1)) // [[1 3]]
	fmt.Println(re.FindAllIndex(content, -1)) // [[1 3] [4 6]]
}

func FindAllString() {
	re := regexp.MustCompile(`a.`)

	fmt.Println(re.FindAllString("paranormal", -1)) // [ar an al]
	fmt.Println(re.FindAllString("paranormal", 2)) // [ar an]
	fmt.Println(re.FindAllString("graal", -1)) // [aa]
	fmt.Println(re.FindAllString("none", -1)) // []
}

func FindAllStringIndex() {
	re := regexp.MustCompile(`a.`)

	fmt.Println(re.FindAllStringIndex("paranormal", -1)) // [[1 3] [3 5] [8 10]]
	fmt.Println(re.FindAllStringIndex("paranormal", 2)) // [[1 3] [3 5]]
	fmt.Println(re.FindAllStringIndex("graal", -1)) // [[2 4]]
	fmt.Println(re.FindAllStringIndex("none", -1)) // []
}

func FindAllStringSubmatchIndex() {
	re := regexp.MustCompile(`a(x*)b`)
	// Indices:
	//    01234567   012345678
	//    -ab-axb-   -axxb-ab-
	fmt.Println("-ab-", "->", re.FindAllStringSubmatchIndex("-ab-", -1)) // [[1 3 2 2]]
	fmt.Println("-axxb-", "->", re.FindAllStringSubmatchIndex("-axxb-", -1)) // [[1 5 2 4]]
	fmt.Println("-ab-axb-", "->", re.FindAllStringSubmatchIndex("-ab-axb-", -1)) // [[1 3 2 2] [4 7 5 6]]
	fmt.Println("-axxb-ab-", "->", re.FindAllStringSubmatchIndex("-axxb-ab-", -1)) // [[1 5 2 4] [6 8 7 7]]
	fmt.Println("-foo-", "->", re.FindAllStringSubmatchIndex("-foo-", -1)) // []
}

func FindAllSubmatch() {
	re := regexp.MustCompile(`foo(.?)`)
	fmt.Printf("%q\n", re.FindAllSubmatch([]byte(`seafood fool`), -1)) // [["food" "d"] ["fool" "l"]]
}

func FindAllSubmatchIndex() {
	content := []byte(`
	# comment line
	option1: value1
	option2: value2
`)
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)
	allIndexes := pattern.FindAllSubmatchIndex(content, -1)

	for _, loc := range allIndexes {
		fmt.Println(loc) // [18 33 18 25 27 33], [35 50 35 42 44 50]
		fmt.Println(string(content[loc[0]:loc[1]])) // option1: value1, option2: value2
		fmt.Println(string(content[loc[2]:loc[3]])) // option1, option2
		fmt.Println(string(content[loc[4]:loc[5]])) // value1, value2
	}
}

func FindIndex() {
	content := []byte(`
	# comment line
	option1: value1
	option2: value2
`)

	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	loc := pattern.FindIndex(content)
	fmt.Println(loc) // [18 33]
	fmt.Println(string(content[loc[0]:loc[1]])) // option1: value1
}

func FindString() {
	re := regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re.FindString("seafood fool")) // "food"
	fmt.Printf("%q\n", re.FindString("meat")) // ""
}

func FindStringIndex() {
	re := regexp.MustCompile(`ab?`)
	fmt.Println(re.FindStringIndex("tablett")) // [1 3]
	fmt.Println(re.FindStringIndex("foo") == nil) // true
}

func FindStringSubmatch() {
	re := regexp.MustCompile(`a(x*)b(y|z)c`)
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-")) // ["axxxbyc" "xxx" "y"]
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-")) // ["abzc" "" "z"]
}

func FindStringSubmatchIndex() {
	content := `
	# comment line
	option1: value1
	option2: value2
`
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)
	allIndexes := pattern.FindStringSubmatchIndex(content)

	fmt.Println(allIndexes) // [18 33 18 25 27 33]
	fmt.Println(content[allIndexes[0]:allIndexes[1]]) // option1: value1
	fmt.Println(content[allIndexes[2]:allIndexes[3]]) // option1
	fmt.Println(content[allIndexes[4]:allIndexes[5]]) // value1
}

func ReplaceAll() {
	re := regexp.MustCompile(`a(x*)b`)

	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("T"))) // -T-T-
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("$1"))) // --xx-
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("$1W"))) // ---
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("${1}W"))) // -W-xxW-
}

func ReplaceAllFunc() {
	re := regexp.MustCompile(`a(x*)b`)
	fn := func(b []byte) []byte {
		return []byte("@")
	}

	fmt.Printf("%s\n", re.ReplaceAllFunc([]byte("-ab-axxb-"), fn)) // -@-@-
}

func ReplaceAllLiteral() {
	re := regexp.MustCompile(`a(x*)b`)
	fmt.Printf("%q\n", re.ReplaceAllLiteral([]byte("-ab-axxb-"), []byte("T"))) // "-T-T-"
	fmt.Printf("%q\n", re.ReplaceAllLiteral([]byte("-ab-axxb-"), []byte("$1"))) // "-$1-$1-"
	fmt.Printf("%q\n", re.ReplaceAllLiteral([]byte("-ab-axxb-"), []byte("${1}"))) // "-${1}-${1}-"
}

func ReplaceAllLiteralString() {
	re := regexp.MustCompile(`a(x*)b`)
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "T")) // -T-T-
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "$1")) // -$1-$1-
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "${1}")) // -${1}-${1}-
}

func ReplaceAllString() {
	re := regexp.MustCompile(`a(x*)b`)

	fmt.Printf("%s\n", re.ReplaceAllString("-ab-axxb-", "T")) // -T-T-
	fmt.Printf("%s\n", re.ReplaceAllString("-ab-axxb-", "$1")) // --xx-
	fmt.Printf("%s\n", re.ReplaceAllString("-ab-axxb-", "$1W")) // ---
	fmt.Printf("%s\n", re.ReplaceAllString("-ab-axxb-", "${1}W")) // -W-xxW-
}

func ReplaceAllStringFunc() {
	re := regexp.MustCompile(`a(x*)b`)
	fn := func(s string) string {
		return "@"
	}

	fmt.Printf("%s\n", re.ReplaceAllStringFunc("-ab-axxb-", fn)) // -@-@-
}

func Split() {
	re := regexp.MustCompile(`a`)

	fmt.Printf("%q\n", re.Split("banana", -1)) // ["b" "n" "n" ""]
}

func Longest() {
	re := regexp.MustCompile(`a(|b)`)
	fmt.Println(re.FindString("ab")) // a
	re.Longest()
	fmt.Println(re.FindString("ab")) // ab

	fmt.Println(re.SubexpNames())
}

func SubexpNames() {
	re := regexp.MustCompile(`(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)`)
	fmt.Println(re.MatchString("Alan Turing")) // true
	fmt.Printf("%q\n", re.SubexpNames()) // ["" "first" "last"]

	reversed := fmt.Sprintf("${%s} ${%s}", re.SubexpNames()[2], re.SubexpNames()[1])
	fmt.Println(reversed) // ${last} ${first}
	fmt.Println(re.ReplaceAllString("Alan Turing", reversed)) // Turing Alan
}

func SubexpIndex() {
	re := regexp.MustCompile(`(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)`)
	fmt.Println(re.MatchString("Alan Turing")) // true
	fmt.Printf("%q\n", re.SubexpNames()) // ["" "first" "last"]

	for _, name := range re.SubexpNames() {
		fmt.Printf("%q -> %d\n", name, re.SubexpIndex(name)) // "" -> -1, "first" -> 1, "last" -> 2
	}
}

func NumSubexp() {
	re := regexp.MustCompile(`(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)`)
	fmt.Printf("%d\n", re.NumSubexp()) // 2

}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}