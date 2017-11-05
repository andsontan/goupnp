package soap

import "regexp"

var xmlCharRx = regexp.MustCompile("[<>&]")

// EscapeXMLText is used by generated code to escape text in XML, but only
// escaping the characters `<`, `>`, and `&`.
//
// This is provided in order to work around SOAP server implementations that
// fail to decode XML correctly, specifically failing to decode `"`, `'`. Note
// that this can only be safely used for injecting into XML text, but not into
// attributes or other contexts.
func EscapeXMLText(s string) string {
	return xmlCharRx.ReplaceAllStringFunc(s, replaceEntity)
}

func replaceEntity(s string) string {
	switch s {
	case "<":
		return "&lt;"
	case ">":
		return "&gt;"
	case "&":
		return "&amp;"
	}
	return s
}
