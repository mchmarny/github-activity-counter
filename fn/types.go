package fn

type gitHubWebHook struct {
	sig     string
	event   string
	id      string
	content []byte
}
