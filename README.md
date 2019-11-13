MQTT to SSE
===

connect a MQTT stream to a SSE request

```
$ go build
$ mqttosse.go
[client]   Connect()
[store]    memorystore initialized
[client]   about to write new connect msg
[client]   socket connected to broker
[client]   Using MQTT 3.1.1 protocol
...
---

$ curl localhost:9000/events
event: workgroup/6346966c27dbfba6d7f81c006527c968/air/temperature
data: {"temperature":15.3}

event: workgroup/6346966c27dbfba6d7f81c006527c968/air/humidity
data: {"humidity":52.4}
...
```

### reference

- https://github.com/alexandrevicenzi/go-sse/blob/master/examples/simple.go
- https://github.com/eclipse/paho.mqtt.golang/blob/master/cmd/sample/main.go
- https://github.com/jw3/mqttosse
