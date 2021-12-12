/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4950
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"time"
)

// IamSessionAllOf Definition of the list of properties defined in 'iam.Session', excluding properties defined in parent classes.
type IamSessionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType         string                  `json:"ObjectType"`
	AccountPermissions []IamAccountPermissions `json:"AccountPermissions,omitempty"`
	// Expiration time for the session.
	Expiration *time.Time `json:"Expiration,omitempty"`
	// Failed logins since last login for admin user.
	FailedLogins *int64 `json:"FailedLogins,omitempty"`
	// Idle time expiration for the session.
	IdleTimeExpiration *time.Time `json:"IdleTimeExpiration,omitempty"`
	// The client address from which last login is initiated.
	LastLoginClient *string `json:"LastLoginClient,omitempty"`
	// The last login time for user.
	LastLoginTime *time.Time `json:"LastLoginTime,omitempty"`
	// Session token shared with the user agent which is used to identify the user session when API requests are received to perform authorization.
	SessionId            *string                    `json:"SessionId,omitempty"`
	Permission           *IamPermissionRelationship `json:"Permission,omitempty"`
	User                 *IamUserRelationship       `json:"User,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamSessionAllOf IamSessionAllOf

// NewIamSessionAllOf instantiates a new IamSessionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamSessionAllOf(classId string, objectType string) *IamSessionAllOf {
	this := IamSessionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamSessionAllOfWithDefaults instantiates a new IamSessionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamSessionAllOfWithDefaults() *IamSessionAllOf {
	this := IamSessionAllOf{}
	var classId string = "iam.Session"
	this.ClassId = classId
	var objectType string = "iam.Session"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamSessionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamSessionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamSessionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamSessionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamSessionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamSessionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccountPermissions returns the AccountPermissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamSessionAllOf) GetAccountPermissions() []IamAccountPermissions {
	if o == nil {
		var ret []IamAccountPermissions
		return ret
	}
	return o.AccountPermissions
}

// GetAccountPermissionsOk returns a tuple with the AccountPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamSessionAllOf) GetAccountPermissionsOk() (*[]IamAccountPermissions, bool) {
	if o == nil || o.AccountPermissions == nil {
		return nil, false
	}
	return &o.AccountPermissions, true
}

// HasAccountPermissions returns a boolean if a field has been set.
func (o *IamSessionAllOf) HasAccountPermissions() bool {
	if o != nil && o.AccountPermissions != nil {
		return true
	}

	return false
}

// SetAccountPermissions gets a reference to the given []IamAccountPermissions and assigns it to the AccountPermissions field.
func (o *IamSessionAllOf) SetAccountPermissions(v []IamAccountPermissions) {
	o.AccountPermissions = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *IamSessionAllOf) GetExpiration() time.Time {
	if o == nil || o.Expiration == nil {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSessionAllOf) GetExpirationOk() (*time.Time, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *IamSessionAllOf) HasExpiration() bool {
	if o != nil && o.Expiration != nil {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *IamSessionAllOf) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetFailedLogins returns the FailedLogins field value if set, zero value otherwise.
func (o *IamSessionAllOf) GetFailedLogins() int64 {
	if o == nil || o.FailedLogins == nil {
		var ret int64
		return ret
	}
	return *o.FailedLogins
}

// GetFailedLoginsOk returns a tuple with the FailedLogins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSessionAllOf) GetFailedLoginsOk() (*int64, bool) {
	if o == nil || o.FailedLogins == nil {
		return nil, false
	}
	return o.FailedLogins, true
}

// HasFailedLogins returns a boolean if a field has been set.
func (o *IamSessionAllOf) HasFailedLogins() bool {
	if o != nil && o.FailedLogins != nil {
		return true
	}

	return false
}

// SetFailedLogins gets a reference to the given int64 and assigns it to the FailedLogins field.
func (o *IamSessionAllOf) SetFailedLogins(v int64) {
	o.FailedLogins = &v
}

// GetIdleTimeExpiration returns the IdleTimeExpiration field value if set, zero value otherwise.
func (o *IamSessionAllOf) GetIdleTimeExpiration() time.Time {
	if o == nil || o.IdleTimeExpiration == nil {
		var ret time.Time
		return ret
	}
	return *o.IdleTimeExpiration
}

// GetIdleTimeExpirationOk returns a tuple with the IdleTimeExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSessionAllOf) GetIdleTimeExpirationOk() (*time.Time, bool) {
	if o == nil || o.IdleTimeExpiration == nil {
		return nil, false
	}
	return o.IdleTimeExpiration, true
}

// HasIdleTimeExpiration returns a boolean if a field has been set.
func (o *IamSessionAllOf) HasIdleTimeExpiration() bool {
	if o != nil && o.IdleTimeExpiration != nil {
		return true
	}

	return false
}

// SetIdleTimeExpiration gets a reference to the given time.Time and assigns it to the IdleTimeExpiration field.
func (o *IamSessionAllOf) SetIdleTimeExpiration(v time.Time) {
	o.IdleTimeExpiration = &v
}

// GetLastLoginClient returns the LastLoginClient field value if set, zero value otherwise.
func (o *IamSessionAllOf) GetLastLoginClient() string {
	if o == nil || o.LastLoginClient == nil {
		var ret string
		return ret
	}
	return *o.LastLoginClient
}

// GetLastLoginClientOk returns a tuple with the LastLoginClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSessionAllOf) GetLastLoginClientOk() (*string, bool) {
	if o == nil || o.LastLoginClient == nil {
		return nil, false
	}
	return o.LastLoginClient, true
}

// HasLastLoginClient returns a boolean if a field has been set.
func (o *IamSessionAllOf) HasLastLoginClient() bool {
	if o != nil && o.LastLoginClient != nil {
		return true
	}

	return false
}

// SetLastLoginClient gets a reference to the given string and assigns it to the LastLoginClient field.
func (o *IamSessionAllOf) SetLastLoginClient(v string) {
	o.LastLoginClient = &v
}

// GetLastLoginTime returns the LastLoginTime field value if set, zero value otherwise.
func (o *IamSessionAllOf) GetLastLoginTime() time.Time {
	if o == nil || o.LastLoginTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastLoginTime
}

// GetLastLoginTimeOk returns a tuple with the LastLoginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSessionAllOf) GetLastLoginTimeOk() (*time.Time, bool) {
	if o == nil || o.LastLoginTime == nil {
		return nil, false
	}
	return o.LastLoginTime, true
}

// HasLastLoginTime returns a boolean if a field has been set.
func (o *IamSessionAllOf) HasLastLoginTime() bool {
	if o != nil && o.LastLoginTime != nil {
		return true
	}

	return false
}

// SetLastLoginTime gets a reference to the given time.Time and assigns it to the LastLoginTime field.
func (o *IamSessionAllOf) SetLastLoginTime(v time.Time) {
	o.LastLoginTime = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *IamSessionAllOf) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSessionAllOf) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *IamSessionAllOf) HasSessionId() bool {
	if o != nil && o.SessionId != nil {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *IamSessionAllOf) SetSessionId(v string) {
	o.SessionId = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *IamSessionAllOf) GetPermission() IamPermissionRelationship {
	if o == nil || o.Permission == nil {
		var ret IamPermissionRelationship
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSessionAllOf) GetPermissionOk() (*IamPermissionRelationship, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *IamSessionAllOf) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given IamPermissionRelationship and assigns it to the Permission field.
func (o *IamSessionAllOf) SetPermission(v IamPermissionRelationship) {
	o.Permission = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *IamSessionAllOf) GetUser() IamUserRelationship {
	if o == nil || o.User == nil {
		var ret IamUserRelationship
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSessionAllOf) GetUserOk() (*IamUserRelationship, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *IamSessionAllOf) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given IamUserRelationship and assigns it to the User field.
func (o *IamSessionAllOf) SetUser(v IamUserRelationship) {
	o.User = &v
}

func (o IamSessionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AccountPermissions != nil {
		toSerialize["AccountPermissions"] = o.AccountPermissions
	}
	if o.Expiration != nil {
		toSerialize["Expiration"] = o.Expiration
	}
	if o.FailedLogins != nil {
		toSerialize["FailedLogins"] = o.FailedLogins
	}
	if o.IdleTimeExpiration != nil {
		toSerialize["IdleTimeExpiration"] = o.IdleTimeExpiration
	}
	if o.LastLoginClient != nil {
		toSerialize["LastLoginClient"] = o.LastLoginClient
	}
	if o.LastLoginTime != nil {
		toSerialize["LastLoginTime"] = o.LastLoginTime
	}
	if o.SessionId != nil {
		toSerialize["SessionId"] = o.SessionId
	}
	if o.Permission != nil {
		toSerialize["Permission"] = o.Permission
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamSessionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamSessionAllOf := _IamSessionAllOf{}

	if err = json.Unmarshal(bytes, &varIamSessionAllOf); err == nil {
		*o = IamSessionAllOf(varIamSessionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccountPermissions")
		delete(additionalProperties, "Expiration")
		delete(additionalProperties, "FailedLogins")
		delete(additionalProperties, "IdleTimeExpiration")
		delete(additionalProperties, "LastLoginClient")
		delete(additionalProperties, "LastLoginTime")
		delete(additionalProperties, "SessionId")
		delete(additionalProperties, "Permission")
		delete(additionalProperties, "User")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamSessionAllOf struct {
	value *IamSessionAllOf
	isSet bool
}

func (v NullableIamSessionAllOf) Get() *IamSessionAllOf {
	return v.value
}

func (v *NullableIamSessionAllOf) Set(val *IamSessionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamSessionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamSessionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamSessionAllOf(val *IamSessionAllOf) *NullableIamSessionAllOf {
	return &NullableIamSessionAllOf{value: val, isSet: true}
}

func (v NullableIamSessionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamSessionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
