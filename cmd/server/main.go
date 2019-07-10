package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/AndrewBurian/eventsource"
	"github.com/RossMerr/Breeze/cmd/server/config"
	"github.com/RossMerr/Breeze/cmd/server/handler"
	"github.com/RossMerr/Breeze/cmd/server/livereload"
)

var env config.Env

func init() {
	env = config.Env{}

	flag.IntVar(&env.Port, "port", 8080, "Port for http server")
	flag.StringVar(&env.WASM, "wasm", "WebAssembly.wasm", "WASM file name")
	flag.StringVar(&env.Src, "src", "/src/", "Directory to run build command in")
	flag.StringVar(&env.Directory, "directory", "/www/", "Directory to host")
	flag.StringVar(&env.Build, "build", "go build", "Command to rebuild after changes")
}

func main() {
	flag.Parse()

	env.Src = filepath.Clean(env.Src) + "/"
	env.Directory = filepath.Clean(env.Directory) + "/"

	watcher, err := livereload.New(1 * time.Second)
	defer watcher.Close()

	if err != nil {
		log.Fatal(err)
	}

	stream := eventsource.NewStream()

	go livereload.Events(watcher, stream, env)

	if err := watcher.Add(env.Src); err != nil {
		log.Fatal("watcher.Add():", err, env.Src)
	}

	http.Handle("/", handler.Index(http.StripPrefix("/", http.FileServer(http.Dir(env.Directory))), env, stream))
	http.Handle("/hotreload", stream)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(env.Port), nil))
}
