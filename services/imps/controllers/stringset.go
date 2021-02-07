// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.

package controllers

type StringSet []string

func (s StringSet) Has(needle string) bool {
	for _, setMember := range s {
		if setMember == needle {
			return true
		}
	}
	return false
}
