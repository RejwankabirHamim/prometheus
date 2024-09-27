# prometheus

```shell
docker build -t my-app-prometheus .
```

```shell
docker run -d \
  -p 8181:8181 \
  -p 9090:9090 \
  --name my-app-prometheus \
  my-app-prometheus
```

Monitoring
```shell
localhost:8181/ping
localhost:8181/metrics
http://localhost:9090

```

Run this command to execute 10 pings to the server. Then check the counter
increases on the metrics url. 
```shell
./ping_test.sh
```
