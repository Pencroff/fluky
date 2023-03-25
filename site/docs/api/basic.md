# Basic

## Bool

`Bool` returns a random bool value. By default, values have equal probability.

```go
flk.Bool() // true
```

### Parameters

#### WithLikelihood

`WithLikelihood` parameter allows to set the likelihood of the `True` result.
It accepts `v`, float64 value in the range [0, 1].

```go
flk.Bool(f.WithLikelihood(0.25)) // 25% chance to return true
```

## Integer

`Integer` returns a random integer value.
By default, the range is from -2^63^ to 2^63^-1.
The max value not included in the range.

```go
flk.Integer() // 3203168211198807973
// above is same as
flk.Integer(f.WithIntRange(math.MinInt64, math.MaxInt64))
```

### Parameters

#### WithIntRange

`WithIntRange` parameter allows to set the range of the result.
It accepts `min` and `max` values, `int`.
The `max` value not included in the range.

```go
flk.Integer(f.WithIntRange(0, 10)) // 8, 0 <= result < 10
```

## Float

`Float` returns a random float64 value.
By default, the range is from 0 to 1 and precision not limited (`-1`).
The max value not included in the range.

```go
flk.Float() // 0.5322073040624193
// above is same as
flk.Float(f.WithFloatRange(0, 1), f.WithPrecision(-1))
```

### Parameters

#### WithFloatRange

`WithFloatRange` parameter allows to set the range of the result.
It accepts `min` and `max` values, `float64`.
The `max` value not included in the range.

```go
flk.Float(f.WithFloatRange(0, 10)) // 8.5322073040624193, 0 <= result < 10
```

#### WithPrecision

`WithPrecision` parameter allows to set the precision of the result.
It accepts `p`, `int8` value.
The `p` value is the number of digits after the decimal point.
The negative value does not change the result.
Max value is 20.

```go
flk.Float(f.WithPrecision(0)) // 1
flk.Float(f.WithPrecision(2)) // 0.53
flk.Float(f.WithPrecision(-1)) // 0.5322073040624193
```

