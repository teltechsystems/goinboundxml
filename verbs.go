package inboundxml

import (
	"bytes"
	"fmt"
	"strings"
)

type Verb interface {
	fmt.Stringer
}

type DialAddresser interface {
	Verb
	Addresser()
}

// DIAL VERB
type Dial struct {
	Addresser DialAddresser
	attrs     *DialAttrs
}

type DialAttrs struct {
	Action            string
	Method            string
	Record            bool
	RecordTrimSilence bool
	RecordCallbackUrl string
	RecordLifetime    int
	RecordDirection   string
	RecordFileFormat  string
	Timeout           int
	TimeLimit         int
	CallerId          string
	CallerName        string
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

func (d *DialAttrs) getCallbackUrl() string {
	return d.CallbackUrl
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

		if len(d.attrs.CallerName) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" callerName=\"%s\"", d.attrs.CallerName))
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

		if len(d.attrs.RecordFileFormat) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" recordFormat=\"%s\"", d.attrs.RecordFileFormat))
		}

		attr_buffer.WriteString(fmt.Sprintf(" record=\"%t\"", d.attrs.Record))
		attr_buffer.WriteString(fmt.Sprintf(" recordTrimSilence=\"%t\"", d.attrs.RecordTrimSilence))
		attr_buffer.WriteString(fmt.Sprintf(" hideCallerId=\"%t\"", d.attrs.HideCallerId))
		attr_buffer.WriteString(fmt.Sprintf(" straightToVm=\"%t\"", d.attrs.StraightToVm))

	}

	return "<Dial" + attr_buffer.String() + ">" + d.Addresser.String() + "</Dial>"
}

func NewDial(addresser DialAddresser, attrs *DialAttrs) *Dial {
	return &Dial{
		Addresser: addresser,
		attrs:     attrs,
	}
}

// MULTI VERB
type MultiVerb struct {
	verbs []Verb
}

func (v *MultiVerb) String() string {
	b := bytes.NewBuffer([]byte{})
	for i := range v.verbs {
		b.WriteString(v.verbs[i].String())
	}

	return b.String()
}

func NewMultiVerb(verbs ...Verb) *MultiVerb {
	return &MultiVerb{
		verbs: verbs,
	}
}

// NUMBER VERB
type Number struct {
	noun  string
	attrs *NumberAttrs
}

type NumberAttrs struct {
	SendDigits      string
	SendOnPreanswer bool
}

func (n *Number) Addresser() {}

func (n *Number) String() string {
	attr_buffer := bytes.NewBuffer([]byte{})

	if n.attrs != nil {
		if len(n.attrs.SendDigits) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" sendDigits=\"%s\"", n.attrs.SendDigits))
		}

		attr_buffer.WriteString(fmt.Sprintf(" sendOnPreanswer=\"%t\"", n.attrs.SendOnPreanswer))
	}

	return "<Number" + attr_buffer.String() + ">" + n.noun + "</Number>"
}

func NewNumber(noun string, attrs *NumberAttrs) *Number {
	return &Number{
		noun:  noun,
		attrs: attrs,
	}
}

// SIP VERB
type Sip struct {
	noun  string
	attrs *SipAttrs
}

type SipAttrs struct {
	Username string
	Password string
}

func (s *Sip) Addresser() {}

func (s *Sip) String() string {
	attr_buffer := bytes.NewBuffer([]byte{})

	if s.attrs != nil {
		if len(s.attrs.Username) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" username=\"%s\"", s.attrs.Username))
		}

		if len(s.attrs.Password) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" password=\"%s\"", s.attrs.Password))
		}
	}

	return "<Sip" + attr_buffer.String() + ">" + s.noun + "</Sip>"
}

func NewSip(noun string, attrs *SipAttrs) *Sip {
	return &Sip{
		noun:  noun,
		attrs: attrs,
	}
}

