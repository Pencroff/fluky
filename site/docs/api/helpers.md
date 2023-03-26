
# Fluki helpers

## Weighted

`Weighted` is a helper that allows you to pick value from slice base on element weight.
It is useful when you want to pick value from slice with different probability.
Accepts two arguments: slice of values (`[]interface{}`) and slice of weights (`[]uint`).

```go
flk.Weighted([]interface{}{"a", "b", "c"}, []uint{1, 2, 3})
```


## PickOne

`PickOne` is a helper that allows you to pick one value from slice.
Accepts one argument: slice of values (`[]interface{}`).

```go
flk.PickOne([]interface{}{"a", "b", "c"})
```
