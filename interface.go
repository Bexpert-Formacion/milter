package milter

import (
	"net"
	"net/textproto"
)

// SessionHandler is an interface for milter callbacks handlers
type SessionHandler interface {
	// Init is called on begin og a new Mail, before Connect() and before MailFrom()
	// Can be used to Reset session state
	// On MailFrom mailID is avaliable
	Init(sessionID, mailID string)

	// Connect is called to provide a SMTP connection data for incoming message
	// supress with NoContent
	Connect(host string, family string, port uint16, addr net.IP, m *Modifier) (Response, error)

	// Helo is called to process any HELO/EHLO related filters
	// supress with NoHelo
	Helo(name string, m *Modifier) (Response, error)

	// MailFrom is called to process filters on envelope FROM address
	// supress with NoMailFrom
	MailFrom(from string, m *Modifier) (Response, error)

	// RcptTo is called to process filters on envelope TO address
	// supress with NoRcptTo
	RcptTo(rcptTo string, m *Modifier) (Response, error)

	// Header is called once for each header in incoming message
	// supress with NoHeaders
	Header(name string, value string, m *Modifier) (Response, error)

	// Headers is called when all message headers have been processed
	// supress with NoHeaders
	Headers(h textproto.MIMEHeader, m *Modifier) (Response, error)

	// BodyChunk is called to process next message body chunk data (up to 64 KB in size)
	// supress with NoBody
	BodyChunk(chunk []byte, m *Modifier) (Response, error)

	// Body is called at the end of each message
	// all changes to message's content & attributes must be done here
	Body(m *Modifier) (Response, error)

	// Disconnect is called at the end of each message handling loop
	Disconnect()
}
