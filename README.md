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



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->
# Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

## Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
