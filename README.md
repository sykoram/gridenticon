# identicon

A SVG identicon generator.

## Table of Contents

- [Setup](#setup)
- [Usage](#usage)


## Setup

[Go](https://golang.org/) and [Git](https://git-scm.com/) have to be installed.

Clone or download this repo somewhere into `$GOPATH`, preferably into `$GOPATH/src/github.com/sykoram/identicon`.

Download and install all dependencies:
```sh
go get ./...
```

And install the program:
```sh
go install
```

This creates an executable inside `$GOBIN` (usually `$GOPATH/bin`), and program should work now. You can try to run:

```sh
identicon -h
```


## Usage

Use flag `-h` to display the help.

The most basic command would be:
```sh
identicon STRING
```
This generates a SVG file (the identicon) using hash of the string.

- `-out OUT_FILE`: Specifies the output file (default: out.svg).
- `-defs DEFS_FILE`: Specifies the defs file containing [design](#custom-design) (default: default.defs)

### Custom Design

The identicon is a 8x8 grid. Each tile is chosen based on a part of hash (generated for the string). There are some default tile designs, but you can specify yours.

To do that you need to have a "defs file". It may look like this:
```
<g id="0"></g>
<g id="1"></g>
<g id="2"></g>
<g id="3"></g>
<g id="4"></g>
<g id="5"><circle cx="5.00" cy="5.00" r="2.00" /></g>
<g id="6"><circle cx="5.00" cy="5.00" r="2.00" /></g>
<g id="7"><circle cx="5.00" cy="5.00" r="2.00" /></g>
<g id="8"><circle cx="5.00" cy="5.00" r="2.00" /></g>
<g id="9"><circle cx="5.00" cy="5.00" r="2.00" /></g>
<g id="a"><circle cx="5.00" cy="5.00" r="2.00" /></g>
<g id="b"><circle cx="5.00" cy="5.00" r="4.00" /></g>
<g id="c"><circle cx="5.00" cy="5.00" r="4.00" /></g>
<g id="d"><circle cx="5.00" cy="5.00" r="4.00" /></g>
<g id="e"><circle cx="5.00" cy="5.00" r="4.00" /></g>
<g id="f"><circle cx="5.00" cy="5.00" r="4.00" /></g>
```

It has to contain all the ids from `0` to `f` (hex). You can put your objects into the groups to create custom tile designs.

