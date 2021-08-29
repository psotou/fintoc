package main

import (
	"fmt"
	"os"

	"github.com/psotou/fintoc/fintoc"
)

func main() {
	secret := os.Getenv("FINTOC_SECRET_KEY")
	client, _ := fintoc.NewClient(secret)

	// links := client.LinkAll()
	link := client.LinkOne("")
	// for _, v := range links {
	// 	fmt.Println(v)
	// }
	fmt.Println(link)

	// we use the link object to get the accounts
	// account := link.AccounOne("kpOQlyLTzAvK13qW")
	// accounts := link.AccountAll()

	// fmt.Println(account)

	// for _, v := range accounts {
	// 	fmt.Println(v)
	// }

	// movements := account.MovementAll()
	// for _, v := range movements {
	// 	fmt.Println(v)
	// }
	// fmt.Println(movements)
	// // movement := account.MovementOne("mov_j4yeaPH0rXa68XOP")
	// fmt.Println(movement)
}
