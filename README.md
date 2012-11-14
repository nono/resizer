# Resizer

Resizer is a small command-line utility written in [Go](http://golang.org/).

Use it to pad images (ie add bleed) to a standard (4:3, 3:2, 5:4)
ratio so that it could be printed by popular online print services safely and
cost-effectively.


## TL;DR

Indeed, online photo print services (apple, negatifplus) follow a rather
stupid logic that leads to image loss when supplied with non-standard ratio
images: they just crop the picture to fit the chosen standard ratio.

`Resizer` nicely selects the closest standard ratio if none is given as an
argument.


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

## TODOs & FIXMEs

* DONE : ~~use goroutines to parallelize image resizing.~~
* DONE : ~~don't do anything if the image has a perfect ratio.~~
* DONE : ~~don't use `sync` package since @nono doesn't like it.~~
* TODO : add relevant suffix to the bleeded image.
* TODO : have the complementary rectangle somehow like a black&white
  checkerboard instead of plain white/black to ease cropping
* TODO : provide max print size at 300dpi (in output filename suffix ?)
* TODO : rename to "bleeder" instead of "resizer". May need to change github
  project name too.


## How to use it?

[Install Go 1](http://golang.org/doc/install) and run this command:

    go run resizer.go -dir=resized -color=white -ratio=40:60 image1 image2 ...


## Credits

♡2012 by Jérôme Andrieux and Bruno Michel. Copying is an act of love. Please copy and share.

Released under the MIT license
