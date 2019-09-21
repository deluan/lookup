# LookUp
[![Build Status](https://github.com/deluan/lookup/workflows/CI/badge.svg)](https://github.com/deluan/lookup/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/deluan/lookup)](https://goreportcard.com/report/github.com/deluan/lookup)
[![Coverage](http://gocover.io/_badge/github.com/deluan/lookup)](http://gocover.io/github.com/deluan/lookup) 
[![Documentation](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat)](https://godoc.org/github.com/deluan/lookup) 


**NOTE**: This is alpha quality code. The API is not finalized yet and will most likely change.

It is a nice, simple and fast library which helps you to lookup objects on a screen. It also includes 
OCR functionality. Using Lookup you can do OCR tricks like recognizing any information in your Robot application. 
Which can be useful for debugging or automating things.

This library is a straight port of the [Java LookUp library](https://github.com/iamshajeer/lookup) to GoLang.
Details on NCC (Normalized cross correlation) used by this library can be found in the original 
library's ['docs'](https://github.com/corintio/lookup/tree/master/docs) folder (a lot of math).

While there is no documentation, take a look at the [lookup_test.go](lookup_test.go) and [ocr_test.go](ocr_test.go) files
for usage examples.

### To Do:
- ~~Add basic LookUp function~~
- ~~Implement OCR~~
- ~~Optimize for speed~~
- ~~Clean-up API~~
- Implement Scaling
- Better docs
