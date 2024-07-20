package nft

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/JoshVarga/svgparser"
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/ubiq/bishop-discord/contracts/chimp"
	"github.com/ubiq/bishop-discord/contracts/gb89"
	"github.com/ubiq/bishop-discord/contracts/n"
	"github.com/ubiq/go-ubiq/v7/common"
	"github.com/ubiq/go-ubiq/v7/ethclient"
)

type NFT struct {
	Owner      common.Address
	Attributes map[string]string
	Picture    []byte
}

type GenericSVGURI struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Image       string              `json:"image"`
	Attributes  []map[string]string `json:"attributes"`
}

type genericURI struct {
	Attributes []map[string]string `json:"attributes"`
}

type Collection struct {
	Name            string
	Description     string
	ImageBaseURL    string
	MetadataBaseURL string
	JawacampBaseURL string
	CountMax        int
	Attributes      bool
	Revealed        bool
}

func (v *Collection) HandleCollection(TokenID int64) NFT {
	var nft NFT

	res, err := http.Get(fmt.Sprintf("%s/%d", v.ImageBaseURL, TokenID))
	if err != nil {
		log.Printf("http.Get -> %v", err)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("io.ReadAll -> %v", err)
	}
	res.Body.Close()

	nft.Picture = data

	// Attributes
	res, err = http.Get(fmt.Sprintf("%s/%d", v.MetadataBaseURL, TokenID))
	if err != nil {
		log.Printf("http.Get -> %v", err)
	}
	data, err = io.ReadAll(res.Body)
	if err != nil {
		log.Printf("io.ReadAll -> %v", err)
	}
	res.Body.Close()

	token := genericURI{}
	json.Unmarshal(data, &token)
	attributes := make(map[string]string)

	for _, v := range token.Attributes {
		for kA, vA := range v {
			if kA == "trait_type" {
				attributes[vA] = v["value"]
			}
		}
	}
	nft.Attributes = attributes

	return nft
}

func HandleChimp(RpcURL string, TokenID *big.Int) NFT {
	client, err := ethclient.Dial(RpcURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	// ERC721 address
	contractAddress := common.HexToAddress("0x1e1628A35C82169F876F97C9CE5B6533895c2B22")
	instance, err := chimp.NewChimpCaller(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	var nft NFT
	ownerOf, err := instance.OwnerOf(nil, TokenID)
	if err != nil {
		log.Println(err)
		return nft
	}
	nft.Owner = ownerOf

	// Picture
	tokenURIbase64, err := instance.TokenURI(nil, TokenID)
	if err != nil {
		log.Println(err)
	}
	chimpURI := decodeGenericSVGTokenURIbase64(tokenURIbase64)
	nft.Picture = svgResizeToImageLibvips(chimpURI.Image)

	// Attributes
	attributes := make(map[string]string)
	attributes["Owner"] = ownerOf.String()
	nft.Attributes = attributes

	return nft
}

func HandleGB89(RpcURL string, TokenID *big.Int) NFT {
	client, err := ethclient.Dial(RpcURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	// ERC721 address
	contractAddress := common.HexToAddress("0x0e2fbBA88C5E526f5160Af1b9Ad79a20130b2216")
	instance, err := gb89.NewGb89Caller(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	var nft NFT
	ownerOf, err := instance.OwnerOf(nil, TokenID)
	if err != nil {
		log.Println(err)
		return nft
	}
	nft.Owner = ownerOf

	// Picture
	tokenURIbase64, err := instance.TokenURI(nil, TokenID)
	if err != nil {
		log.Println(err)
	}
	gb89URI := decodeGenericSVGTokenURIbase64(tokenURIbase64)
	nft.Picture = svgToImageLibvips(gb89URI.Image)

	// Attributes
	attributes := make(map[string]string)
	attributes["Owner"] = ownerOf.String()
	for _, v := range gb89URI.Attributes {
		for kA, vA := range v {
			if kA == "trait_type" {
				attributes[vA] = v["value"]
			}
		}
	}
	nft.Attributes = attributes

	return nft
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
	nft.Owner = ownerOf

	// Picture
	tokenURIbase64, err := instance.TokenURI(nil, TokenID)
	if err != nil {
		log.Println(err)
	}
	nCeptionURI := decodeGenericSVGTokenURIbase64(tokenURIbase64)
	nft.Picture = svgToImageLibvips(nCeptionURI.Image)

	// Attributes
	// nCeption does not have any attributes. We generate a sequence attribute.
	reader := strings.NewReader(nCeptionURI.Image)
	e, _ := svgparser.Parse(reader, false)
	sequence := []string{e.Children[2].Content, e.Children[3].Content, e.Children[4].Content, e.Children[5].Content, e.Children[6].Content, e.Children[7].Content, e.Children[8].Content, e.Children[9].Content}
	sequenceAttribute := strings.Join(sequence[:], ", ")
	attributes := make(map[string]string)
	attributes["Owner"] = ownerOf.String()
	attributes["sequence"] = sequenceAttribute
	nft.Attributes = attributes

	return nft
}

func decodeGenericSVGTokenURIbase64(tokenURIbase64 string) GenericSVGURI {
	tokenURI, _ := base64.StdEncoding.DecodeString(strings.TrimPrefix(tokenURIbase64, "data:application/json;base64,"))

	token := GenericSVGURI{}
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

func svgResizeToImageLibvips(svgImage string) []byte {
	image, err := vips.NewImageFromReader(strings.NewReader(svgImage))
	if err != nil {
		log.Println("vips.NewImageFromReader:", err)
		return nil
	}

	encodedPNG := vips.NewDefaultPNGExportParams()
	// Resize takes a scale factor so hardcoded to 32 * 16px = 512
	_ = image.Resize(32, vips.KernelNearest)
	image1bytes, _, _ := image.Export(encodedPNG)
	return image1bytes
}
