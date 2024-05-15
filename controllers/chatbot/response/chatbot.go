package response

type ChatbotResponse struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}