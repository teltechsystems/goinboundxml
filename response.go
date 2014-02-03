package inboundxml

import (
	"bytes"
)

type Response struct {
	verbs []Verb
}

func (r *Response) Hangup() {
	r.verbs = append(r.verbs, NewHangup())
}

func (r *Response) Play(noun string, attrs *PlayAttrs) {
	r.verbs = append(r.verbs, NewPlay(noun, attrs))
}

func (r *Response) Record(attrs *RecordAttrs) {
	r.verbs = append(r.verbs, NewRecord(attrs))
}

func (r *Response) Reset() {
	r.verbs = nil
}

func (r *Response) Say(noun string, attrs *SayAttrs) {
	r.verbs = append(r.verbs, NewSay(noun, attrs))
}

func (r *Response) String() string {
	buffer := bytes.NewBuffer([]byte{})
	buffer.WriteString("<?xml version=\"1.0\" encoding=\"utf-8\"?>")
	buffer.WriteString("<Response>")

	for _, verb := range r.verbs {
		buffer.WriteString(verb.String())
	}

	buffer.WriteString("</Response>")

	return buffer.String()
}

func NewResponse() *Response {
	return &Response{}
}
