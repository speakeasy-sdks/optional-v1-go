# RfcDate

### Available Operations

* [SendRfc1123Date](#sendrfc1123date) - sendRfc1123Date
* [SendRfc1123DateArray](#sendrfc1123datearray) - sendRfc1123DateArray
* [SendRfc1123DateMap](#sendrfc1123datemap) - sendRfc1123DateMap

## SendRfc1123Date

sendRfc1123Date

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
    res, err := s.RfcDate.SendRfc1123Date(ctx, operations.SendRfc1123DateRequest{
        Rfc1123date: shared.Rfc1123date{
            DateTime: optionalandnullable.String("repellendus"),
            DateTime1: "sapiente",
        },
        Field: "quo",
        SetToNull: false,
        UnSet: false,
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
| `request`                                                                              | [operations.SendRfc1123DateRequest](../../models/operations/sendrfc1123daterequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.SendRfc1123DateResponse](../../models/operations/sendrfc1123dateresponse.md), error**


## SendRfc1123DateArray

sendRfc1123DateArray

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
    res, err := s.RfcDate.SendRfc1123DateArray(ctx, operations.SendRfc1123DateArrayRequest{
        Rfc1123datearray: shared.Rfc1123datearray{
            DateTime: []string{
                "at",
            },
            DateTime1: []string{
                "maiores",
                "molestiae",
                "quod",
                "quod",
            },
        },
        Field: "esse",
        SetToNull: false,
        UnSet: false,
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.SendRfc1123DateArrayRequest](../../models/operations/sendrfc1123datearrayrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.SendRfc1123DateArrayResponse](../../models/operations/sendrfc1123datearrayresponse.md), error**


## SendRfc1123DateMap

sendRfc1123DateMap

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
    res, err := s.RfcDate.SendRfc1123DateMap(ctx, operations.SendRfc1123DateMapRequest{
        Rfc1123datemap: shared.Rfc1123datemap{
            DateTime: map[string]string{
                "porro": "dolorum",
                "dicta": "nam",
                "officia": "occaecati",
            },
            DateTime1: map[string]string{
                "deleniti": "hic",
            },
        },
        Field: "optio",
        SetToNull: false,
        UnSet: false,
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.SendRfc1123DateMapRequest](../../models/operations/sendrfc1123datemaprequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.SendRfc1123DateMapResponse](../../models/operations/sendrfc1123datemapresponse.md), error**

