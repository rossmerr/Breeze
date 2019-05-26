package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"syscall"

	"github.com/AndrewBurian/eventsource"
	"github.com/fsnotify/fsnotify"
)

var port int
var public string
var directory string
var messages chan string

func init() {
	flag.IntVar(&port, "port", 8080, "Port for http server")
	flag.StringVar(&public, "public", "../../public/", "Public folder for WASM")
	flag.StringVar(&directory, "directory", "../../public/", "Directory to watch for changes")
}

func main() {
	flag.Parse()

	messages = make(chan string)

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(""))))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(public))))

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	if err := watcher.Add(directory); err != nil {
		log.Fatal("watcher.Add():", err)
	}

	go events(watcher)

	stream := eventsource.NewStream()

	go func(s *eventsource.Stream) {
		for t := range messages {
			msg := strings.ReplaceAll(t, public, "/public/")
			if strings.EqualFold(path.Ext(msg), ".wasm") {
				stream.Broadcast(eventsource.DataEvent(msg))
			}
		}
	}(stream)

	http.Handle("/stream", stream)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func events(watcher *fsnotify.Watcher) {
	for {
		select {
		case ev := <-watcher.Events:
			if ev.Op&fsnotify.Remove == fsnotify.Remove || ev.Op&fsnotify.Write == fsnotify.Write || ev.Op&fsnotify.Create == fsnotify.Create {
				messages <- ev.Name
			}

		case err := <-watcher.Errors:
			if v, ok := err.(*os.SyscallError); ok {
				if v.Err == syscall.EINTR {
					continue
				}
				log.Fatal("watcher.Error: SyscallError:", v)
			}
			log.Fatal("watcher.Error:", err)
		}
	}

}
