package main

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestMyController(t *testing.T) {
	// 创建一个新的HTTP服务器，用于处理请求
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	server := &http.Server{Addr: ":8080", Handler: mux}

	// 启动HTTP服务器
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			t.Errorf("Error starting server: %v", err)
		}
	}()

	// 等待服务器启动
	time.Sleep(1 * time.Second)

	// 创建一个HTTP客户端
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// 发送一个GET请求到服务器
	resp, err := client.Get("http://localhost:8080")
	if err != nil {
		t.Errorf("Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()

	// 发送一个POST请求到服务器
	resp, err = client.Post("http://localhost:8080", "application/json", nil)
	if err != nil {
		t.Errorf("Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()

	// 检查响应是否成功
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code: %d", resp.StatusCode)
		return
	}

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		return
	}

	// 检查响应内容是否正确
	expected := "Hello, World!"
	if string(body) != expected {
		t.Errorf("Unexpected response body: %s", string(body))
		return
	}

	// 关闭HTTP服务器
	server.Close()
}
