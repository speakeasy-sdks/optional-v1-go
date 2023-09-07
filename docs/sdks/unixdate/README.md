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
        Field: "nam",
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
        Field: "officia",
        SetToNull: false,
        UnSet: false,
        Unixdatearray: shared.Unixdatearray{
            DateTime: []float64{
                5820.2,
            },
            DateTime1: []float64{
                1433.53,
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
        Field: "deleniti",
        SetToNull: false,
        UnSet: false,
        Unixdatemap: shared.Unixdatemap{
            DateTime: map[string]float64{
                "hic": 7586.16,
            },
            DateTime1: map[string]float64{
                "totam": 1059.07,
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

