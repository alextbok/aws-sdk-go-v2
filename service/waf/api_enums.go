// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

type ChangeAction string

// Enum values for ChangeAction
const (
	ChangeActionInsert ChangeAction = "INSERT"
	ChangeActionDelete ChangeAction = "DELETE"
)

func (enum ChangeAction) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChangeAction) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ChangeTokenStatus string

// Enum values for ChangeTokenStatus
const (
	ChangeTokenStatusProvisioned ChangeTokenStatus = "PROVISIONED"
	ChangeTokenStatusPending     ChangeTokenStatus = "PENDING"
	ChangeTokenStatusInsync      ChangeTokenStatus = "INSYNC"
)

func (enum ChangeTokenStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChangeTokenStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ComparisonOperator string

// Enum values for ComparisonOperator
const (
	ComparisonOperatorEq ComparisonOperator = "EQ"
	ComparisonOperatorNe ComparisonOperator = "NE"
	ComparisonOperatorLe ComparisonOperator = "LE"
	ComparisonOperatorLt ComparisonOperator = "LT"
	ComparisonOperatorGe ComparisonOperator = "GE"
	ComparisonOperatorGt ComparisonOperator = "GT"
)

func (enum ComparisonOperator) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ComparisonOperator) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type GeoMatchConstraintType string

// Enum values for GeoMatchConstraintType
const (
	GeoMatchConstraintTypeCountry GeoMatchConstraintType = "Country"
)

