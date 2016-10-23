package more

import . "github.com/pkg4go/assert"
import "testing"

// TODO - test cache

func TestBasic(t *testing.T) {
	expect := `<!DOCTYPE html>
<html>
  <meta charset="UTF-8">
  <title>index</title>
  <header>hello</header>
  <body>more</body>
  <footer>world</footer>
</html>
`

	a := A{t}

	opts := map[string]interface{}{
		"dir":   "fixture/basic",
		"ext":   "html",
		"cache": true,
	}

	r := New(opts)
	a.NotNil(r)

	a.Equal(r.Render("index.html", ""), expect)
	a.Equal(r.Render("index", ""), expect)
}
