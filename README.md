# watcher

Package watcher implements a simple file monitor that calls a callback function when a file is modified, typically used to automatically load configuration files.

## simple

```go
package main

import (
	"fmt"
	"log"

	"github.com/zs5460/jmc"
	"github.com/zs5460/watcher"
)

type config struct {
	Listen  string
	Version string
}

var cfg config

func main() {
	w, err := watcher.New("config.json", loadConfig)
	if err != nil {
		log.Fatal(err)
	}
	w.Start()

	<-make(chan struct{})
}

func loadConfig() {
	jmc.MustLoadConfig("config.json", &cfg)
	fmt.Printf("%#v\n", cfg)
}

```

## License

Released under MIT license, see [LICENSE](LICENSE) for details.