func (enum GeoMatchConstraintType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum GeoMatchConstraintType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type GeoMatchConstraintValue string

// Enum values for GeoMatchConstraintValue
const (
	GeoMatchConstraintValueAf GeoMatchConstraintValue = "AF"
	GeoMatchConstraintValueAx GeoMatchConstraintValue = "AX"
	GeoMatchConstraintValueAl GeoMatchConstraintValue = "AL"
	GeoMatchConstraintValueDz GeoMatchConstraintValue = "DZ"
	GeoMatchConstraintValueAs GeoMatchConstraintValue = "AS"
	GeoMatchConstraintValueAd GeoMatchConstraintValue = "AD"
	GeoMatchConstraintValueAo GeoMatchConstraintValue = "AO"
	GeoMatchConstraintValueAi GeoMatchConstraintValue = "AI"
	GeoMatchConstraintValueAq GeoMatchConstraintValue = "AQ"
	GeoMatchConstraintValueAg GeoMatchConstraintValue = "AG"
	GeoMatchConstraintValueAr GeoMatchConstraintValue = "AR"
	GeoMatchConstraintValueAm GeoMatchConstraintValue = "AM"
	GeoMatchConstraintValueAw GeoMatchConstraintValue = "AW"
	GeoMatchConstraintValueAu GeoMatchConstraintValue = "AU"
	GeoMatchConstraintValueAt GeoMatchConstraintValue = "AT"
	GeoMatchConstraintValueAz GeoMatchConstraintValue = "AZ"
	GeoMatchConstraintValueBs GeoMatchConstraintValue = "BS"
	GeoMatchConstraintValueBh GeoMatchConstraintValue = "BH"
	GeoMatchConstraintValueBd GeoMatchConstraintValue = "BD"
	GeoMatchConstraintValueBb GeoMatchConstraintValue = "BB"
	GeoMatchConstraintValueBy GeoMatchConstraintValue = "BY"
	GeoMatchConstraintValueBe GeoMatchConstraintValue = "BE"
	GeoMatchConstraintValueBz GeoMatchConstraintValue = "BZ"
	GeoMatchConstraintValueBj GeoMatchConstraintValue = "BJ"
	GeoMatchConstraintValueBm GeoMatchConstraintValue = "BM"
	GeoMatchConstraintValueBt GeoMatchConstraintValue = "BT"
	GeoMatchConstraintValueBo GeoMatchConstraintValue = "BO"
	GeoMatchConstraintValueBq GeoMatchConstraintValue = "BQ"
	GeoMatchConstraintValueBa GeoMatchConstraintValue = "BA"
	GeoMatchConstraintValueBw GeoMatchConstraintValue = "BW"
	GeoMatchConstraintValueBv GeoMatchConstraintValue = "BV"
	GeoMatchConstraintValueBr GeoMatchConstraintValue = "BR"
	GeoMatchConstraintValueIo GeoMatchConstraintValue = "IO"
	GeoMatchConstraintValueBn GeoMatchConstraintValue = "BN"
	GeoMatchConstraintValueBg GeoMatchConstraintValue = "BG"
	GeoMatchConstraintValueBf GeoMatchConstraintValue = "BF"
	GeoMatchConstraintValueBi GeoMatchConstraintValue = "BI"
	GeoMatchConstraintValueKh GeoMatchConstraintValue = "KH"
	GeoMatchConstraintValueCm GeoMatchConstraintValue = "CM"
	GeoMatchConstraintValueCa GeoMatchConstraintValue = "CA"
	GeoMatchConstraintValueCv GeoMatchConstraintValue = "CV"
	GeoMatchConstraintValueKy GeoMatchConstraintValue = "KY"
	GeoMatchConstraintValueCf GeoMatchConstraintValue = "CF"
	GeoMatchConstraintValueTd GeoMatchConstraintValue = "TD"
	GeoMatchConstraintValueCl GeoMatchConstraintValue = "CL"
	GeoMatchConstraintValueCn GeoMatchConstraintValue = "CN"
	GeoMatchConstraintValueCx GeoMatchConstraintValue = "CX"
	GeoMatchConstraintValueCc GeoMatchConstraintValue = "CC"
	GeoMatchConstraintValueCo GeoMatchConstraintValue = "CO"
	GeoMatchConstraintValueKm GeoMatchConstraintValue = "KM"
	GeoMatchConstraintValueCg GeoMatchConstraintValue = "CG"
	GeoMatchConstraintValueCd GeoMatchConstraintValue = "CD"
	GeoMatchConstraintValueCk GeoMatchConstraintValue = "CK"
	GeoMatchConstraintValueCr GeoMatchConstraintValue = "CR"
	GeoMatchConstraintValueCi GeoMatchConstraintValue = "CI"
	GeoMatchConstraintValueHr GeoMatchConstraintValue = "HR"
	GeoMatchConstraintValueCu GeoMatchConstraintValue = "CU"
	GeoMatchConstraintValueCw GeoMatchConstraintValue = "CW"
	GeoMatchConstraintValueCy GeoMatchConstraintValue = "CY"
	GeoMatchConstraintValueCz GeoMatchConstraintValue = "CZ"
	GeoMatchConstraintValueDk GeoMatchConstraintValue = "DK"
	GeoMatchConstraintValueDj GeoMatchConstraintValue = "DJ"
	GeoMatchConstraintValueDm GeoMatchConstraintValue = "DM"
	GeoMatchConstraintValueDo GeoMatchConstraintValue = "DO"
	GeoMatchConstraintValueEc GeoMatchConstraintValue = "EC"
	GeoMatchConstraintValueEg GeoMatchConstraintValue = "EG"
	GeoMatchConstraintValueSv GeoMatchConstraintValue = "SV"
	GeoMatchConstraintValueGq GeoMatchConstraintValue = "GQ"
	GeoMatchConstraintValueEr GeoMatchConstraintValue = "ER"
	GeoMatchConstraintValueEe GeoMatchConstraintValue = "EE"
	GeoMatchConstraintValueEt GeoMatchConstraintValue = "ET"
	GeoMatchConstraintValueFk GeoMatchConstraintValue = "FK"
	GeoMatchConstraintValueFo GeoMatchConstraintValue = "FO"
	GeoMatchConstraintValueFj GeoMatchConstraintValue = "FJ"
	GeoMatchConstraintValueFi GeoMatchConstraintValue = "FI"
	GeoMatchConstraintValueFr GeoMatchConstraintValue = "FR"
	GeoMatchConstraintValueGf GeoMatchConstraintValue = "GF"
	GeoMatchConstraintValuePf GeoMatchConstraintValue = "PF"
	GeoMatchConstraintValueTf GeoMatchConstraintValue = "TF"
	GeoMatchConstraintValueGa GeoMatchConstraintValue = "GA"
	GeoMatchConstraintValueGm GeoMatchConstraintValue = "GM"
	GeoMatchConstraintValueGe GeoMatchConstraintValue = "GE"
	GeoMatchConstraintValueDe GeoMatchConstraintValue = "DE"
	GeoMatchConstraintValueGh GeoMatchConstraintValue = "GH"
	GeoMatchConstraintValueGi GeoMatchConstraintValue = "GI"
	GeoMatchConstraintValueGr GeoMatchConstraintValue = "GR"
	GeoMatchConstraintValueGl GeoMatchConstraintValue = "GL"
	GeoMatchConstraintValueGd GeoMatchConstraintValue = "GD"
	GeoMatchConstraintValueGp GeoMatchConstraintValue = "GP"
	GeoMatchConstraintValueGu GeoMatchConstraintValue = "GU"
	GeoMatchConstraintValueGt GeoMatchConstraintValue = "GT"
	GeoMatchConstraintValueGg GeoMatchConstraintValue = "GG"
	GeoMatchConstraintValueGn GeoMatchConstraintValue = "GN"
	GeoMatchConstraintValueGw GeoMatchConstraintValue = "GW"
	GeoMatchConstraintValueGy GeoMatchConstraintValue = "GY"
	GeoMatchConstraintValueHt GeoMatchConstraintValue = "HT"
	GeoMatchConstraintValueHm GeoMatchConstraintValue = "HM"
	GeoMatchConstraintValueVa GeoMatchConstraintValue = "VA"
	GeoMatchConstraintValueHn GeoMatchConstraintValue = "HN"
	GeoMatchConstraintValueHk GeoMatchConstraintValue = "HK"
	GeoMatchConstraintValueHu GeoMatchConstraintValue = "HU"
	GeoMatchConstraintValueIs GeoMatchConstraintValue = "IS"
	GeoMatchConstraintValueIn GeoMatchConstraintValue = "IN"
	GeoMatchConstraintValueId GeoMatchConstraintValue = "ID"
	GeoMatchConstraintValueIr GeoMatchConstraintValue = "IR"
	GeoMatchConstraintValueIq GeoMatchConstraintValue = "IQ"
	GeoMatchConstraintValueIe GeoMatchConstraintValue = "IE"
	GeoMatchConstraintValueIm GeoMatchConstraintValue = "IM"
	GeoMatchConstraintValueIl GeoMatchConstraintValue = "IL"
	GeoMatchConstraintValueIt GeoMatchConstraintValue = "IT"
	GeoMatchConstraintValueJm GeoMatchConstraintValue = "JM"
	GeoMatchConstraintValueJp GeoMatchConstraintValue = "JP"
	GeoMatchConstraintValueJe GeoMatchConstraintValue = "JE"
	GeoMatchConstraintValueJo GeoMatchConstraintValue = "JO"
	GeoMatchConstraintValueKz GeoMatchConstraintValue = "KZ"
	GeoMatchConstraintValueKe GeoMatchConstraintValue = "KE"
	GeoMatchConstraintValueKi GeoMatchConstraintValue = "KI"
	GeoMatchConstraintValueKp GeoMatchConstraintValue = "KP"
	GeoMatchConstraintValueKr GeoMatchConstraintValue = "KR"
	GeoMatchConstraintValueKw GeoMatchConstraintValue = "KW"
	GeoMatchConstraintValueKg GeoMatchConstraintValue = "KG"
	GeoMatchConstraintValueLa GeoMatchConstraintValue = "LA"
	GeoMatchConstraintValueLv GeoMatchConstraintValue = "LV"
	GeoMatchConstraintValueLb GeoMatchConstraintValue = "LB"
	GeoMatchConstraintValueLs GeoMatchConstraintValue = "LS"
	GeoMatchConstraintValueLr GeoMatchConstraintValue = "LR"
	GeoMatchConstraintValueLy GeoMatchConstraintValue = "LY"
	GeoMatchConstraintValueLi GeoMatchConstraintValue = "LI"
	GeoMatchConstraintValueLt GeoMatchConstraintValue = "LT"
	GeoMatchConstraintValueLu GeoMatchConstraintValue = "LU"
	GeoMatchConstraintValueMo GeoMatchConstraintValue = "MO"
	GeoMatchConstraintValueMk GeoMatchConstraintValue = "MK"
	GeoMatchConstraintValueMg GeoMatchConstraintValue = "MG"
	GeoMatchConstraintValueMw GeoMatchConstraintValue = "MW"
	GeoMatchConstraintValueMy GeoMatchConstraintValue = "MY"
	GeoMatchConstraintValueMv GeoMatchConstraintValue = "MV"
	GeoMatchConstraintValueMl GeoMatchConstraintValue = "ML"
	GeoMatchConstraintValueMt GeoMatchConstraintValue = "MT"
	GeoMatchConstraintValueMh GeoMatchConstraintValue = "MH"
	GeoMatchConstraintValueMq GeoMatchConstraintValue = "MQ"
	GeoMatchConstraintValueMr GeoMatchConstraintValue = "MR"
	GeoMatchConstraintValueMu GeoMatchConstraintValue = "MU"
	GeoMatchConstraintValueYt GeoMatchConstraintValue = "YT"
	GeoMatchConstraintValueMx GeoMatchConstraintValue = "MX"
	GeoMatchConstraintValueFm GeoMatchConstraintValue = "FM"
	GeoMatchConstraintValueMd GeoMatchConstraintValue = "MD"
	GeoMatchConstraintValueMc GeoMatchConstraintValue = "MC"
	GeoMatchConstraintValueMn GeoMatchConstraintValue = "MN"
	GeoMatchConstraintValueMe GeoMatchConstraintValue = "ME"
	GeoMatchConstraintValueMs GeoMatchConstraintValue = "MS"
	GeoMatchConstraintValueMa GeoMatchConstraintValue = "MA"
	GeoMatchConstraintValueMz GeoMatchConstraintValue = "MZ"
	GeoMatchConstraintValueMm GeoMatchConstraintValue = "MM"
	GeoMatchConstraintValueNa GeoMatchConstraintValue = "NA"
	GeoMatchConstraintValueNr GeoMatchConstraintValue = "NR"
	GeoMatchConstraintValueNp GeoMatchConstraintValue = "NP"
	GeoMatchConstraintValueNl GeoMatchConstraintValue = "NL"
	GeoMatchConstraintValueNc GeoMatchConstraintValue = "NC"
	GeoMatchConstraintValueNz GeoMatchConstraintValue = "NZ"
	GeoMatchConstraintValueNi GeoMatchConstraintValue = "NI"
	GeoMatchConstraintValueNe GeoMatchConstraintValue = "NE"
	GeoMatchConstraintValueNg GeoMatchConstraintValue = "NG"
	GeoMatchConstraintValueNu GeoMatchConstraintValue = "NU"
	GeoMatchConstraintValueNf GeoMatchConstraintValue = "NF"
	GeoMatchConstraintValueMp GeoMatchConstraintValue = "MP"
	GeoMatchConstraintValueNo GeoMatchConstraintValue = "NO"
	GeoMatchConstraintValueOm GeoMatchConstraintValue = "OM"
	GeoMatchConstraintValuePk GeoMatchConstraintValue = "PK"
	GeoMatchConstraintValuePw GeoMatchConstraintValue = "PW"
	GeoMatchConstraintValuePs GeoMatchConstraintValue = "PS"
	GeoMatchConstraintValuePa GeoMatchConstraintValue = "PA"
	GeoMatchConstraintValuePg GeoMatchConstraintValue = "PG"
	GeoMatchConstraintValuePy GeoMatchConstraintValue = "PY"
	GeoMatchConstraintValuePe GeoMatchConstraintValue = "PE"
	GeoMatchConstraintValuePh GeoMatchConstraintValue = "PH"
	GeoMatchConstraintValuePn GeoMatchConstraintValue = "PN"
	GeoMatchConstraintValuePl GeoMatchConstraintValue = "PL"
	GeoMatchConstraintValuePt GeoMatchConstraintValue = "PT"
	GeoMatchConstraintValuePr GeoMatchConstraintValue = "PR"
	GeoMatchConstraintValueQa GeoMatchConstraintValue = "QA"
	GeoMatchConstraintValueRe GeoMatchConstraintValue = "RE"
	GeoMatchConstraintValueRo GeoMatchConstraintValue = "RO"
	GeoMatchConstraintValueRu GeoMatchConstraintValue = "RU"
	GeoMatchConstraintValueRw GeoMatchConstraintValue = "RW"
	GeoMatchConstraintValueBl GeoMatchConstraintValue = "BL"
	GeoMatchConstraintValueSh GeoMatchConstraintValue = "SH"
	GeoMatchConstraintValueKn GeoMatchConstraintValue = "KN"
	GeoMatchConstraintValueLc GeoMatchConstraintValue = "LC"
	GeoMatchConstraintValueMf GeoMatchConstraintValue = "MF"
	GeoMatchConstraintValuePm GeoMatchConstraintValue = "PM"
	GeoMatchConstraintValueVc GeoMatchConstraintValue = "VC"
	GeoMatchConstraintValueWs GeoMatchConstraintValue = "WS"
	GeoMatchConstraintValueSm GeoMatchConstraintValue = "SM"
	GeoMatchConstraintValueSt GeoMatchConstraintValue = "ST"
	GeoMatchConstraintValueSa GeoMatchConstraintValue = "SA"
	GeoMatchConstraintValueSn GeoMatchConstraintValue = "SN"
	GeoMatchConstraintValueRs GeoMatchConstraintValue = "RS"
	GeoMatchConstraintValueSc GeoMatchConstraintValue = "SC"
	GeoMatchConstraintValueSl GeoMatchConstraintValue = "SL"
	GeoMatchConstraintValueSg GeoMatchConstraintValue = "SG"
	GeoMatchConstraintValueSx GeoMatchConstraintValue = "SX"
	GeoMatchConstraintValueSk GeoMatchConstraintValue = "SK"
	GeoMatchConstraintValueSi GeoMatchConstraintValue = "SI"
	GeoMatchConstraintValueSb GeoMatchConstraintValue = "SB"
	GeoMatchConstraintValueSo GeoMatchConstraintValue = "SO"
	GeoMatchConstraintValueZa GeoMatchConstraintValue = "ZA"
	GeoMatchConstraintValueGs GeoMatchConstraintValue = "GS"
	GeoMatchConstraintValueSs GeoMatchConstraintValue = "SS"
	GeoMatchConstraintValueEs GeoMatchConstraintValue = "ES"
	GeoMatchConstraintValueLk GeoMatchConstraintValue = "LK"
	GeoMatchConstraintValueSd GeoMatchConstraintValue = "SD"
	GeoMatchConstraintValueSr GeoMatchConstraintValue = "SR"
	GeoMatchConstraintValueSj GeoMatchConstraintValue = "SJ"
	GeoMatchConstraintValueSz GeoMatchConstraintValue = "SZ"
	GeoMatchConstraintValueSe GeoMatchConstraintValue = "SE"
	GeoMatchConstraintValueCh GeoMatchConstraintValue = "CH"
	GeoMatchConstraintValueSy GeoMatchConstraintValue = "SY"
	GeoMatchConstraintValueTw GeoMatchConstraintValue = "TW"
	GeoMatchConstraintValueTj GeoMatchConstraintValue = "TJ"
	GeoMatchConstraintValueTz GeoMatchConstraintValue = "TZ"
	GeoMatchConstraintValueTh GeoMatchConstraintValue = "TH"
	GeoMatchConstraintValueTl GeoMatchConstraintValue = "TL"
	GeoMatchConstraintValueTg GeoMatchConstraintValue = "TG"
	GeoMatchConstraintValueTk GeoMatchConstraintValue = "TK"
	GeoMatchConstraintValueTo GeoMatchConstraintValue = "TO"
	GeoMatchConstraintValueTt GeoMatchConstraintValue = "TT"
	GeoMatchConstraintValueTn GeoMatchConstraintValue = "TN"
	GeoMatchConstraintValueTr GeoMatchConstraintValue = "TR"
	GeoMatchConstraintValueTm GeoMatchConstraintValue = "TM"
	GeoMatchConstraintValueTc GeoMatchConstraintValue = "TC"
	GeoMatchConstraintValueTv GeoMatchConstraintValue = "TV"
	GeoMatchConstraintValueUg GeoMatchConstraintValue = "UG"
	GeoMatchConstraintValueUa GeoMatchConstraintValue = "UA"
	GeoMatchConstraintValueAe GeoMatchConstraintValue = "AE"
	GeoMatchConstraintValueGb GeoMatchConstraintValue = "GB"
	GeoMatchConstraintValueUs GeoMatchConstraintValue = "US"
	GeoMatchConstraintValueUm GeoMatchConstraintValue = "UM"
	GeoMatchConstraintValueUy GeoMatchConstraintValue = "UY"
	GeoMatchConstraintValueUz GeoMatchConstraintValue = "UZ"
	GeoMatchConstraintValueVu GeoMatchConstraintValue = "VU"
	GeoMatchConstraintValueVe GeoMatchConstraintValue = "VE"
	GeoMatchConstraintValueVn GeoMatchConstraintValue = "VN"
	GeoMatchConstraintValueVg GeoMatchConstraintValue = "VG"
	GeoMatchConstraintValueVi GeoMatchConstraintValue = "VI"
	GeoMatchConstraintValueWf GeoMatchConstraintValue = "WF"
	GeoMatchConstraintValueEh GeoMatchConstraintValue = "EH"
	GeoMatchConstraintValueYe GeoMatchConstraintValue = "YE"
	GeoMatchConstraintValueZm GeoMatchConstraintValue = "ZM"
	GeoMatchConstraintValueZw GeoMatchConstraintValue = "ZW"
)

func (enum GeoMatchConstraintValue) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum GeoMatchConstraintValue) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type IPSetDescriptorType string

