# sp

Generate PHPUnit Yaml Dataset row format from TSV.

## Installation

```sh
$ go get github.com/genkiroid/sp/...
```

## Usage

```sh
$ sp < data.tsv
```

See [example](https://github.com/genkiroid/sp/blob/master/sp_test.go) about output.

---

As Sequel Pro Bundle.(Copy as yaml.)

Command:

```sh
#! /bin/bash
cat | /path/to/sp | __CF_USER_TEXT_ENCODING=$UID:0x8000100:0x8000100 pbcopy
```

## Author

[genkiroid](https://github.com/genkiroid)

