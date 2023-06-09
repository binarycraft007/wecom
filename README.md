


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

###  account

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /cgi-bin/kf/account/add | [add](#add) |  |
| POST | /cgi-bin/kf/account/del | [delete](#delete) |  |
| POST | /cgi-bin/kf/account/list | [list](#list) |  |
| POST | /cgi-bin/kf/account/update | [update](#update) |  |
  


###  token

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /cgi-bin/gettoken | [get](#get) |  |
  


## Paths

### <span id="add"></span> add (*add*)

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
| [200](#add-200) | OK | add customer service account response |  | [schema](#add-200-schema) |
| [default](#add-default) | | unexpected error |  | [schema](#add-default-schema) |

#### Responses


##### <span id="add-200"></span> 200 - add customer service account response
Status: OK

###### <span id="add-200-schema"></span> Schema
   
  

[AddCustomerServiceAccountReponse](#add-customer-service-account-reponse)

##### <span id="add-default"></span> Default Response
unexpected error

###### <span id="add-default-schema"></span> Schema

  

[ErrorModel](#error-model)

### <span id="delete"></span> delete (*delete*)

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
| [200](#delete-200) | OK | delete customer service account response |  | [schema](#delete-200-schema) |
| [default](#delete-default) | | unexpected error |  | [schema](#delete-default-schema) |

#### Responses


##### <span id="delete-200"></span> 200 - delete customer service account response
Status: OK

###### <span id="delete-200-schema"></span> Schema
   
  

[DelCustomerServiceAccountReponse](#del-customer-service-account-reponse)

##### <span id="delete-default"></span> Default Response
unexpected error

###### <span id="delete-default-schema"></span> Schema

  

[ErrorModel](#error-model)

### <span id="get"></span> get (*get*)

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
| [200](#get-200) | OK | get access token response |  | [schema](#get-200-schema) |
| [default](#get-default) | | unexpected error |  | [schema](#get-default-schema) |

#### Responses


##### <span id="get-200"></span> 200 - get access token response
Status: OK

###### <span id="get-200-schema"></span> Schema
   
  

[GetAccessTokenReponse](#get-access-token-reponse)

##### <span id="get-default"></span> Default Response
unexpected error

###### <span id="get-default-schema"></span> Schema

  

[ErrorModel](#error-model)

### <span id="list"></span> list (*list*)

```
POST /cgi-bin/kf/account/list
```

list customer service account

#### Produces
  * application/json

#### Security Requirements
  * access_token

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| listCustomerServiceAccountRequest | `body` | [ListCustomerServiceAccountRequest](#list-customer-service-account-request) | `models.ListCustomerServiceAccountRequest` | | ✓ | | list customer service account request |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-200) | OK | list customer service account response |  | [schema](#list-200-schema) |
| [default](#list-default) | | unexpected error |  | [schema](#list-default-schema) |

#### Responses


##### <span id="list-200"></span> 200 - list customer service account response
Status: OK

###### <span id="list-200-schema"></span> Schema
   
  

[ListCustomerServiceAccountReponse](#list-customer-service-account-reponse)

##### <span id="list-default"></span> Default Response
unexpected error

###### <span id="list-default-schema"></span> Schema

  

[ErrorModel](#error-model)

### <span id="update"></span> update (*update*)

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
| [200](#update-200) | OK | update customer service account response |  | [schema](#update-200-schema) |
| [default](#update-default) | | unexpected error |  | [schema](#update-default-schema) |

#### Responses


##### <span id="update-200"></span> 200 - update customer service account response
Status: OK

###### <span id="update-200-schema"></span> Schema
   
  

[UpdateCustomerServiceAccountReponse](#update-customer-service-account-reponse)

##### <span id="update-default"></span> Default Response
unexpected error

###### <span id="update-default-schema"></span> Schema

  

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



### <span id="customer-service-account"></span> CustomerServiceAccount


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| avatar | string| `string` | ✓ | |  |  |
| name | string| `string` | ✓ | |  |  |
| open_kfid | string| `string` | ✓ | |  |  |



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
| errcode | int32 (formatted integer)| `int32` | ✓ | |  |  |
| errmsg | string| `string` | ✓ | |  |  |



### <span id="get-access-token-reponse"></span> GetAccessTokenReponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| access_token | string| `string` |  | |  |  |
| errcode | int32 (formatted integer)| `int32` | ✓ | |  |  |
| errmsg | string| `string` | ✓ | |  |  |
| expires_in | int32 (formatted integer)| `int32` |  | |  |  |



### <span id="list-customer-service-account-reponse"></span> ListCustomerServiceAccountReponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| account_list | [][CustomerServiceAccount](#customer-service-account)| `[]*CustomerServiceAccount` |  | |  |  |
| errcode | int32 (formatted integer)| `int32` | ✓ | |  |  |
| errmsg | string| `string` | ✓ | |  |  |



### <span id="list-customer-service-account-request"></span> ListCustomerServiceAccountRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| limit | int32 (formatted integer)| `int32` |  | |  |  |
| offset | int32 (formatted integer)| `int32` |  | |  |  |



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


