UnGoImport
=========

Ridiculous I know!
But sometimes you want to convert that file formatted with `goimports` in to files formatted with `gofmt`.
This or just have `goimports` sort your imports out in the editor of your choosing.
Then revert it back to a file formatted with `gofmt`.

## Usage
Say you want to format all your go files, which have been formatted with `goimport`

```bash
    find . -name "*.go" | xargs -I{} sh -c ‘cat "{}" | ungoimport > "{}.repl"’ -- {}
```

This will generate you a file with the extension `.repl` for each of your go files. 
You can just `mv` all those files over the old ones if you’re happy.

e.g. 

```go
package main

import (
    "fmt"
    "log"

    "github.com/Some/Third/Party/tool"
)

func main() {
 //...
}
```

becomes:

```go
package main

import (
    "fmt"
    "github.com/Some/Third/Party/tool"
    "log"
)

func main() {
 //...
}
```


