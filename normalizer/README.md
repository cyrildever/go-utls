# normalizer

This is a Go normalization library for contact data (address, phones, etc.) adapted from [Edgewhere](https://www.edgewhere.fr)'s Empreinte Sociométrique normalizers.

### Motivation

When it comes to hashing data, the necessary uniqueness of the source sometimes makes it hard to compare two hashed data, mostly if dealing with postal address. This library helps normalizing any contact data (the postal or e-mail address, the mobile or landline phone number, the title, names or date of birth of an individual) that would feed any hashing algorithm to make hash comparison always trustable.
It's based on the work for the Empreinte Sociométrique&trade; patented by Cyril Dever for Edgewhere. For more information on the latter, please [contact me](mailto:cdever@edgewhere.fr).

### Usage

```
go get github.com/cyrildever/go-utls

import "github.com/cyrildever/go-utls/normalizer"
```

*IMPORTANT*: as of this version, most normalizers are for use on French data.

To get a normalized string, you simply need to use the `normalizer.Normalize()` function passing it the data, a normalizer function and eventual arguments.
There are currently eleven specific normalizer functions and a generic one:
* `normalizer.Any`: the generic normalizer should be used if no specific normalizer exists;
* `normalizer.AddressLine`: pass any address line through it to get a normalized address, eg. `8, rue Henner` becomes `8 RUE HENNER`;
* `normalizer.City`: for normalizing city names (it removes any Cedex mention in French address, for instance);
* `normalizer.CodePostalFrance`: for French zip code;
* `normalizer.DateOfBirth`: pass a date and up to two parameters: the input format (eg. `YYYY-MM-DD`) and the output format wished (eg. `normalizer.ISO_FORMAT`), the default values being respectively the ISO format and the French date format;
* `normalizer.DepartementFrance`: extract the French departement number (out of a code postal, for instance);
* `normalizer.Email`: validates the passed e-mail and returns it in lower-case;
* `normalizer.FirstName`: pass a first name and get a normalized one (it uses an enlarged French dictionay of first names to process it making it possible to pass from `"J.-Louis"` to `"JEAN LOUIS"`);
* `normalizer.Mobile`: to validate a French mobile phone;
* `normalizer.PhoneNumber`: to normalize a French phone or fax number in the international format, eg. `+33 (0) 123 456 789`;
* `normalizer.StreetNumber`: parses the passed field to normalize it the Empreinte Sociométrique&trade;'s way, eg. `1bis` or `1 bis` becomes `1B`;
* `normalizer.Title`: returns a code depending on the passed string (1 for gentlemen, 2 for ladies, 0 when undefined or unknown).

```golang
import (
    "fmt"
    
    "github.com/cyrildever/go-utls/normalizer"
)

anyStringNormalized := normalizer.Normalize("#This is any String(). ", normalizer.Any)
// THIS IS ANY STRING
fmt.Println(anyStringNormalized)

const addressLine4Normalized = normalizer.Normalize("24, rué de Maubeuge", normalizer.AddressLine)
// 24 RUE MAUBEUGE
fmt.Println(addressLine4Normalized)

const cityNormalized = normalizer.Normalize("Paris Cedex 09", normalizer.City)
// PARIS
fmt.Println(cityNormalized)

const ddnNormalized = normalizer.Normalize("70/12/01", normalizer.DateOfBirth, "YY/MM/DD", normalizer.FRENCH_DATE)
// 01/12/1970
fmt.Println(ddnNormalized)

const dptNormalized = normalizer.Normalize(" 75009 ", normalizer.DepartementFrance)
// 75
fmt.Println(dptNormalized)

const emailNormalized = normalizer.Normalize(" Contact@GMAIL.com", normalizer.Email)
// contact@gmail.com
fmt.Println(emailNormalized)

const firstNameNormalized = normalizer.Normalize("J.-Louis", normalizer.FirstName)
// JEAN LOUIS
fmt.Println(firstNameNormalized)

const mobileNormalized = normalizer.Normalize("06.23.45.67.89", normalizer.Mobile)
// +33 (0) 123 456 789
fmt.Println(mobileNormalized)

const phoneNormalized = normalizer.Normalize("0123456789", normalizer.PhoneNumber)
// +33 (0) 123 456 789
fmt.Println(phoneNormalized)

const streetNumberNormalized = normalizer.Normalize("1bis", normalizer.StreetNumber)
// 1B
fmt.Println(streetNumberNormalized)

const titleNormalized = normalizer.Normalize("Mademoiselle", normalizer.Title)
// 2
fmt.Println(titleNormalized)
```


### JavaScript library

You might want to check out the corresponding library for JavaScript developed in TypeScript: [`es-normalizer`](https://www.npmjs.com/package/es-normalizer).


### License

This module is distributed under a MIT license.
See the [LICENSE](LICENSE) file.


<hr />
&copy; 2018-2020 Cyril Dever. All rights reserved.