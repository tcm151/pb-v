# pb-v
### Serve a Vue web app (or any other static files) with pocketbase

1. Compile your app into a single directory
2. Copy those files into a folder called static (can be whatever)
3. Build the pocketbase go file below, navigate to 'https://localhost:YOUR_PORT/' and *voila*!

```go
package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	pb := pocketbase.New()

	pb.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Static("/", "static")
		return nil
	})

	if err := pb.Start(); err != nil {
		log.Fatal(err)
	}

}
```
