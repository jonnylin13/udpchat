package protocol

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

// PacketHandshake returns a handshake packet.
func PacketHandshake(name string) Packet {
	p := Packet{}
	p.opcode = Opcodes()["handshake"]
	p.payload = PackString(name)
	return p
}

// PacketHandshakeAck returns a handshake ack packet.
func PacketHandshakeAck() Packet {
	p := Packet{}
	p.opcode = Opcodes()["handshake_ack"]
	p.payload = []byte{}
	return p
}

// PacketLeave returns a leave packet.
func PacketLeave(name string) Packet {
	p := Packet{}
	p.opcode = Opcodes()["leave"]
	p.payload = PackString(name)
	return p
}

// PacketLeaveAck returns a leave ack packet.
func PacketLeaveAck() Packet {
	p := Packet{}
	p.opcode = Opcodes()["leave_ack"]
	p.payload = []byte{}
	return p
}

// PacketUnknownRequestAck returns an unknown request ack packet.
func PacketUnknownRequestAck() Packet {
	p := Packet{}
	p.opcode = Opcodes()["unknown_ack"]
	p.payload = []byte{}
	return p
}

// PacketMessage returns a message packet.
func PacketMessage(name string, msg string) Packet {
	p := Packet{}
	p.opcode = Opcodes()["message"]
	nameField := PackString(name)
	if len(msg) > 1023-len(nameField) {
		msg = msg[0 : 1023-len(nameField)]
	}
	msgField := PackString(msg)
	p.payload = append(nameField, msgField...)
	return p
}

// PacketMessageAck returns a message ack packet.
func PacketMessageAck() Packet {
	p := Packet{}
	p.opcode = Opcodes()["message_ack"]
	p.payload = []byte{}
	return p
}
