# Muddle
[![Build Status](https://travis-ci.com/Spaceface16518/Muddle.svg?branch=master)](https://travis-ci.com/Spaceface16518/Muddle)  
Mess up text through translation with ease!

One of my favorite activities on long bus rides is putting song lyrics through several layers of [Google Translate](https://translate.google.com/) and then seeing how the lyrics have changed.

Translation of natural language is a huge feat of human technology and should be taken seriously, as those who contribute to translation software enable global interconnectedness and all of the benefits that come with it.

Unfortunately, this software is not always perfect. And until it is, we will continue to mock it.

* * *

Muddle is a command-line tool that helps us get to the fun part of this activity, faster. Simply pipe some raw text through standard input into the program and watch it get translated to and from different languages and then back to English.

## Installation

To get Muddle, either download a prebuilt binary from the [Releases page](https://github.com/Spaceface16518/Muddle/releases) of this repository, or build it from source

### Building from source

This [Go](https://golang.org/) project uses Go modules

To build from source, you will need Go version `1.12.2` or higher (specified by [`go.mod`](https://github.com/Spaceface16518/Muddle/blob/master/go.mod))

Simply run

```
$ go build
```

to build, or

```
$ go run muddle
```

to run it straight from the Go cli. Remember to specifiy flags and pipe standard input if you choose `go run`.

## Usage

Muddle is a simple command line application.

For the following examples, muddle is run as though on a unix system. Other systems may have to run `muddle.exe` instead of `./muddle`

It only accepts raw textual input. The easiest way to use it is to copy and paste your favorite song lyrics into a `.txt` file and pipe it into the program; for example

```
$ cat my-lyrics.txt | ./muddle
```

This will output the translation to standard output, while outputing log entries to standard error. There is no way to silence the log right now but this functionality will likely be implemented soon.

For the help menu, run

```
$ ./muddle -h
```

or

```
$ ./muddle --help
```

Other important options are those involving saving and path information. Muddle can save the output to a file instead of standard output, enabled by the `-save` flag. If you enable this, you must provide a directory to save in using the `-path` flag; for example

```
$ cat my-lyrics.txt | ./muddle -save -path ./translated-lyrics-dir/my-lyrics/
```

## Credentials

If you get an error about credentials, you will need to get a bit more technical. This program does not in fact use Google Translate, but instead uses IBM Watson's Language Translator service. You will need an `ibm-credentials.env` file in your working directory.

This section is a **work in progress**. Come back for more details soon; consider `star`ring or `watch`ing this repository for updates!
