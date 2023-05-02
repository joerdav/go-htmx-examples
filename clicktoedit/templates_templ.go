// Code generated by templ@(devel) DO NOT EDIT.

package clicktoedit

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

// GoExpression
import "examples/shared"

func Display(u user) templ.Component {
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
		// Element (standard)
		_, err = templBuffer.WriteString("<div")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" hx-target=\"this\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" hx-swap=\"outerHTML\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<div>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<label>")
		if err != nil {
			return err
		}
		// Text
		var_2 := `First Name`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label>")
		if err != nil {
			return err
		}
		// Text
		var_3 := `: `
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		// StringExpression
		var var_4 string = u.firstName
		_, err = templBuffer.WriteString(templ.EscapeString(var_4))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<div>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<label>")
		if err != nil {
			return err
		}
		// Text
		var_5 := `Last Name`
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label>")
		if err != nil {
			return err
		}
		// Text
		var_6 := `: `
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		// StringExpression
		var var_7 string = u.lastName
		_, err = templBuffer.WriteString(templ.EscapeString(var_7))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<div>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<label>")
		if err != nil {
			return err
		}
		// Text
		var_8 := `Email`
		_, err = templBuffer.WriteString(var_8)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label>")
		if err != nil {
			return err
		}
		// Text
		var_9 := `: `
		_, err = templBuffer.WriteString(var_9)
		if err != nil {
			return err
		}
		// StringExpression
		var var_10 string = u.email
		_, err = templBuffer.WriteString(templ.EscapeString(var_10))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<button")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" hx-get=\"/click-to-edit/contact/1/edit\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" class=\"button is-black\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Text
		var_11 := `Click To Edit`
		_, err = templBuffer.WriteString(var_11)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func Form(u user) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_12 := templ.GetChildren(ctx)
		if var_12 == nil {
			var_12 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		// Element (standard)
		_, err = templBuffer.WriteString("<form")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" hx-put=\"/click-to-edit/contact/1\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" hx-target=\"this\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" hx-swap=\"outerHTML\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<div")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"field\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<div")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"control\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<label>")
		if err != nil {
			return err
		}
		// Text
		var_13 := `First Name`
		_, err = templBuffer.WriteString(var_13)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label>")
		if err != nil {
			return err
		}
		// Element (void)
		_, err = templBuffer.WriteString("<input")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"input\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" type=\"text\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" name=\"firstName\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" value=")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(u.firstName))
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
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<div")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"field\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<div")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"control\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<label>")
		if err != nil {
			return err
		}
		// Text
		var_14 := `Last Name`
		_, err = templBuffer.WriteString(var_14)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label>")
		if err != nil {
			return err
		}
		// Element (void)
		_, err = templBuffer.WriteString("<input")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"input\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" type=\"text\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" name=\"lastName\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" value=")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(u.lastName))
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
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<div")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"field\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<div")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"control\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<label>")
		if err != nil {
			return err
		}
		// Text
		var_15 := `Email Address`
		_, err = templBuffer.WriteString(var_15)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label>")
		if err != nil {
			return err
		}
		// Element (void)
		_, err = templBuffer.WriteString("<input")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"input\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" type=\"email\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" name=\"email\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" value=")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(u.email))
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
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<div")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"field is-grouped\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<div")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"control\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<button")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"button is-black\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Text
		var_16 := `Submit`
		_, err = templBuffer.WriteString(var_16)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<div")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"control\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<button")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"button\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" hx-get=\"/click-to-edit\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Text
		var_17 := `Cancel`
		_, err = templBuffer.WriteString(var_17)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</form>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func Index(u user) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_18 := templ.GetChildren(ctx)
		if var_18 == nil {
			var_18 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		// TemplElement
		var_19 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
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
			var_20 := `Click to Edit`
			_, err = templBuffer.WriteString(var_20)
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
			// TemplElement
			err = Display(u).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = shared.Layout("Click to Edit").Render(templ.WithChildren(ctx, var_19), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

