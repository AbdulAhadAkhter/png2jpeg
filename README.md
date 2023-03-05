# png2jpeg

A small utility to convert png images to jpeg in an output directory

```
A small utility to convert png images to jpeg in an output directory

Usage:
  png2jpeg

Available Commands:
  help        Help about any command

Flags:
  -h, --help            help for png2jpeg
  -o, --output string   Output path to convert images to
  -b, --batch string    Path to images
  -s, --single string   Path to image
  -q, --quality int     Output path to convert images to
  -t, --toggle          Help message for toggle
  
Use "png2jpeg --help" for more information about a command
```

Convert command takes the following flags:

```
--single -s <required>  /path/to/image/file.png

example: png2jpeg --single ./images/test.png

                  or

--batch -b <required>  /path/to/image/files

example: png2jpeg --batch ./images/

--output -o [optional] /path/to/output/images

example: png2jpeg --path ./images/ --output ./images/test

--quality -q [optional] 80 [Default]

example: png2jpeg --path ./images/ --quality 100

```

## TODO

* Play around with github actions to build tag and release (DONEISH) 
* Add tests
