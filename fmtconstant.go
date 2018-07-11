package main

import "fmt"
import uuid "github.com/nu7hatch/gouuid"

func fmtconstantpage(name string) string {
	random, _ := uuid.NewV4()

	format1 := fmt.Sprintf(
		`package %s
		`, name+"model")
	format2 := fmt.Sprintf(
		`	XAuth            = "%s"
		`, random.String())
	format3 := fmt.Sprintf(`
	DatabaseString   = "%s"
)
		`, name+"db")

	return format1 + constantpage + format2 + format3
}
