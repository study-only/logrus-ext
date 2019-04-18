# logrus-ext

## Usage

```go
var log = Get("usecase", &Option{WithFunc: true})

// Output:
//   time="2019-04-18T15:36:37+08:00" level=error msg="query task error: id=1" error="too many connection" 
//   func="github.com/liamylian/logrus-ext.(*db).QueryTask" name=usecase
err := errors.New("too many connection")
log.WithError(err).Errorf("query task error: id=%d", id)
```