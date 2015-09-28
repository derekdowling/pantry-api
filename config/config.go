package config

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/jacobstr/confer"
)

func init() {
	LoadConfig()
}

// Exported Config Namespace
var App *confer.Config

// So we don't overwrite our existing configs
var loaded = false

// LoadConfig will lood a configuration if it hasn't already been loaded.
func LoadConfig() {
	if loaded == false {
		load()
	}

	loaded = true
}

// Load configuration data.
func load() *confer.Config {
	// Load the config
	App = confer.NewConfig()
	App.SetRootPath(getLoadPath())
	if errs := App.ReadPaths(getConfigPaths()...); errs != nil {
		fmt.Println(errs)
	}
	return App
}

// Determines the root path for our configuration data.
func getLoadPath() string {
	// Some magic to get the abs path of the file
	_, filename, _, _ := runtime.Caller(1)
	baseDir := strings.Join([]string{path.Dir(filename), "yml"}, "/")
	return baseDir
}

// Determines the applicable set of config files to load based on our
// current environment.
func getConfigPaths() []string {
	bursaEnv := os.Getenv("PANTRY_ENV")

	var paths []string

	paths = append(paths, "server.yml")

	// Server configuration
	if bursaEnv != "" {
		paths = append(paths, fmt.Sprintf("server.%s.yml", bursaEnv))
	} else {
		paths = append(paths, "server.development.yml")
	}

	// Database configuration
	return append(paths, "database.yml")
}
