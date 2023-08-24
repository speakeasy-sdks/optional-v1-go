# UnixDate

### Available Operations

* [SendUnixDate](#sendunixdate) - sendUnixDate
* [SendUnixDateArray](#sendunixdatearray) - sendUnixDateArray
* [SendUnixDateMap](#sendunixdatemap) - sendUnixDateMap

## SendUnixDate

sendUnixDate

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/optional-v1-go"
	"github.com/speakeasy-sdks/optional-v1-go/pkg/models/operations"
	"github.com/speakeasy-sdks/optional-v1-go/pkg/models/shared"
)

func main() {
    s := optionalandnullable.New()

    ctx := context.Background()
    res, err := s.UnixDate.SendUnixDate(ctx, operations.SendUnixDateRequest{
        Field: "natus",
        SetToNull: false,
        UnSet: false,
        Unixdate: shared.Unixdate{
            DateTime: optionalandnullable.Float64(1484719381),
            DateTime1: 1484719381,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.SendUnixDateRequest](../../models/operations/sendunixdaterequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.SendUnixDateResponse](../../models/operations/sendunixdateresponse.md), error**


## SendUnixDateArray

sendUnixDateArray

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/optional-v1-go"
	"github.com/speakeasy-sdks/optional-v1-go/pkg/models/operations"
	"github.com/speakeasy-sdks/optional-v1-go/pkg/models/shared"
)

func main() {
    s := optionalandnullable.New()

    ctx := context.Background()
    res, err := s.UnixDate.SendUnixDateArray(ctx, operations.SendUnixDateArrayRequest{
        Field: "sed",
        SetToNull: false,
        UnSet: false,
        Unixdatearray: shared.Unixdatearray{
            DateTime: []float64{
                2223.21,
                6169.34,
                3864.89,
            },
            DateTime1: []float64{
                9025.99,
                6818.2,
                4499.5,
                3595.08,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.SendUnixDateArrayRequest](../../models/operations/sendunixdatearrayrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.SendUnixDateArrayResponse](../../models/operations/sendunixdatearrayresponse.md), error**


## SendUnixDateMap

sendUnixDateMap

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/optional-v1-go"
	"github.com/speakeasy-sdks/optional-v1-go/pkg/models/operations"
	"github.com/speakeasy-sdks/optional-v1-go/pkg/models/shared"
)

func main() {
    s := optionalandnullable.New()

    ctx := context.Background()
    res, err := s.UnixDate.SendUnixDateMap(ctx, operations.SendUnixDateMapRequest{
        Field: "iste",
        SetToNull: false,
        UnSet: false,
        Unixdatemap: shared.Unixdatemap{
            DateTime: map[string]float64{
                "saepe": 6976.31,
                "architecto": 602.25,
            },
            DateTime1: map[string]float64{
                "est": 6531.4,
                "laborum": 1709.09,
                "dolorem": 3581.52,
                "explicabo": 7506.86,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.SendUnixDateMapRequest](../../models/operations/sendunixdatemaprequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.SendUnixDateMapResponse](../../models/operations/sendunixdatemapresponse.md), error**

