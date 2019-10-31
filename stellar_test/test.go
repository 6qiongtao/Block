package main

import (
	"fmt"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/keypair"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main1() {
	// pair is the pair that was generated from previous example, or create a pair based on
	// existing keys.
	//生成密钥对
	pair, err := keypair.Random()
	if err != nil {
		log.Println("err:", err)
	}

	log.Println(pair.Seed())
	log.Println(pair.Address())

	//创建test account
	address := pair.Address()
	resp, err := http.Get("https://friendbot.stellar.org/?addr=" + address)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

	//查询余额
	account, err := horizon.DefaultTestNetClient.LoadAccount(pair.Address())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balances for account:", pair.Address())
	for _, balance := range account.Balances {
		log.Println("balance:", balance)
	}
}

func test_xorm() {
//psql -U stellar -h 127.0.0.1 -d stellar
//	密码：C8UQxCrXoXWJVjW0

}

func main() {

	//main1()
	test_xorm()

}