package internal

type (
	MenuButtonClickedEventHandler    func()
	SubmitButtonClickedEventHandler  func(bool)
	FileOpenClickedEventHandler      func()
	AnalyzeButtonClickedEventHandler func()
	OnGetPullRequestEvent            func(pullRequest string)
)
