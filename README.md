# numver

A library for extracting numeric parts from version strings

## Usage

`numver` only consider numeric parts in version strings, and ignore any others.

```go
v1 := numver.Parse("1.2.3b05_20240924")
v2 := numver.Version{1, 2, 3, 5, 20240924}
// v1 and v2 are the same
v1.Compare(v2) // = 0
```

## Credits

GUO YANKE, MIT License
