package kernel

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/derekdowling/pantry-api/api"
	"github.com/derekdowling/pantry-api/config"
	"github.com/gorilla/mux"
)

func init() {

	// loads our config into Confer so it can be used anywhere
	logMode := config.App.GetString("logging.mode")
	if logMode == "production" {
		// Log as JSON instead of the default ASCII formatter.
		log.SetFormatter(&log.JSONFormatter{})
		// log.SetOutput(logstash)
	} else {

		log.SetLevel(log.DebugLevel)

		// gives our logger file/line/stack traces
		log.SetFormatter(&log.TextFormatter{})
	}

}

// Start handles starting up our web kernel. It'll load our routes, controllers, and
// middleware.
func Start(production bool) {

	// get our stack rolling
	stack := buildStack(production)

	// figure out what port we need to be on
	port := config.App.GetString("ports.http")

	// output to help notify that the server is loaded
	log.WithFields(log.Fields{"port": port}).Info("Ready for requests with:")

	// start and log server output
	log.Fatal(http.ListenAndServe(port, stack))
}

// Handle's putting the whole stack together
func buildStack(production bool) *negroni.Negroni {
	// Build our contraption middleware and add the router
	// as the last piece
	stack := negroni.New()

	// define our list of production middleware here for now
	if production {
		// Turns on production API Keys
		config.App.Set("production", true)
		// Secure middleware has a Negroni integration, hence the wonky syntax
		// stack.Use(negroni.HandlerFunc(secureMiddleware().HandlerFuncWithNext))
	}

	stack.Use(negroni.NewLogger())

	// Builds our router and gives it routes
	router := buildRouter()

	// Serve static assets that the website requests
	// staticRoutes := config.App.GetStringMapString("static_routes")

	// log.Warn(staticRoutes)

	// for url, local := range staticRoutes {

	// log.WithFields(log.Fields{
	// "route": url,
	// "path":  local,
	// }).Info("Asset Path:")

	// router.PathPrefix(url).Handler(
	// http.FileServer(http.Dir(local)),
	// )
	// }

	stack.UseHandler(router)
	return stack
}

// Builds our routes
// http://www.gorillatoolkit.org/pkg/mux
func buildRouter() *mux.Router {

	// Create a Gorilla Mux Router
	router := mux.NewRouter()

	router.Queries("email", "")

	// Website Routes
	router.HandleFunc("/lists", api.GetLists).Methods("GET")
	router.HandleFunc("/lists", api.CreateList).Methods("POST")

	// Our 404 Handler
	// router.NotFoundHandler = http.HandlerFunc(home.Handle404)

	// API Routes
	// TODO: Rest Layer

	return router
}

// Sets our secure middleware based on what mode we are in
// func secureMiddleware() *secure.Secure {
// secureMiddleware := secure.New(secure.Options{
// AllowedHosts:          config.App.GetStringSlice("server.Allowed_Hosts"),
// SSLRedirect:           true,
// SSLHost:               config.App.GetString("server.SSL_Host"),
// SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
// STSSeconds:            315360000,
// STSIncludeSubdomains:  true,
// FrameDeny:             true,
// ContentTypeNosniff:    true,
// BrowserXssFilter:      true,
// ContentSecurityPolicy: "default-src 'self'",
// })
// return secureMiddleware
// }
