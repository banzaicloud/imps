// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.

package pullsecrets

import (
	"fmt"
	"hash/crc32"
	"regexp"
	"strings"
)

var invalidCharacterRegexp = regexp.MustCompile("[^a-z0-9.-]+")

// Total limit is 253
const maxSantiziedLength = 100

func SecretNameFromURL(prefix, url string) string {
	sanitizedName := strings.ToLower(url)
	sanitizedName = invalidCharacterRegexp.ReplaceAllString(sanitizedName, "-")
	if len(sanitizedName) > maxSantiziedLength {
		sanitizedName = sanitizedName[0:maxSantiziedLength]
	}
	sanitizedName = strings.Trim(sanitizedName, "-")

	urlCRC := crc32.ChecksumIEEE([]byte(url))

	return fmt.Sprintf("%s-%s-pull-secret-%08x", prefix, sanitizedName, urlCRC)
}
