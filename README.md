# mktable

[![Go](https://github.com/keyneston/mktable/actions/workflows/go.yml/badge.svg)](https://github.com/keyneston/mktable/actions/workflows/go.yml)

`mktable` is a simplish cli tool which consumes content via STDIN and converts
it into a pretty aligned table ready for use in markdown.

# Installation

```shell
brew tap keyneston/tap
brew install mktable
```

With go installed:

```shell
go install github.com/keyneston/mktable@latest
```

Or download the [latest release](https://github.com/keyneston/mktable/releases/) and add to your path.


# Use

```
$ netstat -i | head -5 | mktable -s ' +'
| Name | Mtu   | Network     | Address   | Ipkts   | Ierrs   | Opkts   | Oerrs | Coll |
| ---- | ----- | ----------- | --------- | ------- | ------- | ------- | ----- | ---- |
| lo0  | 16384 | <Link#1>    | 9583572   | 0       | 9583572 | 0       | 0     |      |
| lo0  | 16384 | 127         | localhost | 9583572 | -       | 9583572 | -     | -    |
| lo0  | 16384 | localhost   | ::1       | 9583572 | -       | 9583572 | -     | -    |
| lo0  | 16384 | keyport.loc | fe80:1::1 | 9583572 | -       | 9583572 | -     | -    |
```

## Flags

| Flag               | Default          | Description                                                    |
| ------------------ | ---------------- | -------------------------------------------------------------- |
| `-s`               | `[ \t]*\t[ \t]*` | Regexp used to set the delimiter                               |
| `-no-headers`      | `false`          | Skip printing headers                                          |
| `-r` / `-reformat` | `false`          | Reformat existing markdown table. Alias for `-format mk`       |
| `-a`               | none             | Sets the alignment. See [Alignment](#alignment)                |
| `-f` / `-format`   | `regexp`         | Sets the input format. See [Formats](#format) for more details |

## Formats

| Format            | Description                                         |
| ----------------- | --------------------------------------------------- |
| `re` / `regexp`   | Regular Expression Delimiter                        |
| `csv`             | Comma Separated List                                |
| `mk` / `markdown` | Consume an existing markdown table for reformatting |


## Alignment

**EXPERIMENTAL**

The alignment flag can be passed multiple times and/or passed as a comma
seperated list.

It takes a number indicating a zero indexed column, and a character
indicating how to align the column.

| char | alignment |
| ---- | --------- |
| `>`  | right     |
| `<`  | left      |
| `=`  | center    |

### Examples

| Example              | Description                                             |
| -------------------- | ------------------------------------------------------- |
| `-a '0<'`            | left align the 1st column                               |
| `-a '1='`            | center align the 2nd column                             |
| `-a '1=,2>'`         | center align the 2nd column, and right align the second |
| `-a '1=,2>' -a '3='` | as above, but also center align the 4th column          |


## Reformat existing tables

```
$ mktable -r <<EOF
| Foo      | Bar |
| a | b
| a  | b |
EOF
| Foo | Bar |
| --- | --- |
| a   | b   |
| a   | b   |
```

This can be used with vim to easily fix tables. First select the the lines you want (`Shift-V`) and then pipe it `:!mktable -r`.
