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
	"github.com/speakeasy-sdks/optional-v1-go"
	"github.com/speakeasy-sdks/optional-v1-go/pkg/models/operations"
	"github.com/speakeasy-sdks/optional-v1-go/pkg/models/shared"
)

func main() {
    s := optionalandnullable.New()

    ctx := context.Background()
    res, err := s.Body.SendChild(ctx, operations.SendChildRequest{
        ChildClass: shared.ChildClass{
            BigDecimal: optionalandnullable.String("delectus"),
            ChildClassArray: []shared.ChildClass{
                shared.ChildClass{},
            },
            GrandParentOptional: optionalandnullable.String("tempora"),
            GrandParentRequired: "suscipit",
            GrandParentRequiredNullable: "molestiae",
            Optional: optionalandnullable.String("minus"),
            OptionalNullable: optionalandnullable.String("placeat"),
            OptionalNullableWithDefaultValue: optionalandnullable.String("voluptatum"),
            ParentOptional: optionalandnullable.String("iusto"),
            ParentOptionalNullableWithDefaultValue: optionalandnullable.String("excepturi"),
            ParentRequired: "nisi",
            ParentRequiredNullable: "recusandae",
            Required: "temporibus",
            RequiredNullable: "ab",
            Class: optionalandnullable.Int(337396),
            Precision: optionalandnullable.Float64(871.29),
        },
        Field: "deserunt",
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

