package main

import (
	"log"
	"testing"
)

// go test -v
func TestGetCandidatesFromSDP(m *testing.T) {
	sDP := `v=0
o=- 2001876963261635747 1698317880 IN IP4 0.0.0.0
s=-
t=0 0
a=fingerprint:sha-256 30:9F:65:41:12:B3:0D:29:25:26:0A:E6:75:52:76:4D:16:74:A8:77:59:CC:EB:FE:30:AD:02:76:59:DF:ED:F5
m=audio 0 UDP/TLS/RTP/SAVPF 0
c=IN IP4 0.0.0.0
a=candidate:1689164903 1 udp 2130706431 172.22.0.4 45183 typ host
a=candidate:1689164903 2 udp 2130706431 172.22.0.4 45183 typ host
a=candidate:3605751344 1 udp 16777215 1.1.1.1 14090 typ relay raddr 172.22.0.4 rport 47994
a=candidate:3605751344 2 udp 16777215 1.1.1.1 14090 typ relay raddr 172.22.0.4 rport 47994
a=end-of-candidates`

	// log.Println(sDP)

	/*for _, line := range strings.Split(sDP, "\n") {
		if strings.Contains(line, "a=candidate:") {
			c := line[2:]
			log.Printf("c: '%s'\n", c)
		}
		// log.Println(line)
	}*/

	log.Printf("%+v\n", getCandidatesFromSDP(sDP))
}
