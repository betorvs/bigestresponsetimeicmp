# Biggest Response Time

TravisCI: [![Build Status](https://travis-ci.org/betorvs/biggestresponsetimeicmp.svg?branch=master)](https://travis-ci.org/betorvs/biggestresponsetimeicmp)

This is a simple traceroute command, that print all tested response time until reach the remote server, and show the output of "bigest response time" separately.

## Build

```
$ dep ensure
$ go build
```

## Usage

You must be root to run this command!

```
$ sudo biggestresponsetime REMOTE_HOST
```

### Special Args

Check Version command
```
$ biggestresponsetime version
biggestresponsetime Version:  0.1.0
```

Check help command
```
$ biggestresponsetime help
Usage: sudo biggestresponsetime REMOTE_HOST
```

## Example of Usage

```
$ sudo ./biggestresponsetimeicmp google.com
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

## To test it

```
$ cd test
$ sudo go test
Password:
Traceroute to www.google.com (216.58.206.4), 20 hops max, 52 byte packets
1   192.168.178.1 (192.168.178.1)  1.491047ms
2   me60.stadtnetz-bamberg.de. (185.76.96.1)  5.144797ms
3   ge-1-0-0-0.bam20.core-backbone.com. (80.255.14.253)  5.541081ms
4   ae10-2021.fra20.core-backbone.com. (80.255.14.6)  8.722421ms
5   de-cix.fra.google.com. (80.81.193.108)  8.901392ms
6   108.170.252.65 (108.170.252.65)  10.281466ms
7   216.239.56.151 (216.239.56.151)  9.021359ms
8   fra16s20-in-f4.1e100.net. (216.58.206.4)  9.39704ms

The biggest response time is:  108.170.252.65 with 10.281466ms
Traceroute to 172.217.168.228 (172.217.168.228), 20 hops max, 52 byte packets
1   192.168.178.1 (192.168.178.1)  1.428368ms
2   me60.stadtnetz-bamberg.de. (185.76.96.1)  4.628091ms
3   de-cix.fra.google.com. (80.81.193.108)  8.834956ms
4   108.170.251.145 (108.170.251.145)  9.262449ms
5   209.85.242.79 (209.85.242.79)  10.155882ms
6   209.85.244.158 (209.85.244.158)  15.583812ms
7   216.239.42.210 (216.239.42.210)  14.996354ms
8   108.170.241.129 (108.170.241.129)  14.89738ms
9   209.85.252.245 (209.85.252.245)  15.531189ms
The biggest response time is:  209.85.244.158 with 15.583812ms
10  ams15s40-in-f4.1e100.net. (172.217.168.228)  15.734017ms

error: lookup isbrobous.net: no such host
ok      github.com/betorvs/bigestresponsetimeicmp/test  0.444s
```


## To Release

```
export GITHUB_TOKEN="aa55"
git tag -a 0.1.0 -m "First release"
git push origin 0.1.0
goreleaser --rm-dist
```