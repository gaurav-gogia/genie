package main

import "fmt"

func fmtmainpage(name, username string) string {
	import1 := fmt.Sprintf(
		`	"github.com/%s/%s/%s"
		`, username, name, name+"help")
	import2 := fmt.Sprintf(
		`
	"github.com/%s/%s/%s"
		`, username, name, name+"route")
	import3 := fmt.Sprintf(
		`
	"github.com/%s/%s/%s"
)
		`, username, name, name+"model")

	format1 := fmt.Sprintf(
		`
	%s.Session = %s.ConnectDB()
	defer %s.Session.Close()
		`, name+"help", name+"help", name+"help")

	format2 := fmt.Sprintf(`		Addr:         %s.Port,`, name+"model")
	format3 := fmt.Sprintf(
		`
	fmt.Println("Server is running at port", %s.Port)
	log.Println("Server is running at port", %s.Port)
	srv.ListenAndServe()
}
		`, name+"model", name+"model")

	format4 := fmt.Sprintf(
		`
func registerRoutes(r *mux.Router) {
	r.HandleFunc("/create", %s.Create).Methods("POST")
	r.HandleFunc("/read", %s.Read).Methods("GET")
	r.HandleFunc("/update", %s.Update).Methods("POST")
	r.HandleFunc("/delete", %s.Delete).Methods("POST")
	
	r.NotFoundHandler = http.NotFoundHandler()		
}
		`, name+"route", name+"route", name+"route", name+"route")

	return mainpage1 + import1 + import2 + import3 + mainpage2 + format1 + mainpage3 + format2 + mainpage4 + format3 + format4
}
