# https://developer.mozilla.org/en-US/docs/WebAssembly/Rust_to_wasm

default:
	echo "Hello, world!"

clean:
	rm -rf pkg/*

all: clean build run-server

deps:
	rustup target add wasm32-unknown-emscripten
	rustup target add wasm32-unknown-unknown
	cargo install wasm-bindgen-cli
	cargo add wasm-bindgen web-sys
	cargo install wasm-pack

run-server:
	echo "visit http://127.0.0.1:9124 in your web browser" && cd web && python3 -m http.server 9124

build:
	wasm-pack build --target web --out-dir web/pkg

_build_old:
	cargo build --target wasm32-unknown-unknown --release
	wasm-bindgen target/wasm32-unknown-unknown/release/wasm_demo.wasm --out-dir ../static --no-typescript