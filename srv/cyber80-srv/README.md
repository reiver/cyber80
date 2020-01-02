# cyber80-srv

**cyber80-srv** is a tool used by the developers of **cyber80**, to make their development work easier.

**NOTE***: _cyber80-srv_ is NOT _cyber80_.

## Usage

Run:
```
./cyber80-srv
```

And then, in a web browser on the same computer, go to:
* http://localhost:8080/

## Explanation

**cyber80-srv** is a _web server_ — an HTTP server — that serves:
* http://localhost:8080/
* http://localhost:8080/wasm_exec.js
* http://localhost:8080/cyber80/cyber80.wasm

A human would just go to:
* http://localhost:8080/

And `wasm_exec.js` and `cyber80/cyber80.wasm` automatically be downloaded and pulled-in behind-the-scenes.
