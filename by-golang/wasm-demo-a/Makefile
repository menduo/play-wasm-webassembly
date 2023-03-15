default:
	echo "Hello, world!"

all: clean build-wasm copy-wasm-exec-js run-server

clean:
	rm -rf wasms/main-go.wasm

build-wasm:
	GO111MODULE=auto GOOS=js GOARCH=wasm go build -o wasms/main-go.wasm

copy-wasm-exec-js:
	cp "$(GOROOT)/misc/wasm/wasm_exec.js" static

run-server:
	echo "visit http://127.0.0.1:9123 in your web browser" &&  python3 -m http.server 9123