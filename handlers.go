package main

import (
	"net/http"
	"fmt"
	"os"
	"strings"
	"log"
)

// Redirect all traffic to HTTPS
func (app *App) redirectHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Inside redirectHandler redirecting to '%s' \n", "https://"+app.domain+":"+app.httpsport+"/index.html")
	http.Redirect(w, r, "https://"+app.domain+":"+app.httpsport+"/index.html", http.StatusMovedPermanently)
}

// Serve your index file
func (app *App) indexHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	log.Printf("Inside redirectHandler serving from '%s' \n", app.StaticDir+"/index.html")
	http.ServeFile(w, r, app.StaticDir+"/index.html")
}

// Handle Static Files (files containing a .extension)
// Send a not found error for any requests containing a .extension where the file does not exist
// Redirect any requests that were not for files to your to your SPA
func (app *App) staticHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	requestPath := r.URL.Path
	fileSystemPath := app.StaticDir + r.URL.Path
	endURIPath := strings.Split(requestPath, "/")[len(strings.Split(requestPath, "/"))-1]
	splitPath := strings.Split(endURIPath, ".")
	log.Printf("Inside staticHandler Step 1. values for  requestPath: '%s' fileSystemPath: '%s' endURIPath: '%s' splitPath: '%s' \n",requestPath, fileSystemPath, endURIPath, splitPath )
	if len(splitPath) > 1 {
	  if f, err := os.Stat(fileSystemPath); err == nil && !f.IsDir() {
		log.Printf("Inside staticHandler Step 2. value fileSystemPath: '%s' \n", fileSystemPath)
		http.ServeFile(w, r, fileSystemPath)
		return
	  }
	  //If not found just send back to index.html
	  //http.NotFound(w, r)
	  log.Printf("Inside staticHandler Step 3. serving from '%s' \n", app.StaticDir+"/index.html")
	  http.ServeFile(w, r, app.StaticDir+"/index.html")
	  return
	}
	//fmt.Fprintf(os.Stdout,"Inside staticHandler 4. serving from   %s ",app.StaticDir+"/index.html")
	log.Printf("Inside staticHandler Step 4. serving from '%s' \n", app.StaticDir+"/index.html")
	http.ServeFile(w, r, app.StaticDir+"/index.html")
  }

// Handle a simple healthcheck url for loadbalancers or whatnot
func (app *App) healthCheckHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	log.Printf("Inside healthCheckHandler sending an OK \n", )
	fmt.Fprintf(w, "OK!\r")
}

// Handler to redirect all NoFound to Angular index.html
func (app *App) angularHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Inside angularHandler serving from '%s' \n", app.StaticDir+"/index.html")
	//fmt.Fprintf(os.Stdout,"Inside angularHandler ")
	http.ServeFile(w, r, app.StaticDir+"/index.html")
}