package inboundxml

import (
	"bytes"
	"fmt"
)

type Verb interface {
	fmt.Stringer
}

type Hangup struct{}

func (v *Hangup) String() string {
	return "<Hangup />"
}

func NewHangup() *Hangup {
	return &Hangup{}
}

type SayAttrs struct {
	Loop  int
	Voice string
}

type Say struct {
	noun  string
	attrs *SayAttrs
}

func (v *Say) String() string {
	attr_buffer := bytes.NewBuffer([]byte{})

	if v.attrs != nil {
		if v.attrs.Loop >= 1 {
			attr_buffer.WriteString(fmt.Sprintf(" loop=\"%d\"", v.attrs.Loop))
		}

		if len(v.attrs.Voice) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" voice=\"%s\"", v.attrs.Voice))
		}
	}

	return fmt.Sprintf("<Say%s>%s</Say>", attr_buffer.String(), v.noun)
}

func NewSay(noun string, attrs *SayAttrs) *Say {
	return &Say{
		noun:  noun,
		attrs: attrs,
	}
}

type PlayAttrs struct {
	Loop int
}

type Play struct {
	noun  string
	attrs *PlayAttrs
}

func (v *Play) String() string {
	attr_buffer := bytes.NewBuffer([]byte{})

	if v.attrs != nil {
		if v.attrs.Loop >= 1 {
			attr_buffer.WriteString(fmt.Sprintf(" loop=\"%d\"", v.attrs.Loop))
		}
	}

	return fmt.Sprintf("<Play%s>%s</Play>", attr_buffer.String(), v.noun)
}

func NewPlay(noun string, attrs *PlayAttrs) *Play {
	return &Play{
		noun:  noun,
		attrs: attrs,
	}
}

type RecordAttrs struct {
	Action      string
	Background  bool
	Direction   string
	FinishOnKey string
	MaxLength   int
	Method      string
	PlayBeep    bool
}

type Record struct {
	attrs *RecordAttrs
}

func (v *Record) String() string {
	attr_buffer := bytes.NewBuffer([]byte{})

	if v.attrs != nil {
		if len(v.attrs.Action) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" action=\"%s\"", v.attrs.Action))
		}

		if len(v.attrs.Method) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" method=\"%s\"", v.attrs.Method))
		}

		if len(v.attrs.FinishOnKey) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" finishOnKey=\"%s\"", v.attrs.FinishOnKey))
		}

		if v.attrs.MaxLength > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" maxLength=\"%d\"", v.attrs.MaxLength))
		}

		if len(v.attrs.Direction) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" direction=\"%s\"", v.attrs.Direction))
		}

		attr_buffer.WriteString(fmt.Sprintf(" playBeep=\"%t\"", v.attrs.PlayBeep))
		attr_buffer.WriteString(fmt.Sprintf(" background=\"%t\"", v.attrs.Background))
	}

	return fmt.Sprintf("<Record%s />", attr_buffer.String())
}

func NewRecord(attrs *RecordAttrs) *Record {
	return &Record{
		attrs: attrs,
	}
}
