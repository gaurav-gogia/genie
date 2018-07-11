package main

import "fmt"

func fmthelpage(name string) string {
	format := fmt.Sprintf(
		`package %s
		`, name+"help")

	return format + helppage
}
