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
            BigDecimal: optionalandnullable.String("tempora"),
            ChildClassArray: []shared.ChildClass{
                shared.ChildClass{},
                shared.ChildClass{},
            },
            GrandParentOptional: optionalandnullable.String("molestiae"),
            GrandParentRequired: "minus",
            GrandParentRequiredNullable: "placeat",
            Optional: optionalandnullable.String("voluptatum"),
            OptionalNullable: optionalandnullable.String("iusto"),
            OptionalNullableWithDefaultValue: optionalandnullable.String("excepturi"),
            ParentOptional: optionalandnullable.String("nisi"),
            ParentOptionalNullableWithDefaultValue: optionalandnullable.String("recusandae"),
            ParentRequired: "temporibus",
            ParentRequiredNullable: "ab",
            Required: "quis",
            RequiredNullable: "veritatis",
            Class: optionalandnullable.Int(648172),
            Precision: optionalandnullable.Float64(202.18),
        },
        Field: "ipsam",
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

