package internal

type (
	OnMenuButtonClickedEvent          func()
	OnSubmitButtonClickedEvent        func(bool)
	OnPullRequestMenuItemClickedEvent func()
	OnAnalyzeButtonClickedEvent       func()
	OnGetPullRequestEvent             func(gptModel, pullRequest string)
)
