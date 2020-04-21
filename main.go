package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	//"github.com/dimfeld/httptreemux"
	//"strings"
)

/* // obsolete variables with configuratino details, now delivered via config file
var httpAddr string
var httpsAddr string
var tlsCertPath string
var tlsKeyPath string
var staticDir string
*/

var (
	Version             string     // Version set during compile time, e.g. v0.1.42
	Commit              string     // git commit, set during compiletime
	Branch              string     // git branch, set during compile time
	Buildtime           string     // compile timestamp
	usr                 *user.User // variable that holds the user environment
	ConfigFilePath      string     // holds path to config file
	ConfigFileLocations []string   // holds default locations to look for a config file
	cfg                 *Config    // pointer to configuration slice
	err                 error      // global error variable
	showVersion         bool       // if true programm prints version, commit, branch and buildtime, then exit
)

func init() {
	// determine current user
	usr, err := user.Current()
	if err != nil {
		log.Fatal("Can not determine current user: ", err)
	}
	// fill default config locations
	ConfigFileLocations = []string{
		"/etc/TLSWebserver/config.yml",
		"/usr/local/etc/TLSWebServer/config.yml",
		filepath.Join(usr.HomeDir, ".config/TLSWebServer/config.yml"),
		"./config.yml",
	}
	// create a new config with default values
	cfg = NewConfig("")

	flag.StringVar(&ConfigFilePath, "conf", "./config.yml", "path to config file")
	flag.BoolVar(&showVersion, "version", false, "shows version information and exists")
	//flag.StringVar(&cfg.HttpAddr, "http", "", "http to https redirector listen address")
	//flag.StringVar(&cfg.HttpsAddr, "https", ":443", "https listen address")
	//flag.StringVar(&cfg.TLSCertPath, "cert", "./tls/cert.pem", "tls certificate PEM file")
	//flag.StringVar(&cfg.TLSKeyPath, "key", "./tls/key.pem", "tls key PEM file")
	//flag.StringVar(&cfg.StaticDir, "staticDir", "", "directory with static webcontent")
}

func main() {

	flag.Parse()
	if showVersion {
		fmt.Printf("Version: %s, based on branch: %s (commit: %s), Buildtime: %s\n", Version, Branch, Commit, Buildtime)
		os.Exit(1)
	}
	if ConfigFilePath != "" {
		cfg = NewConfig(ConfigFilePath)
	}
	log.Printf("config:\n%+v\n", cfg)
	if cfg.StaticDir == "" {
		log.Fatal("no staticDir given but required")
	}

	if cfg.TLSCertPath == "" {
		log.Fatal("no certificate file given, but required")
	}

	if cfg.TLSKeyPath == "" {
		log.Fatal("no key file given but required")
	}
	if cfg.httpsport == "" {
		log.Fatal("no HTTPS given but required")
	}
	if cfg.httpport == "" {
		log.Fatal("no HTTP given but required")
	}
	if cfg.domain == "" {
		log.Fatal("no Domain given but required")
	}

	app := &App{
		domain: 	cfg.domain,
		httpport: 	cfg.httpport,
		httpsport: 	cfg.httpsport,
		StaticDir: 	cfg.StaticDir,
		TLSCert:   	cfg.TLSCertPath,
		TLSKey:    	cfg.TLSKeyPath,
	}

go func() {
	log.Printf("Listening for HTTPS connections at: https://%v:%v", cfg.domain, cfg.httpsport)
		app.RunTLSServer()
	}()
	log.Printf("Starting Prime Frontend Server")
	log.Printf("Listening for HTTP connections at: http://%v:%v", cfg.domain, cfg.httpport)
	if err := http.ListenAndServe(cfg.domain+":"+cfg.httpport, http.HandlerFunc(app.redirectHandler)); err != nil {
  		log.Fatalf("ListenAndServe error: %v", err)
	}
}
