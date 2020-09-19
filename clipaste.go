package main

import (
	"fmt"

	"github.com/atotto/clipboard"
)

//linuxFedora環境ではdnfでxcilpをインストールする必要有り
//aaaa
func main() {

	//clipboard.WriteAll("日本語")
	cliptext, _err := clipboard.ReadAll()

	fmt.Println(_err)
	fmt.Println(cliptext)

}
