package parser

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	statusCode := "200"
	io.WriteString(w, "welcome to healthz handler!\n")
	io.WriteString(w, "http status is: "+statusCode+"\n")
}

func HeaderHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header().Set(k, v[0])
	}

	var GO_PROXY string
	GO_PROXY = os.Getenv("GOPROXY")
	w.Header().Set("Goproxy", GO_PROXY)

	clientIP := ""
	xForwardFor := r.Header.Get("X-Forwarded-For")
	if xForwardFor != "" {
		proxyIPs := strings.Split(xForwardFor, ",")
		if len(proxyIPs) > 0 {
			clientIP = proxyIPs[0]
			log.Printf("get", clientIP, "from x-forwarded-for")
		}
	} else {
		xRealIP := r.Header.Get("X-Real-IP")
		if xRealIP != "" {
			clientIP = xRealIP
			log.Printf("get", clientIP, "from x-real-ip")
		} else {
			// ip, port, err := net.SplitHostPort(r.RemoteAddr)
			// if err != nil {
			// 	io.WriteString(w, "not a true ip and port!")
			// }
			// fmt.Printf("ip is: %d\n", ip)
			// fmt.Printf("port is: %d\n", port)
			clientIP = r.RemoteAddr
			log.Printf("get", clientIP, "from remote address")
		}
	}
	io.WriteString(w, "welcome to header handler!\n")
	io.WriteString(w, "client IP is: "+clientIP+"\n")
}
