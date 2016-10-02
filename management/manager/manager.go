package manager

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/nilslice/cms/management/editor"
)

const managerHTML = `
<div class="card editor">
    <form method="post" action="/admin/edit">
		<input type="hidden" name="id" value="{{.ID}}"/>
		<input type="hidden" name="type" value="{{.Kind}}"/>
		{{ .Editor }}
	</form>
</div>
`

type manager struct {
	ID     int
	Kind   string
	Editor template.HTML
}

// Manage ...
func Manage(e editor.Editable, typeName string) ([]byte, error) {
	v, err := e.MarshalEditor()
	if err != nil {
		return nil, fmt.Errorf("Couldn't marshal editor for content %T. %s", e, err.Error())
	}

	m := manager{
		ID:     e.ContentID(),
		Kind:   typeName,
		Editor: template.HTML(v),
	}

	// execute html template into buffer for func return val
	buf := &bytes.Buffer{}
	tmpl := template.Must(template.New("manager").Parse(managerHTML))
	tmpl.Execute(buf, m)

	return buf.Bytes(), nil
}
