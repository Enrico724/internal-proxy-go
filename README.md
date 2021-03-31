# internal-proxy-go

A tool which will let you redirect all traffic of your webserver to the service.

### usage

Import all needed modules 
```
import (
    "fmt"
    "github.com/Enrico724/internal-proxy-go"
    "net/http"
)
```
Create your handlers
```
func helloWorld(w http.ResponseWriter, _ *http.Request)  {
    _, _ = w.Write([]byte("Hello World!"))
}
```
Write the main
```
func main() {
    http.HandleFunc("/", helloWorld)
    err := internal_proxy.Serve()
    if err != nil {
        fmt.Printf("%v\n",err)
    }
}
```
And run your code
```
go run example/main.go
```

You can find this example inside example directory
