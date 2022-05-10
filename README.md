![logo](logo.png)

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)  [![GoDoc](https://godoc.org/github.com/cloudflare/cfssl?status.svg)](https://github.com/nisainan/gredcon)

A Redis compatible server framework for Go based on [gnet](https://github.com/panjf2000/gnet)

## Features

- Create a [Fast](https://github.com/nisainan/gredcon#benchmarks) custom Redis compatible server in Go
- Simple interface. One function `ListenAndServe` and two types `Conn` & `Command`
- Support for pipelining and telnet commands
- Works with Redis clients such as [redigo](https://github.com/garyburd/redigo), [redis-py](https://github.com/andymccurdy/redis-py), [node_redis](https://github.com/NodeRedis/node_redis), and [jedis](https://github.com/xetorthio/jedis)
- Compatible pub/sub support
- Multithreaded

## Installing

~~~she
go get -u github.com/nisainan/gredcon
~~~

## Example

Here's a full example of a Redis clone that accepts:

- SET key value
- GET key
- DEL key
- PING

You can run this example from a terminal:

~~~shell
go run example/main.go
~~~

## Benchmarks

**Redis**: Single-threaded, no disk persistence.

```shell
$ redis-server --port 6379 --appendonly no
redis-benchmark -p 6379 -t set,get -n 10000000 -q -P 512 -c 512
SET: 941265.12 requests per second
GET: 1189909.50 requests per second
```

**Redcon**: Single-threaded, no disk persistence.

```shell
$ GOMAXPROCS=1 go run example/clone.go
redis-benchmark -p 6380 -t set,get -n 10000000 -q -P 512 -c 512
SET: 2018570.88 requests per second
GET: 2403846.25 requests per second
```

**Redcon**: Multi-threaded, no disk persistence.

```shell
$ GOMAXPROCS=0 go run example/clone.go
$ redis-benchmark -p 6380 -t set,get -n 10000000 -q -P 512 -c 512
SET: 1944390.38 requests per second
GET: 3993610.25 requests per second
```

*Running on a MacBook Pro 15" 2.8 GHz Intel Core i7 using Go 1.7*

**GRedcon**: multicore, no disk persistence.

~~~shell
$ redis-benchmark -p 9876 -c 120 -n 20000000  -t get,set -P 2000 -q
SET: 5630630.50 requests per second
GET: 11428028.00 requests per second
~~~

*Running on a Ubuntu20.04  i7-9700 CPU @ 3.00GHz using Go 1.7*

## License

GRedcon source code is available under the MIT [License](https://github.com/nisainan/gredcon/blob/master/LICENSE).
