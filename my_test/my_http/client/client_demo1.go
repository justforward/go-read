package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

func main() {

	cachedTransport := newTransport()

	// 构建client 传入构建的transportCache
	client := &http.Client{
		// 定义每次需要从缓存区获取请求的行为
		Transport: cachedTransport,
		Timeout:   time.Second * 5,
	}

	cacheClearTicker := time.NewTicker(time.Second * 5)

	// 间隔一秒的时间进行请求
	reqTicker := time.NewTicker(time.Second * 1)

	terminateChannel := make(chan os.Signal, 1)

	// 将输入信号转发到c
	// 以上代码告诉 signal ，将对应的信号通知 terminateChannel，然后在 for 循环中针对不同信号做不同的处理， for 循环为死循环。
	signal.Notify(terminateChannel, syscall.SIGTERM, syscall.SIGHUP)

	// 构建请求的时候，需要携带http 否则请求不到
	request, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080", strings.NewReader("xxxxx"))
	if err != nil {
		fmt.Println(err)
	}

	for {
		select {
		// 5s 进行刷新transport cached中的数据
		case <-cacheClearTicker.C:
			cachedTransport.Clear()

			// 获取到对应的信号做出不同的处理
		case <-terminateChannel:
			cacheClearTicker.Stop()
			reqTicker.Stop()
			return

		case <-reqTicker.C:
			// 一秒请求一次
			resp, err := client.Do(request)
			if err != nil {
				log.Printf("An error occurred.... %v", err)
				continue
			}

			// response 的body 是ReadCloser 类型 需要ioutil.ReadAll  来获取数据
			buf, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				log.Printf("An error occurred.... %v", err)
				continue
			}

			fmt.Printf("The body of the response is \"%s\" \n\n", string(buf))

		}

	}
}

func cacheKey(r *http.Request) string {
	return r.URL.String()
}

//构建
type TransportCache struct {
	//保存url请求的缓存
	data map[string]string
	// RWMutex：读写锁，RWMutex 基于 Mutex 实现
	mu sync.RWMutex

	originalTransport http.RoundTripper
}

func (c *TransportCache) Set(r *http.Request, val string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	fmt.Println("cache key:", cacheKey(r))
	fmt.Println("cache value:", val)
	c.data[cacheKey(r)] = val
}

func (c *TransportCache) Get(r *http.Request) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if val, exist := c.data[cacheKey(r)]; exist {
		return val, nil
	}
	return "", errors.New("url not found in cache")
}

/*
对于http客户端，
可以使用不同的实现了 RoundTripper 接口的Transport实现来配置它的行为。
RoundTripper 有点像 http.Client 的中间件
*/

func (c *TransportCache) RoundTrip(r *http.Request) (*http.Response, error) {

	// Check if we have the response cached..
	// If yes, we don't have to hit the server
	// We just return it as is from the cache store.
	if val, err := c.Get(r); err == nil {
		fmt.Println("Fetching the response from the cache")
		return cachedResponse([]byte(val), r)
	}

	// Ok, we don't have the response cached, the store was probably cleared.
	// Make the request to the server.
	// 从roundTrip 中获取response的结果
	resp, err := c.originalTransport.RoundTrip(r)

	if err != nil {
		return nil, err
	}

	// Get the body of the response so we can save it in the cache for the next request.
	// 这个请求会转存相应
	buf, err := httputil.DumpResponse(resp, true)

	if err != nil {
		return nil, err
	}

	// Saving it to the cache store
	// 将结果缓存在客户端
	c.Set(r, string(buf))

	fmt.Println("Fetching the data from the real source")
	return resp, nil
}

func cachedResponse(b []byte, r *http.Request) (*http.Response, error) {
	buf := bytes.NewBuffer(b)
	return http.ReadResponse(bufio.NewReader(buf), r)
}

func newTransport() *TransportCache {

	return &TransportCache{
		data:              make(map[string]string),
		originalTransport: http.DefaultTransport,
	}
}

func (c *TransportCache) Clear() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data = make(map[string]string)
	return nil
}
