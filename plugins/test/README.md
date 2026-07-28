# Test plugin

This Go component implements:

```text
chzcake:metadata/metadata.get@1.0.0
```

Regenerate the WIT bindings after changing `../../api/metadata.wit`:

```sh
go tool componentize-go \
  -d ../../api/metadata.wit \
  -w plugin \
  bindings \
  --format
```

Build the component:

```sh
go tool componentize-go \
  -d ../../api/metadata.wit \
  -w plugin \
  build \
  -o test.wasm
```
