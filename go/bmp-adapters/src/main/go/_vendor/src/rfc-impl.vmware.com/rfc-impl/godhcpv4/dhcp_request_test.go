package dhcpv4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test dispatch to ReplyWriter
func TestDHCPRequestWriteReply(t *testing.T) {
	rw := &testReplyWriter{}

	req := DHCPRequest{
		Packet:      NewPacket(BootRequest),
		ReplyWriter: rw,
	}

	reps := []Reply{
		CreateDHCPAck(req),
		CreateDHCPNak(req),
	}

	for _, rep := range reps {
		rw.wrote = false
		req.WriteReply(rep)
		assert.True(t, rw.wrote)
	}
}
