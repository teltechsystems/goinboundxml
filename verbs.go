package inboundxml

import (
	"bytes"
	"fmt"
)

type Verb interface {
	fmt.Stringer
}

type dialAddresser interface {
	Verb
	addresser()
}

// DIAL VERB
type Dial struct {
	addresser dialAddresser
	attrs     *DialAttrs
}

type DialAttrs struct {
	Action            string
	Method            string
	Record            bool
	RecordCallbackUrl string
	RecordLifetime    int
	RecordDirection   string
	Timeout           int
	TimeLimit         int
	CallerId          string
	HideCallerId      bool
	DialMusic         string
	CallbackUrl       string
	CallbackMethod    string
	ConfirmSound      string
	DigitsMatch       string
	StraightToVm      bool
	HeartbeatUrl      string
	HeartbeatMethod   string
	ForwardedFrom     string
	IfMachine         string `continue,redirect,hangup`
	IfMachineUrl      string
	IfMachineMethod   string
}

func (d *Dial) String() string {
	attr_buffer := bytes.NewBuffer([]byte{})

	if d.attrs != nil {
		if len(d.attrs.Action) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" action=\"%s\"", d.attrs.Action))
		}

		if len(d.attrs.Method) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" method=\"%s\"", d.attrs.Method))
		}

		if d.attrs.Timeout > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" timeout=\"%d\"", d.attrs.Timeout))
		}

		if d.attrs.TimeLimit > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" timeLimit=\"%d\"", d.attrs.TimeLimit))
		}

		if len(d.attrs.CallerId) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" callerId=\"%s\"", d.attrs.CallerId))
		}

		if len(d.attrs.DialMusic) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" dialMusic=\"%s\"", d.attrs.DialMusic))
		}

		if len(d.attrs.CallbackUrl) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" callbackUrl=\"%s\"", d.attrs.CallbackUrl))
		}

		if len(d.attrs.CallbackMethod) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" callbackMethod=\"%s\"", d.attrs.CallbackMethod))
		}

		if len(d.attrs.ConfirmSound) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" confirmSound=\"%s\"", d.attrs.ConfirmSound))
		}

		if len(d.attrs.DigitsMatch) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" digitsMatch=\"%s\"", d.attrs.DigitsMatch))
		}

		if len(d.attrs.HeartbeatUrl) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" heartbeatUrl=\"%s\"", d.attrs.HeartbeatUrl))
		}

		if len(d.attrs.HeartbeatMethod) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" heartbeatMethod=\"%s\"", d.attrs.HeartbeatMethod))
		}

		if len(d.attrs.ForwardedFrom) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" forwardedFrom=\"%s\"", d.attrs.ForwardedFrom))
		}

		if len(d.attrs.IfMachine) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" ifMachine=\"%s\"", d.attrs.IfMachine))
		}

		if len(d.attrs.IfMachineUrl) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" ifMachineUrl=\"%s\"", d.attrs.IfMachineUrl))
		}

		if len(d.attrs.IfMachineMethod) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" ifMachineMethod=\"%s\"", d.attrs.IfMachineMethod))
		}

		if len(d.attrs.RecordCallbackUrl) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" recordCallbackUrl=\"%s\"", d.attrs.RecordCallbackUrl))
		}

		if d.attrs.RecordLifetime > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" recordLifetime=\"%d\"", d.attrs.RecordLifetime))
		}

		if len(d.attrs.RecordDirection) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" recordDirection=\"%s\"", d.attrs.RecordDirection))
		}

		attr_buffer.WriteString(fmt.Sprintf(" record=\"%t\"", d.attrs.Record))
		attr_buffer.WriteString(fmt.Sprintf(" hideCallerId=\"%t\"", d.attrs.HideCallerId))
		attr_buffer.WriteString(fmt.Sprintf(" straightToVm=\"%t\"", d.attrs.StraightToVm))

	}

	return "<Dial" + attr_buffer.String() + ">" + d.addresser.String() + "</Dial>"
}

