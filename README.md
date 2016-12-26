# Easy Option

Easy Option (ezopt) invokes functions with arguments which parsed from string.

Here is simple usage:

```go
fn1 := func(s string, n int) error { ... }

// OK
ezopt.Run(fn1, "abc", "123")

// FAILED: 2nd arg should be parsed as int.
ezopt.Run(fn1, "abc", "def")
```
