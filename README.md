# Go HTMX Examples

<https://htmx.org/examples> implemented in Go and templ.

## Tasks

### run

Runs the app and looks for changes.

```
reflex -r '.*\.(go|templ)' -R '.*_templ\.go' -s -- sh -c 'templ generate && go run .'
```
