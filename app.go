package app

import (
	"fmt"
	"net/http"
)

// 初期化処理
func init() {
	http.HandleFunc("/", handler)
}

// リクエスト受付関数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
