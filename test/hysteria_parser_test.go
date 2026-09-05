package test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/sagernet/sing/common/byteformats"
)

func TestUnmarshalNetworkBytesCompact(t *testing.T) {
	up := &byteformats.NetworkBytesCompat{}
	upmbps := "100 Mbps"
	err := json.Unmarshal(fmt.Appendf(nil, `"%s"`, upmbps), up)
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}
	log.Printf("%+v", up)
}
