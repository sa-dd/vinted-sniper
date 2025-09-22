package discord

import (

	"fmt"
	"github.com/bwmarrin/discordgo"
	"vinted-sniper/client"
)

var (
	embedSlice []*discordgo.MessageEmbed
	embedFields []*discordgo.MessageEmbedField
)

func CreateEmbed(items []client.Item) []*discordgo.MessageEmbed{

	for _, item := range items {
		embedFields = []*discordgo.MessageEmbedField{
			{
			Name: "Status",
			Value: item.Status,
			Inline: true,
			},
			{
			Name: "Favorites",
			Value: fmt.Sprintf("%d", item.FavoriteCount),
			Inline: true,
			},
			{
			Name: "Price",
			Value: item.Price.Amount,
			Inline: true,
			},
			{
			Name: "Size",
			Value: item.SizeTitle,
			Inline: true,
			},
		}

		embed := &discordgo.MessageEmbed{
			Title: item.Title,
			Author: &discordgo.MessageEmbedAuthor{
				Name: item.User.Name,
				IconURL: item.User.Pic.Url,
			},
			URL: item.Url,
			Fields: embedFields,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: item.Photo.Thumbnails[0].Url,
			},
		}
		embedSlice = append(embedSlice, embed)

	}

	return embedSlice
}
