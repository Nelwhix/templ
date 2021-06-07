// Code generated by templ DO NOT EDIT.

package testcall

import "github.com/a-h/templ"
import "context"
import "io"

func personTemplate(p person) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, "<div>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<h1>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(p.name))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</h1>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<div")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " style=\"font-family: &#39;sans-serif&#39;\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " id=\"test\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " data-contents=")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(`something with "quotes" and a <tag>`))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		err = email(p.email).Render(ctx, w)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		return err
	})
}

func email(s string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, "<div>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString("email:"))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<a")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " href=")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		var var_1 templ.SafeURL = templ.URL("mailto: " + s)
		_, err = io.WriteString(w, templ.EscapeString(string(var_1)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(s))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</a>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		return err
	})
}

