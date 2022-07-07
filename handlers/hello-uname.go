package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/circle-makotom/hello-uname/uname"
)

type HelloUnameHandler struct {
	mu sync.Mutex // guards n
	n  int
}

type HelloUnameMessage struct {
	Message string    `json:"message"`
	At      time.Time `json:"at"`
	Uname   string    `json:"uname"`
	Ctr     int       `json:"ctr"`
}

func (h *HelloUnameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	unameStr, _ := uname.GetUnameString()

	if res, err := json.Marshal(h.BuildResponse(time.Now(), unameStr)); err != nil {
		fmt.Fprintln(w, "{}") // Empty JSON'ed object
	} else {
		fmt.Fprintln(w, string(res))
	}
}

func (h *HelloUnameHandler) BuildResponse(at time.Time, uname string) HelloUnameMessage {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++

	return HelloUnameMessage{
		Message: "Hello world!",
		At:      at,
		Uname:   uname,
		Ctr:     h.n,
	}
}
