package proxy

import (
	"net/http"
	"net/textproto"
	"strings"
)

// All code below is copied from the Go standard library and subject
// to the following license:
// 		Copyright 2009 The Go Authors. All rights reserved.
// 		Use of this source code is governed by a BSD-style
// 		license that can be found in the LICENSE file.
//
// We copied this code to remove the isCookieNameValid check, in order
// to support the bad habbit of PHP applications to produce invalid cookie
// names (see https://github.com/khulnasoft/devpod/issues/2470).
//

// readCookies parses all "Cookie" values from the header h and
// returns the successfully parsed Cookies.
//
// if filter isn't empty, only cookies of that name are returned.
func readCookies(h http.Header, filter string) []*http.Cookie {
	lines := h["Cookie"]
	if len(lines) == 0 {
		return []*http.Cookie{}
	}

	cookies := make([]*http.Cookie, 0, len(lines)+strings.Count(lines[0], ";"))
	for _, line := range lines {
		line = textproto.TrimString(line)

		var part string
		for len(line) > 0 { // continue since we have rest
			if splitIndex := strings.Index(line, ";"); splitIndex > 0 {
				part, line = line[:splitIndex], line[splitIndex+1:]
			} else {
				part, line = line, ""
			}
			part = textproto.TrimString(part)
			if len(part) == 0 {
				continue
			}
			name, val := part, ""
			if j := strings.Index(part, "="); j >= 0 {
				name, val = name[:j], name[j+1:]
			}
			if filter != "" && filter != name {
				continue
			}
			val, quoted, ok := parseCookieValue(val, true)
			if !ok {
				continue
			}
			cookies = append(cookies, &http.Cookie{Name: name, Value: val, Quoted: quoted})
		}
	}
	return cookies
}

// parseCookieValue parses a cookie value according to RFC 6265.
// If allowDoubleQuote is true, parseCookieValue will consider that it
// is parsing the cookie-value;
// otherwise, it will consider that it is parsing a cookie-av value
// (cookie attribute-value).
//
// It returns the parsed cookie value, a boolean indicating whether the
// parsing was successful, and a boolean indicating whether the parsed
// value was enclosed in double quotes.
func parseCookieValue(raw string, allowDoubleQuote bool) (value string, quoted, ok bool) {
	// Strip the quotes, if present.
	if allowDoubleQuote && len(raw) > 1 && raw[0] == '"' && raw[len(raw)-1] == '"' {
		raw = raw[1 : len(raw)-1]
		quoted = true
	}
	for i := 0; i < len(raw); i++ {
		if !validCookieValueByte(raw[i]) {
			return "", quoted, false
		}
	}
	return raw, quoted, true
}

func validCookieValueByte(b byte) bool {
	return 0x20 <= b && b < 0x7f && b != '"' && b != ';' && b != '\\'
}
