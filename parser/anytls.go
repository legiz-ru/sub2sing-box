package parser

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/bestnite/sub2sing-box/constant"
	"github.com/bestnite/sub2sing-box/model"
	"github.com/sagernet/sing-box/option"
)

func ParseAnytls(proxy string) (model.Outbound, error) {
	if !strings.HasPrefix(proxy, constant.AnytlsPrefix) {
		return model.Outbound{}, &ParseError{Type: ErrInvalidPrefix, Raw: proxy}
	}

	link, err := url.Parse(proxy)
	if err != nil {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "url parse error",
			Raw:     proxy,
		}
	}

	username := link.User.Username()
	password, exist := link.User.Password()
	if !exist {
		password = username
	}

	query := link.Query()
	server := link.Hostname()
	if server == "" {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing server host",
			Raw:     proxy,
		}
	}
	portStr := link.Port()
	port, err := ParsePort(portStr)
	if err != nil {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidPort,
			Message: err.Error(),
			Raw:     proxy,
		}
	}
	insecure, sni := query.Get("insecure"), query.Get("sni")
	insecureBool := insecure == "1"
	remarks := link.Fragment
	if remarks == "" {
		remarks = fmt.Sprintf("%s:%s", server, portStr)
	}
	remarks = strings.TrimSpace(remarks)

	opts := option.AnyTLSOutboundOptions{
		ServerOptions: option.ServerOptions{
			Server:     server,
			ServerPort: port,
		},
		Password: password,
	}
	if sni != "" {
		opts.OutboundTLSOptionsContainer = option.OutboundTLSOptionsContainer{
			TLS: &option.OutboundTLSOptions{
				Enabled:    true,
				ServerName: sni,
				Insecure:   insecureBool,
			},
		}
	}

	return model.Outbound{
		Type:    "anytls",
		Tag:     remarks,
		Options: opts,
	}, nil
}
