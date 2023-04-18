package shared

import (
	"bytes"
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/yosssi/gohtml"
)

func Raw() templ.ComponentFunc {
	return func(ctx context.Context, w io.Writer) error {
		b := new(bytes.Buffer)
		if err := templ.GetChildren(ctx).Render(ctx, b); err != nil {
			return err
		}
		gohtml.Condense = true
		str := templ.EscapeString(gohtml.Format(b.String()))
		if _, err := w.Write([]byte(str)); err != nil {
			return err
		}
		return nil
	}
}
