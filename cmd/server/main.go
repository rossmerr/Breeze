package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"golang.org/x/net/websocket"
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
	http.Handle("/socket", websocket.Handler(echo))

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	if err := watcher.Add(directory); err != nil {
		log.Fatal("watcher.Add():", err)
	}

	go events(watcher)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func echo(ws *websocket.Conn) {
	for t := range messages {
		fmt.Println("Message", t)
		if err := websocket.Message.Send(ws, t); err != nil {
			fmt.Println("Can't send", err)
			break
		}
	}
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
