# SkyWalking Go APIs

This repository contains the Go files generated from
the [data collect protocol](http://github.com/apache/skywalking-data-collect-protocol) and
the [query protocol](http://github.com/apache/skywalking-query-protocol) for convenient use.

You can use the following commands to install this module.

```shell
go get skywalking.apache.org/repo/goapi
```

## Data Collect Protocol

To use the Go files generated from
the [data collect protocol](http://github.com/apache/skywalking-data-collect-protocol), use the import
path `skywalking.apache.org/repo/goapi/collect`, for example,

```go
package main

import (
	"fmt"

	"skywalking.apache.org/repo/goapi/collect/event/v3"
)

func main() {
	event := &v3.Event{
		Uuid:       "",
		Source:     nil,
		Name:       "",
		Type:       0,
		Message:    "",
		Parameters: nil,
		StartTime:  0,
		EndTime:    0,
	}

	fmt.Printf("+%v", event)
}
```

## Query Protocol

To use the Go files generated from the [query protocol](http://github.com/apache/skywalking-query-protocol), use the
import path `skywalking.apache.org/repo/goapi/query`, for example,

```go
package main

import (
	"fmt"

	"skywalking.apache.org/repo/goapi/query"
)

func main() {
	events := query.Events{
		Events: nil,
		Total:  0,
	}

	fmt.Printf("+%v", events)
}
```

## Development

To update this repo, update the commit sha in the `dependencies.sh` file, and run `make` to regenerate the Go files,
then commit and open a pull request.
