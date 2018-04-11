package authen

import (
	"sync"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type Owner struct {
	// define here the fields that are global
	// and shared to all clients.
	sessionsManager *sessions.Sessions
}

// this package-level variable "application" will be used inside context to communicate with our global Application.
var owner = &Owner{
	sessionsManager: sessions.New(sessions.Config{Cookie: "mysessioncookie"}),
}

// Context is our custom context.
// Let's implement a context which will give us access
// to the client's Session with a trivial `ctx.Session()` call.
type Context struct {
	iris.Context
	session *sessions.Session
}

// Session returns the current client's session.
func (ctx *Context) Session() *sessions.Session {
	// this help us if we call `Session()` multiple times in the same handler
	if ctx.session == nil {
		// start a new session if not created before.
		ctx.session = owner.sessionsManager.Start(ctx.Context)
	}

	return ctx.session
}

// Handler will convert our handler of func(*Context) to an iris Handler,
// in order to be compatible with the HTTP API.
func Handler(h func(*Context)) iris.Handler {
	return func(original iris.Context) {
		ctx := acquire(original)
		h(ctx)
		release(ctx)
	}
}
func acquire(original iris.Context) *Context {
	ctx := contextPool.Get().(*Context)
	ctx.Context = original // set the context to the original one in order to have access to iris's implementation.
	ctx.session = nil      // reset the session
	return ctx
}

func release(ctx *Context) {
	contextPool.Put(ctx)
}

var contextPool = sync.Pool{New: func() interface{} {
	return &Context{}
}}

func CreateSession(name string, value string, original iris.Context) {
	ctx := acquire(original)
	var sess = ctx.Session()
	sess.Set(name, value)
}

func RemoveSession(name string, original iris.Context) {
	ctx := acquire(original)
	var sess = ctx.Session()
	sess.Delete(name)
}
