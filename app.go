package main

// App holds the settings for tls web server
type App struct {
	domain	  string
	httpport  string
	httpsport string
	StaticDir string
	TLSCert   string
	TLSKey    string
}
