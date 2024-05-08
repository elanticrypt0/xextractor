# Extract email and domains from one file.

Simple and super fast

## run

```bash
go run . -file [PATH_TO_FILE]
```

```bash
go run . -file [PATH_TO_FILE] -o [PATH] -ex [MODE: email|domain|x]
```

## set output path

```bash
go run . -file [PATH_TO_FILE] -o [PATH]
```

## set buffer size

Default 200 MB

```bash
go run . -file [PATH_TO_FILE] -buffer [SIZE_IN_MB]
```

## Delete duplicates

```bash
go run . -clean [PATH_TO_FILE]
```

