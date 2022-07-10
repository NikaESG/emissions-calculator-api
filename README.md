# no-code-db-customized


Docker:

```
docker build -t xxx:xx .
docker tag xxxx xxxxx
docker push xxxx
```

1. Get Data

```
http://localhost:8081/api/v1/nocodedb/get_data

{
    "table_name": "Projects",
    "field_name":"CO2E"
}

Response:

{
    "Ret": 0,
    "Msg": "Get Data successfully.",
    "data": [
        "220.0000000000",
        "110.0000000000",
        "330.0000000000"
    ]
}
```