# LookUp
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/deluan/lookup?label=latest)](https://github.com/deluan/lookup/releases)
[![Documentation](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat)](https://godoc.org/github.com/deluan/lookup) 
[![Build Status](https://github.com/deluan/lookup/workflows/CI/badge.svg)](https://github.com/deluan/lookup/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/deluan/lookup)](https://goreportcard.com/report/github.com/deluan/lookup)
[![Coverage](http://gocover.io/_badge/github.com/deluan/lookup)](http://gocover.io/github.com/deluan/lookup) 
[![Maintainability](https://api.codeclimate.com/v1/badges/d4ff0afbc348c6b9291e/maintainability)](https://codeclimate.com/github/deluan/lookup/maintainability)


It is a nice, simple and fast library which helps you to lookup objects on a screen. It also includes 
OCR functionality. Using Lookup you can do OCR tricks like recognizing any information in your Robot
application. Which can be useful for debugging or automating things.

This library is a port of this [Java Lookup library](https://gitlab.com/axet/lookup) 
to GoLang. Details on NCC (Normalized cross correlation) used by this library can be found in the 
original library's ['docs'](https://gitlab.com/axet/lookup/tree/master/docs) folder (a lot of math).

### Usage

Add this library to your project with:
```shell script
go get github.com/deluan/lookup
```

To learn how to use it, take a look at the examples files for [Lookup](examples_lookup_test.go) and 
[OCR](examples_ocr_test.go). All images used in the examples are available in the [testdata](testdata) folder. 
For more details check the full [documentation](https://godoc.org/github.com/deluan/lookup).

### To Do:
- ~~Add basic LookUp function~~
- ~~Implement OCR~~
- ~~Optimize for speed~~
- ~~Clean-up API~~
- ~~Better docs~~
- Implement Scaling
