# goloc
LOC counter for languages using `//`, `/**/` comments.

### Build
```
go build
```

### Usage
```
./goloc file
# -> number of LOC
```

### Known issues

Does not handle this case when there is white space (represented as ".") after newline. This will be counted as LOC.

```
// Comment
.\n                     <- will be counted
```