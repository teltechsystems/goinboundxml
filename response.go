package inboundxml

import (
	"bytes"
)

type Response struct {
	verbs []Verb
}

func (r *Response) Add(verb Verb) {
	r.verbs = append(r.verbs, verb)
}

func (r *Response) Dial(addresser DialAddresser, attrs *DialAttrs) {
	r.verbs = append(r.verbs, NewDial(addresser, attrs))
}

func (r *Response) Gather(innerVerb Verb, attrs *GatherAttrs) {
	r.verbs = append(r.verbs, NewGather(innerVerb, attrs))
}

func (r *Response) Hangup(attrs *HangupAttrs) {
	r.verbs = append(r.verbs, NewHangup(attrs))
}

func (r *Response) Pause(timeout int) {
	r.verbs = append(r.verbs, NewPause(timeout))
}

func (r *Response) Play(noun string, attrs *PlayAttrs) {
	r.verbs = append(r.verbs, NewPlay(noun, attrs))
}

func (r *Response) Ping(noun string, method string) {
	r.verbs = append(r.verbs, NewPing(noun, method))
}

func (r *Response) PreAnswer(innerVerbs ...Verb) {
	r.verbs = append(r.verbs, NewPreAnswer(innerVerbs...))
}

func (r *Response) Record(attrs *RecordAttrs) {
	r.verbs = append(r.verbs, NewRecord(attrs))
}

func (r *Response) Redirect(noun string, attrs *RedirectAttrs) {
	r.verbs = append(r.verbs, NewRedirect(noun, attrs))
}

func (r *Response) Reject() {
	r.verbs = append(r.verbs, NewReject())
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
