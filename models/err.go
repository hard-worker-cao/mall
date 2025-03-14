package models

import "fmt"

func ErrHandler(err error) {
	fmt.Printf("%+v\n", err)
}
