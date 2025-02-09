//go:build linux
// +build linux

package probing

// Returns the length of an ICMP message.
func (p *Pinger) getMessageLength() int {
	//return p.Size + 8
	if p.Privileged() {
		return p.Size + 8
	}
	if p.ipv4 {
		return p.Size + 8 + ipv4.HeaderLen
	}
	return p.Size + 8 + ipv6.HeaderLen
}

// Attempts to match the ID of an ICMP packet.
func (p *Pinger) matchID(ID int) bool {
	// On Linux we can only match ID if we are privileged.
	if p.protocol == "icmp" {
		if ID != p.id {
			return false
		}
	}
	return true
}