// CONFERENCE VERB
type Conference struct {
	noun  string
	attrs *ConferenceAttrs
}

type ConferenceAttrs struct {
	Muted                  bool
	Beep                   bool
	StartConferenceOnEnter bool
	EndConferenceOnExit    bool
	MaxParticipants        int //1to 40
	WaitUrl                string
	WaitMethod             string //POST OR GET
	HangupOnStar           bool
	CallbackUrl            string
	CallbackMethod         string
	WaitSound              string
	WaitSoundMethod        string //POST OR GET\
	DigitsMatch            string //if the user inputs didgits and then match then send a callback
	StayAlone              bool   //if they can stay alone in the conference
	Record                 bool
	RecordCallbackUrl      string
	RecordFileFormat       string //mp3 or wav
}

func (c *Conference) Addresser() {}

func (c *Conference) String() string {
	attr_buffer := bytes.NewBuffer([]byte{})

	if c.attrs != nil {
		if len(c.attrs.WaitUrl) > 0 {
			// The param is inconsistent with attribute due to an update on Zang
			attr_buffer.WriteString(fmt.Sprintf(" waitSound=\"%s\"", c.attrs.WaitUrl))
		}
		if len(c.attrs.WaitMethod) > 0 {
			// The param is inconsistent with attribute due to an update on Zang
			attr_buffer.WriteString(fmt.Sprintf(" waitSoundMethod=\"%s\"", c.attrs.WaitMethod))
		}
		if len(c.attrs.CallbackUrl) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" callbackUrl=\"%s\"", c.attrs.CallbackUrl))
		}
		if len(c.attrs.CallbackMethod) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" callbackMethod=\"%s\"", c.attrs.CallbackMethod))
		}
		if len(c.attrs.WaitSound) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" waitSound=\"%s\"", c.attrs.WaitSound))
		}
		if len(c.attrs.WaitSoundMethod) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" waitSoundMethod=\"%s\"", c.attrs.WaitSoundMethod))
		}
		if len(c.attrs.DigitsMatch) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" digitsMatch=\"%s\"", c.attrs.DigitsMatch))
		}
		if len(c.attrs.RecordCallbackUrl) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" recordCallbackUrl=\"%s\"", c.attrs.RecordCallbackUrl))
		}
		if len(c.attrs.RecordFileFormat) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" recordFileFormat=\"%s\"", c.attrs.RecordFileFormat))
		}
		if c.attrs.MaxParticipants > 0 && c.attrs.MaxParticipants < 41 {
			attr_buffer.WriteString(fmt.Sprintf(" maxParticipants=\"%d\"", c.attrs.MaxParticipants))
		}

		attr_buffer.WriteString(fmt.Sprintf(" muted=\"%t\"", c.attrs.Muted))
		attr_buffer.WriteString(fmt.Sprintf(" beep=\"%t\"", c.attrs.Beep))
		attr_buffer.WriteString(fmt.Sprintf(" startConferenceOnEnter=\"%t\"", c.attrs.StartConferenceOnEnter))
		attr_buffer.WriteString(fmt.Sprintf(" endConferenceOnExit=\"%t\"", c.attrs.EndConferenceOnExit))
		attr_buffer.WriteString(fmt.Sprintf(" hangupOnStar=\"%t\"", c.attrs.HangupOnStar))
		attr_buffer.WriteString(fmt.Sprintf(" stayAlone=\"%t\"", c.attrs.StayAlone))
		attr_buffer.WriteString(fmt.Sprintf(" record=\"%t\"", c.attrs.Record))

	}

	return "<Conference" + attr_buffer.String() + ">" + c.noun + "</Conference>"
}

func NewConference(noun string, attrs *ConferenceAttrs) *Conference {
	return &Conference{
		noun:  noun,
		attrs: attrs,
	}
}

// GATHER VERB
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
type Hangup struct {
	attrs *HangupAttrs
}

type HangupAttrs struct {
	Schedule int
	Reason   string
}

