# Go Plugins

## Compiling the Plugins
The plugin package is compiled using the normal Go toolchain. The only requirement is to use the buildmode=plugin compilation flag as shown below:

```
go build -buildmode=plugin -o samples/uiHooks/*.so samples/uiHooks/*.go

```   
