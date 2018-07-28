package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: genie <project_name>")
		return
	}

	path, _ := os.Getwd()
	data := strings.Split(path, "/")

	fmt.Println("Generating Folders....")
	generateFolders(os.Args[1])

	fmt.Println("Generating Boiler Plate....")
	generatePages(os.Args[1], data[len(data)-1])

	fmt.Println("Done ^.^")
}

func generatePages(name, username string) {
	page := fmtmainpage(name, username)
	createpage(page, "./"+name+"/main.go")

	createpage("package "+name+"db", "./"+name+"/"+name+"db/dbi.go")
	createpage("package "+name+"db", "./"+name+"/"+name+"db/dbo.go")
	createpage("package "+name+"db", "./"+name+"/"+name+"db/dbr.go")

	createpage("package "+name+"link", "./"+name+"/"+name+"link/link.go")

	page = fmtgetpage(name, username)
	createpage(page, "./"+name+"/"+name+"route/get.go")
	page = fmtpostpage(name, username)
	createpage(page, "./"+name+"/"+name+"route/post.go")

	createpage("package "+name+"model", "./"+name+"/"+name+"model/model.go")
	page = fmtconstantpage(name)
	createpage(page, "./"+name+"/"+name+"model/constant.go")
	page = fmtsessionpage(name, username)
	createpage(page, "./"+name+"/"+name+"help/session.go")

	page = fmthelpage(name)
	createpage(page, "./"+name+"/"+name+"help/help.go")
}

func generateFolders(name string) {
	os.Mkdir(name, os.ModePerm)
	os.Mkdir("./"+name+"/"+name+"db", os.ModePerm)
	os.Mkdir("./"+name+"/"+name+"link", os.ModePerm)
	os.Mkdir("./"+name+"/"+name+"help", os.ModePerm)
	os.Mkdir("./"+name+"/"+name+"model", os.ModePerm)
	os.Mkdir("./"+name+"/"+name+"route", os.ModePerm)
}

func createpage(page, path string) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(page)
	if err != nil {
		panic(err)
	}
}
