package prompts

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

const SysExecStatusIsFinished = `You are an AI assistant that determines the execution status of a coding AI's plan for a programming task. Analyze the AI's latest message to determine whether the plan is finished. 

The plan is finished if all the plan's tasks and subtasks have been completed. When a plan is finished, the coding AI will say something like "All tasks have been completed." If the response is a list of tasks, then the plan is not finished. If the response is a list of tasks and a message saying that all tasks have been completed, then the plan is finished.

Return a JSON object with the 'finished' key set to true or false. Only call the 'planIsFinished' function in your response. Don't call any other function.`

func GetExecStatusIsFinishedPrompt(conversation []openai.ChatCompletionMessage) string {
	s := ""
	for _, m := range conversation {
		s += m.Role + ":\n"
		s += m.Content + "\n"
	}

	return SysExecStatusIsFinished + "\n\nConversation:\n" + s
}

var PlanIsFinishedFn = openai.FunctionDefinition{
	Name: "planIsFinished",
	Parameters: &jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"finished": {
				Type:        jsonschema.Boolean,
				Description: "Whether the plan is finished.",
			},
		},
		Required: []string{"finished"},
	},
}

const SysExecStatusNeedsInput = `You are an AI assistant that determines the execution status of a coding AI's plan for a programming task. Analyze the AI's latest message to determine whether the plan needs more input. The plan needs more input if the coding AI requires the user to add more context, provide information, or answer questions the AI has asked.

When the coding AI needs more input, it will say something like "I need more information or context to make a plan for this task."

If the coding AI says or implies that additional information would be helpful or useful, but that information isn't *required* to continue the plan, then the plan *does not* need more input. It only needs more input if the AI says or implies that more information is necessary and required to continue. Return a JSON object with the 'needs_input' key set to true or false. Only call the 'planNeedsInput' function in your response. Don't call any other function.`

func GetExecStatusNeedsInputPrompt(message *openai.ChatCompletionMessage) string {
	return SysExecStatusNeedsInput + "\nLatest message from coding AI:\n" + message.Content
}

var PlanNeedsInputFn = openai.FunctionDefinition{
	Name: "planNeedsInput",
	Parameters: &jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"needs_input": {
				Type:        jsonschema.Boolean,
				Description: "Whether the plan needs more input. If ambiguous or unclear, assume the plan does not need more input.",
			},
		},
		Required: []string{"needs_input"},
	},
}
