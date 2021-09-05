## 题目

### 1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

## 测试命令

```bash
redis-benchmark -t get,set -q -d 10
SET: 123762.38 requests per second, p50=0.207 msec
GET: 128040.97 requests per second, p50=0.199 msec

redis-benchmark -t get,set -q -d 20
SET: 129032.27 requests per second, p50=0.199 msec
GET: 129198.97 requests per second, p50=0.199 msec

redis-benchmark -t get,set -q -d 50
SET: 128369.71 requests per second, p50=0.199 msec
GET: 131406.05 requests per second, p50=0.199 msec

redis-benchmark -t get,set -q -d 100
SET: 127877.23 requests per second, p50=0.199 msec
GET: 128040.97 requests per second, p50=0.199 msec

redis-benchmark -t get,set -q -d 200
SET: 131406.05 requests per second, p50=0.199 msec
GET: 130890.05 requests per second, p50=0.199 msec

redis-benchmark -t get,set -q -d 1024
SET: 130378.09 requests per second, p50=0.199 msec
GET: 128700.12 requests per second, p50=0.199 msec

redis-benchmark -t get,set -q -d 5120
SET: 121802.68 requests per second, p50=0.207 msec
GET: 123456.79 requests per second, p50=0.207 msec

``` 

## 分析

从每秒请求数看，10 20 50 100 200 1k 5k 字节 value 大小，对于redis get set 性能影响不大

### 2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

老师，这个怎么操作，然后分析呢
