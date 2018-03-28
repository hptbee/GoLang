// file: web/middleware/basicauth.go

package authen

import (
	"github.com/kataras/iris/middleware/basicauth"
)

// BasicAuth middleware sample.
var BasicAuth = basicauth.New(basicauth.Config{
	Users: map[string]string{
		"admin": "123123",
	},
	Realm:   "Authorization Required", // defaults to "Authorization Required"
	Expires: 0,
})
