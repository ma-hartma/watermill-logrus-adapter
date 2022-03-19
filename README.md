# Watermill Logrus Adapter
An adapter for using [watermill](https://github.com/ThreeDotsLabs/watermill) with logrus logger.

***

## Installation
```
go get github.com/ma-hartma/watermill-logrus-adapter
```

***

## Use LogrusLoggerAdapter

```go
package main

import (
    wla "github.com/ma-hartma/watermill-logrus-adapter"
	"github.com/sirupsen/logrus"
)

func main() {
    // create a logrus logger instance
    log := logrus.New()

    // pass the logrus logger instance to NewLogrusLogger
    // to get your LoggerAdapter
    var logger watermill.LoggerAdapter
    logger = wla.NewLogrusLogger(log)

    // you can now use the logger with watermill
    ...
}
```

***

## License

[MIT](https://github.com/ma-hartma/watermill-logrus-adapter/raw/main/LICENSE)
