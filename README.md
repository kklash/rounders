## Rounders

Provides utilities for rounding and truncating floating point numbers.

Example:

```go
rounders.RoundToDecimals(12.48881, 2) // 12.49
rounders.TruncToDecimals(12.48881, 2) // 12.48

rounders.RoundToSigFigs(58.11869, 5) // 58.119
rounders.TruncToSigFigs(58.11869, 5) // 58.118
```
