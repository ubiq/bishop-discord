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

// Bot parameters
var (
	GuildID  = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken = flag.String("token", "", "Bot access token")
	RpcURL   = flag.String("rpcURL", "http://127.0.0.1:8588", "RPC URL")
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
		{
			Name:        "subterfuge",
			Description: "Retrieves the Subterfuge NFT based on the supplied Token ID",
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
		"nception": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// NFT handling
			tokenID := big.NewInt(i.ApplicationCommandData().Options[0].IntValue())
			var nCeptionNft nft.NFT
			var msgformat string

			if tokenID.Cmp(big.NewInt(1)) == -1 || tokenID.Cmp(big.NewInt(8888)) == 1 {
				msgformat = fmt.Sprintf("Invalid Token ID: %d\n", tokenID)
			} else {
				nCeptionNft = nft.HandleNception(*RpcURL, tokenID)
				if nCeptionNft.Owner.Hex() == "0x0000000000000000000000000000000000000000" {
					msgformat = "Unclaimed\n"
				}
			}
			var fields []*discordgo.MessageEmbedField
			for key, element := range nCeptionNft.Attributes {
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
				Reader:      bytes.NewReader(nCeptionNft.Picture),
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{&msgembed},
					Files:  []*discordgo.File{&attachment},
				},
			})
		},
		"subterfuge": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// NFT handling
			tokenID := big.NewInt(i.ApplicationCommandData().Options[0].IntValue())
			var subterfugeNFT nft.NFT
			var msgformat string

			if tokenID.Cmp(big.NewInt(1)) == -1 || tokenID.Cmp(big.NewInt(8888)) == 1 {
				msgformat = fmt.Sprintf("Invalid Token ID: %d\n", tokenID)
			} else {
				subterfugeNFT = nft.HandleSubterfuge(*RpcURL, tokenID)
				if subterfugeNFT.Owner.Hex() == "0x0000000000000000000000000000000000000000" {
					msgformat = "Unclaimed\n"
				}
			}
			var fields []*discordgo.MessageEmbedField
			for key, element := range subterfugeNFT.Attributes {
				fields = append(fields, &discordgo.MessageEmbedField{Name: key, Value: element})
			}
			msgembed := discordgo.MessageEmbed{
				Title:       fmt.Sprintf("Subterfuge NFT #%d", tokenID),
				URL:         "https://subterfuge.ubiqsmart.com",
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
				Reader:      bytes.NewReader(subterfugeNFT.Picture),
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

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Gracefully shutdowning")
}
