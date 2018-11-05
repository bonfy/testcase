package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreetHandler(t *testing.T) {
	//创建一个请求
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 我们创建一个 ResponseRecorder (which satisfies http.ResponseWriter)来记录响应
	rec := httptest.NewRecorder()

	//直接使用greetHandler，传入参数rr,req
	greetHandler(rec, req)

	// 检测返回的状态码
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// 检测返回的数据
	expected := `Hello World!`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}
