// string in go are consist of runes(Characters) size of one rune is int32 like java supports natural lang as well as emojis.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
    const emoji = "🐻"      //it is an bear emoji
    const name = "Umesh"
    fmt.Println("rune length ", utf8.RuneCountInString(name),"of Size ", len(name))
    fmt.Println("rune length ", utf8.RuneCountInString(emoji), "of Size ", len(emoji))
}
