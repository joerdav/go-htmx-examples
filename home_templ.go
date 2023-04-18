// Code generated by templ@(devel) DO NOT EDIT.

package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

// GoExpression
import "examples/shared"

func Home(examples []Example) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		// TemplElement
		var_2 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<h2")
			if err != nil {
				return err
			}
			// Element Attributes
			_, err = templBuffer.WriteString(" class=\"title\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(">")
			if err != nil {
				return err
			}
			// Text
			var_3 := `Examples:`
			_, err = templBuffer.WriteString(var_3)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</h2>")
			if err != nil {
				return err
			}
			// Whitespace (normalised)
			_, err = templBuffer.WriteString(` `)
			if err != nil {
				return err
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<table")
			if err != nil {
				return err
			}
			// Element Attributes
			_, err = templBuffer.WriteString(" class=\"table is-fullwidth\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(">")
			if err != nil {
				return err
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<thead>")
			if err != nil {
				return err
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<tr>")
			if err != nil {
				return err
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<th>")
			if err != nil {
				return err
			}
			// Text
			var_4 := `Pattern`
			_, err = templBuffer.WriteString(var_4)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</th>")
			if err != nil {
				return err
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<th>")
			if err != nil {
				return err
			}
			// Text
			var_5 := `Description`
			_, err = templBuffer.WriteString(var_5)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</th>")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</tr>")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</thead>")
			if err != nil {
				return err
			}
			// Element (standard)
			_, err = templBuffer.WriteString("<tbody>")
			if err != nil {
				return err
			}
			// For
			for _, e := range examples {
				// Element (standard)
				_, err = templBuffer.WriteString("<tr>")
				if err != nil {
					return err
				}
				// Element (standard)
				_, err = templBuffer.WriteString("<td>")
				if err != nil {
					return err
				}
				// Element (standard)
				_, err = templBuffer.WriteString("<a")
				if err != nil {
					return err
				}
				// Element Attributes
				_, err = templBuffer.WriteString(" href=")
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("\"")
				if err != nil {
					return err
				}
				var var_6 templ.SafeURL = templ.SafeURL("/"+e.Slug)
				_, err = templBuffer.WriteString(templ.EscapeString(string(var_6)))
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("\"")
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(">")
				if err != nil {
					return err
				}
				// StringExpression
				var var_7 string = e.Name
				_, err = templBuffer.WriteString(templ.EscapeString(var_7))
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</a>")
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</td>")
				if err != nil {
					return err
				}
				// Element (standard)
				_, err = templBuffer.WriteString("<td>")
				if err != nil {
					return err
				}
				// StringExpression
				var var_8 string = e.Desc
				_, err = templBuffer.WriteString(templ.EscapeString(var_8))
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</td>")
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</tr>")
				if err != nil {
					return err
				}
			}
			_, err = templBuffer.WriteString("</tbody>")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</table>")
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = shared.Layout("Home").Render(templ.WithChildren(ctx, var_2), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}
