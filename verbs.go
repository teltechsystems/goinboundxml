package inboundxml

import (
	"bytes"
	"fmt"
)

type Verb interface {
	fmt.Stringer
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
