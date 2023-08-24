# github.com/speakeasy-sdks/optional-v1-go

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/optional-v1-go
```
<!-- End SDK Installation -->

## SDK Example Usage
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Body](docs/sdks/body/README.md)

* [SendChild](docs/sdks/body/README.md#sendchild) - sendChild

### [RfcDate](docs/sdks/rfcdate/README.md)

* [SendRfc1123Date](docs/sdks/rfcdate/README.md#sendrfc1123date) - sendRfc1123Date
* [SendRfc1123DateArray](docs/sdks/rfcdate/README.md#sendrfc1123datearray) - sendRfc1123DateArray
* [SendRfc1123DateMap](docs/sdks/rfcdate/README.md#sendrfc1123datemap) - sendRfc1123DateMap

### [SimpleDate](docs/sdks/simpledate/README.md)

* [SendSimpleDate](docs/sdks/simpledate/README.md#sendsimpledate) - sendSimpleDate
* [SendSimpleDateArray](docs/sdks/simpledate/README.md#sendsimpledatearray) - sendSimpleDateArray

### [UnixDate](docs/sdks/unixdate/README.md)

* [SendUnixDate](docs/sdks/unixdate/README.md#sendunixdate) - sendUnixDate
* [SendUnixDateArray](docs/sdks/unixdate/README.md#sendunixdatearray) - sendUnixDateArray
* [SendUnixDateMap](docs/sdks/unixdate/README.md#sendunixdatemap) - sendUnixDateMap
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
