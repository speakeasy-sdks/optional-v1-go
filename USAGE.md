<!-- Start SDK Example Usage -->


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
            BigDecimal: optionalandnullable.String("corrupti"),
            ChildClassArray: []shared.ChildClass{
                shared.ChildClass{},
                shared.ChildClass{},
                shared.ChildClass{},
            },
            GrandParentOptional: optionalandnullable.String("distinctio"),
            GrandParentRequired: "quibusdam",
            GrandParentRequiredNullable: "unde",
            Optional: optionalandnullable.String("nulla"),
            OptionalNullable: optionalandnullable.String("corrupti"),
            OptionalNullableWithDefaultValue: optionalandnullable.String("illum"),
            ParentOptional: optionalandnullable.String("vel"),
            ParentOptionalNullableWithDefaultValue: optionalandnullable.String("error"),
            ParentRequired: "deserunt",
            ParentRequiredNullable: "suscipit",
            Required: "iure",
            RequiredNullable: "magnam",
            Class: optionalandnullable.Int(891773),
            Precision: optionalandnullable.Float64(567.13),
        },
        Field: "delectus",
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
<!-- End SDK Example Usage -->