func NewDial(addresser dialAddresser, attrs *DialAttrs) *Dial {
	return &Dial{
		addresser: addresser,
		attrs:     attrs,
	}
}

// NUMBER VERB
type Number struct {
	noun  string
	attrs *NumberAttrs
}

type NumberAttrs struct {
	SendDigits string
}

func (n *Number) addresser() {}

func (n *Number) String() string {
	attr_buffer := bytes.NewBuffer([]byte{})

	if n.attrs != nil {
		if len(n.attrs.SendDigits) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" sendDigits=\"%s\"", n.attrs.SendDigits))
		}
	}

	return "<Number" + attr_buffer.String() + ">" + n.noun + "</Number>"
}

func NewNumber(noun string, attrs *NumberAttrs) *Number {
	return &Number{
		noun:  noun,
		attrs: attrs,
	}
}

// DIAL VERB
type Gather struct {
	InnerVerb Verb
	attrs     *GatherAttrs
}

type GatherAttrs struct {
	Action      string
	Method      string
	Timeout     int
	FinishOnKey string
	NumDigits   int
}

func (g *Gather) String() string {
	attr_buffer := bytes.NewBuffer([]byte{})

	if g.attrs != nil {
		if len(g.attrs.Action) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" action=\"%s\"", g.attrs.Action))
		}

		if len(g.attrs.Method) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" method=\"%s\"", g.attrs.Method))
		}

		if g.attrs.Timeout > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" timeout=\"%d\"", g.attrs.Timeout))
		}

		if len(g.attrs.FinishOnKey) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" finishOnKey=\"%s\"", g.attrs.FinishOnKey))
		}

		if g.attrs.NumDigits > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" numDigits=\"%d\"", g.attrs.NumDigits))
		}
	}

	if g.InnerVerb == nil {
		return "<Gather" + attr_buffer.String() + "></Gather>"
	} else {
		return "<Gather" + attr_buffer.String() + ">" + g.InnerVerb.String() + "</Gather>"
	}
}

func NewGather(innerVerb Verb, attrs *GatherAttrs) *Gather {
	return &Gather{
		InnerVerb: innerVerb,
		attrs:     attrs,
	}
}

// HANGUP VERB
type Hangup struct{}

func (v *Hangup) String() string {
	return "<Hangup />"
}

func NewHangup() *Hangup {
	return &Hangup{}
}

// REJECT VERB
type Reject struct{}

func (v *Reject) String() string {
	return "<Reject />"
}

func NewReject() *Reject {
	return &Reject{}
}

// SAY VERB
type Say struct {
	noun  string
	attrs *SayAttrs
}

type SayAttrs struct {
	Loop  int
	Voice string
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

// PLAY VERB
type Play struct {
	noun  string
	attrs *PlayAttrs
}

type PlayAttrs struct {
	Loop int
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

type Pause struct {
	timeout int
}

func (v *Pause) String() string {
	attr_buffer := bytes.NewBuffer([]byte{})

	if v.timeout >= 1 {
		attr_buffer.WriteString(fmt.Sprintf(" length=\"%d\"", v.timeout))
	}

	return fmt.Sprintf("<Pause%s />", attr_buffer.String())
}
func NewPause(timeout int) *Pause {
	return &Pause{
		timeout: timeout,
	}
}

// RECORD VERB
type Record struct {
	attrs *RecordAttrs
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

// REDIRECT VERB
type Redirect struct {
	noun  string
	attrs *RedirectAttrs
}

type RedirectAttrs struct {
	Method string
}

func (r *Redirect) addresser() {}

func (r *Redirect) String() string {
	attr_buffer := bytes.NewBuffer([]byte{})

	if r.attrs != nil {
		if len(r.attrs.Method) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" method=\"%s\"", r.attrs.Method))
		}
	}

	return "<Redirect" + attr_buffer.String() + ">" + r.noun + "</Redirect>"
}

func NewRedirect(noun string, attrs *RedirectAttrs) *Redirect {
	return &Redirect{
		noun:  noun,
		attrs: attrs,
	}
}
