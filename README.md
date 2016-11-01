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

### Extensions
You can mark a patch request parameter with the flag `x-wag-patch`. This is a convenience method that
automatically generates a new data type based on the schema of the patch with all its fields optionals.
This allows you to re-use your data types while also getting nice, backwards-compatible patch methods.

For example:
```
  operationId: wagPatch
  parameters:
    - name: testPatch
      x-wag-patch: true
      in: body
      schema:
        $ref: "#/definitions/Data"
```
```
  Data:
    type: object
    properties:
      name:
        type: string
```

Generates a new data type, PatchData, and changes the interface of wagPatch
```
type PatchData struct {
  Name *string `json:"name"`
}

func WagPatch(ctx context.Context, i *models.DataPatch) error
```


### Unsupported Features
Mime Types

AdditionalProperties: these are supported by the Go server and client, but not the Javascript client

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
