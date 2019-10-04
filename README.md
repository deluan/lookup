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

This library is a straight port of the [Java LookUp library](https://github.com/iamshajeer/lookup) 
to GoLang. Details on NCC (Normalized cross correlation) used by this library can be found in the 
original library's ['doc'](https://github.com/iamshajeer/lookup/tree/master/doc) folder (a lot of math).

### Usage

Take a look at the examples files for [Lookup](examples_lookup_test.go) and [OCR](examples_ocr_test.go) 
for usage samples. For more details check the [documentation](https://godoc.org/github.com/deluan/lookup).

### To Do:
- ~~Add basic LookUp function~~
- ~~Implement OCR~~
- ~~Optimize for speed~~
- ~~Clean-up API~~
- ~~Better docs~~
- Implement Scaling
