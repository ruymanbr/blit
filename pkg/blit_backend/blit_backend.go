/*
 * File:    blit_backend.go
 *
 * Author:  	Ruymán Borges R. (ruyman21@gmail.com)
 * Date:    	1-7-21
 *
 * Summary of File:
 *
 *  This program runs a backend server for Blit package to handle http requests from a react frontend
 *
 */


package blit_backend

import (
	"github.com/ruymanbr/blit/pkg/blit_cli"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
    "os/signal"
    "time"
    "flag"
    "context"
	//"io/ioutil"
)

type File struct {
	isDir string
	lastmod string
	filename string
	size int
}

type Blit_Backend struct {
	path string
	files []File  `json: "omitempty"`
	totFiles int  `json: "totfiles"`
	totSize	int64 `json: "totfsize"`
}

type API interface {
	//DB *gorm.DB
	data Blit_Backend
	Initialize()
	ListHandler()
	ViewHandler()
	Path()
	Files()
	TotFiles()
	TotalSize()
	//pos interface
}

// Starts a backend that starts listening to http://localhost:8080
func Start() {
	var wait time.Duration
    flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
    flag.Parse()

    router := mux.NewRouter()

    srv := &http.Server{
        Addr:         "0.0.0.0:8080",
        // Good practice to set timeouts to avoid Slowloris attacks.
        WriteTimeout: time.Second * 15,
        ReadTimeout:  time.Second * 15,
        IdleTimeout:  time.Second * 60,
        Handler: router, // Pass our instance of gorilla/mux in.
    }
    api := &API{}

	//backend := &Blit_Backend{}
	data := api.Initialize()
	router.HandleFunc("/", api.ListHandler).Methods("GET")
    go func() {
        if err := srv.ListenAndServe(":8080", nil); err != nil {
            log.Println(err)
        }
    }()

	c := make(chan os.Signal, 1)
    // We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
    // SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
    signal.Notify(c, os.Interrupt)

    // Block until we receive our signal.
    <-c

    // Create a deadline to wait for.
    ctx, cancel := context.WithTimeout(context.Background(), wait)
    defer cancel()
    // Doesn't block if no connections, but will otherwise wait
    // until the timeout deadline.
    srv.Shutdown(ctx)
    // Optionally, you could run srv.Shutdown in a goroutine and block on
    // <-ctx.Done() if your application should wait for other services
    // to finalize based on context cancellation.
    log.Println("shutting down")
    os.Exit(0)

	//router.HandleFunc("/", data.handler)
	
	/*
	http.Handle("/", router)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
	*/
}

// handler function of type *Blit_Backend handles request and post methods from and to end points (localhost:8080)
func (backend *Blit_Backend) handler(writter http.ResponseWriter, router *http.Request) {


	// Write to HTTP response.
	writter.WriteHeader(200)
	writter.Write([]byte("Welcome to Blit: Let's list some folders!"))
}

func (api *API) Initialize() *Blit_Backend {
	api.data := &Blit_Backend{}
	return api.data
}

func (api *API) ListHandler(writter http.ResponseWriter, router *http.Request) {
	//resp, err := http.Get("http://localhost:8080")
	/*
	if err != nil {
		log.Fatalln(err)
	}*/
	
	//We Read the response body on the line below.
	/*
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	stringB := string(body)
	//log.Printf(stringB)
	*/

	blitback := Blit_Backend{}

	err := json.NewDecoder(router.Body).Decode(&blitback)
	if err != nil{
		panic(err)
	}

	//Marshal or convert user object back to json and write to response 
	blitbackJson, err := json.Marshal(blitback)
	if err != nil{
		panic(err)
	}	

	// os.Args?? Filter from request??
	blitback.path, _ := blit_cli.GetPath(os.Args)
	fileInfo, pathCorrect, err	:= blit_cli.HandlePath(blitback.path)
	if err != nil{
		panic(err)
	}
	encap_data, err, totSize 	:= blit_cli.EncapData(fileInfo, pathCorrect)
	if err != nil{
		panic(err)
	}
	sizesSli 					:= blit_cli.EncapSizes(fileInfo)

	totFiles 					:= len(sizesSli)
	blitback.totFiles 			:= totFiles

	_, dirList 					:= blit_cli.CleanData(encap_data)

	blit_cli.FileSizeSort(sizesSli, 1)
	sortedSli					:= blit_cli.FastSwitchSli(encap_data, sizesSli, 0)

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response 
	FilesJSON, _ := json.Marshal(sortedSli)
	w.Write(FilesJSON)

	// Select all stars and convert to JSON.
	//api.DB.Find(&stars)
	//FilesJSON, _ := json.Marshal(sortedSli)

	// Write to HTTP response.
	//writter.WriteHeader(200)
	//writter.Write([]byte(FilesJSON))
}

func (api *API) ViewHandler(writter http.ResponseWriter, router *http.Request) {
  var path string

  // Select all stars and convert to JSON.
  //api.DB.Find(&stars)
  starsJSON, _ := json.Marshal(stars)

  // Write to HTTP response.
  writter.WriteHeader(200)
  writter.Write([]byte(starsJSON))
}

func (api *API) Path() string {
	return api.data.path
}

func (api *API) Files() []File {
	return api.data.files
}

func (api *API) TotFiles() int {
	return api.data.totFiles
}

func (api *API) TotalSize() int64 {
	return api.data.totSize
}