func (v *Hangup) String() string {
	attr_buffer := bytes.NewBuffer([]byte{})

	if v.attrs != nil {
		if v.attrs.Schedule > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" schedule=\"%d\"", v.attrs.Schedule))
		}

		if v.attrs.Reason != "" {
			attr_buffer.WriteString(fmt.Sprintf(" reason=\"%s\"", v.attrs.Reason))
		}
	}

	return "<Hangup" + attr_buffer.String() + " />"
}

func NewHangup(attrs *HangupAttrs) *Hangup {
	return &Hangup{
		attrs: attrs,
	}
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
	Loop   int
	Digits string
}

func (v *Play) String() string {
	attr_buffer := bytes.NewBuffer([]byte{})

	if v.attrs != nil {
		if v.attrs.Loop >= 1 {
			attr_buffer.WriteString(fmt.Sprintf(" loop=\"%d\"", v.attrs.Loop))
		}

		if v.attrs.Digits != "" {
			attr_buffer.WriteString(fmt.Sprintf(" digits=\"%s\"", v.attrs.Digits))
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

// DTMF VERB
type DTMF struct {
	tones string
}

func (v *DTMF) String() string {
	attrBuffer := bytes.NewBuffer([]byte{})

	return fmt.Sprintf("<DTMF%s>%s</DTMF>", attrBuffer.String(), v.tones)
}

func NewDTMF(tones string) *DTMF {
	return &DTMF{
		tones:  tones,
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

// GATHER VERB
type PreAnswer struct {
	InnerVerbs []Verb
}

func (pa *PreAnswer) String() string {
	preAnswerBuffer := bytes.NewBuffer([]byte{})

	preAnswerBuffer.WriteString("<PreAnswer>")
	for _, verb := range pa.InnerVerbs {
		preAnswerBuffer.WriteString(verb.String())
	}
	preAnswerBuffer.WriteString("</PreAnswer>")

	return preAnswerBuffer.String()
}

func NewPreAnswer(innerVerbs ...Verb) *PreAnswer {
	return &PreAnswer{
		InnerVerbs: innerVerbs,
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
	Timeout     int
	MaxLength   int
	Method      string
	PlayBeep    bool
	FileFormat  string
	TrimSilence bool
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

		if v.attrs.Timeout > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" timeout=\"%d\"", v.attrs.Timeout))
		}

		if v.attrs.MaxLength > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" maxLength=\"%d\"", v.attrs.MaxLength))
		}

		if len(v.attrs.Direction) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" direction=\"%s\"", v.attrs.Direction))
		}

		if len(v.attrs.FileFormat) > 0 {
			attr_buffer.WriteString(fmt.Sprintf(" fileFormat=\"%s\"", v.attrs.FileFormat))
		}

		attr_buffer.WriteString(fmt.Sprintf(" playBeep=\"%t\"", v.attrs.PlayBeep))
		attr_buffer.WriteString(fmt.Sprintf(" background=\"%t\"", v.attrs.Background))
		attr_buffer.WriteString(fmt.Sprintf(" trimSilence=\"%t\"", v.attrs.TrimSilence))
	}

	return fmt.Sprintf("<Record%s />", attr_buffer.String())
}

func NewRecord(attrs *RecordAttrs) *Record {
	return &Record{
		attrs: attrs,
	}
}

type Ping struct {
	noun   string
	method string
}

func (p *Ping) String() string {
	attr_buffer := bytes.NewBuffer([]byte{})

	if strings.ToUpper(p.method) == "POST" || strings.ToUpper(p.method) == "GET" {
		attr_buffer.WriteString(fmt.Sprintf(" method=\"%s\"", p.method))
	}

	return "<Ping" + attr_buffer.String() + ">" + p.noun + "</Ping>"
}

func NewPing(noun string, method string) *Ping {
	return &Ping{
		noun:   noun,
		method: method,
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
