package protocol

import "encoding/binary"

// Packet represents a 1 Kb packet.
type Packet struct {
	Opcode  byte
	Payload []byte
}

// Pack packs the Packet into a []byte.
func (p Packet) Pack() []byte {
	buf := append([]byte{p.Opcode}, p.Payload[:]...)
	buf = append(buf, make([]byte, 1024-len(buf))...)
	return buf
}

// UnpackString unpacks a string field of a packet.
func UnpackString(buf []byte, start int) (str string, end int) {
	strLength := binary.LittleEndian.Uint16(buf[start : start+2])
	length := int(strLength)
	return string(buf[start+2 : start+2+length]), start + 2 + length
}

// PackString packs a string field of a packet.
func PackString(str string) []byte {
	strBytes := []byte(str)
	lengthBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(lengthBytes, uint16(len(str)))
	nameField := append(lengthBytes, strBytes...)
	return nameField
}
