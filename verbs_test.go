package inboundxml

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDial(t *testing.T) {
	Convey("A dial verb with no attrs", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), nil)

		So(dial.String(), ShouldEqual, "<Dial><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with an action attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{Action: "/testing"})

		So(dial.String(), ShouldEqual, "<Dial action=\"/testing\" record=\"false\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a method attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{Method: "POST"})

		So(dial.String(), ShouldEqual, "<Dial method=\"POST\" record=\"false\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a timeout attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{Timeout: 30})

		So(dial.String(), ShouldEqual, "<Dial timeout=\"30\" record=\"false\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a timeLimit attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{TimeLimit: 60})

		So(dial.String(), ShouldEqual, "<Dial timeLimit=\"60\" record=\"false\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a callerId attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{CallerId: "+15558675309"})

		So(dial.String(), ShouldEqual, "<Dial callerId=\"+15558675309\" record=\"false\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a dialMusic attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{DialMusic: "MUSIC.MP3"})

		So(dial.String(), ShouldEqual, "<Dial dialMusic=\"MUSIC.MP3\" record=\"false\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a callbackUrl attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{CallbackUrl: "/callbackUrl"})

		So(dial.String(), ShouldEqual, "<Dial callbackUrl=\"/callbackUrl\" record=\"false\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a callbackMethod attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{CallbackMethod: "POST"})

		So(dial.String(), ShouldEqual, "<Dial callbackMethod=\"POST\" record=\"false\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a confirmSound attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{ConfirmSound: "/confirmSound"})

		So(dial.String(), ShouldEqual, "<Dial confirmSound=\"/confirmSound\" record=\"false\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	// HERE HERE HERE
	Convey("A dial verb with a digitsMatch attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{DigitsMatch: "123"})

		So(dial.String(), ShouldEqual, "<Dial digitsMatch=\"123\" record=\"false\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a heartbeatUrl attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{HeartbeatUrl: "/heartbeatUrl"})

		So(dial.String(), ShouldEqual, "<Dial heartbeatUrl=\"/heartbeatUrl\" record=\"false\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a heartbeatMethod attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{HeartbeatMethod: "POST"})

		So(dial.String(), ShouldEqual, "<Dial heartbeatMethod=\"POST\" record=\"false\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a forwardedFrom attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{ForwardedFrom: "+15038884341"})

		So(dial.String(), ShouldEqual, "<Dial forwardedFrom=\"+15038884341\" record=\"false\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a ifMachine attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{IfMachine: "continue"})

		So(dial.String(), ShouldEqual, "<Dial ifMachine=\"continue\" record=\"false\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a ifMachineUrl attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{IfMachineUrl: "/ifMachineUrl"})

		So(dial.String(), ShouldEqual, "<Dial ifMachineUrl=\"/ifMachineUrl\" record=\"false\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a ifMachineMethod attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{IfMachineMethod: "POST"})

		So(dial.String(), ShouldEqual, "<Dial ifMachineMethod=\"POST\" record=\"false\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a record attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{Record: true})

		So(dial.String(), ShouldEqual, "<Dial record=\"true\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a recordCallbackUrl attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{Record: true, RecordCallbackUrl: "testing"})

		So(dial.String(), ShouldEqual, "<Dial recordCallbackUrl=\"testing\" record=\"true\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a recordLifetime attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{Record: true, RecordLifetime: 60})

		So(dial.String(), ShouldEqual, "<Dial recordLifetime=\"60\" record=\"true\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})
}

func TestNumber(t *testing.T) {
	number := NewNumber("+15558884341", nil)

	if number_string := number.String(); number_string != "<Number>+15558884341</Number>" {
		t.Errorf("Number string returned an unexpected value : %s", number_string)
	}

	number = NewNumber("+15558884341", &NumberAttrs{SendDigits: "1234"})

	if number_string := number.String(); number_string != "<Number sendDigits=\"1234\" sendOnPreanswer=\"false\">+15558884341</Number>" {
		t.Errorf("Number string returned an unexpected value : %s", number_string)
	}
}

func TestRedirect(t *testing.T) {
	number := NewRedirect("http://www.google.com/", nil)

	if number_string := number.String(); number_string != "<Redirect>http://www.google.com/</Redirect>" {
		t.Errorf("Redirect string returned an unexpected value : %s", number_string)
	}

	number = NewRedirect("http://www.google.com/", &RedirectAttrs{Method: "POST"})

	if number_string := number.String(); number_string != "<Redirect method=\"POST\">http://www.google.com/</Redirect>" {
		t.Errorf("Redirect string returned an unexpected value : %s", number_string)
	}
}

func TestGather(t *testing.T) {
	Convey("A gather verb with no attrs", t, func() {
		gather := NewGather(nil, nil)

		So(gather.String(), ShouldEqual, "<Gather></Gather>")
	})

	Convey("A gather verb with an action attr", t, func() {
		gather := NewGather(NewSay("test", nil), &GatherAttrs{Action: "/testing"})

		So(gather.String(), ShouldEqual, "<Gather action=\"/testing\"><Say>test</Say></Gather>")
	})

	Convey("A gather verb with a method attr", t, func() {
		gather := NewGather(NewSay("test", nil), &GatherAttrs{Method: "POST"})

		So(gather.String(), ShouldEqual, "<Gather method=\"POST\"><Say>test</Say></Gather>")
	})

	Convey("A dial verb with a timeout attr", t, func() {
		gather := NewGather(NewSay("test", nil), &GatherAttrs{Timeout: 30})

		So(gather.String(), ShouldEqual, "<Gather timeout=\"30\"><Say>test</Say></Gather>")
	})

	Convey("A dial verb with a finishOnKey attr", t, func() {
		gather := NewGather(NewSay("test", nil), &GatherAttrs{FinishOnKey: "#"})

		So(gather.String(), ShouldEqual, "<Gather finishOnKey=\"#\"><Say>test</Say></Gather>")
	})

	Convey("A dial verb with a numDigits attr", t, func() {
		gather := NewGather(NewSay("test", nil), &GatherAttrs{NumDigits: 3})

		So(gather.String(), ShouldEqual, "<Gather numDigits=\"3\"><Say>test</Say></Gather>")
	})
}

func TestHangup(t *testing.T) {
	hangup := NewHangup()

	if hangup_string := hangup.String(); hangup_string != "<Hangup />" {
		t.Errorf("Hangup string returned an unexpected value : %s", hangup_string)
	}
}

func TestSay(t *testing.T) {
	var (
		say        *Say
		say_string string
	)

	say = NewSay("Hello World", nil)

	if say_string = say.String(); say_string != "<Say>Hello World</Say>" {
		t.Errorf("Say string returned an unexpected value : %s", say_string)
	}

	say = NewSay("Test Loop", &SayAttrs{
		Loop: 1,
	})

	if say_string = say.String(); say_string != "<Say loop=\"1\">Test Loop</Say>" {
		t.Errorf("Say string returned an unexpected value : %s", say_string)
	}

	say = NewSay("Test Loop & Man Voice", &SayAttrs{
		Loop:  1,
		Voice: "man",
	})

	if say_string = say.String(); say_string != "<Say loop=\"1\" voice=\"man\">Test Loop & Man Voice</Say>" {
		t.Errorf("Say string returned an unexpected value : %s", say_string)
	}

	say = NewSay("Test Woman Voice", &SayAttrs{
		Voice: "woman",
	})

	if say_string = say.String(); say_string != "<Say voice=\"woman\">Test Woman Voice</Say>" {
		t.Errorf("Say string returned an unexpected value : %s", say_string)
	}
}

func TestPlay(t *testing.T) {
	var (
		play        *Play
		play_string string
	)

	play = NewPlay("http://www.example.com/audio.mp3", nil)

	if play_string = play.String(); play_string != "<Play>http://www.example.com/audio.mp3</Play>" {
		t.Errorf("Play string returned an unexpected value : %s", play_string)
	}

	play = NewPlay("http://www.example.com/audio.mp3", &PlayAttrs{
		Loop: 1,
	})

	if play_string = play.String(); play_string != "<Play loop=\"1\">http://www.example.com/audio.mp3</Play>" {
		t.Errorf("Play string returned an unexpected value : %s", play_string)
	}
}

func TestPause(t *testing.T) {
	var (
		pause       *Pause
		pauseString string
	)

	pause = NewPause(3)

	if pauseString = pause.String(); pauseString != "<Pause length=\"3\" />" {
		t.Errorf("pause string returned an unexpected value : %s", pauseString)
	}

}

func TestPing(t *testing.T) {
	var (
		ping       *Ping
		pingString string
	)

	ping = NewPing("www.google.com", "POST")

	if pingString = ping.String(); pingString != "<Ping method=\"POST\">www.google.com</Ping>" {
		t.Errorf("ping string returned an unexpected value : %s", pingString)
	}

	ping = NewPing("www.google.com", "")

	if pingString = ping.String(); pingString != "<Ping>www.google.com</Ping>" {
		t.Errorf("ping string returned an unexpected value : %s", pingString)
	}

}

func TestRecord(t *testing.T) {
	var (
		record        *Record
		record_string string
	)

	record = NewRecord(nil)

	if record_string = record.String(); record_string != "<Record />" {
		t.Errorf("Record string returned an unexpected value : %s", record_string)
	}

	record = NewRecord(&RecordAttrs{
		Action: "http://www.example.com/recording-callback",
	})

	if record_string = record.String(); record_string != "<Record action=\"http://www.example.com/recording-callback\" playBeep=\"false\" background=\"false\" />" {
		t.Errorf("Record string returned an unexpected value : %s", record_string)
	}

	record = NewRecord(&RecordAttrs{
		Action: "http://www.example.com/recording-callback",
		Method: "POST",
	})

	if record_string = record.String(); record_string != "<Record action=\"http://www.example.com/recording-callback\" method=\"POST\" playBeep=\"false\" background=\"false\" />" {
		t.Errorf("Record string returned an unexpected value : %s", record_string)
	}

	record = NewRecord(&RecordAttrs{
		Action:      "http://www.example.com/recording-callback",
		Method:      "POST",
		FinishOnKey: "#",
	})

	if record_string = record.String(); record_string != "<Record action=\"http://www.example.com/recording-callback\" method=\"POST\" finishOnKey=\"#\" playBeep=\"false\" background=\"false\" />" {
		t.Errorf("Record string returned an unexpected value : %s", record_string)
	}

	record = NewRecord(&RecordAttrs{
		Action:  "http://www.example.com/recording-callback",
		Method:  "POST",
		Timeout: 60,
	})

	if record_string = record.String(); record_string != "<Record action=\"http://www.example.com/recording-callback\" method=\"POST\" timeout=\"60\" playBeep=\"false\" background=\"false\" />" {
		t.Errorf("Record string returned an unexpected value : %s", record_string)
	}

	record = NewRecord(&RecordAttrs{
		Action:    "http://www.example.com/recording-callback",
		Method:    "POST",
		MaxLength: 60,
	})

	if record_string = record.String(); record_string != "<Record action=\"http://www.example.com/recording-callback\" method=\"POST\" maxLength=\"60\" playBeep=\"false\" background=\"false\" />" {
		t.Errorf("Record string returned an unexpected value : %s", record_string)
	}

	record = NewRecord(&RecordAttrs{
		Action:    "http://www.example.com/recording-callback",
		Direction: "in",
		PlayBeep:  true,
	})

	if record_string = record.String(); record_string != "<Record action=\"http://www.example.com/recording-callback\" direction=\"in\" playBeep=\"true\" background=\"false\" />" {
		t.Errorf("Record string returned an unexpected value : %s", record_string)
	}

	record = NewRecord(&RecordAttrs{
		Action:     "http://www.example.com/recording-callback",
		Background: true,
	})

	if record_string = record.String(); record_string != "<Record action=\"http://www.example.com/recording-callback\" playBeep=\"false\" background=\"true\" />" {
		t.Errorf("Record string returned an unexpected value : %s", record_string)
	}
}
