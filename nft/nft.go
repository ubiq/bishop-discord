package nft

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"math/big"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"
	n "github.com/ubiq/bishop-discord/contracts"
	"github.com/ubiq/go-ubiq/v5/common"
	"github.com/ubiq/go-ubiq/v5/ethclient"
)

type NFT struct {
	Owner      common.Address
	Attributes map[string]string
	Picture    []byte
}

type NceptionURI struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
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
	// Attributes

	// Picture
	tokenURIbase64, err := instance.TokenURI(nil, TokenID)
	if err != nil {
		log.Println(err)
	}
	log.Println("tokenURI:", tokenURIbase64)
	nCeptionURI := decodeNceptionTokenURIbase64(tokenURIbase64)
	// nft.Picture needs to be in a format which can be displayed by Discord
	nft.Picture = svgToImageLibvips(nCeptionURI.Image)

	return nft
}

func decodeNceptionTokenURIbase64(tokenURIbase64 string) NceptionURI {
	// TODO: Return attributes as well
	tokenURI, _ := base64.StdEncoding.DecodeString(strings.TrimPrefix(tokenURIbase64, "data:application/json;base64,"))

	token := NceptionURI{}
	json.Unmarshal(tokenURI, &token)

	tokenImage, _ := base64.StdEncoding.DecodeString(strings.TrimPrefix(token.Image, "data:image/svg+xml;base64,"))
	token.Image = string(tokenImage)

	return token
}

func svgToImageLibvips(svgImage string) []byte {
	image, err := vips.NewImageFromReader(strings.NewReader(svgImage))
	if err != nil {
		log.Println("vips.NewImageFromReader:", err)
		return nil
	}

	encodedPNG := vips.NewDefaultPNGExportParams()
	image1bytes, _, _ := image.Export(encodedPNG)
	return image1bytes
}
