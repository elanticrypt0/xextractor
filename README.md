# Extract email and domains from one file.


# Install

If you have go installed in your system is very easy. Just run in your console:

```go
go install github.com/k23dev/xextractor@latest
```
If you dont have go. What are you waiting for? go for it [go oficial site](go.dev) download and in windows is just next next next and thats it!

You can add the path and thats it

otherwise you cant download the packet and build or use it like and script running 

```bash
git clone https://github.com/k23dev/xextractor
cd xextractor
go run . [flags]
```

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

## set your rules

create a .txt file with this format:

```text
[rulename1] = [regex1]
[rulename2] = [regex2]
```

```bash
go run . -file [PATH_TO_FILE] -o [PATH] -dic [PATH_TO_DIC_FILE] -ex [YOUR_CUSTOM_RULE]
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

