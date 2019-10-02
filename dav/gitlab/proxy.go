package gitlab

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/SpeedVan/go-common/client/httpclient/gitlab"
	"github.com/SpeedVan/go-common/config"
	"github.com/SpeedVan/proxy-in-dav/dav"
)

// DAVProxy todo
type DAVProxy struct {
	dav.DAV
	GitlabHTTPClient *gitlab.Client
}

// New todo
func New(config config.Config) (*DAVProxy, error) {

	gitlabHTTPClient, err := gitlab.New(config)
	if err != nil {
		return nil, err
	}

	return &DAVProxy{
		GitlabHTTPClient: gitlabHTTPClient,
	}, nil
}

// NewHandleFunc todo
func NewHandleFunc(path string, config config.Config) (string, func(http.ResponseWriter, *http.Request)) {
	o, err := New(config)
	if err != nil {
		log.Fatal(err)
	}
	// http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	// u, p, ok := r.BasicAuth()
	// 	// if !(ok == true && u == wd.Config.WebDav.Username && p == wd.Config.WebDav.Password) {
	// 	// 	w.Header().Set("WWW-Authenticate", `Basic realm="davfs"`)
	// 	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	// 	return
	// 	// }
	// 	// h.ServeHTTP(w, r)
	// })

	//localhost:8887/{protocol:(http|https)}/{domain}/{group}/{project}/{sha}/{path:.*} liunx挂载proxy服务地址

	return path, func(w http.ResponseWriter, r *http.Request) {
		// url := r.URL.Path
		switch r.Method {
		case "HEAD":
			o.Head(w, r)
		case "GET":
			o.Get(w, r)
		case "PROPFIND":
			o.Propfind(w, r)
		case "OPTIONS":
			o.Options(w, r)

		}
	}
}

// Head todo
func (s *DAVProxy) Head(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("head")
	fmt.Println(vars["protocol"])
	fmt.Println(vars["domain"])
	fmt.Println(vars["group"])
	fmt.Println(vars["project"])
	fmt.Println(vars["sha"])
	fmt.Println("Path:" + vars["path"])
	header := w.Header()
	header.Set("Accept-Ranges", "bytes")
	header.Set("Content-Length", "18")
	header.Set("Content-Type", "text/plain; charset=utf-8")
	header.Set("Etag", "\"13442cef32eaa60012\"")
	header.Set("Last-Modified", "Sun, 29 Dec 2013 02:26:31 GMT")
	header.Set("Date", "Mon, 30 Sep 2019 02:08:43 GMT")
	w.WriteHeader(200)
	// r.URL.Path
	// s.GitlabHTTPClient
}

// Get todo
func (s *DAVProxy) Get(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Set("Accept-Ranges", "bytes")
	header.Set("Content-Length", "18")
	header.Set("Content-Type", "text/plain; charset=utf-8")
	header.Set("Etag", "\"13442cef32eaa60012\"")
	header.Set("Last-Modified", "Sun, 29 Dec 2013 02:26:31 GMT")
	header.Set("Date", "Mon, 30 Sep 2019 02:08:43 GMT")
	w.WriteHeader(200)
	w.Write([]byte("# ~/.bash_logout\n\n"))
}

// Propfind todo
func (s *DAVProxy) Propfind(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Set("Content-Type", "text/xml; charset=utf-8")
	header.Set("Transfer-Encoding", "chunked")
	w.WriteHeader(207)
	w.Write([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>"))
	w.Write(data_dav)
}

// Options todo
func (s *DAVProxy) Options(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Set("Allow", "OPTIONS, PROPFIND")
	header.Set("Dav", "1, 2")
	header.Set("Ms-Author-Via", "DAV")
	w.Write(data_dav)
}

// Propfind2 todo
func (s *DAVProxy) Propfind2(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Set("Content-Type", "text/xml; charset=utf-8")
	w.WriteHeader(207)
	w.Write([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>"))
	w.Write(data_dav_bash_logout)
}

// Options2 todo
func (s *DAVProxy) Options2(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Set("Allow", "OPTIONS, GET, HEAD, PROPFIND")
	header.Set("Dav", "1, 2")
	header.Set("Ms-Author-Via", "DAV")
	w.Write(data_dav)
}
