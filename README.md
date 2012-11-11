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


## examples

``` 
square (1:1)
	
	+--------+
	|•••--•••|
	|••----••|
	|••----••|
	|••----••|
	+--------+

==> 4:3

	+-----------+
	|•••--•••   |
	|••----••   |
	|••----••   |
	|••----••   |
	+-----------+

```
or 

```
panoramic (5:2)
	
	+---------------+
	|....::::::.....|
	|.:::::::::::::.|
	+---------------+

==> 3:2

	+---------------+
	|....::::::.....|
	|.:::::::::::::.|
	|               |
	|               |
	+---------------+

```

TODOs & FIXMEs
--------------

* DONE : use goroutines to parallelize image resizing.
* DONE : don't do anything if the image has a perfect ratio.
* TODO : add suffix to the bleeded image.
* TODO : have the complementary rectangle somehow like a black&white
  checkerboard instead of plain white/black to ease cropping
* TODO : provide max print size at 300dpi
* TODO : rename to "bleed" instead of "resize".


How to use it?
--------------

[Install Go 1](http://golang.org/doc/install) and run this command:

    go run resizer.go -dir=resized -color=white -ratio=40:60 image1 image2 ...


Credits
-------

♡2012 by Bruno Michel. Copying is an act of love. Please copy and share.

Released under the MIT license
