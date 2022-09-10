gr
==

Find the numbers which related to the golden ratio.

```
$ gr -o8
5 13

$ gr -r10-100
13
21
33
54
87
```

Usage
-----

```
$ gr OPTION...
Find the numbers which related to the golden ratio.

Predicates:
  -n, --near=N      find numbers near N
  -o, --of=N        find numbers of N
  -r, --range=N-M   find numbers between N and M

Miscellaneous:
  -f, --first=N     set first golden number for N (default: 3.0)
  -d, --decimal     show number as a decimal
      --help        display this help and exit
      --version     display version information and exit
```

Installation
------------

### compiled binary

See [releases](https://github.com/nil-two/gr/releases)

### go get

```
go get github.com/nil-two/gr
```

Options
-------

### -h, --help

Display a help message and exit.

### -v, --version

Display version information and exit..

### -n, --near=N

Find golden numbers near N.

```
$ gr -n9
8 13
(3, 5, [8, 13], 21)

$ gr -n200
141 228
(3, 5, 8, 13, 21, 33, 54, 87, [141, 228], 369)
```

### -o, --of=N

Find golden numbers of N.

```
$ gr -o10
6 16
(10/1.618, 10*1.618)

$ gr -o32
20 52
(32/1.618, 10*1.618)
```

### -r, --range=N-M

Find numbers between N and M.
- When N was omitted, range starts from 0.
- When M was omitted, range continues infinitely.

```
$ gr -r10-100
13
21
33
54
87

$ gr -r-10
3
5
8

$ gr -r100- | head 10
141
228
369
597
966
...

$ gr -r-
3
5
8
13
21
...
```

### -f, --first=N

Set the first golden number for N (default: `3.0`)

```
$ gr -n5
5 8
(3, [5, 8], 13)

$ gr -n5 -f1
4 7
(1, 2, 3, [4, 7], 11)

$ gr -r-10
3
5
8

$ gr -r-10 -f1
1
2
3
4
7
```

### -d, --decimal

Show number as a decimal.

```
$ gr -dn5
4.854101966249685 7.854101966249685

$ gr -do9.5
5.871322893124001 15.371322893124002

$ gr -dr10-100
12.70820393249937
20.562305898749056
33.270509831248425
53.832815729997485
87.10332556124591
```

License
-------

MIT License

Author
------

nil2 <nil2@nil2.org>
