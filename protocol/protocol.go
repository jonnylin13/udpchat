package protocol

// Opcodes returns a map[string]byte of 1-bit opcodes.
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
	p.Opcode = Opcodes()["handshake"]
	p.Payload = PackString(name)
	return p
}

// PacketHandshakeAck returns a handshake ack packet.
func PacketHandshakeAck() Packet {
	p := Packet{}
	p.Opcode = Opcodes()["handshake_ack"]
	p.Payload = []byte{}
	return p
}

// PacketLeave returns a leave packet.
func PacketLeave(name string) Packet {
	p := Packet{}
	p.Opcode = Opcodes()["leave"]
	p.Payload = PackString(name)
	return p
}

// PacketLeaveAck returns a leave ack packet.
func PacketLeaveAck() Packet {
	p := Packet{}
	p.Opcode = Opcodes()["leave_ack"]
	p.Payload = []byte{}
	return p
}

// PacketUnknownRequestAck returns an unknown request ack packet.
func PacketUnknownRequestAck() Packet {
	p := Packet{}
	p.Opcode = Opcodes()["unknown_ack"]
	p.Payload = []byte{}
	return p
}

// PacketMessage returns a message packet.
func PacketMessage(name string, msg string) Packet {
	p := Packet{}
	p.Opcode = Opcodes()["message"]
	nameField := PackString(name)
	if len(msg) > 1023-len(nameField) {
		msg = msg[0 : 1023-len(nameField)]
	}
	msgField := PackString(msg)
	p.Payload = append(nameField, msgField...)
	return p
}

// PacketMessageAck returns a message ack packet.
func PacketMessageAck() Packet {
	p := Packet{}
	p.Opcode = Opcodes()["message_ack"]
	p.Payload = []byte{}
	return p
}
