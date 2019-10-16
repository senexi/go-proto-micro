# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [partners.proto](#partners.proto)
    - [AddPartnerReply](#partners.AddPartnerReply)
    - [Partner](#partners.Partner)
    - [Partner.PhoneNumber](#partners.Partner.PhoneNumber)
    - [PartnerList](#partners.PartnerList)
    - [SearchRequest](#partners.SearchRequest)
  
    - [Partner.PhoneType](#partners.Partner.PhoneType)
  
  
    - [PartnerService](#partners.PartnerService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="partners.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## partners.proto
go:generate protoc -I . -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf partners.proto --gogoslick_out=plugins=grpc:../partners


<a name="partners.AddPartnerReply"></a>

### AddPartnerReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="partners.Partner"></a>

### Partner



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| id | [int32](#int32) |  |  |
| email | [string](#string) |  |  |
| phone | [Partner.PhoneNumber](#partners.Partner.PhoneNumber) | repeated |  |






<a name="partners.Partner.PhoneNumber"></a>

### Partner.PhoneNumber



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| number | [string](#string) |  |  |
| type | [Partner.PhoneType](#partners.Partner.PhoneType) |  |  |






<a name="partners.PartnerList"></a>

### PartnerList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| partners | [Partner](#partners.Partner) | repeated |  |






<a name="partners.SearchRequest"></a>

### SearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| search_string | [string](#string) |  |  |





 


<a name="partners.Partner.PhoneType"></a>

### Partner.PhoneType


| Name | Number | Description |
| ---- | ------ | ----------- |
| MOBILE | 0 |  |
| HOME | 1 |  |
| WORK | 2 |  |


 

 


<a name="partners.PartnerService"></a>

### PartnerService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddPartner | [Partner](#partners.Partner) | [AddPartnerReply](#partners.AddPartnerReply) |  |
| GetPartners | [SearchRequest](#partners.SearchRequest) | [PartnerList](#partners.PartnerList) |  |
| SearchPartner | [SearchRequest](#partners.SearchRequest) | [PartnerList](#partners.PartnerList) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

