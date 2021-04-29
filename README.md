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
