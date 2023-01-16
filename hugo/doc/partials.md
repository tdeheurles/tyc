# Partials

In layouts/partials I can define element to inject into templates files (html)

## Example 1

ie: header.html
```html
<div>
    <h1>{{ .Title }}</h1>
    <hr>
</div>
```
used like this in templates:
```html
<html>
    <body>
        {{ partial "header" . }}
    </body>
</html>
```
The . (dot) represent the context you pass to the partial. Here it means the page.


## Example 2 - passing dict
ie: something.html
```html
<div>
    {{ .foo }}
    <hr>
    {{ .bar }}
</div>
```
used like this in templates:
```html
<html>
    <body>
        {{ partial (dict "foo" "FOO" "bar" "BAR) }}
    </body>
</html>
```
