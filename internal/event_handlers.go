package internal

type (
	MenuButtonClickedEventHandler          func()
	SubmitButtonClickedEventHandler        func(bool)
	PullRequestMenuItemClickedEventHandler func()
	AnalyzeButtonClickedEventHandler       func()
	OnGetPullRequestEvent                  func(pullRequest string)
)
