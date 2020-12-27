package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/parnurzeal/gorequest"
)

const DiscordAPI = "https://discord.com/api"

type CustomOverride struct {
	CustomName    string
	CustomIconURL string
	UseTTS        bool
}

type DWHWebhook struct {
	discordgo.Webhook
	Custom *CustomOverride
}

type requestBody struct {
	Content   string                    `json:"content"`
	Username  string                    `json:"username"`
	AvatarURL string                    `json:"avatar_url"`
	TTS       bool                      `json:"tts"`
	Embeds    []*discordgo.MessageEmbed `json:"embeds"`
}

func (webhook *DWHWebhook) PostEmbeds(embeds []*discordgo.MessageEmbed) []error {
	body := &requestBody{}
	if webhook.Custom != nil {
		if webhook.Custom.CustomName != "" {
			body.Username = webhook.Custom.CustomName
		}

		if webhook.Custom.CustomIconURL != "" {
			body.AvatarURL = webhook.Custom.CustomIconURL
		}

		if webhook.Custom.UseTTS {
			body.TTS = true
		}
	}

	body.Embeds = embeds

	return executeWebhook(webhook, body)
}

func (webhook *DWHWebhook) PostText(text string) []error {
	body := &requestBody{}
	if webhook.Custom != nil {
		if webhook.Custom.CustomName != "" {
			body.Username = webhook.Custom.CustomName
		}

		if webhook.Custom.CustomIconURL != "" {
			body.AvatarURL = webhook.Custom.CustomIconURL
		}

		if webhook.Custom.UseTTS {
			body.TTS = true
		}
	}

	body.Content = text

	return executeWebhook(webhook, body)
}

func (webhook *DWHWebhook) PostTextAndEmbeds(text string, embeds []*discordgo.MessageEmbed) []error {
	body := &requestBody{}
	if webhook.Custom != nil {
		if webhook.Custom.CustomName != "" {
			body.Username = webhook.Custom.CustomName
		}

		if webhook.Custom.CustomIconURL != "" {
			body.AvatarURL = webhook.Custom.CustomIconURL
		}

		if webhook.Custom.UseTTS {
			body.TTS = true
		}
	}

	body.Content = text
	body.Embeds = embeds

	return executeWebhook(webhook, body)
}

func executeWebhook(webhook *DWHWebhook, body *requestBody) []error {
	_, _, errs := gorequest.New().
		Post(fmt.Sprintf("%s/webhooks/%s/%s", DiscordAPI, webhook.ID, webhook.Token)).
		Send(body).
		End()

	return errs
}
