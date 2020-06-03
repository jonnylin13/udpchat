package protocol

import (
	"encoding/binary"
)

// Opcodes returns a map of opcodes
func Opcodes() map[string]byte {
	return map[string]byte{
		"handshake":     0x0,
		"handshake_ack": 0x1,
		"unknown_ack":   0x2,
		"message":       0x3,
		"message_ack":   0x4,
		"leave":         0x5,
		"leave_ack":     0x5,
	}
}

// Packet represents a 1 Kb packet.
type Packet struct {
	opcode  byte
	payload []byte
}

// Pack the Packet into a byte array.
func (p Packet) Pack() []byte {
	buf := append([]byte{p.opcode}, p.payload[:]...)
	buf = append(buf, make([]byte, 1024-len(buf))...)
	return buf
}

// UnpackString unpacks a string field of a packet.
func UnpackString(buf []byte, start int) (str string, end int) {
	strLength := binary.LittleEndian.Uint16(buf[start : start+2])
	return string(buf[start+2 : start+2+int(strLength)]), start + 2 + int(strLength)
}

func packString(str string) []byte {
	strBytes := []byte(str)
	lengthBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(lengthBytes, uint16(len(str)))
	nameField := append(lengthBytes, strBytes...)
	return nameField
}

// NewPacketHandshake returns a handshake packet.
func NewPacketHandshake(name string) Packet {
	p := Packet{}
	p.opcode = Opcodes()["handshake"]
	p.payload = packString(name)
	return p
}

// NewPacketHandshakeAck returns a handshake ack packet.
func NewPacketHandshakeAck() Packet {
	p := Packet{}
	p.opcode = Opcodes()["handshake_ack"]
	p.payload = []byte{}
	return p
}

// NewPacketLeave returns a leave packet.
func NewPacketLeave(name string) Packet {
	p := Packet{}
	p.opcode = Opcodes()["leave"]
	p.payload = packString(name)
	return p
}

// NewPacketLeaveAck returns a leave ack packet.
func NewPacketLeaveAck() Packet {
	p := Packet{}
	p.opcode = Opcodes()["leave_ack"]
	p.payload = []byte{}
	return p
}

// NewPacketUnknownRequestAck returns an unknown request ack packet.
func NewPacketUnknownRequestAck() Packet {
	p := Packet{}
	p.opcode = Opcodes()["unknown_ack"]
	p.payload = []byte{}
	return p
}

// NewPacketMessage returns a message packet.
func NewPacketMessage(name string, msg string) Packet {
	p := Packet{}
	p.opcode = Opcodes()["message"]
	nameField := packString(name)
	if len(msg) > 1023-len(nameField) {
		msg = msg[0 : 1023-len(nameField)]
	}
	msgField := packString(msg)
	p.payload = append(nameField, msgField...)
	return p
}

// NewPacketMessageAck returns a message ack packet.
func NewPacketMessageAck() Packet {
	p := Packet{}
	p.opcode = Opcodes()["message_ack"]
	p.payload = []byte{}
	return p
}
