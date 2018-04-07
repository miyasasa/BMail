package bmail

type Bmail struct {
	From       string
	Recipients []string
	Cc         []string
	Subject    string
	Content    Content
}

type Content struct {
	ContentType string
	Body        string
}
