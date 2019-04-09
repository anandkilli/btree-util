package btree

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
)

var buf strings.Builder

func DrawBtree() {

	var hbuf strings.Builder

	//BtreeToHtml(rootNode)

	out, err := exec.Command("go", "list", "-f", "'{{.Dir}}'", "freestyle-btree/btree").Output()
	handleError(err)

	styles, err := ioutil.ReadFile(string(out)[1:len(string(out))-2] + "/styles.css")
	handleError(err)

	hbuf.WriteString("<html>\n<head>\n")
	hbuf.WriteString(strings.ReplaceAll(string(styles), "%", "%%"))
	hbuf.WriteString("</head>\n<body>\n")
	hbuf.WriteString(buf.String())
	hbuf.WriteString("</body>\n</html>")

	srv := &http.Server{Addr: ":8500"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, hbuf.String())
	})

	open("http://localhost:8500/")
	srv.ListenAndServe()
}

func BtreeToHtml(rootNode node) {

	//var buf strings.Builder

	fmt.Fprintf(&buf, "<div class=\"tree\">\n<ul>%s</ul>\n</div>\n", nodeToHtml(&rootNode))

	//return buf.String()
}

func nodeToHtml(Node *node) string {

	var str strings.Builder

	if Node.Leftnode == nil && Node.Rightnode == nil {

		// A node with no child nodes
		fmt.Fprintf(&str, "<li>\n<a href=\"#\">%s</a>\n</li>\n", Node.Value)
		return str.String()
	} else {
		fmt.Fprintf(&str, "<li>\n<a href=\"#\">%s</a>\n<ul>\n%s\n%s\n</ul>\n</li>",
			Node.Value, nodeToHtml(Node.Leftnode),
			nodeToHtml(Node.Rightnode))
		return str.String()
	}
}

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

//Appends string you want to print in Console to HTML output
func HPrintln(format string, vals ...interface{}) {

	var b strings.Builder
	fmt.Fprintf(&b, format, vals...)
	fmt.Println(b.String())
}

// open opens the specified URL in the default browser of the user.
func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
