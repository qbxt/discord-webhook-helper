## Info
Designed for use with [bwmarrin's discordgo](https://github.com/bwmarrin/discordgo)

## Usage
Use type embedding to cast the `discordgo.Webhook` struct to the `discord-webhook-helper.DWHWebhook` struct. 
Alternatively, create an empty `discord-webhook-helper.DWHWebhook` struct and fill the `ID` and `Token` fields.

```go
func main() {
    // Use type embedding
    webhook := &discord-webhook-helper.DWHWebhook{
        Webhook: myDiscordGoWebhook,
    }

    // Create an empty struct
    webhook := &discord-webhook-helper.DWHWebhook{}
    webhook.ID = "700000001234567890"   
    webhook.Token = "VeryLongWebhookToken"

    // Send text
    if errs := webhook.PostText("hello, world!"); len(errs) != 0 {
        for _, err := range errs {
            fmt.Printf("%s\n", err.Error())
        }
    }

    // Send embeds
    embeds := []*discordgo.MessageEmbed{
        {
            Title: "Hello, ",
            Description: "world!",
        }
    }
    _ = webhook.PostEmbeds(embeds)

    // Send both
    _ = webhook.PostTextAndEmbeds("hello, world!", embeds)        
}
```

## Overrides
You can override the IconURL and Name using the `discord-webhook-helper.CustomOverride` struct.

```go
func main() {

    ...

    webhook.Custom = &discord-webhook-helper.CustomOverride{
        CustomName:    "Fred",
        CustomIconURL: "cdn.example.com/freds-icon.png",
    }

    ...
}
```