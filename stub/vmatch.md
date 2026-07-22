## Usage

Sources:

- [Official upstream documentation](https://github.com/nuko-yokohama/vmatch/blob/6fd255384005a84a3bd8e05dc2e099eb60fa1bdf/readme_jp.txt)
- [Version 1.0 SQL definitions](https://github.com/nuko-yokohama/vmatch/blob/6fd255384005a84a3bd8e05dc2e099eb60fa1bdf/vmatch--1.0.sql)
- [Extension control file](https://github.com/nuko-yokohama/vmatch/blob/6fd255384005a84a3bd8e05dc2e099eb60fa1bdf/vmatch.control)

`vmatch` provides a text-similarity function and the fuzzy boolean operator `/=`. Its custom scoring gives special weight to likely QWERTY-keyboard typos and is intended for small, approximate matching experiments rather than calibrated search ranking.

### Core Workflow

```sql
CREATE EXTENSION vmatch;

SELECT data, get_similar_rate(data, 'postgresql') AS similarity
FROM candidate
WHERE data /= 'postgresql'
ORDER BY similarity DESC;
```

The upstream examples show case-insensitive matches and typo variants such as `PostgresSQL`, `pastgres`, and `pstgres`. Verify the score distribution on the language and data that matter to the application.

### Objects and Semantics

- `get_similar_rate(text, text)` returns a `real` similarity score.
- `vmatch(text, text)` returns the boolean decision used by `/=`.
- `/=` uses a fixed threshold of 0.75 in version `1.0`; there is no documented GUC to change it and no index operator class.

Upstream explicitly lists Japanese-language support, an improved formula, and configurable thresholds as unfinished work, and says the algorithm's practical usefulness has not been validated. The extension is `STRICT` and `IMMUTABLE`, but those declarations do not make its fuzzy result linguistically correct. Do not use `/=` as an equality, uniqueness, or authorization test; normalize text intentionally and test multibyte input, accents, locale, long strings, NULLs, false positives, and false negatives before use.
