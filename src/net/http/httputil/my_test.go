package httputil

import (
	"net/http"
	"net/url"
)

func Example() {
	// 根据传入的url构建反向代理的
	proxy, err := NewHTTPProxy("http:127.0.0.1:8088")
	if err != nil {

	}
	http.Handle("/", proxy)
	http.ListenAndServe(":8081", nil)
}

type HTTPProxy struct {
	// 得到代理，
	proxy *ReverseProxy
}

func NewHTTPProxy(target string) (*HTTPProxy, error) {
	u, err := url.Parse(target)
	if err != nil {
		return nil, err
	}
	// 根据传入的url进行构建，
	return &HTTPProxy{NewSingleHostReverseProxy(u)}, nil
}

// 根据生成的的反向代理进行得到，
func (h *HTTPProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 在上述代码中，HTTPProxy 是一个包含 ReverseProxy 的结构体，
	// 当我们把URL解析成 *url.URL时，
	// 则可以调用httputil.NewSingleHostReverseProxy 函数为目标URL创建一个反向代理，
	// 同时HTTPProxy 需要实现 ServeHTTP 方法，这个方法可以将请求转发到实际代理的HTTP服务器中。
	h.proxy.ServeHTTP(w, r)
}
