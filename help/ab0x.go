// Code generated by fileb0x at "2018-06-29 12:22:44.813082778 +0300 MSK m=+0.018981118" from config file "b0x.toml" DO NOT EDIT.
// modification hash(b2d1e8f0d19675f1cc5c6a021073d7d3.eadbe36098f6442f9c76baa8d818e914)

package help


import (
  "bytes"
  
  "io"
  "net/http"
  "os"
  "path"


  "golang.org/x/net/webdav"
  "context"


)

var ( 
  // CTX is a context for webdav vfs
  CTX = context.Background()

  
  // FS is a virtual memory file system
  FS = webdav.NewMemFS()
  

  // Handler is used to server files through a http handler
  Handler *webdav.Handler

  // HTTP is the http file system
  HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct {}



// FileCreateConfigmapMd is "create/configmap.md"
var FileCreateConfigmapMd = []byte("\x43\x72\x65\x61\x74\x65\x20\x63\x6f\x6e\x66\x69\x67\x6d\x61\x70\x2e\x0a\x43\x6f\x6e\x66\x69\x67\x6d\x61\x70\x20\x69\x73\x20\x61\x20\x66\x69\x6c\x65\x20\x73\x74\x6f\x72\x61\x67\x65\x2c\x20\x77\x68\x69\x63\x68\x20\x63\x61\x6e\x20\x62\x65\x20\x6d\x6f\x75\x6e\x74\x65\x64\x20\x69\x6e\x74\x6f\x20\x61\x20\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x2e\x20\x54\x68\x65\x20\x6d\x6f\x73\x74\x20\x63\x6f\x6d\x6d\x6f\x6e\x20\x75\x73\x61\x67\x65\x20\x6f\x66\x20\x63\x6f\x6e\x66\x69\x67\x6d\x61\x70\x20\x69\x73\x20\x6b\x65\x65\x70\x69\x6e\x67\x20\x63\x6f\x6e\x66\x69\x67\x20\x66\x69\x6c\x65\x73\x2c\x20\x72\x65\x61\x64\x2d\x6f\x6e\x6c\x79\x20\x44\x42\x2c\x20\x61\x6e\x64\x20\x73\x65\x63\x72\x65\x74\x73\x2e\x20\x42\x61\x73\x69\x63\x61\x6c\x6c\x79\x2c\x20\x79\x6f\x75\x20\x63\x61\x6e\x20\x74\x68\x69\x6e\x6b\x20\x61\x62\x6f\x75\x74\x20\x69\x74\x20\x6c\x69\x6b\x65\x20\x61\x62\x6f\x75\x74\x20\x76\x65\x72\x79\x20\x73\x69\x6d\x70\x6c\x65\x20\x6b\x65\x79\x2d\x76\x61\x6c\x75\x65\x20\x73\x74\x6f\x72\x61\x67\x65\x2e\x0a\x0a\x54\x68\x65\x72\x65\x20\x61\x72\x65\x20\x73\x65\x76\x65\x72\x61\x6c\x20\x77\x61\x79\x73\x20\x74\x6f\x20\x63\x6f\x6e\x73\x74\x72\x75\x63\x74\x20\x63\x6f\x6e\x66\x69\x67\x6d\x61\x70\x3a\x0a\x2d\x20\x2d\x2d\x69\x74\x65\x6d\x2d\x73\x74\x72\x69\x6e\x67\x20\x66\x6c\x61\x67\x2c\x20\x66\x6f\x72\x6d\x61\x74\x74\x65\x64\x20\x61\x73\x20\x4b\x45\x59\x3a\x56\x41\x4c\x55\x45\x20\x70\x61\x69\x72\x73\x2e\x20\x54\x68\x65\x20\x56\x41\x4c\x55\x45\x20\x63\x61\x6e\x20\x62\x65\x20\x74\x6f\x6b\x65\x6e\x2c\x20\x73\x68\x6f\x72\x74\x20\x69\x6e\x69\x74\x20\x66\x69\x6c\x65\x2c\x20\x65\x74\x63\x2e\x0a\x2d\x20\x2d\x2d\x69\x74\x65\x6d\x2d\x66\x69\x6c\x65\x20\x66\x6c\x61\x67\x2c\x20\x4b\x45\x59\x3a\x46\x49\x4c\x45\x5f\x50\x41\x54\x48\x20\x6f\x72\x20\x46\x49\x4c\x45\x5f\x50\x41\x54\x48\x20\x28\x66\x69\x6c\x65\x6e\x61\x6d\x65\x20\x77\x69\x6c\x6c\x20\x62\x65\x20\x75\x73\x65\x64\x20\x61\x73\x20\x4b\x45\x59\x29\x0a\x2d\x20\x69\x6e\x74\x65\x72\x61\x63\x74\x69\x76\x65\x20\x77\x69\x7a\x61\x72\x64\x0a\x0a\x55\x73\x65\x20\x74\x68\x65\x20\x2d\x2d\x66\x6f\x72\x63\x65\x20\x66\x6c\x61\x67\x20\x74\x6f\x20\x73\x6b\x69\x70\x20\x77\x69\x7a\x61\x72\x64")

// FileCreateContainerMd is "create/container.md"
var FileCreateContainerMd = []byte("\x41\x64\x64\x20\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x20\x74\x6f\x20\x64\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x20\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x20\x73\x65\x74\x2e\x20\x41\x76\x61\x69\x6c\x61\x62\x6c\x65\x20\x6d\x65\x74\x68\x6f\x64\x73\x20\x74\x6f\x20\x62\x75\x69\x6c\x64\x20\x64\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x3a\x0a\x20\x20\x20\x20\x2d\x20\x66\x72\x6f\x6d\x20\x66\x6c\x61\x67\x73\x0a\x20\x20\x20\x20\x2d\x20\x77\x69\x74\x68\x20\x69\x6e\x74\x65\x72\x61\x63\x74\x69\x76\x65\x20\x63\x6f\x6d\x6d\x61\x6e\x64\x6c\x69\x6e\x65\x20\x77\x69\x7a\x61\x72\x64\x0a\x20\x20\x20\x20\x2d\x20\x66\x72\x6f\x6d\x20\x79\x61\x6d\x6c\x20\x6f\x74\x20\x6a\x73\x6f\x6e\x20\x66\x69\x6c\x65\x0a\x0a\x55\x73\x65\x20\x2d\x2d\x66\x6f\x72\x63\x65\x20\x66\x6c\x61\x67\x20\x74\x6f\x20\x63\x72\x65\x61\x74\x65\x20\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x20\x77\x69\x74\x68\x6f\x75\x74\x20\x69\x6e\x74\x65\x72\x61\x63\x74\x69\x76\x65\x20\x77\x69\x7a\x61\x72\x64\x2e\x0a\x49\x66\x20\x74\x68\x65\x20\x2d\x2d\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x2d\x6e\x61\x6d\x65\x20\x66\x6c\x61\x67\x20\x69\x73\x20\x6e\x6f\x74\x20\x73\x70\x65\x63\x69\x66\x69\x65\x64\x20\x74\x68\x65\x6e\x20\x77\x69\x7a\x61\x72\x64\x20\x67\x65\x6e\x65\x72\x61\x74\x65\x73\x20\x6e\x61\x6d\x65\x20\x52\x41\x4e\x44\x4f\x4d\x5f\x43\x4f\x4c\x4f\x52\x2d\x49\x4d\x41\x47\x45\x2e")

// FileCreateDeploymentMd is "create/deployment.md"
var FileCreateDeploymentMd = []byte("\x43\x72\x65\x61\x74\x65\x20\x64\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x20\x77\x69\x74\x68\x20\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x73\x20\x61\x6e\x64\x20\x72\x65\x70\x6c\x69\x63\x61\x73\x2e\x0a\x41\x76\x61\x69\x6c\x61\x62\x6c\x65\x20\x6d\x65\x74\x68\x6f\x64\x73\x20\x74\x6f\x20\x62\x75\x69\x6c\x64\x20\x64\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x3a\x0a\x2d\x20\x66\x72\x6f\x6d\x20\x66\x6c\x61\x67\x73\x0a\x2d\x20\x77\x69\x74\x68\x20\x69\x6e\x74\x65\x72\x61\x63\x74\x69\x76\x65\x20\x63\x6f\x6d\x6d\x61\x6e\x64\x6c\x69\x6e\x65\x20\x77\x69\x7a\x61\x72\x64\x0a\x2d\x20\x66\x72\x6f\x6d\x20\x79\x61\x6d\x6c\x20\x6f\x74\x20\x6a\x73\x6f\x6e\x20\x66\x69\x6c\x65\x0a\x0a\x55\x73\x65\x20\x2d\x2d\x66\x6f\x72\x63\x65\x20\x66\x6c\x61\x67\x20\x74\x6f\x20\x63\x72\x65\x61\x74\x65\x20\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x20\x77\x69\x74\x68\x6f\x75\x74\x20\x69\x6e\x74\x65\x72\x61\x63\x74\x69\x76\x65\x20\x77\x69\x7a\x61\x72\x64\x2e\x0a\x0a\x54\x68\x65\x72\x65\x20\x61\x72\x65\x20\x73\x65\x76\x65\x72\x61\x6c\x20\x77\x61\x79\x73\x20\x74\x6f\x20\x73\x70\x65\x63\x69\x66\x79\x20\x74\x68\x65\x20\x6e\x61\x6d\x65\x73\x20\x6f\x66\x20\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x73\x20\x77\x69\x74\x68\x20\x66\x6c\x61\x67\x73\x3a\x0a\x2d\x20\x2d\x2d\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x2d\x6e\x61\x6d\x65\x20\x66\x6c\x61\x67\x0a\x2d\x20\x74\x68\x65\x20\x70\x72\x65\x66\x69\x78\x20\x43\x4f\x4e\x54\x41\x49\x4e\x45\x52\x5f\x4e\x41\x4d\x45\x40\x20\x69\x6e\x20\x74\x68\x65\x20\x66\x6c\x61\x67\x73\x20\x2d\x2d\x69\x6d\x61\x67\x65\x2c\x20\x2d\x2d\x6d\x65\x6d\x6f\x72\x79\x2c\x20\x2d\x2d\x63\x70\x75\x2c\x20\x2d\x2d\x65\x6e\x76\x2c\x20\x2d\x2d\x76\x6f\x6c\x75\x6d\x65\x0a\x0a\x49\x66\x20\x74\x68\x65\x20\x2d\x2d\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x2d\x6e\x61\x6d\x65\x20\x66\x6c\x61\x67\x20\x69\x73\x20\x6e\x6f\x74\x20\x73\x70\x65\x63\x69\x66\x69\x65\x64\x20\x61\x6e\x64\x20\x70\x72\x65\x66\x69\x78\x20\x69\x73\x20\x6e\x6f\x74\x20\x75\x73\x65\x64\x20\x69\x6e\x20\x61\x6e\x79\x20\x6f\x66\x20\x74\x68\x65\x20\x66\x6c\x61\x67\x73\x2c\x20\x74\x68\x65\x6e\x20\x77\x69\x7a\x61\x72\x64\x20\x73\x65\x61\x72\x63\x68\x65\x73\x20\x66\x6f\x72\x20\x74\x68\x65\x20\x2d\x2d\x69\x6d\x61\x67\x65\x20\x66\x6c\x61\x67\x73\x20\x77\x69\x74\x68\x6f\x75\x74\x20\x61\x20\x70\x72\x65\x66\x69\x78\x20\x61\x6e\x64\x20\x67\x65\x6e\x65\x72\x61\x74\x65\x73\x20\x6e\x61\x6d\x65\x20\x52\x41\x4e\x44\x4f\x4d\x5f\x43\x4f\x4c\x4f\x52\x2d\x49\x4d\x41\x47\x45\x2e\x0a\x0a\x2a\x2a\x45\x78\x61\x6d\x70\x6c\x65\x73\x3a\x2a\x2a\x0a\x0a\x2d\x2d\x2d\x0a\x2a\x2a\x53\x69\x6e\x67\x6c\x65\x20\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x20\x77\x69\x74\x68\x20\x2d\x2d\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x2d\x6e\x61\x6d\x65\x2a\x2a\x0a\x0a\x60\x60\x60\x62\x61\x73\x68\x0a\x3e\x20\x2e\x2f\x63\x6b\x69\x74\x20\x63\x72\x65\x61\x74\x65\x20\x64\x65\x70\x6c\x20\x5c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2d\x2d\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x2d\x6e\x61\x6d\x65\x20\x64\x6f\x6f\x74\x20\x5c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2d\x2d\x69\x6d\x61\x67\x65\x20\x6e\x67\x69\x6e\x78\x0a\x60\x60\x60\x0a\x0a\x7c\x20\x20\x20\x20\x20\x20\x20\x20\x4c\x41\x42\x45\x4c\x20\x20\x20\x20\x20\x20\x20\x20\x7c\x20\x56\x45\x52\x53\x49\x4f\x4e\x20\x7c\x20\x20\x53\x54\x41\x54\x55\x53\x20\x20\x7c\x20\x20\x43\x4f\x4e\x54\x41\x49\x4e\x45\x52\x53\x20\x20\x7c\x20\x20\x20\x20\x41\x47\x45\x20\x20\x20\x20\x7c\x0a\x7c\x20\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x20\x7c\x20\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x7c\x20\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x20\x7c\x20\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x20\x7c\x20\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x20\x7c\x0a\x7c\x20\x61\x6b\x69\x72\x61\x61\x62\x65\x2d\x68\x65\x69\x73\x65\x6e\x62\x65\x72\x67\x20\x7c\x20\x20\x31\x2e\x30\x2e\x30\x20\x20\x7c\x20\x69\x6e\x61\x63\x74\x69\x76\x65\x20\x7c\x20\x64\x6f\x6f\x74\x20\x5b\x6e\x67\x69\x6e\x78\x5d\x20\x7c\x20\x75\x6e\x64\x65\x66\x69\x6e\x65\x64\x20\x7c\x0a\x0a\x2d\x2d\x2d\x0a\x2a\x2a\x53\x69\x6e\x67\x6c\x65\x20\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x20\x77\x69\x74\x68\x6f\x75\x74\x20\x2d\x2d\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x2d\x6e\x61\x6d\x65\x2a\x2a\x0a\x0a\x60\x60\x60\x62\x61\x73\x68\x0a\x3e\x20\x2e\x2f\x63\x6b\x69\x74\x20\x63\x72\x65\x61\x74\x65\x20\x64\x65\x70\x6c\x20\x5c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2d\x2d\x69\x6d\x61\x67\x65\x20\x6e\x67\x69\x6e\x78\x0a\x60\x60\x60\x0a\x0a\x7c\x20\x20\x20\x20\x20\x20\x20\x20\x4c\x41\x42\x45\x4c\x20\x20\x20\x20\x20\x20\x20\x20\x7c\x20\x56\x45\x52\x53\x49\x4f\x4e\x20\x7c\x20\x20\x53\x54\x41\x54\x55\x53\x20\x20\x7c\x20\x20\x20\x20\x20\x20\x20\x20\x43\x4f\x4e\x54\x41\x49\x4e\x45\x52\x53\x20\x20\x20\x20\x20\x20\x20\x20\x7c\x20\x20\x20\x20\x41\x47\x45\x20\x20\x20\x20\x7c\x0a\x7c\x20\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x20\x7c\x20\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x7c\x20\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x20\x7c\x20\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x20\x7c\x20\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x20\x7c\x0a\x7c\x20\x20\x20\x73\x70\x69\x72\x61\x65\x61\x2d\x6b\x61\x75\x66\x6d\x61\x6e\x20\x20\x20\x7c\x20\x20\x31\x2e\x30\x2e\x30\x20\x20\x7c\x20\x69\x6e\x61\x63\x74\x69\x76\x65\x20\x7c\x20\x61\x71\x75\x61\x6d\x61\x72\x69\x6e\x65\x2d\x6e\x67\x69\x6e\x78\x20\x5b\x6e\x67\x69\x6e\x78\x5d\x20\x7c\x20\x75\x6e\x64\x65\x66\x69\x6e\x65\x64\x20\x7c\x0a\x0a\x2d\x2d\x2d\x0a\x2a\x2a\x4d\x75\x6c\x74\x69\x70\x6c\x65\x20\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x73\x20\x77\x69\x74\x68\x20\x2d\x2d\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x2d\x6e\x61\x6d\x65\x2a\x2a\x0a\x0a\x0a\x60\x60\x60\x62\x61\x73\x68\x0a\x3e\x20\x2e\x2f\x63\x6b\x69\x74\x20\x63\x72\x65\x61\x74\x65\x20\x64\x65\x70\x6c\x20\x5c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2d\x2d\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x2d\x6e\x61\x6d\x65\x20\x67\x61\x74\x65\x77\x61\x79\x20\x5c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2d\x2d\x69\x6d\x61\x67\x65\x20\x6e\x67\x69\x6e\x78\x20\x5c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2d\x2d\x69\x6d\x61\x67\x65\x20\x62\x6c\x6f\x67\x40\x77\x6f\x72\x64\x70\x72\x65\x73\x73\x0a\x60\x60\x60\x0a\x0a\x7c\x20\x20\x20\x20\x20\x20\x20\x20\x4c\x41\x42\x45\x4c\x20\x20\x20\x20\x20\x20\x20\x20\x7c\x20\x56\x45\x52\x53\x49\x4f\x4e\x20\x7c\x20\x20\x53\x54\x41\x54\x55\x53\x20\x20\x7c\x20\x20\x20\x20\x20\x20\x20\x20\x43\x4f\x4e\x54\x41\x49\x4e\x45\x52\x53\x20\x20\x20\x20\x20\x20\x20\x20\x7c\x20\x20\x20\x20\x41\x47\x45\x20\x20\x20\x20\x7c\x0a\x7c\x20\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x20\x7c\x20\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x7c\x20\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x20\x7c\x20\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x20\x7c\x20\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x20\x7c\x0a\x7c\x20\x20\x20\x72\x75\x63\x6b\x65\x72\x73\x2d\x66\x69\x73\x63\x68\x65\x72\x20\x20\x20\x7c\x20\x20\x31\x2e\x30\x2e\x30\x20\x20\x7c\x20\x69\x6e\x61\x63\x74\x69\x76\x65\x20\x7c\x20\x20\x20\x20\x20\x20\x67\x61\x74\x65\x77\x61\x79\x20\x5b\x6e\x67\x69\x6e\x78\x5d\x20\x20\x20\x20\x20\x7c\x20\x75\x6e\x64\x65\x66\x69\x6e\x65\x64\x20\x7c\x0a\x7c\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7c\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7c\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7c\x20\x20\x20\x20\x20\x20\x62\x6c\x6f\x67\x20\x5b\x77\x6f\x72\x64\x70\x72\x65\x73\x73\x5d\x20\x20\x20\x20\x7c\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7c\x0a\x0a\x2d\x2d\x2d\x0a\x2a\x2a\x4d\x75\x6c\x74\x69\x70\x6c\x65\x20\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x73\x20\x77\x69\x74\x68\x6f\x75\x74\x20\x2d\x2d\x63\x6f\x6e\x74\x61\x69\x6e\x65\x72\x2d\x6e\x61\x6d\x65\x2a\x2a\x0a\x60\x60\x60\x62\x61\x73\x68\x0a\x3e\x20\x2e\x2f\x63\x6b\x69\x74\x20\x63\x72\x65\x61\x74\x65\x20\x64\x65\x70\x6c\x20\x5c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2d\x2d\x69\x6d\x61\x67\x65\x20\x6e\x67\x69\x6e\x78\x20\x5c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2d\x2d\x69\x6d\x61\x67\x65\x20\x62\x6c\x6f\x67\x40\x77\x6f\x72\x64\x70\x72\x65\x73\x73\x0a\x60\x60\x60\x0a\x0a\x7c\x20\x20\x20\x20\x20\x20\x20\x20\x4c\x41\x42\x45\x4c\x20\x20\x20\x20\x20\x20\x20\x20\x7c\x20\x56\x45\x52\x53\x49\x4f\x4e\x20\x7c\x20\x20\x53\x54\x41\x54\x55\x53\x20\x20\x7c\x20\x20\x20\x20\x20\x20\x20\x20\x43\x4f\x4e\x54\x41\x49\x4e\x45\x52\x53\x20\x20\x20\x20\x20\x20\x20\x20\x7c\x20\x20\x20\x20\x41\x47\x45\x20\x20\x20\x20\x7c\x0a\x7c\x20\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x20\x7c\x20\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x20\x7c\x20\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x20\x7c\x20\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x20\x7c\x20\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x20\x7c\x0a\x7c\x20\x20\x20\x20\x74\x68\x69\x73\x62\x65\x2d\x6e\x65\x75\x6d\x61\x6e\x6e\x20\x20\x20\x7c\x20\x20\x31\x2e\x30\x2e\x30\x20\x20\x7c\x20\x69\x6e\x61\x63\x74\x69\x76\x65\x20\x7c\x20\x20\x20\x20\x20\x20\x62\x6c\x6f\x67\x20\x5b\x77\x6f\x72\x64\x70\x72\x65\x73\x73\x5d\x20\x20\x20\x20\x7c\x20\x75\x6e\x64\x65\x66\x69\x6e\x65\x64\x20\x7c\x0a\x7c\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7c\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7c\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7c\x20\x20\x20\x20\x67\x61\x72\x6e\x65\x74\x2d\x6e\x67\x69\x6e\x78\x20\x5b\x6e\x67\x69\x6e\x78\x5d\x20\x20\x7c\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7c\x0a")

// FileDeleteDeploymentMd is "delete/deployment.md"
var FileDeleteDeploymentMd = []byte("\x44\x65\x6c\x65\x74\x65\x20\x64\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x2e\x20\x4c\x69\x73\x74\x20\x6f\x66\x20\x64\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x73\x20\x74\x6f\x20\x64\x65\x6c\x65\x74\x65\x20\x6d\x75\x73\x74\x20\x62\x65\x20\x70\x72\x6f\x76\x69\x64\x65\x64\x20\x61\x73\x20\x61\x72\x67\x75\x6d\x65\x6e\x74\x2e\x20\x49\x66\x20\x6c\x69\x73\x74\x20\x69\x73\x20\x65\x6d\x70\x74\x79\x2c\x20\x74\x68\x65\x6e\x20\x63\x68\x6b\x69\x74\x20\x77\x69\x6c\x6c\x20\x73\x74\x61\x72\x74\x20\x69\x6e\x74\x65\x72\x61\x63\x74\x69\x76\x65\x20\x6d\x65\x6e\x75\x2e\x0a\x0a\x2a\x2a\x44\x65\x6c\x65\x74\x65\x20\x6c\x69\x73\x74\x20\x6f\x66\x20\x64\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x73\x20\x77\x69\x74\x68\x6f\x75\x74\x20\x2d\x2d\x66\x6f\x72\x63\x65\x2a\x2a\x0a\x0a\x60\x60\x60\x62\x61\x73\x68\x0a\x3e\x20\x63\x68\x6b\x69\x74\x20\x64\x65\x6c\x65\x74\x65\x20\x64\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x20\x6d\x69\x6d\x6f\x73\x61\x2d\x77\x61\x72\x62\x75\x72\x67\x20\x61\x74\x68\x61\x6d\x61\x6e\x74\x69\x73\x2d\x67\x61\x75\x73\x73\x20\x66\x6c\x6f\x72\x61\x2d\x6f\x6e\x6e\x65\x73\x0a\x41\x72\x65\x20\x79\x6f\x75\x20\x72\x65\x61\x6c\x6c\x79\x20\x77\x61\x6e\x74\x20\x74\x6f\x20\x64\x65\x6c\x65\x74\x65\x20\x74\x6f\x20\x64\x65\x6c\x65\x74\x65\x20\x6d\x69\x6d\x6f\x73\x61\x2d\x77\x61\x72\x62\x75\x72\x67\x2c\x20\x61\x74\x68\x61\x6d\x61\x6e\x74\x69\x73\x2d\x67\x61\x75\x73\x73\x2c\x20\x66\x6c\x6f\x72\x61\x2d\x6f\x6e\x6e\x65\x73\x3f\x20\x5b\x59\x2f\x4e\x5d\x3a\x20\x79\x0a\x44\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x20\x66\x6c\x6f\x72\x61\x2d\x6f\x6e\x6e\x65\x73\x20\x69\x73\x20\x64\x65\x6c\x65\x74\x65\x64\x0a\x44\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x20\x6d\x69\x6d\x6f\x73\x61\x2d\x77\x61\x72\x62\x75\x72\x67\x20\x69\x73\x20\x64\x65\x6c\x65\x74\x65\x64\x0a\x44\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x20\x61\x74\x68\x61\x6d\x61\x6e\x74\x69\x73\x2d\x67\x61\x75\x73\x73\x20\x69\x73\x20\x64\x65\x6c\x65\x74\x65\x64\x0a\x33\x20\x64\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x73\x20\x61\x72\x65\x20\x64\x65\x6c\x65\x74\x65\x64\x0a\x0a\x60\x60\x60\x0a\x0a\x2a\x2a\x44\x65\x6c\x65\x74\x65\x20\x6c\x69\x73\x74\x20\x6f\x66\x20\x64\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x73\x20\x77\x69\x74\x68\x20\x2d\x2d\x66\x6f\x72\x63\x65\x2a\x2a\x0a\x0a\x60\x60\x60\x62\x61\x73\x68\x0a\x3e\x20\x63\x68\x6b\x69\x74\x20\x64\x65\x6c\x65\x74\x65\x20\x64\x65\x70\x6c\x20\x2d\x2d\x66\x6f\x72\x63\x65\x20\x6c\x69\x6e\x64\x65\x6e\x61\x75\x2d\x63\x68\x61\x79\x65\x73\x20\x6d\x61\x6c\x76\x61\x2d\x63\x6c\x61\x72\x6b\x65\x20\x70\x61\x75\x77\x65\x6c\x73\x2d\x74\x6f\x65\x70\x6c\x65\x72\x0a\x44\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x20\x6c\x69\x6e\x64\x65\x6e\x61\x75\x2d\x63\x68\x61\x79\x65\x73\x20\x69\x73\x20\x64\x65\x6c\x65\x74\x65\x64\x0a\x44\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x20\x6d\x61\x6c\x76\x61\x2d\x63\x6c\x61\x72\x6b\x65\x20\x69\x73\x20\x64\x65\x6c\x65\x74\x65\x64\x0a\x44\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x20\x70\x61\x75\x77\x65\x6c\x73\x2d\x74\x6f\x65\x70\x6c\x65\x72\x20\x69\x73\x20\x64\x65\x6c\x65\x74\x65\x64\x0a\x33\x20\x64\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x73\x20\x61\x72\x65\x20\x64\x65\x6c\x65\x74\x65\x64\x0a\x60\x60\x60\x0a\x0a\x2a\x2a\x44\x65\x6c\x65\x74\x65\x20\x64\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x20\x77\x69\x74\x68\x20\x69\x6e\x74\x65\x72\x61\x63\x74\x69\x76\x65\x20\x73\x65\x6c\x65\x63\x74\x69\x6f\x6e\x2a\x2a\x0a\x0a\x60\x60\x60\x62\x61\x73\x68\x0a\x3e\x20\x63\x68\x6b\x69\x74\x20\x64\x65\x6c\x65\x74\x65\x20\x64\x65\x70\x6c\x0a\x53\x65\x6c\x65\x63\x74\x20\x64\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x3a\x0a\x53\x65\x6c\x65\x63\x74\x65\x64\x3a\x0a\x20\x31\x29\x20\x67\x75\x72\x7a\x68\x69\x6a\x2d\x6e\x65\x77\x74\x6f\x6e\x0a\x20\x32\x29\x20\x6a\x61\x63\x6b\x73\x6f\x6e\x2d\x6c\x65\x6e\x61\x72\x64\x0a\x20\x33\x29\x20\x6b\x75\x70\x65\x2d\x6d\x61\x67\x6e\x75\x73\x0a\x20\x34\x29\x20\x6c\x69\x6e\x64\x65\x6e\x61\x75\x2d\x63\x68\x61\x79\x65\x73\x0a\x20\x35\x29\x20\x6d\x61\x6c\x76\x61\x2d\x63\x6c\x61\x72\x6b\x65\x0a\x20\x36\x29\x20\x6d\x61\x72\x63\x68\x69\x73\x2d\x79\x6f\x75\x6e\x67\x0a\x20\x37\x29\x20\x70\x61\x75\x77\x65\x6c\x73\x2d\x74\x6f\x65\x70\x6c\x65\x72\x0a\x20\x38\x29\x20\x72\x65\x62\x65\x6e\x74\x72\x6f\x73\x74\x2d\x74\x68\x61\x6c\x65\x73\x0a\x20\x39\x29\x20\x43\x6f\x6e\x66\x69\x72\x6d\x0a\x43\x68\x6f\x6f\x73\x65\x20\x77\x69\x73\x65\x6c\x79\x3a\x20\x31\x0a\x53\x65\x6c\x65\x63\x74\x20\x64\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x3a\x0a\x53\x65\x6c\x65\x63\x74\x65\x64\x3a\x20\x67\x75\x72\x7a\x68\x69\x6a\x2d\x6e\x65\x77\x74\x6f\x6e\x0a\x20\x31\x29\x20\x6a\x61\x63\x6b\x73\x6f\x6e\x2d\x6c\x65\x6e\x61\x72\x64\x0a\x20\x32\x29\x20\x6b\x75\x70\x65\x2d\x6d\x61\x67\x6e\x75\x73\x0a\x20\x33\x29\x20\x6c\x69\x6e\x64\x65\x6e\x61\x75\x2d\x63\x68\x61\x79\x65\x73\x0a\x20\x34\x29\x20\x6d\x61\x6c\x76\x61\x2d\x63\x6c\x61\x72\x6b\x65\x0a\x20\x35\x29\x20\x6d\x61\x72\x63\x68\x69\x73\x2d\x79\x6f\x75\x6e\x67\x0a\x20\x36\x29\x20\x70\x61\x75\x77\x65\x6c\x73\x2d\x74\x6f\x65\x70\x6c\x65\x72\x0a\x20\x37\x29\x20\x72\x65\x62\x65\x6e\x74\x72\x6f\x73\x74\x2d\x74\x68\x61\x6c\x65\x73\x0a\x20\x38\x29\x20\x43\x6f\x6e\x66\x69\x72\x6d\x0a\x43\x68\x6f\x6f\x73\x65\x20\x77\x69\x73\x65\x6c\x79\x3a\x20\x32\x0a\x53\x65\x6c\x65\x63\x74\x20\x64\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x3a\x0a\x53\x65\x6c\x65\x63\x74\x65\x64\x3a\x20\x67\x75\x72\x7a\x68\x69\x6a\x2d\x6e\x65\x77\x74\x6f\x6e\x20\x6b\x75\x70\x65\x2d\x6d\x61\x67\x6e\x75\x73\x0a\x20\x31\x29\x20\x6a\x61\x63\x6b\x73\x6f\x6e\x2d\x6c\x65\x6e\x61\x72\x64\x0a\x20\x32\x29\x20\x6c\x69\x6e\x64\x65\x6e\x61\x75\x2d\x63\x68\x61\x79\x65\x73\x0a\x20\x33\x29\x20\x6d\x61\x6c\x76\x61\x2d\x63\x6c\x61\x72\x6b\x65\x0a\x20\x34\x29\x20\x6d\x61\x72\x63\x68\x69\x73\x2d\x79\x6f\x75\x6e\x67\x0a\x20\x35\x29\x20\x70\x61\x75\x77\x65\x6c\x73\x2d\x74\x6f\x65\x70\x6c\x65\x72\x0a\x20\x36\x29\x20\x72\x65\x62\x65\x6e\x74\x72\x6f\x73\x74\x2d\x74\x68\x61\x6c\x65\x73\x0a\x20\x37\x29\x20\x43\x6f\x6e\x66\x69\x72\x6d\x0a\x43\x68\x6f\x6f\x73\x65\x20\x77\x69\x73\x65\x6c\x79\x3a\x20\x34\x0a\x53\x65\x6c\x65\x63\x74\x20\x64\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x3a\x0a\x53\x65\x6c\x65\x63\x74\x65\x64\x3a\x20\x67\x75\x72\x7a\x68\x69\x6a\x2d\x6e\x65\x77\x74\x6f\x6e\x20\x6b\x75\x70\x65\x2d\x6d\x61\x67\x6e\x75\x73\x20\x6d\x61\x72\x63\x68\x69\x73\x2d\x79\x6f\x75\x6e\x67\x0a\x20\x31\x29\x20\x6a\x61\x63\x6b\x73\x6f\x6e\x2d\x6c\x65\x6e\x61\x72\x64\x0a\x20\x32\x29\x20\x6c\x69\x6e\x64\x65\x6e\x61\x75\x2d\x63\x68\x61\x79\x65\x73\x0a\x20\x33\x29\x20\x6d\x61\x6c\x76\x61\x2d\x63\x6c\x61\x72\x6b\x65\x0a\x20\x34\x29\x20\x70\x61\x75\x77\x65\x6c\x73\x2d\x74\x6f\x65\x70\x6c\x65\x72\x0a\x20\x35\x29\x20\x72\x65\x62\x65\x6e\x74\x72\x6f\x73\x74\x2d\x74\x68\x61\x6c\x65\x73\x0a\x20\x36\x29\x20\x43\x6f\x6e\x66\x69\x72\x6d\x0a\x43\x68\x6f\x6f\x73\x65\x20\x77\x69\x73\x65\x6c\x79\x3a\x20\x31\x0a\x53\x65\x6c\x65\x63\x74\x20\x64\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x3a\x0a\x53\x65\x6c\x65\x63\x74\x65\x64\x3a\x20\x67\x75\x72\x7a\x68\x69\x6a\x2d\x6e\x65\x77\x74\x6f\x6e\x20\x6b\x75\x70\x65\x2d\x6d\x61\x67\x6e\x75\x73\x20\x6d\x61\x72\x63\x68\x69\x73\x2d\x79\x6f\x75\x6e\x67\x20\x6a\x61\x63\x6b\x73\x6f\x6e\x2d\x6c\x65\x6e\x61\x72\x64\x0a\x20\x31\x29\x20\x6c\x69\x6e\x64\x65\x6e\x61\x75\x2d\x63\x68\x61\x79\x65\x73\x0a\x20\x32\x29\x20\x6d\x61\x6c\x76\x61\x2d\x63\x6c\x61\x72\x6b\x65\x0a\x20\x33\x29\x20\x70\x61\x75\x77\x65\x6c\x73\x2d\x74\x6f\x65\x70\x6c\x65\x72\x0a\x20\x34\x29\x20\x72\x65\x62\x65\x6e\x74\x72\x6f\x73\x74\x2d\x74\x68\x61\x6c\x65\x73\x0a\x20\x35\x29\x20\x43\x6f\x6e\x66\x69\x72\x6d\x0a\x43\x68\x6f\x6f\x73\x65\x20\x77\x69\x73\x65\x6c\x79\x3a\x20\x35\x0a\x41\x72\x65\x20\x79\x6f\x75\x20\x72\x65\x61\x6c\x6c\x79\x20\x77\x61\x6e\x74\x20\x74\x6f\x20\x64\x65\x6c\x65\x74\x65\x20\x74\x6f\x20\x64\x65\x6c\x65\x74\x65\x20\x67\x75\x72\x7a\x68\x69\x6a\x2d\x6e\x65\x77\x74\x6f\x6e\x2c\x20\x6b\x75\x70\x65\x2d\x6d\x61\x67\x6e\x75\x73\x2c\x20\x6d\x61\x72\x63\x68\x69\x73\x2d\x79\x6f\x75\x6e\x67\x2c\x20\x6a\x61\x63\x6b\x73\x6f\x6e\x2d\x6c\x65\x6e\x61\x72\x64\x3f\x20\x5b\x59\x2f\x4e\x5d\x3a\x20\x79\x0a\x44\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x20\x67\x75\x72\x7a\x68\x69\x6a\x2d\x6e\x65\x77\x74\x6f\x6e\x20\x69\x73\x20\x64\x65\x6c\x65\x74\x65\x64\x0a\x44\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x20\x6b\x75\x70\x65\x2d\x6d\x61\x67\x6e\x75\x73\x20\x69\x73\x20\x64\x65\x6c\x65\x74\x65\x64\x0a\x44\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x20\x6d\x61\x72\x63\x68\x69\x73\x2d\x79\x6f\x75\x6e\x67\x20\x69\x73\x20\x64\x65\x6c\x65\x74\x65\x64\x0a\x44\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x20\x6a\x61\x63\x6b\x73\x6f\x6e\x2d\x6c\x65\x6e\x61\x72\x64\x20\x69\x73\x20\x64\x65\x6c\x65\x74\x65\x64\x0a\x34\x20\x64\x65\x70\x6c\x6f\x79\x6d\x65\x6e\x74\x73\x20\x61\x72\x65\x20\x64\x65\x6c\x65\x74\x65\x64\x0a\x0a\x60\x60\x60")

// FileUpdateMd is "update.md"
var FileUpdateMd = []byte("\x55\x70\x64\x61\x74\x65\x20\x63\x68\x6b\x69\x74\x20\x66\x72\x6f\x6d\x20\x67\x69\x74\x68\x75\x62\x20\x6f\x72\x20\x6c\x6f\x63\x61\x6c\x20\x66\x69\x6c\x65\x73\x79\x73\x74\x65\x6d\x2e\x0a\x0a\x4a\x75\x73\x74\x20\x72\x75\x6e\x0a\x0a\x60\x60\x60\x62\x61\x73\x68\x0a\x3e\x20\x63\x68\x6b\x69\x74\x20\x75\x70\x64\x61\x74\x65\x0a\x60\x60\x60")



func init() {
  if CTX.Err() != nil {
		panic(CTX.Err())
	}






  var err error



  
  err = FS.Mkdir(CTX, "delete/", 0777)
  if err != nil && err != os.ErrExist {
    panic(err)
  }
  

  
  err = FS.Mkdir(CTX, "create/", 0777)
  if err != nil && err != os.ErrExist {
    panic(err)
  }
  




  
  var f webdav.File
  

  

  
  

  f, err = FS.OpenFile(CTX, "create/configmap.md", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    panic(err)
  }

  
  _, err = f.Write(FileCreateConfigmapMd)
  if err != nil {
    panic(err)
  }
  

  err = f.Close()
  if err != nil {
    panic(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "create/container.md", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    panic(err)
  }

  
  _, err = f.Write(FileCreateContainerMd)
  if err != nil {
    panic(err)
  }
  

  err = f.Close()
  if err != nil {
    panic(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "create/deployment.md", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    panic(err)
  }

  
  _, err = f.Write(FileCreateDeploymentMd)
  if err != nil {
    panic(err)
  }
  

  err = f.Close()
  if err != nil {
    panic(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "delete/deployment.md", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    panic(err)
  }

  
  _, err = f.Write(FileDeleteDeploymentMd)
  if err != nil {
    panic(err)
  }
  

  err = f.Close()
  if err != nil {
    panic(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "update.md", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    panic(err)
  }

  
  _, err = f.Write(FileUpdateMd)
  if err != nil {
    panic(err)
  }
  

  err = f.Close()
  if err != nil {
    panic(err)
  }
  


  Handler = &webdav.Handler{
    FileSystem: FS,
    LockSystem: webdav.NewMemLS(),
  }


}



// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {


  f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
  if err != nil {
    return nil, err
  }

  return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
  f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
  if err != nil {
    return nil, err
  }

  buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

  // If the buffer overflows, we will get bytes.ErrTooLarge.
  // Return that as an error. Any other panic remains.
  defer func() {
    e := recover()
    if e == nil {
      return
    }
    if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
      err = panicErr
    } else {
      panic(e)
    }
  }()
  _, err = buf.ReadFrom(f)
  return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
  f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
  if err != nil {
    return err
  }
  n, err := f.Write(data)
  if err == nil && n < len(data) {
    err = io.ErrShortWrite
  }
  if err1 := f.Close(); err == nil {
    err = err1
  }
  return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
    return nil, err
  }
  
  err = f.Close()
  if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}


