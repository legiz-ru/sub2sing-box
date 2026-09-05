package model

import (
	"github.com/sagernet/sing-box/option"
	"github.com/sagernet/sing/common/json/badjson"
)

type LogOptions struct {
	Disabled     bool   `json:"disabled,omitempty"`
	Level        string `json:"level,omitempty"`
	Output       string `json:"output,omitempty"`
	Timestamp    bool   `json:"timestamp,omitempty"`
	DisableColor bool   `json:"-"`
}

type StubOptions struct{}

type Endpoint option.Endpoint

func (e *Endpoint) MarshalJSON() ([]byte, error) {
	return badjson.MarshallObjects((*option.Endpoint)(e), e.Options)
}

type Inbound option.Inbound

func (i *Inbound) MarshalJSON() ([]byte, error) {
	return badjson.MarshallObjects((*option.Inbound)(i), i.Options)
}

type Outbound option.Outbound

func (o *Outbound) MarshalJSON() ([]byte, error) {
	return badjson.MarshallObjects((*option.Outbound)(o), o.Options)
}
