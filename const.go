package main

const mainpage1 = `
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
`

const mainpage2 = `
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	r := mux.NewRouter()
	registerRoutes(r)
`

const mainpage3 = `
	srv := configServer(r)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	go startServer(srv)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	cancel()
	srv.Shutdown(ctx)
	fmt.Println("\nServer Shutdown GRACEFULLY")
}

func configServer(r http.Handler) *http.Server {
	srv := &http.Server{
		Handler:      handlers.CORS()(r),
`

const mainpage4 = `
		WriteTimeout: 20 * time.Second,
		ReadTimeout:  20 * time.Second,
	}

	return srv
}

func startServer(srv *http.Server) {
	f, err := os.OpenFile("./log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)
`

const constantpage = `
// Global constnats for entire application
const (
	Port             = ":8080"
	LevelInfo        = "INFO"
	LevelWarning     = "WARNING"
	LevelError       = "ERROR"
	ConnectionString = "mongodb://localhost/"
	AdminModelString = "AdminModel"
	DataModelString  = "DataModel"
	Response         = "200"
	ErrResponse      = "500"
`
const helppage = `
import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
	"time"

	"github.com/nu7hatch/gouuid"
)

// PrintLog function prints out log in a formatted way
func PrintLog(err, level string, time time.Duration, r *http.Request) {
	log.Println("<<Message Start>>")
	log.Println("Level: ", level)
	log.Println("Path: ", r.URL.Path)
	log.Println("Host: ", r.RemoteAddr)
	log.Println("Time: ", time)
	log.Println("Method: ", r.Method)
	log.Println("Err: ", err)
	log.Printf("<<Message End>>\n\n")
}

// HashTxt function returns md5 hash of a string
func HashTxt(txt string) string {
	hasher := md5.New()

	hasher.Write([]byte(txt))
	hashByte := hasher.Sum(nil)
	hash := hex.EncodeToString(hashByte)

	return hash
}

// GenRandom function generates a random key/string based on SHA1 identifier of name
func GenRandom() string {
	key, _ := uuid.NewV4()
	return key.String()
}

// ConvertTime function converts string date-time into timestamp object
// date-time must be in this format: "Fri Feb 9 00:46:11 IST 2018"
func ConvertTime(dateTime string) int64 {
	timeStamp, _ := time.Parse(time.UnixDate, dateTime)
	return timeStamp.Unix()
}

// AddSafeHeaders will add headers to the route
func AddSafeHeaders(w http.ResponseWriter) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("Accept-Encoding", "json")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("X-Frame-Options", "SAMEORIGIN")
	w.Header().Set("Strict-Transport-Security", "max-age=2592000; includeSubDomains")
}`

const postroutepage1 = `
// Create post/delete API for creating something
func Create(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
`

const postroutepage2 = `
// Update post/delete API for updating something
func Update(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
`

const postroutepage3 = `
// Delete post/delete API for deleting something
func Delete(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
`
const getroutepage = `
// Read post/put API for reading something
func Read(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
`
const sessionpage1 = `
// Session global variable for mongodb global session
var Session *mongo.Client

// ConnectDB function creates a global connection to the database
func ConnectDB() *mongo.Client {
`
const sessionpage2 = `
	if err != nil {
		panic(err)
	}

	return session
}
`
