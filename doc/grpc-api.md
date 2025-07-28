# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [store.proto](#store-proto)
    - [CreateDatabaseRequest](#-CreateDatabaseRequest)
    - [ListCollectionsRequest](#-ListCollectionsRequest)
    - [ListCollectionsResponse](#-ListCollectionsResponse)
    - [ListDatabasesRequest](#-ListDatabasesRequest)
    - [ListDatabasesResponse](#-ListDatabasesResponse)
    - [RemoveDatabaseRequest](#-RemoveDatabaseRequest)
    - [StatusResponse](#-StatusResponse)
  
    - [Status](#-Status)
  
    - [Store](#-Store)
  
- [health.proto](#health-proto)
    - [HealthCheckRequest](#-HealthCheckRequest)
    - [HealthCheckResponse](#-HealthCheckResponse)
  
    - [HealthCheckResponse.ServingStatus](#-HealthCheckResponse-ServingStatus)
  
    - [Health](#-Health)
  
- [metric.proto](#metric-proto)
    - [GetMetricRequest](#-GetMetricRequest)
    - [GetMetricResponse](#-GetMetricResponse)
    - [ListMetricRequest](#-ListMetricRequest)
    - [ListMetricResponse](#-ListMetricResponse)
  
    - [Metric](#-Metric)
  
- [config.proto](#config-proto)
    - [GetConfigRequest](#-GetConfigRequest)
    - [GetConfigResponse](#-GetConfigResponse)
    - [GetVersionRequest](#-GetVersionRequest)
    - [GetVersionResponse](#-GetVersionResponse)
    - [ListConfigRequest](#-ListConfigRequest)
    - [ListConfigResponse](#-ListConfigResponse)
  
    - [Config](#-Config)
  
- [Scalar Value Types](#scalar-value-types)



<a name="store-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## store.proto



<a name="-CreateDatabaseRequest"></a>

### CreateDatabaseRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database_name | [string](#string) |  |  |






<a name="-ListCollectionsRequest"></a>

### ListCollectionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database_name | [string](#string) |  |  |






<a name="-ListCollectionsResponse"></a>

### ListCollectionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collections | [string](#string) | repeated |  |






<a name="-ListDatabasesRequest"></a>

### ListDatabasesRequest







<a name="-ListDatabasesResponse"></a>

### ListDatabasesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| databases | [string](#string) | repeated |  |






<a name="-RemoveDatabaseRequest"></a>

### RemoveDatabaseRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database_name | [string](#string) |  |  |






<a name="-StatusResponse"></a>

### StatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#Status) |  |  |
| code | [int32](#int32) |  |  |
| message | [string](#string) |  |  |





 


<a name="-Status"></a>

### Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| OK | 0 |  |
| ERROR | 1 |  |


 

 


<a name="-Store"></a>

### Store


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateDatabase | [.CreateDatabaseRequest](#CreateDatabaseRequest) | [.StatusResponse](#StatusResponse) |  |
| RemoveDatabase | [.RemoveDatabaseRequest](#RemoveDatabaseRequest) | [.StatusResponse](#StatusResponse) |  |
| ListDatabases | [.ListDatabasesRequest](#ListDatabasesRequest) | [.ListDatabasesResponse](#ListDatabasesResponse) |  |
| ListCollections | [.ListCollectionsRequest](#ListCollectionsRequest) | [.ListCollectionsResponse](#ListCollectionsResponse) |  |

 



<a name="health-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## health.proto



<a name="-HealthCheckRequest"></a>

### HealthCheckRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service | [string](#string) |  |  |






<a name="-HealthCheckResponse"></a>

### HealthCheckResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [HealthCheckResponse.ServingStatus](#HealthCheckResponse-ServingStatus) |  |  |





 


<a name="-HealthCheckResponse-ServingStatus"></a>

### HealthCheckResponse.ServingStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 |  |
| SERVING | 1 |  |
| NOT_SERVING | 2 |  |
| SERVICE_UNKNOWN | 3 | Used only by the Watch method. |


 

 


<a name="-Health"></a>

### Health


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Check | [.HealthCheckRequest](#HealthCheckRequest) | [.HealthCheckResponse](#HealthCheckResponse) |  |

 



<a name="metric-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## metric.proto



<a name="-GetMetricRequest"></a>

### GetMetricRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="-GetMetricResponse"></a>

### GetMetricResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| names | [string](#string) | repeated |  |
| values | [string](#string) | repeated |  |






<a name="-ListMetricRequest"></a>

### ListMetricRequest







<a name="-ListMetricResponse"></a>

### ListMetricResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| values | [string](#string) | repeated |  |





 

 

 


<a name="-Metric"></a>

### Metric


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListMetric | [.ListMetricRequest](#ListMetricRequest) | [.ListMetricResponse](#ListMetricResponse) |  |
| GetMetric | [.GetMetricRequest](#GetMetricRequest) | [.GetMetricResponse](#GetMetricResponse) |  |

 



<a name="config-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## config.proto



<a name="-GetConfigRequest"></a>

### GetConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="-GetConfigResponse"></a>

### GetConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  |  |






<a name="-GetVersionRequest"></a>

### GetVersionRequest







<a name="-GetVersionResponse"></a>

### GetVersionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  |  |






<a name="-ListConfigRequest"></a>

### ListConfigRequest







<a name="-ListConfigResponse"></a>

### ListConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| values | [string](#string) | repeated |  |





 

 

 


<a name="-Config"></a>

### Config


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListConfig | [.ListConfigRequest](#ListConfigRequest) | [.ListConfigResponse](#ListConfigResponse) |  |
| GetConfig | [.GetConfigRequest](#GetConfigRequest) | [.GetConfigResponse](#GetConfigResponse) |  |
| GetVersion | [.GetVersionRequest](#GetVersionRequest) | [.GetVersionResponse](#GetVersionResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

