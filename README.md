A Redis compatible server framework for Go based on gnet

Thanks [redcon](https://github.com/tidwall/redcon),[gnet](https://github.com/panjf2000/gnet)

benchmark(multicore: true)

~~~shell
root@ubuntu:/home/taosheng# redis-benchmark -p 9876 -c 120 -n 20000000  -t get,set -P 2000 -q
ERROR: ERR unknown command 'config'
ERROR: failed to fetch CONFIG from 127.0.0.1:9876
WARN: could not fetch server CONFIG
SET: 5630630.50 requests per second
GET: 11428028.00 requests per second
~~~

