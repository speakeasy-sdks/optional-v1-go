# SimpleDate
(*SimpleDate*)

### Available Operations

* [SendSimpleDate](#sendsimpledate) - sendSimpleDate
* [SendSimpleDateArray](#sendsimpledatearray) - sendSimpleDateArray

## SendSimpleDate

sendSimpleDate

### Example Usage

```go
package main

import(
	"context"
	"log"
	optionalv1go "github.com/speakeasy-sdks/optional-v1-go"
	"github.com/speakeasy-sdks/optional-v1-go/pkg/models/operations"
	"github.com/speakeasy-sdks/optional-v1-go/pkg/models/shared"
	"github.com/speakeasy-sdks/optional-v1-go/pkg/types"
)

func main() {
    s := optionalv1go.New()

    ctx := context.Background()
    res, err := s.SimpleDate.SendSimpleDate(ctx, operations.SendSimpleDateRequest{
        Field: "string",
        SetToNull: false,
        Simpledate: shared.Simpledate{
            Date: types.MustDateFromString("1994-02-13"),
            DateNullable: types.MustDateFromString("1994-02-13"),
        },
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.SendSimpleDateRequest](../../models/operations/sendsimpledaterequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.SendSimpleDateResponse](../../models/operations/sendsimpledateresponse.md), error**


## SendSimpleDateArray

sendSimpleDateArray

### Example Usage

```go
package main

import(
	"context"
	"log"
	optionalv1go "github.com/speakeasy-sdks/optional-v1-go"
	"github.com/speakeasy-sdks/optional-v1-go/pkg/models/operations"
	"github.com/speakeasy-sdks/optional-v1-go/pkg/models/shared"
	"github.com/speakeasy-sdks/optional-v1-go/pkg/types"
)

func main() {
    s := optionalv1go.New()

    ctx := context.Background()
    res, err := s.SimpleDate.SendSimpleDateArray(ctx, operations.SendSimpleDateArrayRequest{
        Field: "string",
        SetToNull: false,
        Simpledatearray: shared.Simpledatearray{
            Date: []types.Date{
                types.MustDateFromString("1994-02-13 00:00:00 +0000 UTC"),
                types.MustDateFromString("1994-02-14 00:00:00 +0000 UTC"),
            },
            Date1: []types.Date{
                types.MustDateFromString("1994-02-13 00:00:00 +0000 UTC"),
                types.MustDateFromString("1994-02-14 00:00:00 +0000 UTC"),
            },
        },
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.SendSimpleDateArrayRequest](../../models/operations/sendsimpledatearrayrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.SendSimpleDateArrayResponse](../../models/operations/sendsimpledatearrayresponse.md), error**

