package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var account *Account = &Account{CustomerName: "Gabriel"}
	account.setDetails("412", "current")
	jsonAccount, _ := json.Marshal(account)

	// private class, hidden "details"
	fmt.Println(string(jsonAccount)) //{"CustomerName":"Gabriel"}

	fmt.Println(account.getAccountType()) // currenct
	fmt.Println(account.getId())          // 412
}
