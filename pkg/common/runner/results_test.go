package runner

import (
	"bytes"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"text/template"

	kubev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
	kubetest "k8s.io/client-go/testing"
)

func resultString(files ...string) string {
	var b bytes.Buffer
	resultTemplate.Execute(&b, struct{ Files []string }{files})
	return b.String()
}

func keys(in map[string][]byte) []string {
	out := make([]string, 0, len(in))
	for key := range in {
		out = append(out, key)
	}
	return out
}

var (
	goodResults = map[string][]byte{
		"a":                 []byte("testdata"),
		"b":                 []byte("moretestdata"),
		"c":                 []byte("evenmoretestdata"),
		"junit-results.xml": []byte(goodXML),
	}
	badResults = map[string][]byte{
		"a":                     []byte("testdata"),
		"b":                     []byte("moretestdata"),
		"c":                     []byte("evenmoretestdata"),
		"junit-bad-results.xml": []byte(badXML),
	}
	noXMLResults = map[string][]byte{
		"a": []byte("testdata"),
		"b": []byte("moretestdata"),
		"c": []byte("evenmoretestdata"),
	}
	goodXML = `<testsuite name="Suite" tests="1" failures="0" errors="0" time="0">
    <testcase name="[Suite] testname" classname="classname" time="0">
        <passed>Passed with 0 matches
        </passed>
    </testcase>
</testsuite>`

	badXML = `<testsuite name="Suite" tests="1" failures="0" errors="0" time="0">
    <test`

	resultTemplate = template.Must(template.New("results").Parse(`
<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 3.2 Final//EN"><html>
<title>Directory listing for /</title>
<body>
<h2>Directory listing for /</h2>
<hr>
<ul>
{{ range .Files }}
<li><a href="{{.}}">{{.}}</a></li>
{{end}}
</ul>
<hr>
</body>
</html>
`))
	resultPage       = response(resultString(keys(goodResults)...))
	noXMLResultPage  = response(resultString(keys(noXMLResults)...))
	badXMLResultPage = response(resultString(keys(badResults)...))
)

func TestRetrieveResults(t *testing.T) {
	type testcase struct {
		Name string
		ResultsServerReactor
		Expected    map[string][]byte
		ShouldError bool
	}
	for _, testcase := range []testcase{
		{
			Name:                 "validXMLPresent",
			ResultsServerReactor: ResultsServerReactor{resultPage},
			Expected:             goodResults,
			ShouldError:          false,
		},
		{
			Name:                 "XMLMissing",
			ResultsServerReactor: ResultsServerReactor{noXMLResultPage},
			Expected:             noXMLResults,
			ShouldError:          true,
		},
		{
			Name:                 "invalidXML",
			ResultsServerReactor: ResultsServerReactor{badXMLResultPage},
			Expected:             badResults,
			ShouldError:          true,
		},
	} {
		t.Run(testcase.Name, func(t *testing.T) {
			// setup mock client
			client := fake.NewSimpleClientset()
			client.AddProxyReactor("services", testcase.React)

			// setup runner
			def := *DefaultRunner
			r := &def
			r.Kube = client

			// create results service
			svc, err := r.createService(new(kubev1.Pod))
			if err != nil {
				t.Fatalf("Failed to create example service: %v", err)
			}
			r.svc = svc

			// get results
			results, err := r.RetrieveResults()
			if err != nil && !testcase.ShouldError {
				t.Fatalf("Failed to get results: %v", err)
			} else if err == nil && testcase.ShouldError {
				t.Fatalf("RetrieveResults should have failed")
			}
			if !testcase.ShouldError {
				// compare to expected
				for k, v := range testcase.Expected {
					if actualV, ok := results[k]; !ok {
						t.Fatalf("missing file '%s' in results", k)
					} else if !bytes.Equal(actualV, v) {
						t.Fatalf("file '%s' has been corrupted: want '%s', got '%s'", k, v, actualV)
					}
				}
			}
		})
	}
}

type ResultsServerReactor struct {
	IndexPage rest.ResponseWrapper
}

func (r ResultsServerReactor) React(action kubetest.Action) (handled bool, ret rest.ResponseWrapper, err error) {
	proxyAction := action.(kubetest.ProxyGetActionImpl)

	// only respond to service proxy requests
	if handled = proxyAction.Matches(http.MethodGet, "services"); !handled {
		return
	}

	path := strings.TrimPrefix(proxyAction.Path, "/")
	if path == "" {
		ret = r.IndexPage
	} else if data, ok := goodResults[path]; ok {
		ret = response(data)
	} else {
		ret = response{}
	}

	return
}

func resultsServerReactor(action kubetest.Action) (handled bool, ret rest.ResponseWrapper, err error) {
	proxyAction := action.(kubetest.ProxyGetActionImpl)

	// only respond to service proxy requests
	if handled = proxyAction.Matches(http.MethodGet, "services"); !handled {
		return
	}

	path := strings.TrimPrefix(proxyAction.Path, "/")
	if path == "" {
		ret = resultPage
	} else if data, ok := goodResults[path]; ok {
		ret = response(data)
	} else {
		ret = response{}
	}

	return
}

type response []byte

func (r response) DoRaw(context.Context) ([]byte, error) {
	if len(r) == 0 {
		return nil, errors.New("file does not exist")
	}
	return r, nil
}
func (r response) Stream(context.Context) (io.ReadCloser, error) {
	if len(r) == 0 {
		return nil, errors.New("file does not exist")
	}
	return ioutil.NopCloser(bytes.NewReader(r)), nil
}
