package client

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/x1xo/wind/store"
	"github.com/x1xo/wind/utils"
)

var cmd discordgo.ApplicationCommand = discordgo.ApplicationCommand{
	Name:        "get-key",
	Description: "Will give you a new API key to use.",
}

var client *discordgo.Session

func GetClient() *discordgo.Session {
	return client
}

func InitClient(config *utils.Config) (*discordgo.Session, error) {
	discord, err := discordgo.New("Bot " + config.Discord.Token)
	
	client = discord

	discord.State.TrackPresences = true
	discord.State.TrackMembers = true

	discord.Identify.Intents = discordgo.IntentGuilds | discordgo.IntentsGuildPresences | discordgo.IntentGuildMembers

	if err != nil {
		return nil, err
	}

	//Ready
	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Printf("%s Logged in as %s\n", utils.Format(utils.GREEN, "[INFO]"), s.State.User.Username)
		discord.ApplicationCommandCreate(discord.State.User.ID, "", &cmd)
	})

	//Interactions
	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.ApplicationCommandData().Name == "get-key" {
			key := uuid.NewString()
			err := (*store.GetStore()).SetAPIKey(i.Member.User.ID, key)
			if err != nil {
				fmt.Println(err)
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: 4, //InteractionResponseChannelMessageWithSource
					Data: &discordgo.InteractionResponseData{
						Flags: discordgo.MessageFlagsEphemeral,
						Embeds: []*discordgo.MessageEmbed{
							{
								Description: "Something went wrong while creating a new key.",
								Timestamp:   time.Now().Format(time.RFC3339),

								Footer: &discordgo.MessageEmbedFooter{
									IconURL: s.State.User.AvatarURL(""),
								},
							},
						},
					},
				})
				return
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: 4, //InteractionResponseChannelMessageWithSource
				Data: &discordgo.InteractionResponseData{
					Flags: discordgo.MessageFlagsEphemeral,
					Embeds: []*discordgo.MessageEmbed{
						{
							Title:       "Your new API key!",
							Description: fmt.Sprintf("Successfully created new API key.\n`%s`\n\nKeep in mind that after every `/get-key` command,\n a new API key will be generated.", key),
							Timestamp:   time.Now().Format(time.RFC3339),
							Thumbnail: &discordgo.MessageEmbedThumbnail{
								URL: s.State.User.AvatarURL("512"),
							},
						},
					},
				},
			})

			if err != nil {
				fmt.Println(err)
			}

			return
		}
	})

	err = discord.Open()
	if err != nil {
		return nil, err
	}

	return discord, nil
}
