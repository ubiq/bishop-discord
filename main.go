package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/ubiq/bishop-discord/nft"
)

var (
	// Flags
	GuildID  = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken = flag.String("token", "", "Bot access token")
	RpcURL   = flag.String("rpcURL", "http://127.0.0.1:8588", "RPC URL")

	cryptsters = nft.Collection{
		Name:            "Cryptsters",
		Description:     "Cryptsters is a collection of 888 randomly generated and stylistically curated NFTs that exist on the Ubiq blockchain. Cryptsters is licensed under CC0 - You can modify, distribute and perform the work, even for commercial purposes, all without asking permission.",
		ImageBaseURL:    "https://cryptstersapi.ubiqsmart.com/image",
		MetadataBaseURL: "https://cryptstersapi.ubiqsmart.com/cryptster",
		JawacampBaseURL: "https://jawacamp.ubiqsmart.com/collections/0xc3d585288d35215754e970a23dd8de37a08e71b7",
		CountMax:        888,
		Attributes:      true,
		Revealed:        true,
	}
)

var s *discordgo.Session

func init() { flag.Parse() }

func init() {
	var err error
	s, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "chimp",
			Description: "Retrieves the Chimp NFT based on the supplied Token ID",
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "tokenid",
					Description: "Token ID",
					Required:    true,
				},
			},
		},
		{
			Name:        "cryptster",
			Description: "Retrieves the NFT based on the supplied Token ID",
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "tokenid",
					Description: "Token ID",
					Required:    true,
				},
			},
		},
		{
			Name:        "gb89",
			Description: "Retrieves the GB89 NFT based on the supplied Token ID",
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "tokenid",
					Description: "Token ID",
					Required:    true,
				},
			},
		},
		{
			Name:        "nception",
			Description: "Retrieves the nCeption NFT based on the supplied Token ID",
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "tokenid",
					Description: "Token ID",
					Required:    true,
				},
			},
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"chimp": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			tokenID := big.NewInt(i.ApplicationCommandData().Options[0].IntValue())
			var iNFT nft.NFT
			var msgformat string

			if tokenID.Cmp(big.NewInt(0)) == -1 {
				msgformat = fmt.Sprintf("Invalid Token ID: %d\n", tokenID)
			} else {
				iNFT = nft.HandleChimp(*RpcURL, tokenID)
				if iNFT.Owner.Hex() == "0x0000000000000000000000000000000000000000" {
					msgformat = "Unclaimed\n"
				}
			}
			var fields []*discordgo.MessageEmbedField
			for key, element := range iNFT.Attributes {
				fields = append(fields, &discordgo.MessageEmbedField{Name: key, Value: element})
			}
			msgembed := discordgo.MessageEmbed{
				Title:       fmt.Sprintf("CHIMP NFT #%d", tokenID),
				URL:         "https://chimp.ubiqsmart.com",
				Color:       170,
				Description: msgformat,
				Fields:      fields,
				Image: &discordgo.MessageEmbedImage{
					URL: "attachment://output.png",
				},
			}
			attachment := discordgo.File{
				Name:        "output.png",
				ContentType: "image/png",
				Reader:      bytes.NewReader(iNFT.Picture),
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{&msgembed},
					Files:  []*discordgo.File{&attachment},
				},
			})
		},
		"cryptster": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			tokenID := i.ApplicationCommandData().Options[0].IntValue()
			var iNFT nft.NFT
			var msgformat string

			if tokenID < 1 || tokenID > int64(cryptsters.CountMax) {
				msgformat = fmt.Sprintf("Invalid Token ID: %d\n", tokenID)
			} else {
				imageURL := fmt.Sprintf("%s/%d", cryptsters.ImageBaseURL, tokenID)
				msgformat = fmt.Sprintf("Tweet with hash tag #%s - Image URL: %s", cryptsters.Name, imageURL)
				iNFT = cryptsters.HandleCollection(tokenID)
			}
			var fields []*discordgo.MessageEmbedField
			for key, element := range iNFT.Attributes {
				fields = append(fields, &discordgo.MessageEmbedField{Name: key, Value: element})
			}
			msgembed := discordgo.MessageEmbed{
				Title:       fmt.Sprintf("%s #%d", cryptsters.Name, tokenID),
				URL:         fmt.Sprintf("%s/%d", cryptsters.JawacampBaseURL, tokenID),
				Color:       170,
				Description: msgformat,
				Fields:      fields,
				Image: &discordgo.MessageEmbedImage{
					URL: "attachment://output.png",
				},
			}
			attachment := discordgo.File{
				Name:        "output.png",
				ContentType: "image/png",
				Reader:      bytes.NewReader(iNFT.Picture),
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{&msgembed},
					Files:  []*discordgo.File{&attachment},
				},
			})
		},
		"gb89": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			tokenID := big.NewInt(i.ApplicationCommandData().Options[0].IntValue())
			var iNFT nft.NFT
			var msgformat string

			if tokenID.Cmp(big.NewInt(0)) == -1 || tokenID.Cmp(big.NewInt(255)) == 1 {
				msgformat = fmt.Sprintf("Invalid Token ID: %d\n", tokenID)
			} else {
				iNFT = nft.HandleGB89(*RpcURL, tokenID)
				if iNFT.Owner.Hex() == "0x0000000000000000000000000000000000000000" {
					msgformat = "Unclaimed\n"
				}
			}
			var fields []*discordgo.MessageEmbedField
			for key, element := range iNFT.Attributes {
				fields = append(fields, &discordgo.MessageEmbedField{Name: key, Value: element})
			}
			msgembed := discordgo.MessageEmbed{
				Title:       fmt.Sprintf("GB89 NFT #%d", tokenID),
				URL:         "https://ubiq.github.io/gb89/",
				Color:       170,
				Description: msgformat,
				Fields:      fields,
				Image: &discordgo.MessageEmbedImage{
					URL: "attachment://output.png",
				},
			}
			attachment := discordgo.File{
				Name:        "output.png",
				ContentType: "image/png",
				Reader:      bytes.NewReader(iNFT.Picture),
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{&msgembed},
					Files:  []*discordgo.File{&attachment},
				},
			})
		},
		"nception": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			tokenID := big.NewInt(i.ApplicationCommandData().Options[0].IntValue())
			var iNFT nft.NFT
			var msgformat string

			if tokenID.Cmp(big.NewInt(1)) == -1 || tokenID.Cmp(big.NewInt(8888)) == 1 {
				msgformat = fmt.Sprintf("Invalid Token ID: %d\n", tokenID)
			} else {
				iNFT = nft.HandleNception(*RpcURL, tokenID)
				if iNFT.Owner.Hex() == "0x0000000000000000000000000000000000000000" {
					msgformat = "Unclaimed\n"
				}
			}
			var fields []*discordgo.MessageEmbedField
			for key, element := range iNFT.Attributes {
				fields = append(fields, &discordgo.MessageEmbedField{Name: key, Value: element})
			}
			msgembed := discordgo.MessageEmbed{
				Title:       fmt.Sprintf("nCeption NFT #%d", tokenID),
				URL:         "https://nception.ubiqsmart.com",
				Color:       170,
				Description: msgformat,
				Fields:      fields,
				Image: &discordgo.MessageEmbedImage{
					URL: "attachment://output.png",
				},
			}
			attachment := discordgo.File{
				Name:        "output.png",
				ContentType: "image/png",
				Reader:      bytes.NewReader(iNFT.Picture),
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{&msgembed},
					Files:  []*discordgo.File{&attachment},
				},
			})
		},
	}
)

func init() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is up!")
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	for _, v := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, *GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Gracefully shutting down")
}
