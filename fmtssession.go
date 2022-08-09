package main

import "fmt"

func fmtsessionpage(name string) string {
	format1 := fmt.Sprintf(
		`package %s
		`, name+"help")

	import1 := fmt.Sprintf(`
import (
	"%s/%s"
	"gopkg.in/mgo.v2"
)
	`, name, name+"model")

	format2 := fmt.Sprintf(
		`
	session, err := mgo.Dial(%s.ConnectionString)
		`, name+"model")

	return format1 + import1 + sessionpage1 + format2 + sessionpage2
}
