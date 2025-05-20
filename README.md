# Lookup
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/deluan/lookup?label=latest)](https://github.com/deluan/lookup/releases)
[![Documentation](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat)](https://godoc.org/github.com/deluan/lookup)
[![Build Status](https://img.shields.io/github/actions/workflow/status/deluan/lookup/go.yml?branch=master&logo=github&style=flat-square)](https://github.com/deluan/lookup/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/deluan/lookup)](https://goreportcard.com/report/github.com/deluan/lookup)
[![Maintainability](https://api.codeclimate.com/v1/badges/d4ff0afbc348c6b9291e/maintainability)](https://codeclimate.com/github/deluan/lookup/maintainability)


It is a nice, simple and fast library which helps you to lookup objects on a screen. It also includes 
OCR functionality. Using Lookup you can do OCR tricks like recognizing any information in your Robot
application. Which can be useful for debugging or automating things.

This library is a port of the [Java Lookup library](https://gitlab.com/axet/lookup) 
to GoLang. Details of NCC (Normalized Cross Correlation), used by this library, can be found in the 
original library's ['docs'](https://gitlab.com/axet/lookup/tree/master/docs) folder (a lot of math).

### Usage

Add this library to your project with:
```shell script
go get github.com/deluan/lookup
```

The Go `image` package only knows about the image formats that you register.
If you want to decode PNG, JPEG or other formats, remember to import their
decoders for their side‑effects:

```go
import (
    _ "image/png"
    _ "image/jpeg"
)
```

Without these imports calls like `image.Decode` will fail and may result in nil
images being passed to Lookup.

Passing a `nil` image to Lookup will now panic with a clear message. Make sure
to check the errors returned by `image.Decode` before using the image.

To learn how to use it, take a look at the example files for [Lookup](examples_lookup_test.go) and 
[OCR](examples_ocr_test.go). All images used in the examples are available in the [testdata](testdata) folder. 
For more details check the full [documentation](https://godoc.org/github.com/deluan/lookup).

### To Do:
- ~~Add basic LookUp function~~
- ~~Implement OCR~~
- ~~Optimize for speed~~
- ~~Clean-up API~~
- ~~Better docs~~
- Implement Scaling
