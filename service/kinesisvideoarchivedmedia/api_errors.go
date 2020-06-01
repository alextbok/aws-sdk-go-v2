// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideoarchivedmedia

const (

	// ErrCodeInvalidArgumentException for service response error code
	// "InvalidArgumentException".
	//
	// A specified parameter exceeds its restrictions, is not supported, or can't
	// be used.
	ErrCodeInvalidArgumentException = "InvalidArgumentException"

	// ErrCodeInvalidCodecPrivateDataException for service response error code
	// "InvalidCodecPrivateDataException".
	//
	// The codec private data in at least one of the tracks of the video stream
	// is not valid for this operation.
	ErrCodeInvalidCodecPrivateDataException = "InvalidCodecPrivateDataException"

	// ErrCodeInvalidMediaFrameException for service response error code
	// "InvalidMediaFrameException".
	//
	// One or more frames in the requested clip could not be parsed based on the
	// specified codec.
	ErrCodeInvalidMediaFrameException = "InvalidMediaFrameException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// Kinesis Video Streams has throttled the request because you have exceeded
	// the limit of allowed client calls. Try making the call later.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeMissingCodecPrivateDataException for service response error code
	// "MissingCodecPrivateDataException".
	//
	// No codec private data was found in at least one of tracks of the video stream.
	ErrCodeMissingCodecPrivateDataException = "MissingCodecPrivateDataException"

	// ErrCodeNoDataRetentionException for service response error code
	// "NoDataRetentionException".
	//
	// A streaming session was requested for a stream that does not retain data
	// (that is, has a DataRetentionInHours of 0).
	ErrCodeNoDataRetentionException = "NoDataRetentionException"

	// ErrCodeNotAuthorizedException for service response error code
	// "NotAuthorizedException".
	//
	// Status Code: 403, The caller is not authorized to perform an operation on
	// the given stream, or the token has expired.
	ErrCodeNotAuthorizedException = "NotAuthorizedException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// GetMedia throws this error when Kinesis Video Streams can't find the stream
	// that you specified.
	//
	// GetHLSStreamingSessionURL and GetDASHStreamingSessionURL throw this error
	// if a session with a PlaybackMode of ON_DEMAND or LIVE_REPLAYis requested
	// for a stream that has no fragments within the requested time range, or if
	// a session with a PlaybackMode of LIVE is requested for a stream that has
	// no fragments within the last 30 seconds.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeUnsupportedStreamMediaTypeException for service response error code
	// "UnsupportedStreamMediaTypeException".
	//
	// The type of the media (for example, h.264 or h.265 video or ACC or G.711
	// audio) could not be determined from the codec IDs of the tracks in the first
	// fragment for a playback session. The codec ID for track 1 should be V_MPEG/ISO/AVC
	// and, optionally, the codec ID for track 2 should be A_AAC.
	ErrCodeUnsupportedStreamMediaTypeException = "UnsupportedStreamMediaTypeException"
)
