# sevenzip
A Golang wrapper for [7-Zip](https://www.7-zip.org/)

## Installation

```sh
go get -u github.com/akbarhabiby/sevenzip
```

You should have the a 7-Zip executable available in your system.

## Usage

```go
package main

import (
	"context"

	"github.com/akbarhabiby/sevenzip"
)

func main() {
	ctx := context.Background()
	sw := sevenzip.NewSwitches()

	cmd := sevenzip.ExtractFullContext(ctx, "./archive.7z", "./output/dir/", sw)

	// cmd.Stdout = outReader
	// cmd.Stderr = errReader

	err := cmd.Run()
	if err != nil {
		// ... do something with error
		return
	}
}

```

***
With :heart: from [akbarhabiby](https://akbar.hk)