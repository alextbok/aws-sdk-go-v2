// +build codegen

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"
)

// WaiterAcceptor is the acceptors defined in the model the SDK will use
// to wait on resource states with.
type WaiterAcceptor struct {
	State    string
	Matcher  string
	Argument string
	Expected interface{}
}

// ExpectedString returns the string that was expected by the WaiterAcceptor
func (a *WaiterAcceptor) ExpectedString() string {
	switch a.Expected.(type) {
	case string:
		return fmt.Sprintf("%q", a.Expected)
	default:
		return fmt.Sprintf("%v", a.Expected)
	}
}

// A Waiter is an individual waiter definition.
type Waiter struct {
	Name          string
	Delay         int
	MaxAttempts   int
	OperationName string `json:"operation"`
	Operation     *Operation
	Acceptors     []WaiterAcceptor
}

// WaitersGoCode generates and returns Go code for each of the waiters of
// this API.
func (a *API) WaitersGoCode() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "import (\n%q\n%q\n\n%q\n)",
		"time",
		"context",
		SDKImportRoot+"/aws",
	)

	for _, w := range a.Waiters {
		buf.WriteString(w.GoCode())
	}
	return buf.String()
}

// used for unmarshaling from the waiter JSON file
type waiterDefinitions struct {
	*API
	Waiters map[string]Waiter
}

// AttachWaiters reads a file of waiter definitions, and adds those to the API.
// Will panic if an error occurs.
func (a *API) AttachWaiters(filename string) error {
	p := waiterDefinitions{API: a}

	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return err
	}
	err = json.NewDecoder(f).Decode(&p)
	if err != nil {
		return fmt.Errorf("failed to decode %s, err: %v", filename, err)
	}

	return p.setup()
}

func (p *waiterDefinitions) setup() error {
	p.API.Waiters = []Waiter{}
	i, keys := 0, make([]string, len(p.Waiters))
	for k := range p.Waiters {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	for _, n := range keys {
		e := p.Waiters[n]
		n = p.ExportableName(n)
		e.Name = n
		e.OperationName = p.ExportableName(e.OperationName)
		e.Operation = p.API.Operations[e.OperationName]
		if e.Operation == nil {
			return fmt.Errorf("unknown operation %s for waiter %s",
				e.OperationName, n)
		}
		p.API.Waiters = append(p.API.Waiters, e)
	}

	return nil
}

var waiterTmpls = template.Must(template.New("waiterTmpls").Funcs(
	template.FuncMap{
		"titleCase": func(v string) string {
			return strings.Title(v)
		},
	},
).Parse(`
{{ define "waiter"}}
// WaitUntil{{ .Name }} uses the {{ .Operation.API.NiceName }} API operation
// {{ .OperationName }} to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *{{ .Operation.API.StructName }}) WaitUntil{{ .Name }}(` +
	`ctx context.Context, input {{ .Operation.InputRef.GoType }}, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:    "WaitUntil{{ .Name }}",
		MaxAttempts: {{ .MaxAttempts }},
		Delay: aws.ConstantWaiterDelay({{ .Delay }} * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{{ range $_, $a := .Acceptors }}{
				State:    aws.{{ titleCase .State }}WaiterState,
				Matcher:  aws.{{ titleCase .Matcher }}WaiterMatch,
				{{- if .Argument }}Argument: "{{ .Argument }}",{{ end }}
				Expected: {{ .ExpectedString }},
			},
			{{ end }}
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy {{ .Operation.InputRef.GoType }}
			if input != nil  {
				tmp := *input
				inCpy = &tmp
			}
			req := c.{{ .OperationName }}Request(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}
{{- end }}

{{ define "waiter interface" }}
WaitUntil{{ .Name }}(context.Context, {{ .Operation.InputRef.GoTypeWithPkgName }}, ...aws.WaiterOption) error
{{- end }}
`))

// InterfaceSignature returns a string representing the Waiter's interface
// function signature.
func (w *Waiter) InterfaceSignature() string {
	w.Operation.API.imports[packageImport{Path: "context"}] = true

	var buf bytes.Buffer
	if err := waiterTmpls.ExecuteTemplate(&buf, "waiter interface", w); err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

// GoCode returns the generated Go code for an individual waiter.
func (w *Waiter) GoCode() string {
	var buf bytes.Buffer
	if err := waiterTmpls.ExecuteTemplate(&buf, "waiter", w); err != nil {
		panic(err)
	}

	return buf.String()
}
