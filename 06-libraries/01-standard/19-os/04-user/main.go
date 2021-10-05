package main

import (
	"fmt"
	"os/user"
)

func main() {
	Sample("Current", Current)
	Sample("Lookup", Lookup)
	Sample("LookupGroup", LookupGroup)
}

func Current() {
	u, err := user.Current()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	groups, err := u.GroupIds()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Uid:", u.Uid)
	fmt.Println("Name:", u.Name)
	fmt.Println("Username:", u.Username)
	fmt.Println("HomeDir:", u.HomeDir)
	fmt.Println("Gid:", u.Gid)
	fmt.Println("Groups:", groups)
}

func Lookup() {
	u, err := user.Current()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	result, err := user.Lookup(u.Username)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Found:", result.Uid)
}

func LookupGroup() {
	u, err := user.Current()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	result, err := user.LookupGroupId(u.Gid)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Found:", result.Name)
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}