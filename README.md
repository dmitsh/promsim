# Prometheus target simulator

`promsim` is a Prometheus target simulator, that exposes the following metrics:
- `cpu_percent_used` as gauge
- `http_requests_total` as counter
- `request_duration_seconds` as histogram
- `response_size_bytes` as summary

`promsim` allows to modify number of time series by changing `--set` parameter. Each set generates approximately 1000 time series.

Usage: `promsim target <parameters>`

Parameters:

|Long Option|Short Option|Value      |Description|
|:---------:|:----------:|:---------:|-----------|
|--address  |-a          |address    |Scraping endpoint address. Default `127.0.0.1:80` |
|--metrics  |-m          |string     |Metrics path. Default `/metrics` |
|--job      |-j          |string     |The value of the `job` label (if specified). |
|--sets     |-n          |int        |Number of metric sets. Default `1` |
|--rate     |-r          |duration   |Time interval between two metric updates. Default `1s` |
|--tls.enable|           |bool       |Enable TLS. Default `false`|
|--tls.key   |           |path       |Server key path|
|--tls.cert  |           |path       |Server certificate path|

### Building

```bash
$ make
```

#### Running in Docker container

```bash
docker run --rm -p 8080:8080 docker.io/dmitsh/promsim:0.3 --help
usage: promsim target [<flags>]

Start metrics generating target.

Flags:
      --help                Show context-sensitive help (also try --help-long
                            and --help-man).
  -a, --address=":8080"     scraping endpoint address.
  -m, --metrics="/metrics"  metrics path.
  -j, --job=JOB             job name.
  -n, --sets=1              number of time series sets.
  -r, --rate="1s"           time interval between two metric updates.
      --tls.enabled         enable TLS.
      --tls.key=TLS.KEY     path to the server key.
      --tls.cert=TLS.CERT   path to the server certificate.
```

```bash
docker run --rm -p 8080:8080 docker.io/dmitsh/promsim:0.3
```

#### Probing Prometheus endpoint

