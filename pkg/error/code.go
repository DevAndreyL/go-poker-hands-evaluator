package pokererr

type Code string

const (
	CodeGeneralError    Code = "general_error"
	CodeUnknown         Code = "unknown"
	CodeValidationError Code = "failed_validation_request"
	CodeApiDecoderError Code = "api.decoder.error"
)
