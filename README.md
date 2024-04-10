# go-vast

Go package to parse, manipulate and build Digital Video Ad Serving Templates (VAST).
The currently supported version is VAST 4.2.

[![Test coverage](https://img.shields.io/badge/coverage-87%25-success)](https://github.com/joeig/go-vast/tree/master/.github/testcoverage.yml)
[![Go Report Card](https://goreportcard.com/badge/go.eigsys.de/go-vast)](https://goreportcard.com/report/go.eigsys.de/go-vast)
[![PkgGoDev](https://pkg.go.dev/badge/go.eigsys.de/go-vast)](https://pkg.go.dev/go.eigsys.de/go-vast)

## Usage

```shell
go get -u go.eigsys.de/go-vast
```

### Important note on `CDATA` handling of `encoding/xml`

The `encoding/xml` package merges `CDATA` sections with surrounding text.
This can lead to unexpected results when parsing VAST files.

Consider the following example:

<!-- @formatter:off -->
```xml
          <JavaScriptResource>
            <![CDATA[https://verificationcompany1.com/verification_script1.js]]>
          </JavaScriptResource>
```
<!-- @formatter:on -->

After parsing `JavaScriptResource`, the value is:

<!-- @formatter:off -->
```text
\n            https://verificationcompany1.com/verification_script1.js\n          
```
<!-- @formatter:on -->

If you need the content without the Unicode spaces, consider using [`strings.TrimSpace()`](https://pkg.go.dev/strings#TrimSpace).
For more information, see [#43168](https://github.com/golang/go/issues/43168).

## Examples

A complete list of examples is available in the [package reference](https://pkg.go.dev/go.eigsys.de/go-vast).

### Create empty VAST

```go
package main

import (
	"go.eigsys.de/go-vast"
)

func main() {
	example := vast.New()
	example.Version = vast.VAST42Version
}
```

### Read VAST from file

```go
package main

import (
	"go.eigsys.de/go-vast"
	"log"
	"os"
)

func main() {
	handle, err := os.Open("vast.xml")
	if err != nil {
		log.Fatalf("%v", err)
	}
	
	example, err := vast.Read(handle)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
```

### Marshal VAST

```go
package main

import (
	"go.eigsys.de/go-vast"
	"log"
)

func main() {
	example := vast.New()

	exampleBytes, err := example.Bytes()
	if err != nil {
		log.Fatalf("%v", err)
	}
}
```
