# Go Logger

A simple logging package for Go.

## Usage

```go
package main

import (
    "github.com/tiranosaur/go-logger"
)

func main() {
	array := []string{"sldf", "fffffff"}
	Logger.WARN("warning {} {} {}", array)
	Logger.INFO("info {} {} {}", "InfMessage", 23, true)
	Logger.ERROR("error {} {} {}", "ERRmessage", 23, true)
}

