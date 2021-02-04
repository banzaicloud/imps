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
