package main
import (
	"testing"
)

func TestParse(t *testing.T) {
	var out, want string
	out = parse("./test/body.html")
	want = "<html>hello</html>\n"
	if out != want {
		t.Errorf("failed parse(). out = %s", out)
	}

	want = "goodbye"
	out = parse(want)
	if out != want {
		t.Errorf("failed parse(). out = %s", out)
	}

	setDefaultConfig("config/default.toml")
	setFlag()
	showConfig()
	usage()

	_ = getHeader()
	_ = getBody()
	_ = getAttach()
	// doSendMail(h, b, a)
}
