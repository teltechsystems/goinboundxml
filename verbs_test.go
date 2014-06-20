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

		So(dial.String(), ShouldEqual, "<Dial action=\"/testing\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a method attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{Method: "POST"})

		So(dial.String(), ShouldEqual, "<Dial method=\"POST\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a timeout attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{Timeout: 30})

		So(dial.String(), ShouldEqual, "<Dial timeout=\"30\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a timeLimit attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{TimeLimit: 60})

		So(dial.String(), ShouldEqual, "<Dial timeLimit=\"60\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a callerId attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{CallerId: "+15558675309"})

		So(dial.String(), ShouldEqual, "<Dial callerId=\"+15558675309\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a dialMusic attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{DialMusic: "MUSIC.MP3"})

		So(dial.String(), ShouldEqual, "<Dial dialMusic=\"MUSIC.MP3\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a callbackUrl attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{CallbackUrl: "/callbackUrl"})

		So(dial.String(), ShouldEqual, "<Dial callbackUrl=\"/callbackUrl\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a callbackMethod attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{CallbackMethod: "POST"})

		So(dial.String(), ShouldEqual, "<Dial callbackMethod=\"POST\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a confirmSound attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{ConfirmSound: "/confirmSound"})

		So(dial.String(), ShouldEqual, "<Dial confirmSound=\"/confirmSound\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	// HERE HERE HERE
	Convey("A dial verb with a digitsMatch attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{DigitsMatch: "123"})

		So(dial.String(), ShouldEqual, "<Dial digitsMatch=\"123\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a heartbeatUrl attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{HeartbeatUrl: "/heartbeatUrl"})

		So(dial.String(), ShouldEqual, "<Dial heartbeatUrl=\"/heartbeatUrl\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a heartbeatMethod attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{HeartbeatMethod: "POST"})

		So(dial.String(), ShouldEqual, "<Dial heartbeatMethod=\"POST\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a forwardedFrom attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{ForwardedFrom: "+15038884341"})

		So(dial.String(), ShouldEqual, "<Dial forwardedFrom=\"+15038884341\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a ifMachine attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{IfMachine: "continue"})

		So(dial.String(), ShouldEqual, "<Dial ifMachine=\"continue\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a ifMachineUrl attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{IfMachineUrl: "/ifMachineUrl"})

		So(dial.String(), ShouldEqual, "<Dial ifMachineUrl=\"/ifMachineUrl\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})

	Convey("A dial verb with a ifMachineMethod attr", t, func() {
		dial := NewDial(NewNumber("+15038884341", nil), &DialAttrs{IfMachineMethod: "POST"})

		So(dial.String(), ShouldEqual, "<Dial ifMachineMethod=\"POST\" hideCallerId=\"false\" straightToVm=\"false\"><Number>+15038884341</Number></Dial>")
	})
}

func TestHangup(t *testing.T) {
	var (
		hangup        *Hangup
		hangup_string string
	)

	hangup = NewHangup()

	if hangup_string = hangup.String(); hangup_string != "<Hangup />" {
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
