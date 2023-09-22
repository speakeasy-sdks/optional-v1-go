# Body

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
            BigDecimal: optionalv1go.String("perferendis"),
            ChildClassArray: []shared.ChildClass{
                shared.ChildClass{},
            },
            GrandParentOptional: optionalv1go.String("ipsam"),
            GrandParentRequired: "repellendus",
            GrandParentRequiredNullable: "sapiente",
            Optional: optionalv1go.String("quo"),
            OptionalNullable: optionalv1go.String("odit"),
            OptionalNullableWithDefaultValue: optionalv1go.String("at"),
            ParentOptional: optionalv1go.String("at"),
            ParentOptionalNullableWithDefaultValue: optionalv1go.String("maiores"),
            ParentRequired: "molestiae",
            ParentRequiredNullable: "quod",
            Required: "quod",
            RequiredNullable: "esse",
            Class: optionalv1go.Int(520478),
            Precision: optionalv1go.Float64(7805.29),
        },
        Field: "dolorum",
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

