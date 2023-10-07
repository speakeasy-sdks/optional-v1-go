# Body
(*Body*)

### Available Operations

* [SendChild](#sendchild) - sendChild

## SendChild

sendChild

### Example Usage

```go
package main

import(
	"context"
	"log"
	optionalv1go "github.com/speakeasy-sdks/optional-v1-go"
	"github.com/speakeasy-sdks/optional-v1-go/pkg/models/operations"
	"github.com/speakeasy-sdks/optional-v1-go/pkg/models/shared"
)

func main() {
    s := optionalv1go.New()

    ctx := context.Background()
    res, err := s.Body.SendChild(ctx, operations.SendChildRequest{
        ChildClass: shared.ChildClass{
            ChildClassArray: []shared.ChildClass{
                shared.ChildClass{},
            },
            GrandParentRequired: "Distributed",
            GrandParentRequiredNullable: "transition Berkshire",
            ParentRequired: "program methodology",
            ParentRequiredNullable: "knowledge linear support",
            Required: "standardization Future Northwest",
            RequiredNullable: "East Kia",
        },
        Field: "system Bedfordshire",
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.SendChildRequest](../../models/operations/sendchildrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.SendChildResponse](../../models/operations/sendchildresponse.md), error**

