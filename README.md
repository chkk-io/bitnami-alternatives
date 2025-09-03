# Bitnami Alternatives

This source repository contains alternatives to Bitnami-produced Docker images
and Helm Charts.

## Usage

The `alternatives.json` file contains the set of known Bitnami Docker images
and associated Helm Charts for a large number of open source software projects.

Feel free to consume this file directly by simply fetching the raw file
contents. If you have a Go application, you can make use of the alternatives
data by simply importing the `github.com/chkk-io/bitnami-alternatives` package:

```go
package main

import (
    "fmt"

    bitnamialt "github.com/chkk-io/bitnami-alternatives"
)

func main() {
    for _, entry := range bitnamialt.Alternatives {
        fmt.Println("===============================================")
        fmt.Println("Project:", entry.Project)
        fmt.Println("===============================================")
        fmt.Println("     BITNAMI IMAGES for", entry.Project)
        fmt.Println("===============================================")
        for _, img := range entry.Bitnami.Images {
            fmt.Println(" -", img)
        }
        fmt.Println("===============================================")
        fmt.Println("     ALTERNATIVE IMAGES for", entry.Project)
        fmt.Println("===============================================")
        if len(entry.Alternatives.Images) == 0 {
            fmt.Println("No known alternative Docker images.")
        } else {
            for _, img := range entry.Alternatives.Images {
                fmt.Println(" -", img)
            }
        }
        fmt.Println("===============================================")
        fmt.Println("     BITNAMI CHARTS for", entry.Project)
        fmt.Println("===============================================")
        if len(entry.Bitnami.Charts) == 0 {
            fmt.Println("No Bitnami Helm charts.")
        } else {
            for _, chart := range entry.Bitnami.Charts {
                if chart.URL != "" {
                    fmt.Println(" -", chart.URL)
                } else {
                    fmt.Println(" -", chart.Registry, chart.Name)
                }
            }
        }
        fmt.Println("===============================================")
        fmt.Println("     ALTERNATIVE CHARTS for", entry.Project)
        fmt.Println("===============================================")
        if len(entry.Alternatives.Charts) == 0 {
            fmt.Println("No known alternative Helm charts.")
        } else {
            for _, chart := range entry.Alternatives.Charts {
                if chart.URL != "" {
                    fmt.Println(" -", chart.URL)
                } else {
                    fmt.Println(" -", chart.Registry, chart.Name)
                }
            }
        }
    }
}
```

## Contributions

[Contributions](CONTRIBUTING.md) to `bitnami-alternatives` are welcomed!

Do you know of an alternative Docker image or Helm Chart for a Bitnami-produced
artifact? Edit the `alternatives.json` file and add your alternative!
