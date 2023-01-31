# png2jpeg

A small utility to convert png images to jpeg in an output directory

```
A small utility to convert png images to jpeg in an output directory

Usage:
  png2jpeg [command]

Available Commands:
  help        Help about any command

Flags:
  -h, --help            help for png2jpeg
  -o, --output string   Output path to convert images to
  -p, --path string     Path to images
  -q, --quality int     Output path to convert images to
  -t, --toggle          Help message for toggle
  
Use "png2jpeg [command] --help" for more information about a command
```

Convert command takes the following flags:

```
--path -p [required]  /path/to/image/files

example: png2jpeg --path ./images/

--output -o [optional] /path/to/output/images

example: png2jpeg --path ./images/ --output ./images/test

--quality -q [optional] 80 [Default]

example: png2jpeg --path ./images/ --quality 100

```

## TODO

* Play around with github actions (DONEISH) 
* Add tests
* Build a release pipeline