// Enum values for IPSetDescriptorType
const (
	IPSetDescriptorTypeIpv4 IPSetDescriptorType = "IPV4"
	IPSetDescriptorTypeIpv6 IPSetDescriptorType = "IPV6"
)

func (enum IPSetDescriptorType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum IPSetDescriptorType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MatchFieldType string

// Enum values for MatchFieldType
const (
	MatchFieldTypeUri            MatchFieldType = "URI"
	MatchFieldTypeQueryString    MatchFieldType = "QUERY_STRING"
	MatchFieldTypeHeader         MatchFieldType = "HEADER"
	MatchFieldTypeMethod         MatchFieldType = "METHOD"
	MatchFieldTypeBody           MatchFieldType = "BODY"
	MatchFieldTypeSingleQueryArg MatchFieldType = "SINGLE_QUERY_ARG"
	MatchFieldTypeAllQueryArgs   MatchFieldType = "ALL_QUERY_ARGS"
)

func (enum MatchFieldType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MatchFieldType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MigrationErrorType string

// Enum values for MigrationErrorType
const (
	MigrationErrorTypeEntityNotSupported    MigrationErrorType = "ENTITY_NOT_SUPPORTED"
	MigrationErrorTypeEntityNotFound        MigrationErrorType = "ENTITY_NOT_FOUND"
	MigrationErrorTypeS3BucketNoPermission  MigrationErrorType = "S3_BUCKET_NO_PERMISSION"
	MigrationErrorTypeS3BucketNotAccessible MigrationErrorType = "S3_BUCKET_NOT_ACCESSIBLE"
	MigrationErrorTypeS3BucketNotFound      MigrationErrorType = "S3_BUCKET_NOT_FOUND"
	MigrationErrorTypeS3BucketInvalidRegion MigrationErrorType = "S3_BUCKET_INVALID_REGION"
	MigrationErrorTypeS3InternalError       MigrationErrorType = "S3_INTERNAL_ERROR"
)

func (enum MigrationErrorType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MigrationErrorType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ParameterExceptionField string

// Enum values for ParameterExceptionField
const (
	ParameterExceptionFieldChangeAction                     ParameterExceptionField = "CHANGE_ACTION"
	ParameterExceptionFieldWafAction                        ParameterExceptionField = "WAF_ACTION"
	ParameterExceptionFieldWafOverrideAction                ParameterExceptionField = "WAF_OVERRIDE_ACTION"
	ParameterExceptionFieldPredicateType                    ParameterExceptionField = "PREDICATE_TYPE"
	ParameterExceptionFieldIpsetType                        ParameterExceptionField = "IPSET_TYPE"
	ParameterExceptionFieldByteMatchFieldType               ParameterExceptionField = "BYTE_MATCH_FIELD_TYPE"
	ParameterExceptionFieldSqlInjectionMatchFieldType       ParameterExceptionField = "SQL_INJECTION_MATCH_FIELD_TYPE"
	ParameterExceptionFieldByteMatchTextTransformation      ParameterExceptionField = "BYTE_MATCH_TEXT_TRANSFORMATION"
	ParameterExceptionFieldByteMatchPositionalConstraint    ParameterExceptionField = "BYTE_MATCH_POSITIONAL_CONSTRAINT"
	ParameterExceptionFieldSizeConstraintComparisonOperator ParameterExceptionField = "SIZE_CONSTRAINT_COMPARISON_OPERATOR"
	ParameterExceptionFieldGeoMatchLocationType             ParameterExceptionField = "GEO_MATCH_LOCATION_TYPE"
	ParameterExceptionFieldGeoMatchLocationValue            ParameterExceptionField = "GEO_MATCH_LOCATION_VALUE"
	ParameterExceptionFieldRateKey                          ParameterExceptionField = "RATE_KEY"
	ParameterExceptionFieldRuleType                         ParameterExceptionField = "RULE_TYPE"
	ParameterExceptionFieldNextMarker                       ParameterExceptionField = "NEXT_MARKER"
	ParameterExceptionFieldResourceArn                      ParameterExceptionField = "RESOURCE_ARN"
	ParameterExceptionFieldTags                             ParameterExceptionField = "TAGS"
	ParameterExceptionFieldTagKeys                          ParameterExceptionField = "TAG_KEYS"
)

func (enum ParameterExceptionField) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ParameterExceptionField) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ParameterExceptionReason string

// Enum values for ParameterExceptionReason
const (
	ParameterExceptionReasonInvalidOption      ParameterExceptionReason = "INVALID_OPTION"
	ParameterExceptionReasonIllegalCombination ParameterExceptionReason = "ILLEGAL_COMBINATION"
	ParameterExceptionReasonIllegalArgument    ParameterExceptionReason = "ILLEGAL_ARGUMENT"
	ParameterExceptionReasonInvalidTagKey      ParameterExceptionReason = "INVALID_TAG_KEY"
)

func (enum ParameterExceptionReason) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ParameterExceptionReason) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PositionalConstraint string

// Enum values for PositionalConstraint
const (
	PositionalConstraintExactly      PositionalConstraint = "EXACTLY"
	PositionalConstraintStartsWith   PositionalConstraint = "STARTS_WITH"
	PositionalConstraintEndsWith     PositionalConstraint = "ENDS_WITH"
	PositionalConstraintContains     PositionalConstraint = "CONTAINS"
	PositionalConstraintContainsWord PositionalConstraint = "CONTAINS_WORD"
)

func (enum PositionalConstraint) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PositionalConstraint) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PredicateType string

