package stanza

import (
	"encoding/xml"
)

// ============================================================================
// Ping (XEP-0199)
// https://xmpp.org/extensions/xep-0199.html

const (
	// NSPing defines the namespace for ping IQ stanzas
	NSPing = "urn:xmpp:ping"
)

// ----------
// Namespaces

type Ping struct {
	XMLName xml.Name `xml:"urn:xmpp:ping ping"`
}

// Namespace lets Ping implement the IQPayload interface
func (p *Ping) Namespace() string {
	return p.XMLName.Space
}

// ---------------
// Builder helpers

// Ping builds a default Ping payload
func (iq *IQ) Ping() *Ping {
	d := Ping{
		XMLName: xml.Name{
			Space: NSPing,
			Local: "ping",
		},
	}
	iq.Payload = &d
	return &d
}

// ============================================================================
// Registry init

func init() {
	stanza.TypeRegistry.MapExtension(stanza.PKTIQ, xml.Name{Space: NSPing, Local: "ping"}, Ping{})
}
