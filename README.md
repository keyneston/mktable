# mktable

`mktable` is a simplish cli tool which consumes content via STDIN and converts
it into a pretty aligned table ready for use in markdown.

# Installation

With go installed:

```
go install github.com/keyneston/mktable@latest
```

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

| Flag               | Default          | Description                                     |
| ------------------ | ---------------- | ----------------------------------------------- |
| `-s`               | `[ \t]*\t[ \t]*` | Regexp used to set the delimiter                |
| `-no-headers`      | `false`          | Skip printing headers                           |
| `-r` / `-reformat` | `false`          | Reformat existing markdown table                |
| `-a`               | none             | Sets the alignment. See [Alignment](#alignment) |

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
