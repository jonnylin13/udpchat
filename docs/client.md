# client
--
    import "github.com/jonnylin13/udpchat/client"


## Usage

#### func  Connect

```go
func Connect(addrStr string) (addr *net.UDPAddr, pc *net.UDPConn, err error)
```
Connect to a chatroom.

#### func  Start

```go
func Start(addrStr string)
```
Start the client.
