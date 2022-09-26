package httpx

import (
	"context"
	"log"
	"net/http"
	"text/template"
	"time"

	"alukart32.com/usage/context/pkg/google"
	"alukart32.com/usage/context/pkg/userip"
)

var (
	paths = map[string]http.HandlerFunc{
		"/search": searchHandler,
	}
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	for p, h := range paths {
		mux.HandleFunc(p, h)
	}
	return mux
}

var resultsTemplate = template.Must(template.New("results").Parse(`
<html>
<head/>
<body>
  <ol>
  {{range .Results}}
  	<li>{{.Title}} - <a href="{{.URL}}">{{.URL}}</a></li>
  {{end}}
  </ol>
  <p>{{len .Results}} results in {{.Elapsed}}; timeout {{.Timeout}}</p>
</body>
</html>
`))

func searchHandler(w http.ResponseWriter, r *http.Request) {
	// ctx is the Context for this handler. Calling cancel closes the
	// ctx.Done channel, which is the cancellation signal for requests
	// started by this handler.
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	timeout, err := time.ParseDuration(r.URL.Query()["timeout"][0])
	if err == nil {
		// The request has a timeout, so create a context that is
		// canceled automatically when the timeout expires.
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel()

	query := r.URL.Query()["q"]
	if query == nil {
		http.Error(w, "no query", http.StatusBadRequest)
		return
	}

	// get userIP from request
	userIP, err := userip.FromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// create anew context for userIP
	ctx = userip.NewContext(ctx, userIP)

	// Run the Google search and print the results.
	start := time.Now()
	results, err := google.Search(ctx, query[0])
	elapsed := time.Since(start)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := resultsTemplate.Execute(w, struct {
		Results          google.Results
		Timeout, Elapsed time.Duration
	}{
		Results: results,
		Timeout: timeout,
		Elapsed: elapsed,
	}); err != nil {
		log.Print(err)
		return
	}
}
