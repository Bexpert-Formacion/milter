package milter

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net/textproto"
)

// Modifier provides access to Macros, Headers and Body data to callback handlers. It also defines a
// number of functions that can be used by callback handlers to modify processing of the email message
type Modifier struct {
	Macros      map[string]string
	Headers     textproto.MIMEHeader
	writePacket func(*Message) error
}

// AddRecipient appends a new envelope recipient for current message
func (m *Modifier) AddRecipient(r string) error {
	data := []byte(fmt.Sprintf("<%s>", r) + null)
	return m.writePacket(NewResponse(SMFIR_ADDRCPT, data).Response())
}

// DeleteRecipient removes an envelope recipient address from message
func (m *Modifier) DeleteRecipient(r string) error {
	data := []byte(fmt.Sprintf("<%s>", r) + null)
	return m.writePacket(NewResponse(SMFIR_DELRCPT, data).Response())
}

// ReplaceBody substitutes message body with provided body
func (m *Modifier) ReplaceBody(body []byte) error {
	return m.writePacket(NewResponse(SMFIR_REPLBODY, body).Response())
}

// AddHeader appends a new email message header the message
func (m *Modifier) AddHeader(name, value string) error {
	data := []byte(name + null + value + null)
	return m.writePacket(NewResponse(SMFIR_ADDHEADER, data).Response())
}

// Quarantine a message by giving a reason to hold it
func (m *Modifier) Quarantine(reason string) error {
	return m.writePacket(NewResponse(SMFIR_QUARANTINE, []byte(reason+null)).Response())
}

// ChangeHeader replaces the header at the specified position with a new one
func (m *Modifier) ChangeHeader(index int, name, value string) error {
	buffer := new(bytes.Buffer)
	// encode header index in the beginning
	if err := binary.Write(buffer, binary.BigEndian, uint32(index)); err != nil {
		return err
	}
	// add header name and value to buffer
	data := []byte(name + null + value + null)
	if _, err := buffer.Write(data); err != nil {
		return err
	}
	// prepare and send response packet
	return m.writePacket(NewResponse(SMFIR_CHGHEADER, buffer.Bytes()).Response())
}

// InsertHeader inserts the header at the pecified position
func (m *Modifier) InsertHeader(index int, name, value string) error {
	buffer := new(bytes.Buffer)
	// encode header index in the beginning
	if err := binary.Write(buffer, binary.BigEndian, uint32(index)); err != nil {
		return err
	}
	// add header name and value to buffer
	data := []byte(name + null + value + null)
	if _, err := buffer.Write(data); err != nil {
		return err
	}
	// prepare and send response packet
	return m.writePacket(NewResponse(SMFIR_INSHEADER, buffer.Bytes()).Response())
}

// ChangeFrom replaces the FROM envelope header with a new one
func (m *Modifier) ChangeFrom(value string) error {
	buffer := new(bytes.Buffer)
	// add header name and value to buffer
	data := []byte(value + null)
	if _, err := buffer.Write(data); err != nil {
		return err
	}
	// prepare and send response packet
	return m.writePacket(NewResponse(SMFIR_CHGFROM, buffer.Bytes()).Response())
}

// newModifier creates a new Modifier instance from milterSession
func newModifier(s *milterSession) *Modifier {
	return &Modifier{
		Macros:      s.macros,
		Headers:     s.headers,
		writePacket: s.WritePacket,
	}
}
