package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

// Struct for the OpenAI API request
type OpenAIChatRequest struct {
	Model       string        `json:"model"`
	Messages    []ChatMessage `json:"messages"`
	Temperature float64       `json:"temperature"`
	MaxTokens   int           `json:"max_tokens"`
	TopP        float64       `json:"top_p"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIChatResponse struct {
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
}

// Retrieves the OpenAI API key from the environment variables.
func getOpenAIKey() (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		prompt := promptui.Prompt{
			Label: "OpenAI API Key not found in environment. Please enter your OpenAI API Key",
			Mask:  '*',
		}
		apiKey, err := prompt.Run()
		if err != nil {
			return "", err
		}
		// Optional: Set the environment variable for future use
		os.Setenv("OPENAI_API_KEY", apiKey)
	}
	return apiKey, nil
}

// Retrieves the system information.
func getSystemInfo() (string, string) {
	osInfo := runtime.GOOS
	var shellVersion string
	if osInfo == "windows" {
		shellVersion = getWindowsShellVersion()
	} else {
		shellVersion = getUnixShellVersion()
	}
	return osInfo, shellVersion
}

// Retrieves the PowerShell version on Windows systems.
func getWindowsShellVersion() string {
	cmd := exec.Command("powershell", "$PSVersionTable.PSVersion")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	return string(output)
}

// Retrieves the shell version on Unix-like systems.
func getUnixShellVersion() string {
	cmd := exec.Command("sh", "-c", "echo $SHELL --version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	return string(output)
}

// Encapsulates the OpenAI API call.
func callOpenAI(apiKey string, messages []ChatMessage, model string, maxTokens int) (string, error) {
	requestBody := OpenAIChatRequest{
		Model:       model,
		Messages:    messages,
		Temperature: 1.0,
		MaxTokens:   maxTokens,
		TopP:        1.0,
	}

	requestBodyJson, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBodyJson))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var openaiResp OpenAIChatResponse
	if err := json.Unmarshal(body, &openaiResp); err != nil {
		return "", err
	}

	if len(openaiResp.Choices) > 0 {
		return strings.TrimSpace(openaiResp.Choices[0].Message.Content), nil
	}

	return "", fmt.Errorf("no response from OpenAI")
}

// Generates a command based on user input.
func generateCommand(apiKey, osInfo, shellVersion, prompt string, debug bool) (string, error) {
	systemPrompt := fmt.Sprintf(
		`# Instruction
     - You are an assistant for a %s operating system with the following details:
     - You must only provide the Script, without any additional explanation or text like description.
     - Your responses should be informative, visually appealing, logical and actionable.
     - Your responses should be very simple and complete.

    # Script Creation Rules
     - OS Information: %s
     - Shell Version: %s
     - Based on the user's input, generate a script to accomplish the task.
     - To distinguish between each server, print the hostnames on environment variables.`, osInfo, osInfo, shellVersion)

	messages := []ChatMessage{
		{Role: "system", Content: systemPrompt},
		{Role: "user", Content: prompt},
	}

	response, err := callOpenAI(apiKey, messages, "gpt-4o", 256)
	if err != nil {
		return "", err
	}

	// Extract only the script part from the response
	command := extractScript(response)

	// Print the generated command in debug mode
	if debug {
		color.Yellow("Generated Command: %s", command)
	}

	return command, nil
}

// Interprets the output of a command based on the user's query.
func interpretCommandOutput(apiKey, commandOutput, userQuery string, debug bool) (string, error) {
	interpretationPrompt := fmt.Sprintf(
		`You need to provide a detailed explanation of the results of executing a script.
    The user should not know that the explanation is based on the script's results.
    ---
    User Query: %s
    Script Execution Results: %s
    ---
    Refer to the script execution results to respond simply to the user's query.
    If the query is simply to run a specific program, respond that the program has been executed.
    Always respond in the user's language.`, userQuery, commandOutput)

	messages := []ChatMessage{
		{Role: "system", Content: interpretationPrompt},
	}

	return callOpenAI(apiKey, messages, "gpt-4o", 256)
}

// Generates an initial response to the user's input.
func generateResponse(apiKey, userQuery string, debug bool) (string, error) {
	responsePrompt := fmt.Sprintf(
		`You need to provide a response to the user's input.
    The response should acknowledge the user's request and indicate that you are processing it.
    ---
    User Query: %s
    ---
    Respond in a friendly and concise manner. Always respond in the user's language.`, userQuery)

	messages := []ChatMessage{
		{Role: "system", Content: responsePrompt},
	}

	return callOpenAI(apiKey, messages, "gpt-4o", 100)
}

// Determines whether the user's query is a new task or a simple question.
func isTaskOrQuery(apiKey, query string) (string, error) {
	// Generate the prompt to send to OpenAI API.
	taskOrQueryPrompt := fmt.Sprintf(
		`You need to determine if the user's query requires executing a script, is a simple question, or is potentially dangerous.
    Respond with "y" if it requires executing a script, "n" if it is a simple question, and "w" if it is a potentially dangerous task.
    ---
    User Query: %s`, query)

	messages := []ChatMessage{
		{Role: "system", Content: taskOrQueryPrompt},
	}

	response, err := callOpenAI(apiKey, messages, "gpt-4o", 50)
	if err != nil {
		return "", err
	}

	// The response should be either "y", "n", or "w".
	return strings.TrimSpace(response), nil
}

