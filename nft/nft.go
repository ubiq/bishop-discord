package nft

import (
	"log"
	"math/big"

	n "github.com/ubiq/bishop-discord/contracts"
	"github.com/ubiq/go-ubiq/v5/common"
	"github.com/ubiq/go-ubiq/v5/ethclient"
)

type NFT struct {
	Owner      common.Address
	Attributes map[string]string
	Picture    string
}

func HandleNception(RpcURL string, TokenID *big.Int) NFT {
	client, err := ethclient.Dial(RpcURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	// nCeption ERC721 address
	contractAddress := common.HexToAddress("0x81f1a6e026d49c2260a8D6D8e14Bca1628c1Df43")
	instance, err := n.NewNCaller(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// Stats
	// Title "nCeption #x" Details "Owner, Attributes, Picture"
	var nft NFT
	ownerOf, err := instance.OwnerOf(nil, TokenID)
	if err != nil {
		// TODO:
		// Handle error
		// 2021/10/15 22:37:26 execution reverted: ERC721: owner query for nonexistent token
		// 2021/10/15 22:37:26 ownerOf: 0x0000000000000000000000000000000000000000
		log.Println(err)
		return nft
	}
	log.Println("ownerOf:", ownerOf)
	nft.Owner = ownerOf

	return nft
}
