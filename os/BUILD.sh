
#!/bin/sh

WASM_FILENAME=magic256.wasm

GOOS=js GOARCH=wasm go build -o "${WASM_FILENAME}" && \
mv "${WASM_FILENAME}" ../www/                      && \
cd ../srv/cyber80-srv/                             && \
./cyber80-srv ../../www/
