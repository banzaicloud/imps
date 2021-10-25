// Copyright © 2021 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

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
