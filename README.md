# BigestResponseTime

This is a simple traceroute command, that print all tested response time until reach the remote server, and show the output of "bigest response time" separately.

## Build

```
dep ensure
go build
```

## Usage

You must be root to run this command!

```
$ sudo bigestresponsetime REMOTE_HOST
```

### Special Args

Check Version command
```
$ bigestresponsetime version
bigestresponsetime Version:  1.0.0
```

Check help command
```
$ bigestresponsetime help
Usage: sudo bigestresponsetime REMOTE_HOST
```

