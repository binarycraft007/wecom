


# Wecom customer service api
Wecom customer service api
  

## Informations

### Version

1.0.0

### License

[MIT]()

### Contact

binarycraft  

### Terms Of Service

http://swagger.io/terms/

## Content negotiation

### URI Schemes
  * https

### Consumes
  * application/json
  * application/xml

### Produces
  * application/json
  * application/xml

## Access control

### Security Schemes

#### access_token (query: access_token)



> **Type**: apikey

## All endpoints

###  operations

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /cgi-bin/kf/account/add | [add customer service account](#add-customer-service-account) |  |
| POST | /cgi-bin/kf/account/del | [del customer service account](#del-customer-service-account) |  |
| GET | /cgi-bin/gettoken | [get token](#get-token) |  |
| POST | /cgi-bin/kf/account/update | [update customer service account](#update-customer-service-account) |  |
  


## Paths

### <span id="add-customer-service-account"></span> add customer service account (*addCustomerServiceAccount*)

```
POST /cgi-bin/kf/account/add
```

add customer service account

#### Produces
  * application/json

#### Security Requirements
  * access_token

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| AddCustomerServiceAccountRequest | `body` | [AddCustomerServiceAccountRequest](#add-customer-service-account-request) | `models.AddCustomerServiceAccountRequest` | | ✓ | | add customer service account request |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#add-customer-service-account-200) | OK | add customer service account response |  | [schema](#add-customer-service-account-200-schema) |
| [default](#add-customer-service-account-default) | | unexpected error |  | [schema](#add-customer-service-account-default-schema) |

#### Responses


##### <span id="add-customer-service-account-200"></span> 200 - add customer service account response
Status: OK

###### <span id="add-customer-service-account-200-schema"></span> Schema
   
  

[AddCustomerServiceAccountReponse](#add-customer-service-account-reponse)

##### <span id="add-customer-service-account-default"></span> Default Response
unexpected error

###### <span id="add-customer-service-account-default-schema"></span> Schema

  

[ErrorModel](#error-model)

### <span id="del-customer-service-account"></span> del customer service account (*delCustomerServiceAccount*)

```
POST /cgi-bin/kf/account/del
```

delete customer service account

#### Produces
  * application/json

#### Security Requirements
  * access_token

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| DelCustomerServiceAccountRequest | `body` | [DelCustomerServiceAccountRequest](#del-customer-service-account-request) | `models.DelCustomerServiceAccountRequest` | | ✓ | | delete customer service account request |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#del-customer-service-account-200) | OK | delete customer service account response |  | [schema](#del-customer-service-account-200-schema) |
| [default](#del-customer-service-account-default) | | unexpected error |  | [schema](#del-customer-service-account-default-schema) |

#### Responses


##### <span id="del-customer-service-account-200"></span> 200 - delete customer service account response
Status: OK

###### <span id="del-customer-service-account-200-schema"></span> Schema
   
  

[DelCustomerServiceAccountReponse](#del-customer-service-account-reponse)

##### <span id="del-customer-service-account-default"></span> Default Response
unexpected error

###### <span id="del-customer-service-account-default-schema"></span> Schema

  

[ErrorModel](#error-model)

### <span id="get-token"></span> get token (*getToken*)

```
GET /cgi-bin/gettoken
```

get access token using corpid and secret

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| corpid | `query` | string | `string` |  | ✓ |  | corporation id |
| corpsecret | `query` | string | `string` |  | ✓ |  | corporation secret |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-token-200) | OK | get access token response |  | [schema](#get-token-200-schema) |
| [default](#get-token-default) | | unexpected error |  | [schema](#get-token-default-schema) |

#### Responses


##### <span id="get-token-200"></span> 200 - get access token response
Status: OK

###### <span id="get-token-200-schema"></span> Schema
   
  

[GetAccessTokenReponse](#get-access-token-reponse)

##### <span id="get-token-default"></span> Default Response
unexpected error

###### <span id="get-token-default-schema"></span> Schema

  

[ErrorModel](#error-model)

### <span id="update-customer-service-account"></span> update customer service account (*updateCustomerServiceAccount*)

```
POST /cgi-bin/kf/account/update
```

update customer service account

#### Produces
  * application/json

#### Security Requirements
  * access_token

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| updateCustomerServiceAccountRequest | `body` | [UpdateCustomerServiceAccountRequest](#update-customer-service-account-request) | `models.UpdateCustomerServiceAccountRequest` | | ✓ | | update customer service account request |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-customer-service-account-200) | OK | update customer service account response |  | [schema](#update-customer-service-account-200-schema) |
| [default](#update-customer-service-account-default) | | unexpected error |  | [schema](#update-customer-service-account-default-schema) |

#### Responses


##### <span id="update-customer-service-account-200"></span> 200 - update customer service account response
Status: OK

###### <span id="update-customer-service-account-200-schema"></span> Schema
   
  

[UpdateCustomerServiceAccountReponse](#update-customer-service-account-reponse)

##### <span id="update-customer-service-account-default"></span> Default Response
unexpected error

###### <span id="update-customer-service-account-default-schema"></span> Schema

  

[ErrorModel](#error-model)

## Models

### <span id="add-customer-service-account-reponse"></span> AddCustomerServiceAccountReponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| errcode | int32 (formatted integer)| `int32` | ✓ | |  |  |
| errmsg | string| `string` | ✓ | |  |  |
| open_kfid | string| `string` |  | |  |  |



### <span id="add-customer-service-account-request"></span> AddCustomerServiceAccountRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| media_id | string| `string` | ✓ | |  |  |
| name | string| `string` | ✓ | |  |  |



### <span id="del-customer-service-account-reponse"></span> DelCustomerServiceAccountReponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| errcode | int32 (formatted integer)| `int32` | ✓ | |  |  |
| errmsg | string| `string` | ✓ | |  |  |



### <span id="del-customer-service-account-request"></span> DelCustomerServiceAccountRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| open_kfid | string| `string` | ✓ | |  |  |



### <span id="error-model"></span> ErrorModel


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| code | int32 (formatted integer)| `int32` | ✓ | |  |  |
| message | string| `string` | ✓ | |  |  |



### <span id="get-access-token-reponse"></span> GetAccessTokenReponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| access_token | string| `string` |  | |  |  |
| errcode | int32 (formatted integer)| `int32` | ✓ | |  |  |
| errmsg | string| `string` | ✓ | |  |  |
| expires_in | int32 (formatted integer)| `int32` |  | |  |  |



### <span id="update-customer-service-account-reponse"></span> UpdateCustomerServiceAccountReponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| errcode | int32 (formatted integer)| `int32` | ✓ | |  |  |
| errmsg | string| `string` | ✓ | |  |  |



### <span id="update-customer-service-account-request"></span> UpdateCustomerServiceAccountRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| media_id | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |
| open_kfid | string| `string` | ✓ | |  |  |


