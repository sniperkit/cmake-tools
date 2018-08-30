/*
Sniperkit-Bot
- Status: analyzed
*/

package model

import (
	"fmt"

	. "github.com/sniperkit/snk.fork.daniel-fanjul-alcuten-cmake-tools/token"
)

// an unexpected token
type UnexpectedToken struct {
	Token Token
}

func (u UnexpectedToken) String() string {
	return u.ItemString()
}

func (u UnexpectedToken) ItemString() string {
	return fmt.Sprintf("%T(%v)", u, u.Token.TokenString())
}

func (u UnexpectedToken) Equal(i Item) bool {
	if n, ok := i.(UnexpectedToken); ok {
		return u.Token == n.Token
	}
	return false
}
