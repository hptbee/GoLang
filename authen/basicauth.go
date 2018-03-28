// file: web/middleware/basicauth.go

package authen

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

// BasicAuth middleware sample.
// var BasicAuth = basicauth.New(basicauth.Config{
// 	Users: map[string]string{
// 		"admin": "123123",
// 	},
// 	Realm:   "Authorization Required", // defaults to "Authorization Required"
// 	Expires: 0,
// })

const cookieNameForSessionID = "hptcookie"

var sess = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})

var BasicAuth = Check

func Check(ctx iris.Context) {

	var sesss = sess.Start(ctx)
	var userID = sesss.GetIntDefault("UserID", 0)

	fmt.Println(userID)
	fmt.Println(sesss)
	// Check if user is authenticated
	if userID == 0 {
		ctx.Redirect("/login")
	}

	// Print secret message
}
