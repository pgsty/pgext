## Usage

Sources:

- [Upstream README](https://github.com/palestamp/hamming_distance/blob/22c4c95bc1e8b92bfc32bc0c6024a1bacd5ce4aa/README.md)
- [Extension control file](https://github.com/palestamp/hamming_distance/blob/22c4c95bc1e8b92bfc32bc0c6024a1bacd5ce4aa/hamming_distance.control)
- [SQL installation script](https://github.com/palestamp/hamming_distance/blob/22c4c95bc1e8b92bfc32bc0c6024a1bacd5ce4aa/hamming_distance--0.0.1.sql)
- [C implementation](https://github.com/palestamp/hamming_distance/blob/22c4c95bc1e8b92bfc32bc0c6024a1bacd5ce4aa/hamming_distance.c)

`hamming_distance` version `0.0.1` compares equal-length hexadecimal strings bit by bit. It exposes raw and normalized distance and similarity functions.

### Example

```sql
CREATE EXTENSION hamming_distance;
SELECT hamming_distance('1d3f', '1110');
SELECT hamming_similarity('1d3f', '1110');
SELECT hamming_distance_normalized('1d3f', '1110');
```

Inputs should be prevalidated as nonempty, equal-length, even-length strings matching `^[0-9A-Fa-f]+$`. The old C parser does not reliably reject every non-hex character and does not account for a trailing half-byte in odd-length input. The repository has no current PostgreSQL compatibility statement, so test the exact build and inputs in isolation before relying on results.
