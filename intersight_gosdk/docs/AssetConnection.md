# AssetConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Credential** | Pointer to [**NullableMoBaseComplexType**](MoBaseComplexType.md) | A credential which can be used to authenticate with a managed target. | [optional] 

## Methods

### NewAssetConnection

`func NewAssetConnection(classId string, objectType string, ) *AssetConnection`

NewAssetConnection instantiates a new AssetConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetConnectionWithDefaults

`func NewAssetConnectionWithDefaults() *AssetConnection`

NewAssetConnectionWithDefaults instantiates a new AssetConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetConnection) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetConnection) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetConnection) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetConnection) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetConnection) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetConnection) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCredential

`func (o *AssetConnection) GetCredential() MoBaseComplexType`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *AssetConnection) GetCredentialOk() (*MoBaseComplexType, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *AssetConnection) SetCredential(v MoBaseComplexType)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *AssetConnection) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### SetCredentialNil

`func (o *AssetConnection) SetCredentialNil(b bool)`

 SetCredentialNil sets the value for Credential to be an explicit nil

### UnsetCredential
`func (o *AssetConnection) UnsetCredential()`

UnsetCredential ensures that no value is present for Credential, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


