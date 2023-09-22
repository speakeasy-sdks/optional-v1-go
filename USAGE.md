<!-- Start SDK Example Usage -->


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
            BigDecimal: optionalv1go.String("corrupti"),
            ChildClassArray: []shared.ChildClass{
                shared.ChildClass{},
            },
            GrandParentOptional: optionalv1go.String("provident"),
            GrandParentRequired: "distinctio",
            GrandParentRequiredNullable: "quibusdam",
            Optional: optionalv1go.String("unde"),
            OptionalNullable: optionalv1go.String("nulla"),
            OptionalNullableWithDefaultValue: optionalv1go.String("corrupti"),
            ParentOptional: optionalv1go.String("illum"),
            ParentOptionalNullableWithDefaultValue: optionalv1go.String("vel"),
            ParentRequired: "error",
            ParentRequiredNullable: "deserunt",
            Required: "suscipit",
            RequiredNullable: "iure",
            Class: optionalv1go.Int(297534),
            Precision: optionalv1go.Float64(8917.73),
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