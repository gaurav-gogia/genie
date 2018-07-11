package main

import "fmt"

func fmtgetpage(name, username string) string {
	format1 := fmt.Sprintf(
		`package %s
		`, name+"route")

	import1 := fmt.Sprintf(`
import (		
	"net/http"
	"time"

	"github.com/%s/%s/%s"
	"github.com/%s/%s/%s"	
)
	`, username, name, name+"help", username, name, name+"model")

	format2 := fmt.Sprintf(
		`	go %s.AddSafeHeaders(w)

	// Your Code Here

	go %s.PrintLog("", %s.LevelInfo, time.Since(start), r)
}
		`, name+"help", name+"help", name+"model")

	return format1 + import1 + getroutepage + format2
}

func fmtpostpage(name, username string) string {
	format1 := fmt.Sprintf(
		`package %s
		`, name+"route")

	import1 := fmt.Sprintf(`
import (		
	"net/http"
	"time"
		
	"github.com/%s/%s/%s"	
	"github.com/%s/%s/%s"
)
			`, username, name, name+"help", username, name, name+"model")

	format2 := fmt.Sprintf(
		`	go %s.AddSafeHeaders(w)
	
		// Your Code Here
	
	go %s.PrintLog("", %s.LevelInfo, time.Since(start), r)
}
			`, name+"help", name+"help", name+"model")

	return format1 + import1 + postroutepage1 + format2 + postroutepage2 + format2 + postroutepage3 + format2
}
