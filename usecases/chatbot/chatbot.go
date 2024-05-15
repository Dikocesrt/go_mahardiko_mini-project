package chatbot

import (
	"context"
	"habit/configs"
	"habit/constants"
	chatbotEntities "habit/entities/chatbot"

	"github.com/sashabaranov/go-openai"
)

func getCompletionFromMessages(
	ctx context.Context,
	client *openai.Client,
	messages []openai.ChatCompletionMessage,
	model string,
) (openai.ChatCompletionResponse, error) {
	if model == "" {
		model = openai.GPT3Dot5Turbo // see another model options: https://platform.openai.com/docs/models
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)
	return resp, err
}

func Chat(chat *chatbotEntities.Chatbot) (*chatbotEntities.Chatbot, error) {
	if chat.Content == "" {
		return nil, constants.ErrChatEmptyInput
	}

	ctx := context.Background()

	client := openai.NewClient(configs.InitConfigKeyChatbot())

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are a friendly chatbot.",
		},
	}
	model := openai.GPT3Dot5Turbo

	if chat.Content != "" {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: chat.Content,
		})
		chat.Content = ""
	}

	resp, err := getCompletionFromMessages(ctx, client, messages, model)

	if err != nil {
		return nil, constants.ErrChatbotServer
	}

	answer := openai.ChatCompletionMessage{
		Role:    resp.Choices[0].Message.Role,
		Content: resp.Choices[0].Message.Content,
	}

	chatAnswerEnt := chatbotEntities.Chatbot{
		Role:    answer.Role,
		Content: answer.Content,
	}

	return &chatAnswerEnt, nil
}