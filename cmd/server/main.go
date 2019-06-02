package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/AndrewBurian/eventsource"
	"github.com/fsnotify/fsnotify"
)

var flagPort int
var flagWASM string
var flagDirectory string
var flagSrc string
var flagBuild string

func init() {
	flag.IntVar(&flagPort, "port", 8080, "Port for http server")
	flag.StringVar(&flagWASM, "wasm", "WebAssembly.wasm", "WASM file name")
	flag.StringVar(&flagSrc, "src", "/src/", "Directory to run build command in")
	flag.StringVar(&flagDirectory, "directory", "/www/", "Directory to host")
	flag.StringVar(&flagBuild, "build", "go build", "Command to rebuild after changes")
}

func main() {
	flag.Parse()

	flagSrc = filepath.Clean(flagSrc) + "/"
	flagDirectory = filepath.Clean(flagDirectory) + "/"

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(flagDirectory))))

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	log.Println(flagSrc)
	if err := watcher.Add(flagSrc); err != nil {
		log.Fatal("watcher.Add():", err, flagSrc)
	}
	if err := watcher.Add(flagDirectory); err != nil {
		log.Fatal("watcher.Add():", err, flagDirectory)
	}
	stream := eventsource.NewStream()

	go events(watcher, stream)

	http.Handle("/hotreload", stream)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(flagPort), nil))
}

func events(watcher *fsnotify.Watcher, stream *eventsource.Stream) {
	for {
		select {
		case ev, ok := <-watcher.Events:
			if !ok {
				return
			}
			if ev.Op&fsnotify.Remove == fsnotify.Remove || ev.Op&fsnotify.Write == fsnotify.Write || ev.Op&fsnotify.Create == fsnotify.Create {
				file := filepath.Base(ev.Name)
				dir := filepath.Dir(ev.Name) + "/"
				if dir == flagSrc {
					build()
				} else if dir == flagDirectory {
					if strings.EqualFold(file, flagWASM) {
						stream.Broadcast(eventsource.DataEvent(file))
					}
				}
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

func build() bool {
	log.Println("Running build command!")

	args := strings.Split(flagBuild+" -a -o "+flagDirectory+flagWASM+" .", " ")
	if len(args) == 0 {
		// If the user has specified and empty then we are done.
		return true
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = flagSrc

	output, err := cmd.CombinedOutput()

	if err == nil {
		log.Println("Build ok.")
	} else {
		log.Println("Error while building:\n", string(output))
		log.Println("error", err)
	}

	return err == nil
}