// Extracts only the script part from the OpenAI response.
func extractScript(response string) string {
	lines := strings.Split(response, "\n")
	var scriptLines []string
	for _, line := range lines {
		// Skip lines starting with ```
		if strings.HasPrefix(strings.TrimSpace(line), "```") {
			continue
		}
		scriptLines = append(scriptLines, line)
	}
	return strings.Join(scriptLines, "\n")
}

// Executes the command and returns the output.
func runCommand(osInfo, command string, debug bool) (string, error) {
	var cmd *exec.Cmd
	if osInfo == "windows" {
		cmd = exec.Command("powershell", "-Command", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}

	// Print the command being executed in debug mode
	if debug {
		color.Yellow("Executing Command: %s", command)
	}

	output, err := cmd.CombinedOutput()

	// Print the command execution result in debug mode
	if debug {
		if err != nil {
			color.Red("Command execution error: %v", err)
		}
		color.Yellow("Command Output: %s", string(output))
	}

	return string(output), err
}

func main() {
	debug := flag.Bool("d", false, "Enable debug mode")
	flag.Parse()

	apiKey, err := getOpenAIKey()
	if err != nil {
		fmt.Printf("Failed to retrieve OpenAI API key: %v\n", err)
		return
	}

	osInfo, shellVersion := getSystemInfo()
	color.Cyan("OS: %s\nShell Version: %s", osInfo, shellVersion)

	for {
		prompt := promptui.Prompt{
			Label: fmt.Sprintf("%s %s $", color.YellowString("ShellScribeAI"), color.GreenString(osInfo)),
			// Prompt for the user's input command
		}
		query, err := prompt.Run()
		if err != nil || strings.ToLower(query) == "exit" || strings.ToLower(query) == "quit" {
			// Exit the program on "exit" or "quit" command
			color.Blue("Exiting the program. Goodbye!")
			break
		}

		if strings.TrimSpace(query) == "" {
			continue
		}

		// Determine if the user's query is a task, a simple question, or potentially dangerous
		queryType, err := isTaskOrQuery(apiKey, query)
		if err != nil {
			color.Red("Failed to determine the nature of the query: %v\n", err)
			continue
		}

		switch queryType {
		case "y":
			// Generate a response indicating the task is being processed
			response, err := generateResponse(apiKey, "Processing your request. Please wait a moment.", *debug)
			if err != nil {
				color.Red("Failed to generate response: %v\n", err)
				continue
			}
			color.Green("%s", response)

			// Generate the command to execute
			generatedCommand, err := generateCommand(apiKey, osInfo, shellVersion, query, *debug)
			if err != nil {
				color.Red("Failed to generate command: %v\n", err)
				continue
			}

			// Execute the command and interpret the result
			commandOutput, err := runCommand(osInfo, generatedCommand, *debug)
			if err != nil {
				color.Red("Failed to execute command: %v\n", err)
				continue
			}

			interpretation, err := interpretCommandOutput(apiKey, commandOutput, query, *debug)
			if err != nil {
				color.Red("Failed to interpret command output: %v\n", err)
				continue
			}
			color.Green("%s", interpretation)

		case "n":
			// Generate a response for a simple query
			response, err := generateResponse(apiKey, query, *debug)
			if err != nil {
				color.Red("Failed to generate response: %v\n", err)
				continue
			}
			color.Green("%s", response)

		case "w":
			// Generate a response indicating the task may be dangerous and prompt for confirmation
			response, err := generateResponse(apiKey, "The requested task may be dangerous. Please confirm the command below.", *debug)
			if err != nil {
				color.Red("Failed to generate response: %v\n", err)
				continue
			}
			color.Red("%s", response)

			// Generate the command for the potentially dangerous task
			generatedCommand, err := generateCommand(apiKey, osInfo, shellVersion, query, *debug)
			if err != nil {
				color.Red("Failed to generate command: %v\n", err)
				continue
			}

			color.Yellow("Generated Command: %s", generatedCommand)

			// Prompt the user to confirm execution of the command
			confirmPrompt := promptui.Prompt{
				Label:     "Do you want to execute this command? (yes/no)",
				IsConfirm: true,
			}
			_, err = confirmPrompt.Run()
			if err != nil {
				color.Yellow("Command execution canceled.")
				continue
			}

			// Execute the command and interpret the result
			commandOutput, err := runCommand(osInfo, generatedCommand, *debug)
			if err != nil {
				color.Red("Failed to execute command: %v\n", err)
				continue
			}

			interpretation, err := interpretCommandOutput(apiKey, commandOutput, query, *debug)
			if err != nil {
				color.Red("Failed to interpret command output: %v\n", err)
				continue
			}
			color.Green("%s", interpretation)

		default:
			color.Red("Invalid query type received: %s", queryType)
		}
	}
}
