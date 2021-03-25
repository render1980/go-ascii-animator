# go-ascii-animator

Yet another tool to animate your Ascii graphics.

## Build

`go build go-ascii-animator.go`

## Run

Use help:

```
./go-ascii-animator -h

Usage of ./go-ascii-animator:
  -d duration
    	Delay between frames (milliseconds) (default 500ns)
  -f string
    	File with ASCII graphics to animate
  -l	If the animation should loop”
```

Example:

```
./go-ascii-animator -d 1000 -f samples/cat.txt -l
```

## ASCII files format

- `[yaaa]` at the end of each block
- `[end]` at the end of the file

You can also see examples in samples/ folder.
