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
<!-- End SDK Example Usage -->