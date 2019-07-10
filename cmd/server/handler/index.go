package handler

import (
	"net/http"

	"github.com/AndrewBurian/eventsource"
	"github.com/RossMerr/Breeze/cmd/server/config"
	"github.com/RossMerr/Breeze/cmd/server/livereload"
)

func Index(h http.Handler, env config.Env, stream *eventsource.Stream) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)

		go func() {
			if r.URL.Path == "/" {
				if www, ok := livereload.Build(env); ok {
					stream.Broadcast(eventsource.DataEvent(www))
				}
			}
		}()
	})
}
