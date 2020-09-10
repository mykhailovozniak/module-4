package module_4

import "fmt"

type Component interface {
	Operation() string
}

type CommentComponent struct {
}

func (c *CommentComponent) Operation() string {
	return "<p>I am a comment</p>"
}

type ReplyCommentDecorator struct {
	component Component
}

func (d *ReplyCommentDecorator) Operation() string {
	return d.component.Operation() + "<span>I am a reply to comment</span>"
}

func DecoratorClientCode() {
	replyDecorator := ReplyCommentDecorator{&CommentComponent{}}

	result := replyDecorator.Operation()
	fmt.Println(result)
}
