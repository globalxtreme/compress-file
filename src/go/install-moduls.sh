#!/bin/bash

go get -v github.com/chai2010/webp
go get -v github.com/disintegration/imaging
cp -Rf vendor/globalxtreme/compress-file/src/go/vendor/* vendor
go build -o vendor/globalxtreme/compress-file/src/go/resize-image vendor/globalxtreme/compress-file/src/go/resizeImage.go
