use std::str;
use wasm_bindgen::prelude::*;

#[wasm_bindgen]
pub fn add_by_rust(a: i32, b: i32) -> i32 {
    a + b + 1000
}
