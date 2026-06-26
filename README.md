# memoriessong-com-go

Go metadata and integration helper module for [memoriessong.com](https://memoriessong.com).

This module gives Go scripts, content automation jobs, indexing utilities, and future SDK code a stable import path for Memories Song. It exposes the canonical website URL, documentation URL, source repository, local repository path, MDX content path, and Next.js app path used by the project.

## Website focus

Memories Song focuses on AI keepsake music, memory-song creation, and emotional tribute song workflows. This first release intentionally publishes a lightweight metadata layer that can grow into a fuller API client later.

## Installation

```bash
go get github.com/youram470-art/memoriessong-com-go@v0.1.0
```

## Quick start

```go
package main

import (
    "fmt"

    site "github.com/youram470-art/memoriessong-com-go"
)

func main() {
    fmt.Println(site.Hello())
    fmt.Println(site.Homepage)
    fmt.Printf("%+v\n", site.GetSiteInfo())
}
```

## Links

- Website: https://memoriessong.com
- Documentation: https://memoriessong.com/docs
- Source: https://github.com/youram470-art/memoriessong-com-go
- pkg.go.dev: https://pkg.go.dev/github.com/youram470-art/memoriessong-com-go

## License

MIT
