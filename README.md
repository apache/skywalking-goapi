# SkyWalking Go APIs

This repository contains the Go files generated from
the sniff protocol([data collect protocol](http://github.com/apache/skywalking-data-collect-protocol), the satellite protocol) and
the [query protocol](http://github.com/apache/skywalking-query-protocol) for convenient use.

You can use the following commands to install this module.

```shell
go get skywalking.apache.org/repo/goapi
```

## Data Sniff Protocol

To use the Go files generated from
the [data collect protocol](http://github.com/apache/skywalking-data-collect-protocol) and the satellite protocol, use the import
path `skywalking.apache.org/repo/goapi/collect` and `skywalking.apache.org/repo/goapi/satellite`, for example,

```go
package main

import (
	"fmt"
	"time"

	v3 "skywalking.apache.org/repo/goapi/collect/event/v3"
	v1 "skywalking.apache.org/repo/goapi/satellite/event/v1"
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

	sniffData := &v1.SniffData{
		Timestamp: time.Now().Unix() / 1e6,
		Name:      "Satellite_event",
		Type:      v1.SniffType_EventType,
		Meta:      nil,
		Remote:    true,
		Data: &v1.SniffData_Event{
			Event: event,
		},
	}
	fmt.Printf("+%v", sniffData)
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
