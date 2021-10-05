package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
)

func main() {
	err := steps()

	if err != nil {
		result := err.(*multierror.Error)

		result.ErrorFormat = func(errs []error) string {
			var builder strings.Builder
			for _, err := range errs {
				builder.WriteString(fmt.Sprintf("- %s", err))
				builder.WriteString("\n")
			}
			return builder.String()
		}

		fmt.Println(result)
	}
}

func steps() error {
	var result error

	if err := step1(); err != nil {
		result = multierror.Append(result, err)
	}

	if err := step2(); err != nil {
		result = multierror.Append(result, err)
	}

	return result
}

func step1() error {
	return errors.New("step1 error")
}

func step2() error {
	return errors.New("step2 error")
}