package oapi

func GetRawSpec() (spec []byte, fault error) {
	return decodeSpec()
}
