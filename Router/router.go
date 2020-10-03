package Router

import (
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httputil"
	"net/url"
	"servel/Container"
	"servel/Server"
	"strconv"
	s "strings"
	"time"
)

func splitHost(url string) (hostname string, port uint64) {
	arr := s.Split(url, ":")
	port, _ = strconv.ParseUint(arr[1], 10, 64)
	return arr[0], port
}

func ProxyRoute(w http.ResponseWriter, r *http.Request) {
	path := chi.URLParam(r, "path")
	var environment = make(map[string]Container.Environment)
	host, _ := splitHost(r.Host)
	Server.GetApps(&environment)
	feed := true
	Container.RunExec(environment[host], &feed)
	time.Sleep(time.Millisecond * 50)
	port := environment[host].Port

	remote, err := url.Parse("http://localhost:" + port)

	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.Director = func(req *http.Request) {
		req.Header = r.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = path
	}

	if feed {
		proxy.ServeHTTP(w, r)
	}else{
		w.Write([]byte("END"))
	}

}
