#!/bin/bash

cd vendor/globalxtreme/compress-file/src/go
/usr/local/go/bin/go get -v github.com/chai2010/webp
/usr/local/go/bin/go get -v github.com/disintegration/imaging
cp -Rf vendor/* ../../../../../vendor
/usr/local/go/bin/go build -o resize-image resizeImage.go
/usr/local/go/bin/go build -o resize-video resizeVideo.go
