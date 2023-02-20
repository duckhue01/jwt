package jwt

import (
	"encoding/base64"
	"fmt"
)

const (
	DotTok = '.'
)

func Decode(s string) (string, error) {
	segments := lexSegments(s)
	header, err := DecodeSegment(segments[0])
	if err != nil {
		return "", err
	}

	body, err := DecodeSegment(segments[1])
	if err != nil {
		return "", err
	}

	return fmt.Sprintln(header, body), nil

}

func DecodeSegment(seg string) ([]byte, error) {
	encoding := base64.RawURLEncoding

	// if l := len(seg) % 4; l > 0 {
	// 	seg += strings.Repeat("=", 4-l)
	// }
	// encoding = base64.URLEncoding

	// encoding = encoding.Strict()
	return encoding.DecodeString(seg)
}

func lexSegments(s string) [3]string {
	c := 0
	segments := [3]string{}
	segment := make([]rune, 0, 20)
	for _, v := range s {
		if v == DotTok {
			segments[c] = string(segment)
		}
		segment = append(segment, v)
	}
	return segments
}
