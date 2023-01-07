# Go package

Go needs a `go.mod` file recognize that the directory contains a Go script/s.

## Table of content

[[toc]]

## Managing packages ### Creating a package To create a package type

```bash go mod init example/name

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

To import form local go package is as simple.
If it is a local module (that is not in a repository), we
need to link it to the current module, in order to use it by
Importing it

The project looks like this:

```sh
<root>/
|--greetings/
  --go.mod
  --greeting.go
|--example/
  --go.mod
  --example.go
```

To import the `example.com/greetings` module into the `example.com/example` we
run the following command in the example directory:

```sh
go mod edit -replace example.com/greetings=../greetings
```

Now it can be used in the example module like so:

```go
//...
import (
  ...

  "example.com/greetings"
)
//...
```
