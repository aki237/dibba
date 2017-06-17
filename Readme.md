# dibba - package format reader

[![Go Report Card](https://goreportcard.com/badge/github.com/aki237/dibba)](https://goreportcard.com/report/github.com/aki237/dibba)

Dibba is a small go library for a new TLV based file format for storing just files.
Refer [godoc](https://godoc.org/github.com/aki237/dibba) for more.
See Examples for sample usage.

+ Packager : Packages files into a single dibba file.
  ```
  $ cd $GOPATH/github.com/aki237/dibba/examples/packager/
  $ go build
  $ ./packager outFile.dib example/*
  35
  4259516
  $ ls
  .  ..  example  outFile.dib  package.go  packager
  ```

+ Parser : Parse the dibba file and print out the contents of a given file from the package
  ```
  $ cd $GOPATH/github.com/aki237/dibba/examples/parser/
  $ go build
  $ ./parser ../packager/outFile.dib Readme.md
  # Readme

  This is a sample readme.
  ```
