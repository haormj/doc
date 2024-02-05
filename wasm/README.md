## wasm

### 概述

WebAssembly主要是为了解决在浏览器中执行密集计算任务使用JavaScript低效的问题，从而就新除了一个编译target，可以将其他语言编写的内容编译为该类型，编译后的内容可以在runtime中运行，当然编译后的字节码是可以跨平台的。可以类比java字节码和jre

### 介绍

- runtime
  - https://github.com/bytecodealliance/wasmtime
  - https://github.com/WasmEdge/WasmEdge
- can build to wasm
  - golang https://github.com/tinygo-org/tinygo
  - rust
