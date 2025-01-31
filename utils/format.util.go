package utils

import "github.com/gobeam/stringy"

func ToSnakeCase(str string) (result string) {
	result = stringy.New(str).SnakeCase().ToLower()

	return
}
