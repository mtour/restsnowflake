# restsnowflake
restapi to get autoincrease id  (twitter snowflake)

# get id request
```
http://localhost:8080/snowflake/getid/[worker_id]
the range of worker_id is 0~1023
```

return data
```
{
  "code": 0,
  "message": "success !",
  "data": {
    "id": 907626740979949568
  }
}

```
