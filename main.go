package main


import (
    "fmt"
    "log"
	"flag"
	"embed"
	"io/fs"
	"strings"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	uc "github.com/mahani-software-engineering/website/usecases"
)

//go:embed pages/*
var pages embed.FS

func resourceNotFound(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(struct{Success string}{Success: "The API doesn't have what you are looking for !"})
}

func getRouter() *mux.Router {
	website, _ := fs.Sub(pages, "client/tailwind")
	
	router := mux.NewRouter()
	//---
	router.HandleFunc("/user/login", uc.UserLogin).Methods("POST")
	router.HandleFunc("/user/new", uc.RegisterUser).Methods("POST")
	//pages
	router.PathPrefix("/").Handler( http.FileServer(http.FS(website)) ).Methods("GET")
	//Not found
	router.NotFoundHandler = http.HandlerFunc(resourceNotFound)
	
	return router
}

func main() {
    //++++| os.Args |+++++
    baseAddress := ":443" 
    addr := flag.String("addr", baseAddress, "Web server address") 
    flag.Parse()
    //++++++++++++++++++++
    uc.Init()
    
    fmt.Println("Server listening on port: "+(strings.Split(baseAddress,":")[1])) 
    log.Fatal(http.ListenAndServe(*addr, getRouter()))
}












