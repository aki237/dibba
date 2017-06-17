# dibba - package format reader

Dibba is a small go library for a new TLV based file format for storing just files.
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
