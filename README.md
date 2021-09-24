**hog** is a library which provides a set of extensions on go's standard `http.Client` library,
so all interfaces are just a syntax sugar on the standard ones.

Major additional concepts are:

- *SOE*: Simple, Original, Extended

- Native http client types where is it possible, like: **http.Header**, **url.Values**

- Possibility to set original **http.Client** to painless integration of existing codebase: `h := hog.NewClient(client)`

- `hog.Get("https://httpbin.org/get")` and `hog.Post("https://httpbin.org/post")` to quick result

- Possibility to get original **http.Response**: `hog.Get("https://httpbin.org/get").Response()`

## Usage

### Simple
`Get` request as **string** result:
```go
package main

import (
	"fmt"
	"github.com/aaapi-net/hog"
)

func main() {
	result, err := hog.Get("https://httpbin.org/get").AsString()
	fmt.Println(result, err)
}
```

`Post` request with **json** `body` as **map** result:

```go
package main

import (
	"fmt"
	"net/url"
	"github.com/aaapi-net/hog"
)

func main() {
	result, err := hog.Post("https://httpbin.org/get").
		Form(url.Values{
			"name": {"alice"},
			"age":  {"16"},
		}).
		AsMap()

	fmt.Println(result, err)
}
```


### Original
`Get` request as **http.Response** result:
```go
package main

import (
	"fmt"
	"github.com/aaapi-net/hog"
)

func main() {
	originalResponse, err := hog.Get("https://httpbin.org/get").Response()
	fmt.Println(originalResponse.Status, err)
}
```

`Get` request as **[]byte** and **http.Response** result:
```go
package main

import (
	"fmt"
	"github.com/aaapi-net/hog"
)

func main() {
	bytes, originalResponse, err := hog.Get("https://httpbin.org/get").AsBytesResponse()
	fmt.Println(string(bytes), originalResponse.StatusCode, err)
}
```

### Extended

`Get` Request as **string** result:

```go
package main

import (
	"context"
	"fmt"
	"github.com/aaapi-net/hog"
)

func main() {
	h := hog.New()
    
	result, err := h.
                Context(context.Background()).
                Get("https://httpbin.org/get").
                SetValue("id", "777"). // https://httpbin.org/get?id=777
                SetHeader("go", "lang").
                AsString()
	
	fmt.Println(result, err)
}
```

`Post` Request With **Json** Body as **string** result:

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"github.com/aaapi-net/hog"
)

func main() {
	h := hog.New()

	result, err := h.
		Context(context.Background()).
		Post("https://httpbin.org/post").
		// set all query values in one function
		Query(url.Values{"firstName": []string{"Joe"}, "lastName": []string{"Doe"}}). // https://httpbin.org/post?firstName=Joe&lastName=Doe
		// multiple headers in one function
		Headers(http.Header{"go": {"lang", "is", "awesome"}}).
		// interface{} as json body: {"name": "alice", "age": 16}
		Json(map[string]interface{}{
			"name": "alice",
			"age":  16,
		}).AsString()

	fmt.Println(result, err)
}
```

> **hog** is a usable **go h**ttp.Client


### Alternatives

https://github.com/dghubble/sling

https://github.com/nahid/gohttp

https://github.com/cizixs/gohttp 

https://github.com/BRUHItsABunny/gOkHttp 

https://github.com/franela/goreq 

