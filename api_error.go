package foursquare

// https://developer.foursquare.com/overview/responses
type ApiError struct {
	Meta		MetaInfo
}

type MetaInfo struct {
	Code 		int32
	ErrorDetail	string
	ErrorType	string
}