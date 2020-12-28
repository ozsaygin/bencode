# bencode

bencode is a bencoding package written in Go.

# Usage

Import the repository like this.

Encoding can be achieved like below.

```go

`import "github.com/ozsaygin/"`

// Decoding

var v map[string]interface {}

text := "d3:cow3:moo4:spam4:eggse"
bencode.Decode(v)

// Encoding

text := map[string]string {
    "cow": "moo", 
    "spam": "eggs",
}
data := bencode.Encode(text)

```

Decoding can be achieved like below.

```


# Tests

# Documentation

[bencoder spesification](https://www.bittorrent.org/beps/bep_0003.html)

# Contributing
# License
```
