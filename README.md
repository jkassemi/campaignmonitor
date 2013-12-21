# Campaign Monitor for Go

A go interface to the Campaign Monitor API. This provides a very minimal set of
features at this point in time - but if everyone who wanted to use Campaign
Monitor with their Go application added just _one_ method of their own, this
thing would have... well... at least _one_ more method. 

A few BTC donations would likely get this flushed out a bit more, as well
(https://twitter.com/jkassemi/statuses/397205845898833920).

1N2ALH42LDpqoPvMEpkQNNYgYbFzdd9T3w

## Installation

```bash
go get github.com/jkassemi/campaignmonitor
```

## Usage

Currently limited to API key authentication:

```go
  import (
    cm "github.com/jkassemi/campaignmonitor"
  )

  client := cm.NewApiClient(os.Getenv("CM_API_KEY"))

  // Now use any methods exposed on client

  lists, e := client.GetClientSubscriberLists("abcd")

  if e != nil {
    panic("We had an error! " + e.Error())
  }

  log.Printf("We've got %d lists for that client", len(lists))
```

## Contributions

Please fork, add methods, and issue a pull request!
