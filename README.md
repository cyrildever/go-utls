# go-utls
_Utilities for Go_

![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/cyrildever/go-utls)
![GitHub last commit](https://img.shields.io/github/last-commit/cyrildever/go-utls)
![GitHub issues](https://img.shields.io/github/issues/cyrildever/go-utls)
![GitHub](https://img.shields.io/github/license/cyrildever/go-utls)

go-utls is a small Go repository where I put all the useful stuff I regularly need in my projects.
Feel free to use at your discretion with the appropriate license mentions.

_NB: I've developed the same kind of library for TypeScript, available on NPM: [`ts-utls`](https://www.npmjs.com/package/ts-utls)_


### Usage

```console
go get github.com/cyrildever/go-utls
```

This repository contains the following modules:
- `crypto`: a proxy to Go-Ethereum's ECIES library and to my [`ecies-geth`](https://www.npmjs.com/package/ecies-geth) JavaScript library (including the `Path` type) as well as a small `SSHPublicKey2String()` utility and a working copy of a [`bip32` module](https://github.com/sammyne/bip32) I needed;
- `io`: a light REST client utility on top of `fasthttp` with `Delete`, `Get`, `Patch`, `Post` and `Put` methods;
- `model`: a list of types I frequently use in my projects (such as `Base64` or `Hash` types) all implementing my [`Model`](model/Model.go) interface;
- `normalizer`: the adaptation of my Empreinte Sociométrique&trade; patented work for normalizing contact data (see its specific [README](normalizer/README.md) or use its TypeScript equivalent on NPM: [`es-normalizer`](https://www.npmjs.com/package/es-normalizer));
- a few utility sub-modules:
  * `caller`: to get information about the location of the calling function (file name and line number);
  * `concurrent`: to handle concurrent maps and slices (with faster slice appending when its length is set at instantiation through the `concurrent.NewSlice` function);
  * `email`: my "quick-and-dirty" SMTP client (including examples in tests for use with AWS SES or Gmail);
  * `env`: to know if an environment variable is set and cast it as either a boolean, an integer or a string;
  * `event`: a simple event bus manager;
  * `file`: to find, truncate, know existence, delete or get lines from files;
  * `logger`: a wrapper to the `log` package to output logs to stderr and optionally a file;
  * `ntp`: another small wrapper to handle time with NTP;
  * `packer`: to marshal/unmarshal data (JSON, MessagePack, MongoDB's Bson, &mldr;);
  * `utils`: a bunch of useful utility functions (`Capitalize()`, `Chunk()`, `DateFormat()` from Java notation, `EuclideanDivision()`, `Flatten()`, `FromHex()`/`ToHex()`, back-and-forth conversions of byte arrays (to string, number, etc.), `IsPointer()`/`IsValue()` test methods, `ToUTF8()` string formatting, &mldr;);
  * `xor`: to apply XOR operation to strings or byte arrays.


### License

These modules are distributed under a MIT license (except for the `bip32` submodule which is an [ISC license](crypto/bip32/LICENSE)). \
See the [LICENSE](LICENSE) file.


<hr />
&copy; 2020-2022 Cyril Dever. All rights reserved.