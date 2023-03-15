//go:build js && wasm

package main

//  照着 https://geektutu.com/post/quick-go-wasm.html 来的

import (
	"syscall/js"
)

func jsWrapperSumByGo(this js.Value, args []js.Value) interface{} {
	var va, vb int
	if len(args) >= 2 {
		va = args[0].Int()
		vb = args[1].Int()
	}
	callback := args[2]

	result := va + vb + 1000
	callback.Invoke(result)
	return nil
}

func main() {
	done := make(chan int, 0)
	js.Global().Set("sumByGo", js.FuncOf(jsWrapperSumByGo))
	<-done
}
