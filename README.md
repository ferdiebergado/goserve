# goserve

## Description

Quickly spin up an HTTP web server to serve static files.

## Requirements

- Go 1.22 or higher

## Usage

```
goserve [flags] <path>
```

#### Flags:

- `-a string`
  - network address to bind to (default "localhost")
- `-p string`
  - port to listen on (default "8888")

#### Arguments:

- `<path>` Path to serve (default ".")

## Example

Run without arguments to serve the current directory on localhost port 8888.

```sh
goserve
```

Serve on a particular port:

```sh
goserver -p 8080
```

Serve a particular directory:

```sh
goserve public
```

Serve a directory on a particular port:

```sh
goserve -p 8080 public
```

Serve a directory on the local network at a particular port:

```sh
goserve -a 0.0.0.0 -p 8080 public
```
