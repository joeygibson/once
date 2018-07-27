# once
Like uniq, but doesn't require sorted input

With input like this

```
foo
bar
baz
baz
baz
baz
bar
text
foo
foo
chester
```

running it like this `cat foo.txt | once` will result in this output

```
foo
bar
baz
text
chester
```