```bash
$ curl localhost:8080/metrics
# HELP cpu_percent_used CPU percent used.
# TYPE cpu_percent_used gauge
cpu_percent_used{host=":8080",module="api",set="1"} 12
cpu_percent_used{host=":8080",module="backend",set="1"} 38
cpu_percent_used{host=":8080",module="frontend",set="1"} 62
# HELP http_requests_total Total number of HTTP requests.
# TYPE http_requests_total counter
http_requests_total{host=":8080",module="api",path="/api",set="1",status_code="200"} 207
http_requests_total{host=":8080",module="api",path="/api",set="1",status_code="401"} 28
http_requests_total{host=":8080",module="api",path="/api",set="1",status_code="503"} 12
http_requests_total{host=":8080",module="api",path="/auth",set="1",status_code="200"} 205
http_requests_total{host=":8080",module="api",path="/auth",set="1",status_code="401"} 34
http_requests_total{host=":8080",module="api",path="/auth",set="1",status_code="503"} 8
http_requests_total{host=":8080",module="api",path="/home",set="1",status_code="200"} 200
http_requests_total{host=":8080",module="api",path="/home",set="1",status_code="401"} 34
http_requests_total{host=":8080",module="api",path="/home",set="1",status_code="503"} 13
http_requests_total{host=":8080",module="backend",path="/api",set="1",status_code="200"} 199
http_requests_total{host=":8080",module="backend",path="/api",set="1",status_code="401"} 38
http_requests_total{host=":8080",module="backend",path="/api",set="1",status_code="503"} 10
http_requests_total{host=":8080",module="backend",path="/auth",set="1",status_code="200"} 194
http_requests_total{host=":8080",module="backend",path="/auth",set="1",status_code="401"} 42
http_requests_total{host=":8080",module="backend",path="/auth",set="1",status_code="503"} 11
http_requests_total{host=":8080",module="backend",path="/home",set="1",status_code="200"} 201
http_requests_total{host=":8080",module="backend",path="/home",set="1",status_code="401"} 43
http_requests_total{host=":8080",module="backend",path="/home",set="1",status_code="503"} 3
http_requests_total{host=":8080",module="frontend",path="/api",set="1",status_code="200"} 196
http_requests_total{host=":8080",module="frontend",path="/api",set="1",status_code="401"} 37
http_requests_total{host=":8080",module="frontend",path="/api",set="1",status_code="503"} 14
http_requests_total{host=":8080",module="frontend",path="/auth",set="1",status_code="200"} 201
http_requests_total{host=":8080",module="frontend",path="/auth",set="1",status_code="401"} 32
http_requests_total{host=":8080",module="frontend",path="/auth",set="1",status_code="503"} 14
http_requests_total{host=":8080",module="frontend",path="/home",set="1",status_code="200"} 192
http_requests_total{host=":8080",module="frontend",path="/home",set="1",status_code="401"} 44
http_requests_total{host=":8080",module="frontend",path="/home",set="1",status_code="503"} 11
# HELP request_duration_seconds Time (in seconds) spent serving HTTP requests.
# TYPE request_duration_seconds histogram
request_duration_seconds_bucket{host=":8080",module="api",path="/api",set="1",status_code="200",le="0.005"} 0
request_duration_seconds_bucket{host=":8080",module="api",path="/api",set="1",status_code="200",le="0.01"} 0
request_duration_seconds_bucket{host=":8080",module="api",path="/api",set="1",status_code="200",le="0.025"} 5
request_duration_seconds_bucket{host=":8080",module="api",path="/api",set="1",status_code="200",le="0.05"} 20
request_duration_seconds_bucket{host=":8080",module="api",path="/api",set="1",status_code="200",le="0.1"} 68
request_duration_seconds_bucket{host=":8080",module="api",path="/api",set="1",status_code="200",le="0.25"} 162
request_duration_seconds_bucket{host=":8080",module="api",path="/api",set="1",status_code="200",le="0.5"} 162
request_duration_seconds_bucket{host=":8080",module="api",path="/api",set="1",status_code="200",le="1"} 207
request_duration_seconds_bucket{host=":8080",module="api",path="/api",set="1",status_code="200",le="2.5"} 207
request_duration_seconds_bucket{host=":8080",module="api",path="/api",set="1",status_code="200",le="5"} 207
request_duration_seconds_bucket{host=":8080",module="api",path="/api",set="1",status_code="200",le="10"} 207
request_duration_seconds_bucket{host=":8080",module="api",path="/api",set="1",status_code="200",le="+Inf"} 207
request_duration_seconds_sum{host=":8080",module="api",path="/api",set="1",status_code="200"} 50.84900000000004
request_duration_seconds_count{host=":8080",module="api",path="/api",set="1",status_code="200"} 207
# HELP response_size_bytes Response size in bytes.
# TYPE response_size_bytes summary
response_size_bytes{host=":8080",module="api",path="/api",set="1",status_code="200",quantile="0.5"} 554
response_size_bytes{host=":8080",module="api",path="/api",set="1",status_code="200",quantile="0.9"} 975
response_size_bytes{host=":8080",module="api",path="/api",set="1",status_code="200",quantile="0.99"} 1183
response_size_bytes_sum{host=":8080",module="api",path="/api",set="1",status_code="200"} 120890
response_size_bytes_count{host=":8080",module="api",path="/api",set="1",status_code="200"} 207
response_size_bytes{host=":8080",module="api",path="/api",set="1",status_code="401",quantile="0.5"} 545
response_size_bytes{host=":8080",module="api",path="/api",set="1",status_code="401",quantile="0.9"} 1178
response_size_bytes{host=":8080",module="api",path="/api",set="1",status_code="401",quantile="0.99"} 1187
response_size_bytes_sum{host=":8080",module="api",path="/api",set="1",status_code="401"} 17683
response_size_bytes_count{host=":8080",module="api",path="/api",set="1",status_code="401"} 28
response_size_bytes{host=":8080",module="api",path="/api",set="1",status_code="503",quantile="0.5"} 491
response_size_bytes{host=":8080",module="api",path="/api",set="1",status_code="503",quantile="0.9"} 1066
response_size_bytes{host=":8080",module="api",path="/api",set="1",status_code="503",quantile="0.99"} 1109
response_size_bytes_sum{host=":8080",module="api",path="/api",set="1",status_code="503"} 7095
response_size_bytes_count{host=":8080",module="api",path="/api",set="1",status_code="503"} 12
response_size_bytes{host=":8080",module="api",path="/auth",set="1",status_code="200",quantile="0.5"} 526
response_size_bytes{host=":8080",module="api",path="/auth",set="1",status_code="200",quantile="0.9"} 978
response_size_bytes{host=":8080",module="api",path="/auth",set="1",status_code="200",quantile="0.99"} 1171
response_size_bytes_sum{host=":8080",module="api",path="/auth",set="1",status_code="200"} 120432
response_size_bytes_count{host=":8080",module="api",path="/auth",set="1",status_code="200"} 205
response_size_bytes{host=":8080",module="api",path="/auth",set="1",status_code="401",quantile="0.5"} 528
response_size_bytes{host=":8080",module="api",path="/auth",set="1",status_code="401",quantile="0.9"} 981
response_size_bytes{host=":8080",module="api",path="/auth",set="1",status_code="401",quantile="0.99"} 1180
response_size_bytes_sum{host=":8080",module="api",path="/auth",set="1",status_code="401"} 19521
response_size_bytes_count{host=":8080",module="api",path="/auth",set="1",status_code="401"} 34
response_size_bytes{host=":8080",module="api",path="/auth",set="1",status_code="503",quantile="0.5"} 465
response_size_bytes{host=":8080",module="api",path="/auth",set="1",status_code="503",quantile="0.9"} 785
response_size_bytes{host=":8080",module="api",path="/auth",set="1",status_code="503",quantile="0.99"} 785
response_size_bytes_sum{host=":8080",module="api",path="/auth",set="1",status_code="503"} 3770
response_size_bytes_count{host=":8080",module="api",path="/auth",set="1",status_code="503"} 8
response_size_bytes{host=":8080",module="api",path="/home",set="1",status_code="200",quantile="0.5"} 556
response_size_bytes{host=":8080",module="api",path="/home",set="1",status_code="200",quantile="0.9"} 1027
response_size_bytes{host=":8080",module="api",path="/home",set="1",status_code="200",quantile="0.99"} 1182
response_size_bytes_sum{host=":8080",module="api",path="/home",set="1",status_code="200"} 121477
response_size_bytes_count{host=":8080",module="api",path="/home",set="1",status_code="200"} 200
response_size_bytes{host=":8080",module="api",path="/home",set="1",status_code="401",quantile="0.5"} 525
response_size_bytes{host=":8080",module="api",path="/home",set="1",status_code="401",quantile="0.9"} 1033
response_size_bytes{host=":8080",module="api",path="/home",set="1",status_code="401",quantile="0.99"} 1083
response_size_bytes_sum{host=":8080",module="api",path="/home",set="1",status_code="401"} 20749
response_size_bytes_count{host=":8080",module="api",path="/home",set="1",status_code="401"} 34
response_size_bytes{host=":8080",module="api",path="/home",set="1",status_code="503",quantile="0.5"} 707
response_size_bytes{host=":8080",module="api",path="/home",set="1",status_code="503",quantile="0.9"} 934
response_size_bytes{host=":8080",module="api",path="/home",set="1",status_code="503",quantile="0.99"} 956
```
