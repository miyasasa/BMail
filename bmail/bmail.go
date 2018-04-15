package bmail

type Bmail struct {
	From        string
	Recipients  []string
	Cc          []string
	Subject     string
	Content     Content
	Attachments []Attachment
}

type Content struct {
	ContentType string
	Body        string
}

type Attachment struct {
	Name        string
	ContentType string
	Body        string
}
