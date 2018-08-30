/*
Sniperkit-Bot
- Status: analyzed
*/

package rule

import (
	. "github.com/sniperkit/snk.fork.daniel-fanjul-alcuten-cmake-tools/model"
)

type Error struct {
	// TODO add line numbers
	error
}

type Rule interface {
	Check(items <-chan Item, errs chan<- Error)
	Format(input <-chan Item, output chan<- Item)
}
