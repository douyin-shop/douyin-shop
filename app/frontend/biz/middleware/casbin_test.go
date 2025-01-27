// @Author Adrian.Wang 2025/1/26 23:01:00
package middleware

import "testing"

func TestRegexRouterMatcher(t *testing.T) {
	router := "/cart/test"
	regex := `^/cart/.+$`
	t.Log(RegexRouterMatcher(router, regex))
}
