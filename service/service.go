// Package service provides configured service instances.
package service

import (
	"net/http"

	"github.com/daaku/go.fbapi"
	"github.com/daaku/go.httpcontrol"
	"github.com/daaku/go.redis"
	"github.com/daaku/go.redis/bytecache"
	"github.com/daaku/go.redis/bytestore"
	"github.com/daaku/go.static"
	"github.com/daaku/go.stats/stathatbackend"
	"github.com/daaku/go.xsrf"
)

var (
	Xsrf          = xsrf.ProviderFlag("xsrf")
	Static        = static.HandlerFlag("rell.static")
	Stats         = stathatbackend.EZKeyFlag("rell.stats")
	Redis         = redis.ClientFlag("rell.redis")
	ByteCache     = bytecache.New(Redis)
	ByteStore     = bytestore.New(Redis)
	HttpTransport = httpcontrol.TransportFlag("rell.transport")
	HttpClient    = &http.Client{Transport: HttpTransport}
	FbApiClient   = fbapi.ClientFlag("rell.fbapi")
)

func init() {
	Stats.Client = HttpClient
	FbApiClient.HttpClient = HttpClient
	Redis.Stats = Stats
}
