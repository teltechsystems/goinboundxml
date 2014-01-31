package inboundxml

import "testing"

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
