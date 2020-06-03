package protocol

import "encoding/binary"

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

// PackString packs a string field of a packet.
func PackString(str string) []byte {
	strBytes := []byte(str)
	lengthBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(lengthBytes, uint16(len(str)))
	nameField := append(lengthBytes, strBytes...)
	return nameField
}
