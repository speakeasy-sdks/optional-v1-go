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
            },
            GrandParentOptional: optionalandnullable.String("provident"),
            GrandParentRequired: "distinctio",
            GrandParentRequiredNullable: "quibusdam",
            Optional: optionalandnullable.String("unde"),
            OptionalNullable: optionalandnullable.String("nulla"),
            OptionalNullableWithDefaultValue: optionalandnullable.String("corrupti"),
            ParentOptional: optionalandnullable.String("illum"),
            ParentOptionalNullableWithDefaultValue: optionalandnullable.String("vel"),
            ParentRequired: "error",
            ParentRequiredNullable: "deserunt",
            Required: "suscipit",
            RequiredNullable: "iure",
            Class: optionalandnullable.Int(297534),
            Precision: optionalandnullable.Float64(8917.73),
        },
        Field: "ipsa",
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