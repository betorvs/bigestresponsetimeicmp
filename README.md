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
bigestresponsetime Version:  0.1.0
```

Check help command
```
$ bigestresponsetime help
Usage: sudo bigestresponsetime REMOTE_HOST
```

## Example of Usage

```
$ sudo ./bigestresponsetimeicmp google.com
Password:
Traceroute to google.com (216.58.206.14), 20 hops max, 52 byte packets
1   192.168.178.1 (192.168.178.1)  2.054172ms
2   me60.stadtnetz-bamberg.de. (185.76.96.1)  28.405286ms
3   ge-2-2-0-0.bam10.core-backbone.com. (81.95.15.77)  5.993266ms
4   ae10-2021.fra20.core-backbone.com. (80.255.14.6)  9.410265ms
5   de-cix.fra.google.com. (80.81.193.108)  10.281282ms
6   108.170.252.65 (108.170.252.65)  11.33321ms
7   216.239.56.151 (216.239.56.151)  9.960641ms
8   fra16s20-in-f14.1e100.net. (216.58.206.14)  9.842548ms

The biggest response time is:  185.76.96.1 with 28.405286ms
```

## To Release

```
export GITHUB_TOKEN="aa55"
git tag -a 0.1.0 -m "First release"
git push origin 0.1.0
goreleaser --rm-dist
```