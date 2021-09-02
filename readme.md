# AXGOS

HTTP client with:

* Mocking capabilities
* Timeouts by default
* A single *http.Client (concurrency safe)
* Support for `json` and `msgpack` (and `xml`)
* The only external dependency is to handle `msgpack`

## QUICK-START

### GET request

```go
import (
    "fmt"
    "net/http"
    "gitlab.com/germanDV/axgos/axgos"
)

type BlogPost struct {
    ID     int    `json:"id,omitempty"`
    Title  string `json:"title"`
    Body   string `json:"body"`
    UserID int    `json:"userId"`
}

func main() {
    commonHeaders := make(http.Header)
    commonHeaders.Set("Content-Type", "application/json")
    commonHeaders.Set("Accept", "application/json")

    client := axgos.
        NewBuilder().
        SetBaseURL("https://jsonplaceholder.typicode.com").
        SetHeaders(commonHeaders).
        Build()

    res, err := client.Get("/posts/1")
    if err != nil {
      panic(err)
    }

    if !res.OK() {
      panic(fmt.Sprintf("API returned an error: %d %q\n", res.StatusCode, res.String()))
    }

    var bp BlogPost
    err = res.Unmarshal(&bp)
    if err != nil {
      panic(err)
    }

    fmt.Println(res.StatusCode)
    fmt.Printf("ID: %d\n", bp.ID)
    fmt.Printf("Title: %q\n", bp.Title)
    fmt.Printf("Body: %q\n", bp.Body)
    fmt.Printf("UserID: %d\n", bp.UserID)
}
```

### POST request

```go
import (
    "fmt"
    "net/http"
    "gitlab.com/germanDV/axgos/axgos"
)

type BlogPost struct {
    ID     int    `json:"id,omitempty"`
    Title  string `json:"title"`
    Body   string `json:"body"`
    UserID int    `json:"userId"`
}

func main() {
    commonHeaders := make(http.Header)
    commonHeaders.Set("Content-Type", "application/json")
    commonHeaders.Set("Accept", "application/json")

    client := axgos.
        NewBuilder().
        SetBaseURL("https://jsonplaceholder.typicode.com").
        SetHeaders(commonHeaders).
        Build()

    authHeader := make(http.Header)
    authHeader.Set("Authorization", "Bearer my-token-abc123")

    bp := BlogPost{
        Title:  "My new blog post",
        Body:   "lorem ipsum dolor sit amet.",
        UserID: 99,
    }

    res, err := client.Post("/posts", bp, authHeader)
    if err != nil {
      panic(err)
    }

    if !res.OK() {
      panic(fmt.Sprintf("API returned an error: %d %q\n", res.StatusCode, res.String()))
    }

    fmt.Println(res.StatusCode)
    fmt.Println(res.String())
}
```

## HEADERS

Headers can be set on a global level during the _client_ creation and on a per-request basis.

```go
globalHeaders := make(http.Header)
globalHeaders.Set("Content-Type", "application/json")
globalHeaders.Set("Accept", "application/json")
client := axgos.NewBuilder().SetHeaders(globalHeaders).Build()

requestHeaders := make(http.Header)
requestHeaders.Set("Authorization", "Bearer my-token-abc123")
client.Get("some-url", requestHeaders)
```

## CONTENT TYPE

By default, axgos works with JSON. However, you can control the encoding via the _Content-Type_
and _Accept_ headers.

* The _Content-Type_ header determines the format that the **request body** is going to be sent in.
* The _Accept_ header determines the format that the **response body** is going to be unmarshalled
  to when calling the `Unmarshal` method on the response.

In both cases, the acceptable options are:

* `application/json`
* `application/msgpack`
* `application/xml`

### Support for [msgpack](https://msgpack.org/)

If you wish to use **msgpack**, simply set the appropriate headers:

```go
headers := make(http.Header)
headers.Set("Content-Type", "application/msgpack")
headers.Set("Accept", "application/msgpack")
```

And add `msgpack` tags to your struct, building on the examples above:

```go
type BlogPost struct {
    ID     int    `json:"id,omitempty" msgpack:"id,omitempty"`
    Title  string `json:"title" msgpack:"title"`
    Body   string `json:"body" msgpack:"body"`
    UserID int    `json:"userId" msgpack:"userId"`
}
```

## TIMEOUT AND LIMITS

By default, the following timeouts and connection limits are applied:

```
defaultMaxConnectionsPerHost     = 5
defaultMaxIdleConnectionsPerHost = 5
defaultResponseTimeout           = 5 * time.Second
defaultConnectionTimeout         = 2 * time.Second
```

They can all be configured when building the client with these methods:

```go
SetConnectionTimeout(timeout time.Duration) Builder
SetResponseTimeout(timeout time.Duration) Builder
SetMaxConnectionsPerHost(connections int) Builder
SetMaxIdleConnectionsPerHost(connections int) Builder
```

## TESTING

Some examples are provided in `./examples/*_test.go`.

Basically, requests can be mocked to easily test without making actual HTTP calls.

1. The first step is to enable the `MockServer`:

```go
import (
    ...
    "gitlab.com/germanDV/axgos/mock"
)

mock.MockServer.Enable()
```

2. Create and add a `Mock`:

```go
// In this case, we want to mock an error due to a connection timeout
mock.MockServer.Add(mock.Mock{
    Method: http.MethodGet,
    Url:    "https://jsonplaceholder.typicode.com/posts/1",
    Error:  errors.New("connection timeout"),
})
```

3. Make the request and the assertions:

```go
res, err := Get("https://jsonplaceholder.typicode.com/posts/1")
if res != nil {
  t.Error("Expected no response")
}
if err == nil {
  t.Error("Expected an error")
}
if err.Error() != "connection timeout" {
  t.Errorf("Expected error 'connection timeout', got %q\n", err.Error())
}
```

Mocks are matched by the combination of _Method_, _Url_ and _Body_. So, when a request is made, as
long as `MockServer` is enabled, the _client_ will look for the presence of a `Mock` for that given
request, and return that.

You can also add a `Mock` with a specific status code and response body (instead of an error as above):

```go
// Here we expect the API to return an error
mock.MockServer.Add(mock.Mock{
    Method:     http.MethodGet,
    Url:        "https://jsonplaceholder.typicode.com/posts/9",
    StatusCode: http.StatusForbidden,
    ResBody:    `{"message":"forbidden resource"}`,
})

// Here we expect a successful response
mock.MockServer.Add(mock.Mock{
    Method:     http.MethodGet,
    Url:        "https://jsonplaceholder.typicode.com/posts/4",
    StatusCode: http.StatusOK,
    ResBody:    `{"id": 4, "userId": 1, "title": "eum et est occaecati", "body": "etc"}`,
})
```

## NEXT STEPS

Provide the ability to cancel requests.
