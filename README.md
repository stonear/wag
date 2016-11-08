# wag
sWAGger - Web API Generator

## Usage
Note that WAG requires Go 1.7.
### Generating Code
Create a swagger.yml file with your [service definition](http://editor.swagger.io/#/). Note that WAG supports a [subset](https://github.com/Clever/wag#swagger-spec) of the Swagger spec.
Copy the latest `wag.mk` from the [dev-handbook](https://github.com/Clever/dev-handbook/blob/master/make/wag.mk).
Set up a `generate` target in your `Makefile` that will generate server and client code:

```
include wag.mk

WAG_VERSION := 0.1.0

generate: wag-generate-deps
	$(call wag-generate,./swagger.yml, $(PKG))
```

Define global BadRequest and InternalError response types. These are used internally for validation errors and unknown errors respectively. They must reference a definition with a msg field. For example:

```
responses:
  BadRequest:
    description: Bad Request
    schema:
      $ref: "#/definitions/BadRequestError"
  InternalError:
    description: Internal Error
    schema:
      $ref: "#/definitions/InternalError"

definitions:
  BadRequestError:
    type: object
    properties:
      msg:
        type: string
  InternalError:
    type: object
    properties:
      msg:
        type: string
```

If your yml doesn't define a 400 or a 500 response for any of your operations, Wag will implicitly set it to the globally defined responses, models.BadRequest and models.InternalError respectively.

Also, all your 400+ responses must return a schema type with a Msg field. The msg field is required so that we can covert the response into Go's error type and use the Msg in `Error() string`. Each of the responses must also be unique so that the client and server can convert between status code and type.

Then generate your code:
```
make generate
```

This generates four directories. You should not have to modify any of the generated code:
- gen-go/models: contains all the definitions in your Swagger file as well as API input / output definitions
- gen-go/server: contains the router, middleware, and handler logic
- gen-go/client: contains the Go client library
- gen-js: contains the javascript client library

### Using the Go Server
To use the generated code you need to do two things:
- Implement the controller interface defined in `gen-go/server/interface.go`
- Pass the controller into the Server constructor. For example:
```
  s := server.New(myController, ":8000")
  // Serve should not return
  log.Fatal(s.Serve())
```

All interface methods on the controller take in a `context.Context` object.
This object comes with additional features for you to use in implementing your server logic:

* **Logging**. The [kayvee middleware logger](https://godoc.org/gopkg.in/Clever/kayvee-go.v5/middleware) is automatically added to the context object.
  It can be pulled out of the context object and used via the kayvee `FromContext` method:

```go
import "gopkg.in/Clever/kayvee-go.v5/logger"
...
logger.FromContext(ctx).Info(...)
```

  You should use this logger for all logging within your controller implementation.

* **Success Response Types**
  * Wag generates success types in three ways:
    * If there one success status code and it does have a data type associated with it then Wag generates an interface that takes a pointer to that data type as the first argument.
    * If there one success status code and it doesn't have a data type associated with it then Wag generates an interface with only an error response. If the err is nil then the request succeeded
    * If there is more than one success status code then Wag generates an interface named
    {{.OperationID}}Output which each success type, named {{.OperationID}}{.StatusCode}}Output
    implements.



* **Errors**.
  * Wag supports three types of errors
    * Global error response types
    * Response types for a specific operation
    * Undefined error types

  * Any of the three can be returned from a controller. To return a global or response specific error type use return a pointer to the model defintion for that error type. To return an undefined error type, return anything that implements the Error interface. Wag automatically converts undefined errors into the 500 error type.

  * Errors returned from your controller are logged by the
  autogenerated handler code, so there is no need to separately log errors
  yourself. If you use the `github.com/go-errors/errors` package, the
  stacktrace will also be logged, making debugging easier.

  For undefined error types the best practices are:
    * If you receive an error from an external dependency, use
      `errors.WrapPrefix(err, "foopackage.func", 0)` to return an error with a
     stacktrace and prefix the source of the error (in this example, we received
     an error from the function `func` in the package `foopackage`).
   * If you generate a new error, use `errors.Errorf` or `errors.New` to build
     the error.
   * If you receive an error from an internal function, just return the error
    directly since it should already have stacktrace information (either it is
      a wrapped external error or a `go-errors`-generated internal error).

* **Tracing**: `wag` instruments the context object with tracing-related metadata.
  This is done via [opentracing](http://opentracing.io/).
  In order for it to work, you are required to do two things:

  * Configure your `main()` function to report tracing data to [LightStep](http://lightstep.com/).
     We are testing Lightstep as a way to view tracing-related data:
```go
package main

import (
	lightstep "github.com/lightstep/lightstep-tracer-go"
	opentracing "github.com/opentracing/opentracing-go"
)

func main() {
	tags := make(map[string]interface{})
	tags[lightstep.ComponentNameKey] = "<name of the repo>"
	lightstepTracer := lightstep.NewTracer(lightstep.Options{
	    AccessToken: os.Getenv("LIGHTSTEP_ACCESS_TOKEN"),
	    Tags:        tags,
	})
	defer lightstep.FlushLightStepTracer(lightstepTracer)
	opentracing.InitGlobalTracer(lightstepTracer)
	...
}
```

  * Pass the context object to any subsequent network requests you make.
    Many client libraries accept the context object, e.g.:
    * **net/http**: If you're making HTTP requests, use the [golang.org/x/net/context/ctxhttp](https://godoc.org/golang.org/x/net/context/ctxhttp) package.
    * **wag**: All `wag`-generated clients take in a context object as the first argument.
      If your handler consumes a `wag`-generated client, then pass the context object to these client methods.

### Using the Go Client
Initialize the client with `New`
```
c := client.New("https://url_of_your_service:port")
```

Make an API call
```
books, err := c.GetBooks(ctx, GetBookByIDInput{Authors: []string{"Twain"}})
if err != nil {
  // Do something with the error
}
```

If you're using the client from another WAG-ified service you should pass in the `ctx` object you get in your server handler. Otherwise you can use `context.Background()`

### Using the Javascript Client
You can initialize the client by either passing a url or by using [discovery](https://github.com/Clever/discovery-node).

```javascript
import * as SampleClientLib from 'sample-client-lib-js';

const sampleClient = new SampleClientLib({address: "https://url_of_your_service:port"}); // Explicit url
// OR
const sampleClient = new SampleClientLib({discovery: true}); // Using discovery
```

You may also configure a global timeout for requests when initalizing the client.

```javascript
const sampleClient = new SampleClientLib({discovery: true, timeout: 1000}); // Timeout any requests taking longer than 1 second
```

You may then call methods on the client. Methods support callbacks and promises.

```javascript
// Promises
sampleClient.getBookById("bookID").then((book) => {
  // ...
}).catch((err) => {
  // ...
});

// Callbacks
sampleClient.getBookById("bookID", (err, book) => {
  // ...
});
```

You can also pass an optional options argument. This can have the following options
- `timeout` - overide the global timeout for this specific call
- `span` - Pass an opentracing span to instrument with the call - More on this below

```javascript
const options = {
  timeout: 5000 // Timeout after 5 seconds
}

sampleClient.getBookById("bookID", options, (err, book) => {
  // ...
});
```

#### Tracing

To utilize the `span` option above you need to pass an opentracing span into the request. The below
example shows you how to setup an express app to track requests and any calls made via a wag client.

```bash
npm install lightstep-tracer opentracing@0.11 --save # >=0.12 contains untested breaking changes to the API
```

```javascript
import * as SampleClientLib from 'sample-client-lib-js';
import * as express from 'express';
import * as Tracer from 'opentracing';
import * as LightStep from 'lightstep-tracer';
import * as lodash from 'lodash';
Tracer.initGlobalTracer(LightStep.tracer({
  access_token   : process.env.LIGHTSTEP_ACCESS_TOKEN,
  component_name : 'repo-name',
}));

const sampleClient = new SampleClientLib({discovery: true}); // Using discovery

const app = express();

// Middleware to look for a span from inbound requests
app.use((req, res, next) => {
  const wireCtx = Tracer.extract(Tracer.FORMAT_HTTP_HEADERS, req.headers);
  req.span = Tracer.startSpan(req.url, {childOf: wireCtx});
  req.span.logEvent("request_received");

  // include trace ID in headers so that we can debug slow requests we see in
  // the browser by looking up the trace ID found in response headers
  const responseHeaders = {};
  Tracer.inject(req.span, Tracer.FORMAT_TEXT_MAP, responseHeaders);
  lodash.forOwn(responseHeaders, (value, key) => res.setHeader(key, responseHeaders[key]));

  // finalize the span when the response is completed
  res.on("finish", () => {
    req.span.logEvent("request_finished");
    req.span.finish();
  });

  next();
});

app.get("/my-url", (req, res) => {
  sampleClient.getBookById("bookID", {span: req.span}, (err, book) => {
    // ...
  });
});

```

### Custom String Validation
We've added custom string validation for mongo-ids to avoid repeating: "^[0-9a-f]{24}$"` throughout the swagger.yml. To use it you have must:

- Change you swagger.yml file to have the `mongo-id` format. For example:
```
authorID:
        type: string
        format: mongo-id
```

- Import `github.com/Clever/wag/swagger` and call `swagger.InitCustomFormats()` in your server code.

Note that custom string validation only applies to input parameters and does not have any impact on objects defined in '#/definitions'.

Right now we do not allow user-defined custom strings, but this is something we may add if there's sufficient demand.

## Tests
```
make test
```

## Swagger Spec

Currently, WAG doesn't implement the entire Swagger Spec. A couple things to keep in mind:
- All schemas should reference type definitions in /definitions. Any schemas defined in /paths will cause an error.
- All parameters and definitions should be defined in their specific path object instead of globally.
- Scheme, produces, and consumers can only be defined in the top-level swagger object, not individual operations. On the top level object the scheme must be 'http', produces must be 'application/json' and consumes must be 'application/json'

Below is a more comprehensive list of the features we don't support along with some custom extensions.


### Unsupported Features
Mime Types

Multi-File Swagger Definitions

Schema:
- host
- tags
- scheme (must be http)
- consumes
- produces
- securityDefinitions
- security

Consumes:
- produces (must be application/json)
- consumes (must be application/json)
- schemes
- security

Form parameter type
Parameter:
- file parameter type
- collectionFormat
- global definitions
- possibly the json schema requirements? (uniqueItems, multipleOf, etc...)

Schema object (all these have to be defined in /definitions and are generated by go-swagger)

Discriminators

XML Modeling

Security Objects

Response:
  - Headers
