# Shortcodes

In layouts/shortcodes I can define element to inject into content files

## Example 1 - named parameter

superShort1.html
```html
<div>
    <h1 style="color:{{ .Get `color`}}">My title is in {{ .Get "color" }} color</h1>
</div>
```
is used like this:
```markdown
---
title: bla
date: 2023-01-16
author: thibault

---

This is my content
{{< superShort1 color="blue" >}}
```


## Example 2 - position parameter

superShort2.html
```html
<div>
    <h1 style="color:{{ .Get 0}}">My title is in {{ .Get 0 }} color</h1>
</div>
```
is used like this:
```markdown
---
title: bla
date: 2023-01-16
author: thibault

---

This is my content
{{< superShort2 blue >}}
```

## Example 3 - body

superShort3.html
```html
<div>
    <hr>
    {{ .Inner }}
    <hr>
</div>
```
is used like this:
```markdown
---
title: bla
date: 2023-01-16
author: thibault

---

This is my content
{{< superShort3 >}}
    This is the text that will go inside
{{< /superShort3 >}}
```
