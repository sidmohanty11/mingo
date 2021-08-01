# Mingo
![mingo-httpclient](https://user-images.githubusercontent.com/73601258/127781760-f4c5b525-024f-40dd-9b02-8941861b03b3.jpg)

![madeingo](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

*A minimal yet powerful HTTP client implemented in go (inspired by axios).*

### Usage
Install mingo,

    go get github.com/sidmohanty11/mingo
A simple instance of HTTP-Client,
```go
var  m  = mingo.MakeNewClient().Make() // a simple client initialization
```
or you can customize however you like,
```go
m := mingo.
// required fpr the initialization of the makeClient struct which contains all the features
MakeNewClient().
// if you want to disable timeouts, else the defaults will be set
DisableTimeouts(false).
// set the max idle connections per hour by mentioning an int, for example
SetMaxIdleConnections(5).
// set the connection timeout by mentioning a duration, for example
SetConnectionTimeout(2  * time.Second).
// set the response timeout by mentioning a duration, for example
.SetResponseTimeout(2 * time.Second)
// sets the custom headers which you provide, receives http.Header
.SetHeaders(headers)
// sets a custom header "User-Agent" as key and provided username as it's value
.SetUserAgent(username)
// this creates the client responsible for making requests (GET, POST, PUT, PATCH, DELETE, OPTIONS)
.Make()
```
A simple GET request,
```go
res, err  := m.Get("https://yoururl.com")
if err !=  nil {
return  nil, err
}
// res.StatusCode() => returns the status code as int
// res.Status() => returns the status as a string
// res.Headers() => returns all headers
// res.BodyAsBytes() => returns response body as bytes
// res.BodyAsString() => returns response body as string in json/xml as you mentioned in Content-Type header, the default is json
if res.StatusCode() != 200 {
return nil, errors.New("error")
}
var endp Endpoints
if err := res.UnmarshalJSON(&endp); err != nil {
return nil, err
}
return &endp, nil
```
For writing tests, there is an inbuilt mock test environment.
Writing test for a simple GET request,
```go
mingo.StartMockServer() // starts up a mock server
mingo.FlushMocks() // deletes all previous mock keys
// here you can specify what should the endpoint do in certain cases
mingo.AddMock(mingo.Mock{
Method: http.MethodGet,
Url: "https://yoururl.com",
Error: errors.New("timeout getting the endp"),
}) // will generate a fake error of timeout when called
endp, err := GetEndp() 
// and handle test cases accordingly
```
For more examples, there is a  examples directory which contains some reference code.