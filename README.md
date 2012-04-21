Resizer
=======

Resizer is a small command-line utility written in [Go](http://golang.org/).
It resizes images to a ratio just by expanding it. No deformation or
interpolation are made on the images.


How to use it?
--------------

[Install Go 1](http://golang.org/doc/install) and run this command:

    go run resizer.go -dir=resized -color=white -ratio=40:60 image1 image2 ...


Credits
-------

♡2012 by Bruno Michel. Copying is an act of love. Please copy and share.

Released under the MIT license
