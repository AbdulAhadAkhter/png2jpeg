# png2jpeg

A small utility to convert png images to jpeg in an output directory

```
A small utility to convert png images to jpeg in an output directory

Usage:
  png2jpeg [command]

Available Commands:
  convert     Convert png images to jpeg
  help        Help about any command

Flags:
  -h, --help     help for png2jpeg
  -t, --toggle   Help message for toggle

Use "png2jpeg [command] --help" for more information about a command
```

Convert command takes the following flags:

```
--path -p [required]  /path/to/image/files

example: png2jpeg convert --path ./images/

--output -o [optional] /path/to/output/images

example: png2jpeg convert --path ./images/ --output ./images/test

```
