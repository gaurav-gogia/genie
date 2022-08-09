package main

import "fmt"

func fmtgetpage(name string) string {
	format1 := fmt.Sprintf(
		`package %s
		`, name+"route")

	import1 := fmt.Sprintf(`
import (
	"net/http"
	"time"

	"%s/%s"
	"%s/%s"
)
	`, name, name+"help", name, name+"model")

	format2 := fmt.Sprintf(
		`	go %s.AddSafeHeaders(w)

	// Your Code Here

	go %s.PrintLog("", %s.LevelInfo, time.Since(start), r)
}
		`, name+"help", name+"help", name+"model")

	return format1 + import1 + getroutepage + format2
}

func fmtpostpage(name string) string {
	format1 := fmt.Sprintf(
		`package %s
		`, name+"route")

	import1 := fmt.Sprintf(`
import (
	"net/http"
	"time"

	"%s/%s"
	"%s/%s"
)
			`, name, name+"help", name, name+"model")

	format2 := fmt.Sprintf(
		`	go %s.AddSafeHeaders(w)

		// Your Code Here

	go %s.PrintLog("", %s.LevelInfo, time.Since(start), r)
}
			`, name+"help", name+"help", name+"model")

	return format1 + import1 + postroutepage1 + format2 + postroutepage2 + format2 + postroutepage3 + format2
}
