# CURRENCY APIs

# APIs Contract
APIs contract for currency-api service

## Create new currency
### Request
`POST /api/v1/currency/create/`

```shell
curl --location --request POST 'http://localhost:8081/api/v1/currency/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id":1,
    "name":"test"
}'
```

#### Payload
- id (int, required)
- name (string, required)

### Response
```shell
# Success
{
    "status": "success",
    "code": 200,
    "data": "",
    "message": "success"
}

# Bad request
{
    "status": "error",
    "code": 400,
    "message": " Name is required"
}
```
