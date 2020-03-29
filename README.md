# transmission

This repository it's a wrapper for [`transmission RPC API`](https://github.com/transmission/transmission/blob/master/extras/rpc-spec.txt).

## Available methods

- [x] torrent-start
- [x] torrent-start-now
- [x] torrent-stop
- [x] torrent-verify
- [x] torrent-reannounce
- [x] torrent-set
- [x] torrent-get
- [x] torrent-add
- [x] torrent-remove
- [x] torrent-set-location
- [x] torrent-rename-path
- [x] session-close
- [x] session-get
- [x] session-set
- [x] session-stats
- [x] queue-move-top
- [x] queue-move-up
- [x] queue-move-down
- [x] queue-move-bottom
- [x] free-space
- [x] port-test
- [x] blocklist-update

For more information read [spec file](https://github.com/transmission/transmission/blob/master/extras/rpc-spec.txt).

## Installation

```bash
go get -u github.com/mfuentesg/transmission
```

## Examples

> Check connectivity with transmission service

```go
package main

import (
    "context"
    "log"

    "github.com/mfuentesg/transmission"
)

func main() {
	client := transmission.New(
		transmission.WithURL("http://service-url.com/tranmission/rpc"),
		transmission.WithBasicAuth("username", "password"),
	)

    // this method is not part of the spec (just check connectivity)
    if err := client.Ping(context.Background()); err != nil {
        log.Fatalf("could not connect to transmission service: %+v", err)
    }

    log.Println("connected to the transmission service")
}
```

> Get list of torrents
```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/mfuentesg/transmission"
)

func main() {
	client := transmission.New(
		transmission.WithURL("http://service-url.com/tranmission/rpc"),
		transmission.WithBasicAuth("username", "password"),
	)

    // get list of torrents
    // For know more about the available fields, take a look to the below link
    // https://github.com/transmission/transmission/blob/20119f006ca0f3a13245b379c74254c92f372910/extras/rpc-spec.txt#L111
    torrents, err := client.TorrentGet(context.Background(), transmission.TorrentGet{
        Ids: 17, // search by id, you can use hashString as well
        Fields: []string{"id", "hashString"},
    })

    if err != nil {
        log.Fatalf("could not get torrent list: %+v", err)
    }

    for _, torrent := range torrents {
    	fmt.Printf("torrent with id %d and hashString %s\n", torrent.Id, torrent.HashString)
    }
}
```

## TODO

- [ ] Write benchmark tests
- [ ] Write unit tests
- [ ] Add CI/CD for linting and tests
- [ ] Improve README file
- [ ] Add documentation to each function and reference to transmission fields