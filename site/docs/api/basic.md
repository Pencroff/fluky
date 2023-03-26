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
By default, the range is from -2^62^ to 2^62^-1.


```go
flk.Integer() // 3203168211198807973
// above is same as
flk.Integer(f.WithIntRange(math.MinInt64, math.MaxInt64))
```

### Parameters

#### WithIntRange

`WithIntRange` parameter allows to set the range of the result.
It accepts `min` and `max` values, `int`.
The `min` and `max` values included in the range.

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
flk.Float(f.WithFloatRange(7.5, 12.7)) // 8.5322073, 7.5 <= result < 12.7
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

## String

`String` returns a random string value.
By default, the length is between 5 and 20 and
the charset is:

* latin lower (`abcdefghijklmnopqrstuvwxyz`)
* latin upper (`ABCDEFGHIJKLMNOPQRSTUVWXYZ`)
* numbers (`0123456789`)
* symbols (`!@#$%^&*+=_-`).

```go
flk.String() // peSnT7UeH0fq4sm
// above is same as
flk.String(f.WithLengthRange(5, 20), f.WithLatinLowerAlphabet(), f.AndLatinUpperAlphabet(), f.AndNumericAlphabet(), AndSymbolsAlphabet())
```
### Parameters

#### WithLength

`WithLength` parameter allows to set the length of the result.
It accepts `l`, `uint8` value.

```go
flk.String(f.WithLength(10)) // 10 characters long
```

#### WithLengthRange

`WithLengthRange` parameter allows to set the length range of the result.
It accepts `min` and `max` values, `uint8`. The `min` and `max` values included in the range.

```go
flk.String(f.WithLengthRange(5, 10)) // 5 <= length <= 10
```

#### WithAlphabet

`WithAlphabet` parameter allows to set the charset of the result.
It accepts `alphabet`, `string` value.

```go
flk.String(f.WithAlphabet("abc"))
```

#### AndAlphabet

`AndAlphabet` parameter allows to extend the charset of the result by the given alphabet.

```go
flk.String(f.WithAlphabet("12345"), f.AndAlphabet("abcde"))
```

#### WithNumericAlphabet

`WithNumericAlphabet` parameter allows to set the charset of the result to numbers.

```go
flk.String(f.WithNumericAlphabet())
```

#### AndNumericAlphabet

`AndNumericAlphabet` parameter allows to extend the charset of the result by numbers.

```go
flk.String(f.WithAlphabet("abcde"), f.AndNumericAlphabet())
```

#### WithLatinLowerAlphabet

`WithLatinLowerAlphabet` parameter allows to set the charset of the result to lower case letters.

```go
flk.String(f.WithLatinLowerAlphabet())
```

#### AndLowerAlphabet

`AndLatinLowerAlphabet` parameter allows to extend the charset of the result by lower case letters.

```go
flk.String(f.WithAlphabet("ABC"), f.AndLatinLowerAlphabet())
```

#### WithLatinUpperAlphabet

`WithLatinUpperAlphabet` parameter allows to set the charset of the result to upper case letters.

```go
flk.String(f.WithLatinUpperAlphabet())
```

#### AndLatinUpperAlphabet

`AndLatinUpperAlphabet` parameter allows to extend the charset of the result by upper case letters.

```go
flk.String(f.WithAlphabet("12345"), f.AndLatinUpperAlphabet())
```

#### WithSymbolsAlphabet

`WithSymbolsAlphabet` parameter allows to set the charset of the result to symbols.
The symbols are: `!@#$%^&*+=_-`.

```go
flk.String(f.WithSymbolsAlphabet())
```

#### AndSymbolsAlphabet

`AndSymbolsAlphabet` parameter allows to extend the charset of the result by symbols.
The symbols are: `!@#$%^&*+=_-`.

```go
flk.String(f.WithAlphabet("abcde"), f.AndSymbolsAlphabet())
```

#### WithSymbolsUrlSafeAlphabet

`WithSymbolsUrlSafeAlphabet` parameter allows to set the charset of the result to url safe symbols.
The symbols are: `_-`.

```go
flk.String(f.WithSymbolsUrlSafeAlphabet())
```

#### AndSymbolsUrlSafeAlphabet

`AndSymbolsUrlSafeAlphabet` parameter allows to extend the charset of the result by url safe symbols.
The symbols are: `_-`.

```go
flk.String(f.WithAlphabet("abcde"), f.AndSymbolsUrlSafeAlphabet())
```

#### WithUrlSafeAlphabet

`WithUrlSafeAlphabet` parameter allows to set the charset of the result to url safe symbols and letters.
The symbols are: `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-`.

```go
flk.String(f.WithUrlSafeAlphabet())
```

#### WithHexAlphabet

`WithHexAlphabet` parameter allows to set the charset of the result to hex symbols.
The symbols are: `0123456789abcdef`.

```go
flk.String(f.WithHexAlphabet())
```

#### WithUkrainianLowerAlphabet

`WithUkrainianLowerAlphabet` parameter allows to set the charset of the result to ukrainian lower case letters.

```go
flk.String(f.WithUkrainianLowerAlphabet())
```

#### AndUkrainianLowerAlphabet

`AndUkrainianLowerAlphabet` parameter allows to extend the charset of the result by ukrainian lower case letters.

```go
flk.String(f.WithAlphabet("0123456789"), f.AndUkrainianLowerAlphabet())
```

#### WithUkrainianUpperAlphabet

`WithUkrainianUpperAlphabet` parameter allows to set the charset of the result to ukrainian upper case letters.

```go
flk.String(f.WithUkrainianUpperAlphabet())
```

#### AndUkrainianUpperAlphabet

`AndUkrainianUpperAlphabet` parameter allows to extend the charset of the result by ukrainian upper case letters.

```go
flk.String(f.WithAlphabet("0123456789"), f.AndUkrainianUpperAlphabet())
```

#### WithGreekLowerAlphabet

`WithGreekLowerAlphabet` parameter allows to set the charset of the result to greek lower case letters.

```go
flk.String(f.WithGreekLowerAlphabet())
```

#### AndGreekLowerAlphabet

`AndGreekLowerAlphabet` parameter allows to extend the charset of the result by greek lower case letters.

```go
flk.String(f.WithAlphabet("0123456789"), f.AndGreekLowerAlphabet())
```

#### WithGreekUpperAlphabet

`WithGreekUpperAlphabet` parameter allows to set the charset of the result to greek upper case letters.

```go
flk.String(f.WithGreekUpperAlphabet())
```

#### AndGreekUpperAlphabet

`AndGreekUpperAlphabet` parameter allows to extend the charset of the result by greek upper case letters.

```go
flk.String(f.WithAlphabet("0123456789"), f.AndGreekUpperAlphabet())
```
