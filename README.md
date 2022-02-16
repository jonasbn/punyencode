# punyencode

A simple CLI tool to encode a string to a punycode encoded string

## Usage

Encoding a single string, meaning conversion from human readable text to punycode

```bash
punyencode blåbærgrød
```

Will emit

```text
xn--blbrgrd-fxak7p
```

As an alternative to provided arguments, you can pipe text into `punyencode`

```bash
echo blåbærgrød | punydencode
```

Will emit

```text
xn--blbrgrd-fxak7p
```

## Installation

Installation is easy using Go

```bash
go install github.com/jonasbn/punyencode@latest
```

If you want a particular version, please see [Go Modules Reference][MOD]

## Description

If you want to decode from punycode, see [`punydecode`][punydecode].

## Diagnostics

## Exit Status

- `0` success, provided string has been decoded and printed

- `1` failure not argument provided or data from STDIN

- `2` failure reading from STDIN

## Dependencies

This utility requires:

- [golang.org/x/net/idna][goidna]

In addition to a few of the standard libraries

## Bugs and Limitations

There a no known bugs, please see the GitHub repository [issues section](https://github.com/jonasbn/punyencode/issues) for a up to date listing.

### Only support for Unicode

The utility is limited to decoding to Unicode (version 13) from Punycode.

Please see [golang.org/x/net/idna][goidna] for details.

### Only a single argument

`punyencode` only takes a single argument.

```bash
punyencode xn--blbrgrd-fxak7p
```

## Author

- jonasbn

## Motivation

This utility was created, when in the process of learning Go. I have worked in the DNS and domain name business for a decade so it was only natural to work on something I know when learning Go.

This particular repository touched the following topics:

1. Learning to make CLI tools
1. Making an executable distributable and installable component
1. Reading data from the CLI
1. Reading data from STDIN
1. Testing a CLI tool / Main function in Go

See the resources and references below for resources on the above topics.

## Resources and References

1. [Wikipedia: Punycode](https://en.wikipedia.org/wiki/Punycode)
1. [Go Modules Reference][MOD]
1. [GitHub: punyencode][punyencode]
1. [golang.org/x/net/idna][goidna]
1. [yourbasic.org/golang: Read a file (stdin) line by line](https://yourbasic.org/golang/read-file-line-by-line/)
1. [Blog post: Test the main function in Go by Johannes Malsam](https://mj-go.in/golang/test-the-main-function-in-go)

## License and Copyright

Copyright Jonas Brømsø (jonasbn) 2022

MIT License, see separate `LICENSE` file

[MOD]: https://go.dev/ref/mod#go-install
[punydecode]: https://github.com/jonasbn/punydecode
[goidna]: https://pkg.go.dev/golang.org/x/net/idna
