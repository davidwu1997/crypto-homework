# crypto-homework
---

# Local run

```
make deploy-local
```

---
# API
---
## Create currency

```
curl --location --request POST '127.0.0.1:8080/currency/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "currency_code": "rmb",
    "currency_name": "rmb"
}'

curl --location --request POST '127.0.0.1:8080/currency/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "currency_code": "usd",
    "currency_name": "usd"
}'
```

## Create user

```
curl --location --request POST '127.0.0.1:8080/member-api/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "login": "david",
    "pass_word": "111111",
    "name": "david"
}'

curl --location --request POST '127.0.0.1:8080/member-api/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "login": "crypto",
    "pass_word": "111111",
    "name": "crypto"
}'
```

## Get user wallet list

```
curl --location --request GET '127.0.0.1:8080/member-api/wallet/list?member_login=david'

curl --location --request GET '127.0.0.1:8080/member-api/wallet/list?member_login=crypto'
```

## User wallet deposit

```
curl --location --request POST '127.0.0.1:8080/member-api/wallet/deposit' \
--header 'Content-Type: application/json' \
--data-raw '{
    "wallet_id": "4473857808588588741",
    "transfer_amount": 100
}'
```

- wallet_id can get from `Get user wallet list API`

## User wallet withdraw

```
curl --location --request POST '127.0.0.1:8080/member-api/wallet/withdraw' \
--header 'Content-Type: application/json' \
--data-raw '{
    "wallet_id": "4473857808588588741",
    "transfer_amount": 20
}'
```

- wallet_id can get from `Get user wallet list API`


## User wallet transfer
```
curl --location --request POST '127.0.0.1:8080/member-api/wallet/transfer' \
--header 'Content-Type: application/json' \
--data-raw '{
    "from_wallet_id": "4473857808588588741",
    "transfer_amount": 20,
    "to_wallet_id": "2321150099253688741"
}'
```
- from_wallet_id can get from `Get user wallet list API`
- to_wallet_id can get from `Get user wallet list API`

## User wallet deposit record
```
curl --location --request GET '127.0.0.1:8080/member-api/report/deposit-log?member_login=david'
```
- member_login: member account

## User wallet withdraw record
```
curl --location --request GET '127.0.0.1:8080/member-api/report/withdraw-log?member_login=david'
```
- member_login: member account

## User wallet transfer record
```
curl --location --request GET '127.0.0.1:8080/member-api/report/transfer-log?member_login=david'
```
- member_login: member account

## User wallet transaction record
```
curl --location --request GET '127.0.0.1:8080/member-api/report/transaction-log?member_login=david'
```
- member_login: member account

---

# Optimize

- 完善登入驗證的 token 機制
- 將出款、入款、轉帳流程轉為透過 MQ
    - API -> service -> Nats streaming
    - Nats streaming -> service -> mysql
    - 透過持久化 MQ 的導入，進行流量消峰
    - 將服務設計為 cluster 模式，可水平擴充分散流量

---
