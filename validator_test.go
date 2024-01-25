package validator

import "testing"

func TestIsAlpha(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"Ⅸ", false},
		{"0", false},
		{" ", false},
		{".", false},
		{"-", false},
		{"+", false},
		{"\n", false},
		{"\r", false},
		{"소", false},
		{"-0", false},
		{"++", false},
		{"+1", false},
		{"--", false},
		{"1¾", false},
		{"123", false},
		{"-1¾", false},
		{"1++", false},
		{"1+1", false},
		{"1--", false},
		{"1-1", false},
		{"〥〩", false},
		{"모자", false},
		{"0123", false},
		{"abc1", false},
		{"소주", false},
		{"۳۵۶۰", false},
		{"abc〩", false},
		{"소aBC", false},
		{"\u0026", false}, // UTF-8(ASCII): &
		{"\u0030", false}, // UTF-8(ASCII): 0
		{"-00123", false},
		{"\ufff0", false},
		{"abc!!!", false},
		{"〩Hours", false},
		{"123.123", false},
		{"달기&Co.", false},
		{"   fooo   ", false},

		{"", true},
		{"ix", true},
		{"abc", true},
		{"ABC", true},
		{"FoObAr", true},
		{"\u0070", true}, // UTF-8(ASCII): p
	}
	for _, test := range tests {
		actual := IsAlpha(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsAlpha(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsAlphanumeric(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"Ⅸ", false},
		{" ", false},
		{".", false},
		{"-", false},
		{"+", false},
		{"\n", false},
		{"\r", false},
		{"--", false},
		{"++", false},
		{"+1", false},
		{"1¾", false},
		{"소", false},
		{"-0", false},
		{"1++", false},
		{"1+1", false},
		{"1--", false},
		{"1-1", false},
		{"-1¾", false},
		{"〥〩", false},
		{"모자", false},
		{"소aBC", false},
		{"۳۵۶۰", false},
		{"abc〩", false},
		{"abc!!!", false},
		{"-00123", false},
		{"\ufff0", false},
		{"\u0026", false}, // UTF-8(ASCII): &
		{"123.123", false},
		{"〩Hours", false},
		{"달기&Co.", false},
		{"   fooo   ", false},

		{"", true},
		{"0", true},
		{"ix", true},
		{"abc", true},
		{"123", true},
		{"ABC", true},
		{"0123", true},
		{"abc1", true},
		{"소주", false},
		{"FoObAr", true},
		{"abc123", true},
		{"ABC111", true},
		{"\u0070", true}, // UTF-8(ASCII): p
		{"\u0030", true}, // UTF-8(ASCII): 0
	}
	for _, test := range tests {
		actual := IsAlphanumeric(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsAlphanumeric(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsEmail(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"invalid.com", false},
		{"@invalid.com", false},
		{"invalidemail@", false},
		{"foo@bar.coffee..coffee", false},

		{"x@x.x", true},
		{"foo@bar.com", true},
		{"foo@bar.中文网", true},
		{"foo@bar.com.au", true},
		{"foo@bar.coffee", true},
		{"foo+bar@bar.com", true},
		{"hans@m端ller.com", true},
		{"foo@bar.bar.coffee", true},
		{"test|123@m端ller.com", true},
		{"hans.m端ller@test.com", true},
		{"NathAn.daVIeS@DomaIn.cOM", true},
		{"NATHAN.DAVIES@DOMAIN.CO.UK", true},
	}
	for _, test := range tests {
		actual := IsEmail(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsEmail(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsExistingEmail(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"invalid.com", false},
		{"@invalid.com", false},
		{"invalidemail@", false},
		{"foo@bar.coffee..coffee", false},
		{"[prasun.joshi]@DomaIn.cOM", false},
		{"nosuchdomain@bar.nosuchdomainsuffix", false},
		{"sizeofuserismorethansixtyfour0123sizeofuserismorethansixtyfour0123@DOMAIN.CO.UK", false},

		{"foo@bar.com", true},
		{"foo@bar.com.au", true},
		{"foo+bar@bar.com", true},
		{"prasun.joshi@localhost", true},
		{"NathAn.daVIeS@DomaIn.cOM", true},
		{"NATHAN.DAVIES@DOMAIN.CO.UK", true},
	}
	for _, test := range tests {
		actual := IsExistingEmail(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsExistingEmail(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsNull(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"abacaba", false},
		{"", true},
	}
	for _, test := range tests {
		actual := IsNull(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsNull(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsNotNull(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"abacaba", true},
	}
	for _, test := range tests {
		actual := IsNotNull(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsNull(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsNumeric(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{

		{" ", false},
		{"Ⅸ", false},
		{".", false},
		{"-", false},
		{"+", false},
		{"\n", false},
		{"\r", false},
		{"소", false},
		{"1¾", false},
		{"--", false},
		{"++", false},
		{"+1", false},
		{"ix", false},
		{"-0", false},
		{"ABC", false},
		{"1--", false},
		{"1-1", false},
		{"1++", false},
		{"-1¾", false},
		{"1+1", false},
		{"abc", false},
		{"abc1", false},
		{"소주", false},
		{"〥〩", false},
		{"모자", false},
		{"12𐅪3", false},
		{"۳۵۶۰", false},
		{"소aBC", false},
		{"abc〩", false},
		{"abc!!!", false},
		{"FoObAr", false},
		{"\ufff0", false},
		{"-00123", false},
		{"+00123", false},
		{"\u0070", false}, // UTF-8(ASCII): p
		{"\u0026", false}, // UTF-8(ASCII): &
		{"123.123", false},
		{"〩Hours", false},
		{"달기&Co.", false},
		{"   fooo   ", false},

		{"", true},
		{"0", true},
		{"123", true},
		{"0123", true},
		{"\u0030", true}, // UTF-8(ASCII): 0
	}
	for _, test := range tests {
		actual := IsNumeric(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsNumeric(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsURL(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		// {"invalid.", false}, is it false like "localhost."?

		{"", false},
		{"foo", false},
		{".com", false},
		{"google", false},
		{",foo.com", false},
		{"/abs/test/dir", false},
		{"./rel/test/dir", false},
		{"http://.foo.com", false},
		{"http://,foo.com", false},
		{"xyz://foobar.com", false},
		{"rtmp://foobar.com", false},
		{"http://foo^bar.org", false},
		{"http://foo&bar.org", false},
		{"http://foo bar.org", false},
		{"http://_foobar.com", false},
		{"http://foo&*bar.org", false},
		{"http://[::1]:909388", false},
		{"http://foobar.c_o_m", false},
		{"foo_bar-fizz-buzz:13:13", false},
		{"http://www.-foobar.com/", false},
		{"foo_bar-fizz-buzz://1313", false},
		{"http://www.foo---bar.com/", false},
		{"irc://irc.server.org/channel", false},
		{"1200::AB00:1234::2552:7777:1313", false},
		{"http://cant_end_with_underescore_", false},
		{"http://_cant_start_with_underescore", false},
		{"http://cant-end-with-hyphen-.example.com", false},
		{"http://-cant-start-with-hyphen.example.com", false},

		{"foobar.com", true},
		{"ftp.foo.bar", true},
		{"ftp://foobar.ru/", true},
		{"http://127.0.0.1/", true},
		{"http://[::1]:9093", true},
		{"http://foobar.com", true},
		{"http://foobar.ORG", true},
		{"http://foo.bar.org", true},
		{"http://foobar.org/", true},
		{"http://foo.bar#com", true},
		{"http://foo_bar.com", true},
		{"https://foobar.com", true},
		{"foo_bar.example.com", true},
		{"http://foobar.com/a-", true},
		{"http://www.foo.co.uk", true},
		{"http://me.example.com", true},
		{"http://foobar.coffee/", true},
		{"http://foobar.中文网/", true},
		{"foo_bar-fizz-buzz:1313", true},
		{"http://www.foo.bar.org", true},
		{"http://localhost:3000/", true},
		{"irc://#channel@network", true},
		{"http://foobar.org:8080/", true},
		{"http://myservice.:9093/", true},
		{"http://foobar.پاکستان/", true},
		{"http://www.me.example.com", true},
		{"http://foobar.com?foo=bar", true},
		{"http://foobar.com#baz=qux", true},
		{"mailto:someone@example.com", true},
		{"http://foo_bar.example.com", true},
		{"https://foo_bar.example.com", true},
		{"http://m.abcd.com/test.html", true},
		{"http://duckduckgo.com/?q=%2F", true},
		{"http://www.foobar.com/~foobar", true},
		{"foo_bar_fizz_buzz.example.com", true},
		{"http://hello_world.example.com", true},
		{"https://farm6.static.flickr.com", true},
		{"http://user:pass@www.foobar.com/", true},
		{"http://foobar.com/t$-_.+!*\\'(),", true},
		{"http://www.xn--froschgrn-x9a.net/", true},
		{"https://127.0.0.1/a/b/c?a=v&c=11d", true},
		{"http://foobar.com/?foo=bar#baz=qux", true},
		{"http://foo_bar_fizz_buzz.example.com", true},
		{"http://www.domain-can-have-dashes.com", true},
		{"http://r6---snnvoxuioq6.googlevideo.com", true},
		{"http://user:pass@foo_bar_bar.bar_foo.com", true},
		{"http://[2001:db8:a0b:12f0::1]/index.html", true},
		{"http://hyphenated-host-name.example.co.in", true},
		{"http://user:pass@www.foobar.com/path/file", true},
		{"aio1_alertmanager_container-63376c45:9093", true},
		{"http://user:pass@[::1]:9093/a/b/c/?a=v#abc", true},
		{"http://prometheus-alertmanager.service.q:9093", true},
		{"http://m.abcd.com/a/b/c/d/test.html?args=a&b=c", true},
		{"http://[1200:0000:AB00:1234:0000:2552:7777:1313]", true},
		{"https://zh.wikipedia.org/wiki/Wikipedia:%E9%A6%96%E9%A1%B5", true},
		{"https://pbs.twimg.com/profile_images/560826135676588032/j8fWrmYY_normal.jpeg", true},
		{"https://www.logn-123-123.url.with.sigle.letter.d:12345/url/path/foo?bar=zzz#user", true},
	}
	for _, test := range tests {
		actual := IsURL(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsURL(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsRequestURL(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"", false},
		{".com", false},
		{"invalid.", false},
		{"foobar.com", false},
		{"/abs/test/dir", false},
		{"./rel/test/dir", false},

		{"ftp://foobar.ru/", true},
		{"xyz://foobar.com", true},
		{"rtmp://foobar.com", true},
		{"http://127.0.0.1/", true},
		{"http://foobar.com", true},
		{"http://foobar.org/", true},
		{"https://foobar.com", true},
		{"http://foo.bar/#com", true},
		{"http://foobar.coffee/", true},
		{"http://foobar.中文网/", true},
		{"http://localhost:3000/", true},
		{"http://localhost:3000/", true},
		{"http://www.foo_bar.com/", true},
		{"http://foobar.org:8080/", true},
		{"http://www.-foobar.com/", true},
		{"http://foobar.com?foo=bar", true},
		{"http://www.foo---bar.com/", true},
		{"mailto:someone@example.com", true},
		{"http://foobar.com/#baz=qux", true},
		{"irc://irc.server.org/channel", true},
		{"http://duckduckgo.com/?q=%2F", true},
		{"http://www.foobar.com/~foobar", true},
		{"http://foobar.com/t$-_.+!*\\'(),", true},
		{"http://user:pass@www.foobar.com/", true},
		{"http://www.xn--froschgrn-x9a.net/", true},
		{"http://foobar.com/?foo=bar#baz=qux", true},
	}
	for _, test := range tests {
		actual := IsRequestURL(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsRequestURL(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsRequestURI(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{".com", false},
		{"invalid.", false},
		{"foobar.com", false},
		{"./rel/test/dir", false},

		{"/abs/test/dir", true},
		{"xyz://foobar.com", true},
		{"ftp://foobar.ru/", true},
		{"rtmp://foobar.com", true},
		{"http://127.0.0.1/", true},
		{"http://foobar.com", true},
		{"http://foobar.org/", true},
		{"https://foobar.com", true},
		{"http://foo.bar/#com", true},
		{"http://foobar.coffee/", true},
		{"http://foobar.中文网/", true},
		{"http://localhost:3000/", true},
		{"http://localhost:3000/", true},
		{"http://foobar.org:8080/", true},
		{"http://www.-foobar.com/", true},
		{"http://www.foo_bar.com/", true},
		{"http://www.foo---bar.com/", true},
		{"http://foobar.com?foo=bar", true},
		{"mailto:someone@example.com", true},
		{"http://foobar.com/#baz=qux", true},
		{"http://duckduckgo.com/?q=%2F", true},
		{"irc://irc.server.org/channel", true},
		{"http://www.foobar.com/~foobar", true},
		{"http://foobar.com/t$-_.+!*\\'(),", true},
		{"http://user:pass@www.foobar.com/", true},
		{"http://www.xn--froschgrn-x9a.net/", true},
		{"http://foobar.com/?foo=bar#baz=qux", true},
	}
	for _, test := range tests {
		actual := IsRequestURI(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsRequestURI(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsUTFDigit(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"+", false},
		{" ", false},
		{".", false},
		{"Ⅸ", false},
		{"\n", false},
		{"\r", false},
		{"++", false},
		{"1¾", false},
		{"ix", false},
		{"소", false},
		{"--0", false},
		{"-0-", false},
		{"1++", false},
		{"1+1", false},
		{"-1¾", false},
		{"ABC", false},
		{"abc", false},
		{"abc1", false},
		{"소주", false},
		{"〥〩", false},
		{"모자", false},
		{"12𐅪3", false},
		{"abc〩", false},
		{"소aBC", false},
		{"abc!!!", false},
		{"FoObAr", false},
		{"\ufff0", false},
		{"\u0070", false}, // UTF-8(ASCII): p
		{"\u0026", false}, // UTF-8(ASCII): &
		{"123.123", false},
		{"〩Hours", false},
		{"달기&Co.", false},
		{"   fooo   ", false},

		{"", true},
		{"0", true},
		{"-0", true},
		{"+1", true},
		{"-29", true},
		{"123", true},
		{"0123", true},
		{"۳۵۶۰", true},
		{"۳۵۶۰", true},
		{"-00123", true},
		{"\u0030", true}, // UTF-8(ASCII): 0
		{"1483920", true},
	}
	for _, test := range tests {
		actual := IsUTFDigit(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsUTFDigit(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsUTFLetter(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"Ⅸ", false},
		{"0", false},
		{" ", false},
		{".", false},
		{"-", false},
		{"+", false},
		{"\n", false},
		{"\r", false},
		{"-0", false},
		{"1¾", false},
		{"--", false},
		{"++", false},
		{"+1", false},
		{"-1¾", false},
		{"1--", false},
		{"1-1", false},
		{"1++", false},
		{"1+1", false},
		{"123", false},
		{"〥〩", false},
		{"0123", false},
		{"abc1", false},
		{"abc〩", false},
		{"۳۵۶۰", false},
		{"abc!!!", false},
		{"\ufff0", false},
		{"-00123", false},
		{"\u0026", false}, // UTF-8(ASCII): &
		{"\u0030", false}, // UTF-8(ASCII): 0
		{"〩Hours", false},
		{"123.123", false},
		{"달기&Co.", false},
		{"   fooo   ", false},

		{"", true},
		{"소", true},
		{"ix", true},
		{"abc", true},
		{"ABC", true},
		{"모자", true},
		{"소주", true},
		{"소aBC", true},
		{"FoObAr", true},
		{"\u0070", true}, // UTF-8(ASCII): p
	}
	for _, test := range tests {
		actual := IsUTFLetter(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsUTFLetter(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsUTFLetterNumeric(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{" ", false},
		{".", false},
		{"-", false},
		{"+", false},
		{"\n", false},
		{"\r", false},
		{"-0", false},
		{"--", false},
		{"++", false},
		{"+1", false},
		{"-1¾", false},
		{"1--", false},
		{"1-1", false},
		{"1++", false},
		{"1+1", false},
		{"abc!!!", false},
		{"\ufff0", false},
		{"-00123", false},
		{"\u0026", false}, // UTF-8(ASCII): &
		{"123.123", false},
		{"달기&Co.", false},
		{"   fooo   ", false},

		{"", true},
		{"Ⅸ", true},
		{"0", true},
		{"1¾", true},
		{"소", true},
		{"ix", true},
		{"123", true},
		{"abc", true},
		{"ABC", true},
		{"〥〩", true},
		{"모자", true},
		{"소주", true},
		{"0123", true},
		{"abc1", true},
		{"۳۵۶۰", true},
		{"abc〩", true},
		{"소aBC", true},
		{"FoObAr", true},
		{"\u0070", true}, // UTF-8(ASCII): p
		{"\u0030", true}, // UTF-8(ASCII): 0
		{"〩Hours", true},
	}
	for _, test := range tests {
		actual := IsUTFLetterNumeric(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsUTFLetterNumeric(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsUTFNumeric(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{

		{"+", false},
		{" ", false},
		{".", false},
		{"소", false},
		{"ix", false},
		{"++", false},
		{"\n", false},
		{"\r", false},
		{"1++", false},
		{"1+1", false},
		{"--0", false},
		{"-0-", false},
		{"abc", false},
		{"ABC", false},
		{"모자", false},
		{"abc1", false},
		{"소주", false},
		{"abc〩", false},
		{"소aBC", false},
		{"abc!!!", false},
		{"FoObAr", false},
		{"\ufff0", false},
		{"\u0070", false}, // UTF-8(ASCII): p
		{"\u0026", false}, // UTF-8(ASCII): &
		{"123.123", false},
		{"〩Hours", false},
		{"달기&Co.", false},
		{"   fooo   ", false},

		{"", true},
		{"Ⅸ", true},
		{"0", true},
		{"-0", true},
		{"+1", true},
		{"1¾", true},
		{"-1¾", true},
		{"123", true},
		{"0123", true},
		{"〥〩", true},
		{"12𐅪3", true},
		{"۳۵۶۰", true},
		{"-00123", true},
		{"\u0030", true}, // UTF-8(ASCII): 0
	}
	for _, test := range tests {
		actual := IsUTFNumeric(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsUTFNumeric(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}
