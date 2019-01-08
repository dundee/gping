Simple Ping util written in Go which accepts hostnames and URLs as well.

Installation
===========

```
go get github.com/Dundee/gping
sudo cp $GOPATH/bin/gping /usr/bin/gping
sudo setcap cap_net_raw+ep /usr/bin/gping
```

Usage
=====

```
$ gping https://google.com/
PING google.com (216.58.201.110):
62 bytes from 216.58.201.110: icmp_seq=0 time=1.965933ms
62 bytes from 216.58.201.110: icmp_seq=1 time=2.489417ms
62 bytes from 216.58.201.110: icmp_seq=2 time=1.810876ms
62 bytes from 216.58.201.110: icmp_seq=3 time=1.62561ms
62 bytes from 216.58.201.110: icmp_seq=4 time=2.950999ms
^C
--- google.com ping statistics ---
5 packets transmitted, 5 packets received, 0% packet loss
round-trip min/avg/max/stddev = 1.62561ms/2.168567ms/2.950999ms/485.582µs
```
