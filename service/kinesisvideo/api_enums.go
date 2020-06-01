// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideo

type APIName string

// Enum values for APIName
const (
	APINamePutMedia                   APIName = "PUT_MEDIA"
	APINameGetMedia                   APIName = "GET_MEDIA"
	APINameListFragments              APIName = "LIST_FRAGMENTS"
	APINameGetMediaForFragmentList    APIName = "GET_MEDIA_FOR_FRAGMENT_LIST"
	APINameGetHlsStreamingSessionUrl  APIName = "GET_HLS_STREAMING_SESSION_URL"
	APINameGetDashStreamingSessionUrl APIName = "GET_DASH_STREAMING_SESSION_URL"
	APINameGetClip                    APIName = "GET_CLIP"
)

func (enum APIName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum APIName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ChannelProtocol string

// Enum values for ChannelProtocol
const (
	ChannelProtocolWss   ChannelProtocol = "WSS"
	ChannelProtocolHttps ChannelProtocol = "HTTPS"
)

func (enum ChannelProtocol) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChannelProtocol) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ChannelRole string

// Enum values for ChannelRole
const (
	ChannelRoleMaster ChannelRole = "MASTER"
	ChannelRoleViewer ChannelRole = "VIEWER"
)

func (enum ChannelRole) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChannelRole) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ChannelType string

// Enum values for ChannelType
const (
	ChannelTypeSingleMaster ChannelType = "SINGLE_MASTER"
)

func (enum ChannelType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChannelType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ComparisonOperator string

// Enum values for ComparisonOperator
const (
	ComparisonOperatorBeginsWith ComparisonOperator = "BEGINS_WITH"
)

func (enum ComparisonOperator) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ComparisonOperator) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Status string

// Enum values for Status
const (
	StatusCreating Status = "CREATING"
	StatusActive   Status = "ACTIVE"
	StatusUpdating Status = "UPDATING"
	StatusDeleting Status = "DELETING"
)

func (enum Status) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Status) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type UpdateDataRetentionOperation string

// Enum values for UpdateDataRetentionOperation
const (
	UpdateDataRetentionOperationIncreaseDataRetention UpdateDataRetentionOperation = "INCREASE_DATA_RETENTION"
	UpdateDataRetentionOperationDecreaseDataRetention UpdateDataRetentionOperation = "DECREASE_DATA_RETENTION"
)

func (enum UpdateDataRetentionOperation) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UpdateDataRetentionOperation) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
