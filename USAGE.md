<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	optionalv1go "github.com/speakeasy-sdks/optional-v1-go"
	"github.com/speakeasy-sdks/optional-v1-go/pkg/models/operations"
	"github.com/speakeasy-sdks/optional-v1-go/pkg/models/shared"
	"log"
)

func main() {
	s := optionalv1go.New()

	ctx := context.Background()
	res, err := s.Body.SendChild(ctx, operations.SendChildRequest{
		ChildClass: shared.ChildClass{
			ChildClassArray: []shared.ChildClass{
				shared.ChildClass{},
			},
			GrandParentRequired:         "string",
			GrandParentRequiredNullable: "string",
			ParentRequired:              "string",
			ParentRequiredNullable:      "string",
			Required:                    "string",
			RequiredNullable:            "string",
		},
		Field:     "string",
		SetToNull: false,
		UnSet:     false,
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