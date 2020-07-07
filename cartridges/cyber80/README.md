# Building cyber80

To build, run a command similar to:
```
GOOS=js GOARCH=wasm go build -o cyber80.wasm
```

The resulting `cyber80.wasm` file that is created is used by `index.html` in the parent directory.
