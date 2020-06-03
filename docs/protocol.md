# protocol
--
    import "github.com/jonnylin13/udpchat/protocol"


## Usage

#### func  Opcodes

```go
func Opcodes() map[string]byte
```
Opcodes returns a map[string]byte of 1-bit opcodes.

#### func  PackString

```go
func PackString(str string) []byte
```
PackString packs a string field of a packet.

#### func  UnpackString

```go
func UnpackString(buf []byte, start int) (str string, end int)
```
UnpackString unpacks a string field of a packet.

#### type Packet

```go
type Packet struct {
	Opcode  byte
	Payload []byte
}
```

Packet represents a 1 Kb packet.

#### func  PacketHandshake

```go
func PacketHandshake(name string) Packet
```
PacketHandshake returns a handshake packet.

#### func  PacketHandshakeAck

```go
func PacketHandshakeAck() Packet
```
PacketHandshakeAck returns a handshake ack packet.

#### func  PacketLeave

```go
func PacketLeave(name string) Packet
```
PacketLeave returns a leave packet.

#### func  PacketLeaveAck

```go
func PacketLeaveAck() Packet
```
PacketLeaveAck returns a leave ack packet.

#### func  PacketMessage

```go
func PacketMessage(name string, msg string) Packet
```
PacketMessage returns a message packet.

#### func  PacketMessageAck

```go
func PacketMessageAck() Packet
```
PacketMessageAck returns a message ack packet.

#### func  PacketUnknownRequestAck

```go
func PacketUnknownRequestAck() Packet
```
PacketUnknownRequestAck returns an unknown request ack packet.

#### func (Packet) Pack

```go
func (p Packet) Pack() []byte
```
Pack packs the Packet into a []byte.
