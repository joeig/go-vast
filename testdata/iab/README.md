# IAB VAST 4.2 samples

* Source: https://github.com/InteractiveAdvertisingBureau/VAST_Samples/tree/master/VAST%204.2%20Samples
* License: [LICENSE](LICENSE)

## Changes

* Removed unused `xmlns:xs="http://www.w3.org/2001/XMLSchema"` attributes from `VAST` tags since XMLNS prefixes are currently unsupported by `encoding/xml` ([#48641](https://github.com/golang/go/pull/48641)).
* Removed surrounding Unicode spaces from tags that contain `CDATA` ([#43168](https://github.com/golang/go/issues/43168)).
* Changed the `Pricing` from `25.00` to `25.12` and removed surrounding spaces in order to properly test floating point values.
* `Inline_Companion_Tag-test.xml` and `Inline_Non-Linear_Tag-test.xml`: Removed the `CDATA` tag from `AdTitle` to comply with the specification.
