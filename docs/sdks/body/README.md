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
            BigDecimal: optionalv1go.String("Distributed"),
            ChildClassArray: []shared.ChildClass{
                shared.ChildClass{},
            },
            GrandParentOptional: optionalv1go.String("transition Berkshire"),
            GrandParentRequired: "program methodology",
            GrandParentRequiredNullable: "knowledge linear support",
            Optional: optionalv1go.String("standardization Future Northwest"),
            OptionalNullable: optionalv1go.String("East Kia"),
            OptionalNullableWithDefaultValue: optionalv1go.String("system Bedfordshire"),
            ParentOptional: optionalv1go.String("rich"),
            ParentOptionalNullableWithDefaultValue: optionalv1go.String("Gasoline Electronic"),
            ParentRequired: "SSL Frozen Baby",
            ParentRequiredNullable: "Sterling East whiteboard",
            Required: "Engineer Dakota",
            RequiredNullable: "scalable AI Bacon",
            Class: optionalv1go.Int(462184),
            Precision: optionalv1go.Float64(9784.81),
        },
        Field: "lest East",
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

