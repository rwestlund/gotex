# gotex
A simple Go library for rendering LaTeX documents

# Install
```bash
go get -u github.com/rwestlund/gotex
```

# Documentation
See the documentation at [https://godoc.org/github.com/rwestlund/gotex](https://godoc.org/github.com/rwestlund/gotex)

# Example
```go
package main

import "github.com/rwestlund/gotex"

func main() {
    var document = `
        \documentclass[12pt]{article}
        \begin{document}
        This is a LaTeX document.
        \end{document}
        `
    var pdf, err = gotex.Render(document,
        gotex.Options{Command: "/usr/bin/pdflatex", Runs: 1})

    if err != nil {
        log.Println("render failed ", err)
    } else {
        // Do something with the PDF file, like send it to an HTTP client
        // or write it to a file.
        sendSomewhere(pdf)
    }
}

```

# License
This code is under the BSD-2-Clause license.
