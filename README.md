# LookUp
[![Build Status](https://github.com/deluan/go-lookup/workflows/CI/badge.svg)](https://github.com/deluan/go-lookup/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/deluan/go-lookup)](https://goreportcard.com/report/github.com/deluan/go-lookup)
[![Coverage](http://gocover.io/_badge/github.com/deluan/go-lookup)](http://gocover.io/github.com/deluan/go-lookup) 

**NOTE**: This is alpha quality code, and the API is not finalized yet and will most likely change.

It is a fast, nice, simple and friendly to use library which helps you to lookup objects on a screen. It also includes 
OCR functionality. Using Lookup you can do OCR tricks like recognizing any information in your Robot application. 
Which can be useful for debugging or automating things.

This library is a straight port of the [Java LookUp library](https://github.com/iamshajeer/lookup) to GoLang.
Details on NCC (Normalized cross correlation) used by this library can be found in the original 
library's ['docs'](https://github.com/corintio/lookup/tree/master/docs) folder (a lot of math).

While there is no documentation, take a look at the [ncc_test.go](ncc_test.go) and [ocr_test.go](ocr_test.go) files
for usage examples.

### To Do:
- ~~Add basic LookUp function~~
- ~~Implement OCR~~
- Optimize for speed
- Clean-up API
- Implement Scaling
- Better docs
