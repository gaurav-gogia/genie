package main

import "fmt"

func fmtsessionpage(name string) string {
	format1 := fmt.Sprintf(
		`package %s
		`, name+"help")

	import1 := fmt.Sprintf(`
import (
	"context"
	"%s/%s"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
	`, name, name+"model")
	format2 := fmt.Sprintf(
		`ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(%s.ConnectionString)
	session, err := mongo.Connect(ctx, opts)
		`, name+"model")

	return format1 + import1 + sessionpage1 + format2 + sessionpage2
}
