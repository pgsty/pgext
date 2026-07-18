## Usage

Sources:

- [Pinned Japanese normalization rules](https://github.com/nuko-yokohama/ntext/blob/4ec273e7ac4a6f45a25d8eb3b5c12602a7dd919d/readme_ja.txt)
- [Pinned SQL type and operator definitions](https://github.com/nuko-yokohama/ntext/blob/4ec273e7ac4a6f45a25d8eb3b5c12602a7dd919d/ntext--1.0.sql)
- [Pinned extension control file](https://github.com/nuko-yokohama/ntext/blob/4ec273e7ac4a6f45a25d8eb3b5c12602a7dd919d/ntext.control)

`ntext` defines a UTF-8 Japanese normalized-comparison text type. It normalizes hiragana, katakana, and half-width katakana to katakana; folds full-width and half-width Latin letters to lower-case half-width form; normalizes symbol width; and applies additional Japanese-specific substitutions before comparison.

```sql
CREATE EXTENSION ntext;

SELECT 'かな'::ntext = 'カナ'::ntext AS same_after_normalization;
SELECT pg_ntext_normalize('ＡＢＣ'::ntext);
SELECT pg_similar_rate('東京都'::ntext, '東京'::ntext);
```

The extension provides comparison operators, a default B-tree operator class, `pg_ntext_normalize(ntext)`, `pg_similar_rate(ntext,ntext)`, and approximate-match operator `/=`. Index and equality semantics therefore follow the normalized representation rather than bytewise source text.

Some rules are deliberately lossy: the upstream document maps the `ヴァ` row to the `バ` row, maps historic kana such as `ゐ`/`ヰ` and `ゑ`/`ヱ`, rewrites `を`/`ヲ`, and removes a trailing prolonged-sound mark. Validate those choices against application search policy and retain the original text separately. The reviewed version is `1.0`; upstream publishes neither a license statement nor a PostgreSQL compatibility matrix.
