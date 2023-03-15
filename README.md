# Play WASM(WebAssembly)

This is a simple example of how to play with the WASM by using go/rust/walang, and compare the size of the wasm file by 
different languages.

# what they do

very simple: 

```javascript
function sum(a, b) {
    return a + b + 1000;
}
```

## go

```shell
$ cd by-golang/wasm-demo-a
$ make -s all;
```

visit <http://127.0.0.1:9123>

## rust

```shell
$ cd by-rust/wasm_demo;
$ make -s deps; 
$ make -s all;
```

visit <http://127.0.0.1:9124>

## walang

todo.

