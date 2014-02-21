package inboundxml

import "testing"

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

	if record_string = record.String(); record_string != "<Record action=\"http://www.example.com/recording-callback\" playBeep=\"false\" />" {
		t.Errorf("Record string returned an unexpected value : %s", record_string)
	}

	record = NewRecord(&RecordAttrs{
		Action: "http://www.example.com/recording-callback",
		Method: "POST",
	})

	if record_string = record.String(); record_string != "<Record action=\"http://www.example.com/recording-callback\" method=\"POST\" playBeep=\"false\" />" {
		t.Errorf("Record string returned an unexpected value : %s", record_string)
	}

	record = NewRecord(&RecordAttrs{
		Action:      "http://www.example.com/recording-callback",
		Method:      "POST",
		FinishOnKey: "#",
	})

	if record_string = record.String(); record_string != "<Record action=\"http://www.example.com/recording-callback\" method=\"POST\" finishOnKey=\"#\" playBeep=\"false\" />" {
		t.Errorf("Record string returned an unexpected value : %s", record_string)
	}

	record = NewRecord(&RecordAttrs{
		Action:    "http://www.example.com/recording-callback",
		Method:    "POST",
		MaxLength: 60,
	})

	if record_string = record.String(); record_string != "<Record action=\"http://www.example.com/recording-callback\" method=\"POST\" maxLength=\"60\" playBeep=\"false\" />" {
		t.Errorf("Record string returned an unexpected value : %s", record_string)
	}

	record = NewRecord(&RecordAttrs{
		Action:    "http://www.example.com/recording-callback",
		Direction: "in",
		PlayBeep:  true,
	})

	if record_string = record.String(); record_string != "<Record action=\"http://www.example.com/recording-callback\" direction=\"in\" playBeep=\"true\" />" {
		t.Errorf("Record string returned an unexpected value : %s", record_string)
	}
}
