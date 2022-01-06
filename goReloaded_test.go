package main

import "testing"

func TestProcessString(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{
			"handles kood case 1",
			"it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			"It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
		{
			"handles kood case 2",
			"Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			"Simply add 66 and 2 and you will see the result is 68.",
		},
		{
			"handles kood case 3",
			"There is no greater agony than bearing a untold story inside you.",
			"There is no greater agony than bearing an untold story inside you.",
		},
		{
			"handles kood case 4",
			"Punctuation tests are ... kinda boring ,don't you think !?",
			"Punctuation tests are... kinda boring, don't you think!?",
		},
		{
			"handles kood audit 1",
			"If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			"If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
		},
		{
			"handles kood audit 2",
			"I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			"I have to pack 5 outfits. Packed 26 just to be sure",
		},
		{
			"handles kood audit 3",
			"Don not be sad ,because sad backwards is das . And das not good",
			"Don not be sad, because sad backwards is das. And das not good",
		},
		{
			"handles kood audit 4",
			"harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			"Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
		},
		{
			"only spaces",
			"  ",
			"",
		},
		{
			"combines operations, articles, punctuation, operations and quotes",
			"flapjacks (cap) are' the best (up, 2) 'snacks (ever)    ... aren't they a oat (cap, 2) treat",
			"Flapjacks are 'THE BEST' snacks (ever)... aren't they An Oat treat",
		},
		{
			"multiple punctuation",
			".,;",
			".,;",
		},
		{
			"single words",
			"Hello",
			"Hello",
		},
		{
			"multiple words",
			"City of San Marino",
			"City of San Marino",
		},
		{
			"special characters",
			"City of San Marino",
			"City of San Marino",
		},
		{
			"Operation to no words",
			"(cap, 3)",
			"",
		},
		{
			"Overflow operation",
			"Hello World (low, 4)",
			"hello world",
		},
		{
			"Operations to specific count",
			"My favourite colour is orange (up, 3)",
			"My favourite COLOUR IS ORANGE",
		},
		{
			"single operation",
			"I am from liverpool (cap)",
			"I am from Liverpool",
		},
		{
			"duplicate operations",
			"I am delighted (up) (up) (up)",
			"I am DELIGHTED",
		},
		{
			"many specific operations",
			"writing TESTS takes a REALLY long (low, 2) (cap) time but I love it.",
			"writing TESTS takes a really Long time but I love it.",
		},
		{
			"many operations",
			"thIs (cap) should be ThIs, and thIs (low) (cap) should be This",
			"ThIs should be ThIs, and This should be This",
		},
		{
			"operation at start",
			"(hex) 101010",
			"101010",
		},
		{
			"operation before punctuation",
			"101010 (bin):",
			"42:",
		},
		{
			"operation at end",
			"101010 (bin).",
			"42.",
		},
		{
			"invalid operation",
			"yo (bin, 0)",
			"yo (bin, 0)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := processString(tt.input)
			if got != tt.want {
				t.Fatalf("\ngot\n[%s]\nwant\n[%s]", got, tt.want)
			}
		})
	}
}
