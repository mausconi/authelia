package handlers

import (
	"github.com/valyala/fasthttp"
)

// PortalGet is the handler serving Authelia's portal.
var PortalGet = fasthttp.FSHandler("./public_html", 0)
