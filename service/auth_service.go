package service

import (
	"context"
	"sacramento/config"
	"sacramento/utils"
)

func CreateCustomToken(uid string) string {
	client, err := FirebaseAdmin.Auth(context.Background())
	if err != nil {
		utils.SugarLogger.Errorln("error getting Auth client: %v\n", err)
		return ""
	}
	token, err := client.CustomToken(context.Background(), uid)
	if err != nil {
		utils.SugarLogger.Errorln("error minting custom token: %v\n", err)
		return ""
	}
	utils.SugarLogger.Info("Got custom token for user \"" + uid + "\": " + token[:16] + "..." + token[len(token)-16:])
	go Discord.ChannelMessageSend(config.DiscordChannel, "User \""+uid+"\" just logged in with token: `"+token[:16]+"..."+token[len(token)-16:]+"`")
	return token
}
