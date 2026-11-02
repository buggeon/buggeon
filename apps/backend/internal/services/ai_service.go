package services

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type AIService struct {
	client         openai.Client
	projectService *ProjectService
	memberService  *MemberService
}

func NewAIService(projectService *ProjectService, memberService *MemberService) *AIService {

	insecureHTTPClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	return &AIService{
		client: openai.NewClient(
			option.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
			option.WithBaseURL("https://api.giga.chat/v1"),
			option.WithHTTPClient(insecureHTTPClient),
		),
		projectService: projectService,
		memberService:  memberService,
	}
}

func (s *AIService) makeRequest(ctx context.Context, request string) (string, error) {

	result, err := s.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			{
				OfUser: &openai.ChatCompletionUserMessageParam{
					Content: openai.ChatCompletionUserMessageParamContentUnion{
						OfString: openai.String(request),
					},
				},
			},
		},
		Model: os.Getenv("OPENAI_MODEL"),
	})

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return result.Choices[0].Message.Content, nil

}

func (s *AIService) FullAnalytic(ctx *gin.Context, projectID string, request string) (string, error) {

	projectInfo, err := s.fetchFullProjectInfoViaGraphQL(ctx, projectID)

	if err != nil {
		return "", err
	}

	return s.makeRequest(ctx, fmt.Sprintf("%s\n\nИнформация о проекте: %s", request, projectInfo))

}

func (s *AIService) fetchFullProjectInfoViaGraphQL(ctx *gin.Context, projectID string) (string, error) {

	query := `
		query GetProject($projectId: ID!) {
			project(projectId: $projectId) {
				id,
				name,
				progress,
				lead{
					id,
					role,
					directions,

					user{
						name
					}
				},
				members {
					id,
					role,
					directions,

					user{
						name
					}
				},
				boards{
					id,
					name,
					cardsStatus,

					cards {
						id,
						title,
						content,
						priority,
						createdAt,
						updatedAt,
						dueDate,
						status,

						assignees{
							id,
							role,
							directions,

							user{
								name
							}
						}

						messages{
							id,
							content,
							createdAt,
							updatedAt,
							
							sender{
								name
							}
						}
					}
				}
			}
		}
	`

	fmt.Println("ProjectId: " + projectID)

	payload := map[string]any{
		"query":     query,
		"variables": map[string]any{"projectId": projectID},
	}

	bodyBytes, _ := json.Marshal(payload)

	req, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:9187/api/query", bytes.NewReader(bodyBytes))

	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", ctx.GetHeader("Authorization"))

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	respBytes, _ := io.ReadAll(resp.Body)

	return string(respBytes), nil

}
