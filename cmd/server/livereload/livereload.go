package livereload

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/AndrewBurian/eventsource"
	"github.com/RossMerr/Breeze/cmd/server/config"
	"github.com/fsnotify/fsnotify"
)

func Events(watcher *Batcher, stream *eventsource.Stream, env config.Env) {

	for {
		select {
		case events := <-watcher.Events:
			evs := map[string]fsnotify.Event{}

			for _, ev := range events {
				if ev.Op&fsnotify.Remove == fsnotify.Remove || ev.Op&fsnotify.Write == fsnotify.Write || ev.Op&fsnotify.Create == fsnotify.Create {
					evs[ev.Name] = ev
				}
			}

			if len(evs) > 0 {
				if www, ok := Build(env); ok {
					stream.Broadcast(eventsource.DataEvent(www))
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

func Build(env config.Env) (string, bool) {
	log.Println("Running build command!")

	args := strings.Split(env.Build+" -a -o "+env.Directory+env.WASM+" .", " ")
	if len(args) == 0 {
		return "", false
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = env.Src
	cmd.Env = []string{"GOOS=js", "GOARCH=wasm", "GOPATH=" + os.Getenv("GOPATH"), "HOME=" + os.Getenv("HOME")}
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Println("Error while building:\n", string(output))
		log.Println("error", err)
		return "", false
	}

	log.Println("Build ok.")
	return env.WASM, true
}
