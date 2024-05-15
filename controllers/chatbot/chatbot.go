package chatbot

import (
	"habit/controllers/chatbot/request"
	"habit/controllers/chatbot/response"
	chatbotEntities "habit/entities/chatbot"
	"habit/utilities/base"
	"net/http"

	"github.com/labstack/echo/v4"
	openai "github.com/sashabaranov/go-openai"
)

type ChatbotController struct {
	chatbotUseCase chatbotEntities.UseCaseInterface
}

func NewChatbotController(chatbotUseCase chatbotEntities.UseCaseInterface) *ChatbotController {
	return &ChatbotController{
		chatbotUseCase: chatbotUseCase,
	}
}

func (chatbotController *ChatbotController) Chat(c echo.Context) error {
	var chatbotRequest request.ChatbotRequest
	c.Bind(&chatbotRequest)

	chatbotEnt := chatbotEntities.Chatbot{
		Content: chatbotRequest.Content,
	}

	chatbotEnt.Role = openai.ChatMessageRoleUser

	chatbot, err := chatbotController.chatbotUseCase.Chat(&chatbotEnt)

	chatbotResp := response.ChatbotResponse{
		Role:    chatbot.Role,
		Content: chatbot.Content,
	}
	
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success", chatbotResp))
}

