Resizer
=======

Resizer is a small command-line utility written in [Go](http://golang.org/).

The primary intended use for resizer is to pad images to a standard 4:3 or 3:2
ratio so that it could be printed through popular online print services.

Indeed, online photo print services such as apple, negatifplus follow a rather
stupid logic that leads to loss of image information when supplied with
non-standard ratio images: they just crop the picture to fit the chosen
standard ratio.

It nicely selects the closest standard ratio if none is given as an argument.
Oh, and default "gravity" is not center, so we don't have to crop twice ;-)


How to use it?
--------------

[Install Go 1](http://golang.org/doc/install) and run this command:

    go run resizer.go -dir=resized -color=white -ratio=40:60 image1 image2 ...


Credits
-------

♡2012 by Bruno Michel. Copying is an act of love. Please copy and share.

Released under the MIT license
