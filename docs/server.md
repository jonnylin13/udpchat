# server
--
    import "github.com/jonnylin13/udpchat/server"


## Usage

#### func  GetNames

```go
func GetNames(userList []User) []string
```
GetNames returns a []string of names from []User.

#### func  Start

```go
func Start(port string)
```
Start the server.

#### type User

```go
type User struct {
	Name string
	Addr net.Addr
}
```

User represents a user.

#### func  RemoveUser

```go
func RemoveUser(userList []User, name string) []User
```
RemoveUser returns a []User with the removed name.