// Enum values for PredicateType
const (
	PredicateTypeIpmatch           PredicateType = "IPMatch"
	PredicateTypeByteMatch         PredicateType = "ByteMatch"
	PredicateTypeSqlInjectionMatch PredicateType = "SqlInjectionMatch"
	PredicateTypeGeoMatch          PredicateType = "GeoMatch"
	PredicateTypeSizeConstraint    PredicateType = "SizeConstraint"
	PredicateTypeXssMatch          PredicateType = "XssMatch"
	PredicateTypeRegexMatch        PredicateType = "RegexMatch"
)

func (enum PredicateType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PredicateType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RateKey string

// Enum values for RateKey
const (
	RateKeyIp RateKey = "IP"
)

func (enum RateKey) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RateKey) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TextTransformation string

// Enum values for TextTransformation
const (
	TextTransformationNone               TextTransformation = "NONE"
	TextTransformationCompressWhiteSpace TextTransformation = "COMPRESS_WHITE_SPACE"
	TextTransformationHtmlEntityDecode   TextTransformation = "HTML_ENTITY_DECODE"
	TextTransformationLowercase          TextTransformation = "LOWERCASE"
	TextTransformationCmdLine            TextTransformation = "CMD_LINE"
	TextTransformationUrlDecode          TextTransformation = "URL_DECODE"
)

func (enum TextTransformation) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TextTransformation) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type WafActionType string

// Enum values for WafActionType
const (
	WafActionTypeBlock WafActionType = "BLOCK"
	WafActionTypeAllow WafActionType = "ALLOW"
	WafActionTypeCount WafActionType = "COUNT"
)

func (enum WafActionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum WafActionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type WafOverrideActionType string

// Enum values for WafOverrideActionType
const (
	WafOverrideActionTypeNone  WafOverrideActionType = "NONE"
	WafOverrideActionTypeCount WafOverrideActionType = "COUNT"
)

func (enum WafOverrideActionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum WafOverrideActionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type WafRuleType string

// Enum values for WafRuleType
const (
	WafRuleTypeRegular   WafRuleType = "REGULAR"
	WafRuleTypeRateBased WafRuleType = "RATE_BASED"
	WafRuleTypeGroup     WafRuleType = "GROUP"
)

func (enum WafRuleType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum WafRuleType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
