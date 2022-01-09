// 编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：

// 接收客户端 request，并将 request 中带的 header 写入 response header
// 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
// Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
// 当访问 localhost/healthz 时，应返回 200
// 提交地址： https://jinshuju.net/f/fWHV2q
// 截止日期：2022 年 1 月 9 日 23:59

package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/header", headerHandler)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	statusCode := "200"
	io.WriteString(w, "welcome to healthz handler!\n")
	io.WriteString(w, "http status is: "+statusCode+"\n")
}

func headerHandler(w http.ResponseWriter, r *http.Request) {
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
