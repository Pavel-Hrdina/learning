# Go packages

Go needs a `go.mod` file recognize that the directory contains a Go script/s.

## Table of content

[[toc]]

## Managing packages

### Creating a package

To create a package type:

```bash
go mod init example/name
```

Go creates a file in the current directory named: `go.mod` witch contains:

```go
module example/name

go 1.19
```

### Creating a Go script

Creating a go script, is as easy as to create a file. Create a file any way you like example script creation using bash:

```bash
#!bin/bash

touch hello.go
```

Simple hello world script:

```go{8}
// filename: hello.go

package main

import "fmt"

func main() {
 fmt.Println("Hello, World!")
}
```

### Importing packages from another file
