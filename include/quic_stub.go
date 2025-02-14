//go:build !with_quic

package include

import (
	"context"

	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing-box/common/tls"
	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/option"
	"github.com/sagernet/sing-box/transport/v2ray"
	"github.com/sagernet/sing-dns"
	E "github.com/sagernet/sing/common/exceptions"
	M "github.com/sagernet/sing/common/metadata"
	N "github.com/sagernet/sing/common/network"
)

const WithQUIC = false

func init() {
	dns.RegisterTransport([]string{"quic", "h3"}, func(ctx context.Context, dialer N.Dialer, link string) (dns.Transport, error) {
		return nil, C.ErrQUICNotIncluded
	})
	v2ray.RegisterQUICConstructor(
		func(ctx context.Context, options option.V2RayQUICOptions, tlsConfig tls.Config, handler N.TCPConnectionHandler, errorHandler E.Handler) (adapter.V2RayServerTransport, error) {
			return nil, C.ErrQUICNotIncluded
		},
		func(ctx context.Context, dialer N.Dialer, serverAddr M.Socksaddr, options option.V2RayQUICOptions, tlsConfig tls.Config) (adapter.V2RayClientTransport, error) {
			return nil, C.ErrQUICNotIncluded
		},
	)
}
