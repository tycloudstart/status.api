# status.api
api for falcon status

## introduction
api interface between falcon and tymeissa

## interface
```bash
# query one judge.status

request: HTTP.GET, args in url
curl -s "$httpprex/api/judge/status/$uuid/$endpoint/$counter"

result:ok
{
    "data": {
        "counter": "test.metric/t0=tag0,t1=tag1",
        "endpoint": "test.endpoint",
        "status": 0,
        "uuid": "e_1986"
    },
    "msg": "success"
}
result:error
{
    "data": "reason of error",
    "msg": "error"
}


# query multi judge.statuses

request: HTTP.POST, args in body
curl -s -X POST -d "[{\"uuid\":\"$u\",\"endpoint\":\"$e\",\"counter\":\"$c\"}]"  "$httpprex/api/judge/statuses"

result:ok
{
    "data": [
        {
            "counter": "test.metric/t0=tag0,t1=tag1",
            "endpoint": "test.endpoint",
            "status": 0,
            "uuid": "e_1986"
        },
        {
            "counter": "test.metric/t2=tag2",
            "endpoint": "test.endpoint",
            "status": 0,
            "uuid": "e_1987"
        }
    ],
    "msg": "success"
}
result:error
{
    "data": "reason of error",
    "msg": "error"
}

```

## install
```bash
# get src
git clone https://github.com/tycloudstart/status.api
cd status.api
go get ./...

# change config
mv cfg.example.json cfg.json
vim cfg.json
...

# build
./control build

# start
./control start

...

# stop
./control stop

```
## config
```json
{
    "http": {
        "enable": true,
        "listen": "0.0.0.0:19001"
    },
    "judge": {
        "enable": true,
        "cluster": {
            "judge-00" : "127.0.0.1:6080"
        }
    }
}
```

## debug
use ```./test/debug``` to debug this service

```bash
cd /home/to/status.api

# log
./test/debug tail

# get internal statistics
./test/debug/counter

```

