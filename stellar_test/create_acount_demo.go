package main

import (
	"fmt"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/txnbuild"
	"io/ioutil"
	"log"
	"net/http"
)


//创建账户
func createAccount() (*keypair.Full, error) {
	pair, err := keypair.Random()
	if err != nil {
		log.Fatal(err)
	}
	//私钥
	log.Println(pair.Seed())
	//公钥
	log.Println(pair.Address())
	return pair, err
}

func sendPubKey(pubKey string) {
	// pair is the pair that was generated from previous example, or create a pair based on
	// existing keys.
	address := pubKey
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
}

//查询余额
func getBlance (a *horizonclient.Client, add string) {

	account, err := a.AccountDetail(horizonclient.AccountRequest{AccountID: add})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balances for account:",add)
	for _, balance := range account.Balances {
		fmt.Printf("balance: %+v \n", balance)
	}
}

//向Friendbot发送您创建的公共密钥。它将使用该公钥作为帐户ID创建新帐户并为其提供资金
func sendPubKey2(a *horizonclient.Client, kp *keypair.Full) {
	// Get information about the account we just created
	accountRequest := horizonclient.AccountRequest{AccountID: kp.Address()}
	hAccount0, err := a.AccountDetail(accountRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Generate a second randomly generated address
	kp2, err := createAccount()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Seed    B:", kp2.Seed())
	log.Println("Address B:", kp2.Address())

	// Construct the operation
	createAccountOp := txnbuild.CreateAccount{
		Destination: kp2.Address(),
		Amount:      "10",
	}

	// Construct the transaction that will carry the operation
	tx := txnbuild.Transaction{
		SourceAccount: &hAccount0,
		Operations:    []txnbuild.Operation{&createAccountOp},
		Timebounds:    txnbuild.NewTimeout(300),
		Network:       network.TestNetworkPassphrase,
	}

	// Sign the transaction, serialise it to XDR, and base 64 encode it
	txeBase64, err := tx.BuildSignEncode(kp)
	log.Println("Transaction base64: ", txeBase64)

	// Submit the transaction
	resp, err := a.SubmitTransactionXDR(txeBase64)
	if err != nil {
		hError := err.(*horizonclient.Error)
		log.Fatal("Error submitting transaction:", hError)
	}
	log.Println("\nTransaction response: ", resp)
}

/*
	2019/10/15 18:40:08 Seed    A: SCWK6NI3ELM7C6CKFK2ZGE6XJVBO2LOHQPLEZ2NAZ2QZCPLJTCGGXICZ
	2019/10/15 18:40:08 Address A: GDT5NJNXEE2XZPQMMY5SSA6BWX6VLP4OEIJQ3PPR3IP5QSKNCEJMI3JW
	2019/10/15 18:40:12 Seed    B: SBE4LNICHUZKY335FP4QCHZBE3B4XZTJENWM2F3GHSV7A7QIA7KGHLXQ
	2019/10/15 18:40:12 Address B: GBV7TSUF5EPPH566LFEUY4OKAAHAYWC5JLHLVMXAJE5RYWW2ZUGC3QZH

*/

//公钥
var APubAddr  = "GDT5NJNXEE2XZPQMMY5SSA6BWX6VLP4OEIJQ3PPR3IP5QSKNCEJMI3JW"
//私钥，不对外公开
var APrivSeed = "SCWK6NI3ELM7C6CKFK2ZGE6XJVBO2LOHQPLEZ2NAZ2QZCPLJTCGGXICZ"

var BPubAddr  = "GBV7TSUF5EPPH566LFEUY4OKAAHAYWC5JLHLVMXAJE5RYWW2ZUGC3QZH"
var BPrivSeed = "SBE4LNICHUZKY335FP4QCHZBE3B4XZTJENWM2F3GHSV7A7QIA7KGHLXQ"


func CreateClint() *horizonclient.Client {
	horizonUrl := "http://3.15.200.138:8000"

	client := &horizonclient.Client{
		HorizonURL: horizonUrl,
		HTTP:		http.DefaultClient,
		AppName:	"GoTest",
		AppVersion:	"1.0",
	}
	return client
}

func test1() {

	// Generate a new randomly generated address
	//kp1, err  := createAccount()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("Seed    A:", kp1.Seed())
	//log.Println("Address A:", kp1.Address())

	// Create and fund the address on TestNet, using friendbot
	clientA := horizonclient.DefaultTestNetClient
	//clientA.Fund(kp1.Address())
	clientA.Fund("GDT5NJNXEE2XZPQMMY5SSA6BWX6VLP4OEIJQ3PPR3IP5QSKNCEJMI3JW")

	//查询余额
	getBlance(clientA, "GDT5NJNXEE2XZPQMMY5SSA6BWX6VLP4OEIJQ3PPR3IP5QSKNCEJMI3JW")

	fmt.Println("-----------------")
	//创建账户2
	//kp2, err  := createAccount()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("Seed    B:", kp2.Seed())
	//log.Println("Address B:", kp2.Address())

	//使用friendbot TestNet 创建测试Client
	// Create and fund the address on TestNet, using friendbot
	clientB := horizonclient.DefaultTestNetClient
	clientB.Fund("GBV7TSUF5EPPH566LFEUY4OKAAHAYWC5JLHLVMXAJE5RYWW2ZUGC3QZH")

	//查询余额2
	getBlance(clientB, "GBV7TSUF5EPPH566LFEUY4OKAAHAYWC5JLHLVMXAJE5RYWW2ZUGC3QZH")

	//转账



}

func main() {
	//创建客户端
	client := CreateClint()
	fmt.Printf("client: %+v \n", client)
	//创建账户

	//txSuccess, err := client.Fund(APubAddr)
	//if err != nil {
	//	fmt.Printf("client.Fund err: %v \n", err)
	//}
	//fmt.Printf("txSuccess: %+v \n", txSuccess)
	//getBlance(client, APubAddr)


}