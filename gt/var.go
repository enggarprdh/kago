package gt

import "github.com/eaciit/toolkit"

//IntData get random int
func IntData(max int) int {
	return toolkit.RandInt(max)
}

//StringData get random string
func StringData(length int) string {
	return toolkit.RandomString(length)
}
