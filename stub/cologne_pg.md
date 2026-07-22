## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/DrSloth/cologne_pg/blob/6871c07bb24c9a05e63b5e2b034bf34eda64426c/README.md)
- [Rust implementation](https://github.com/DrSloth/cologne_pg/blob/6871c07bb24c9a05e63b5e2b034bf34eda64426c/src/lib.rs)
- [Rust package manifest](https://github.com/DrSloth/cologne_pg/blob/6871c07bb24c9a05e63b5e2b034bf34eda64426c/Cargo.toml)

`cologne_pg` exposes the Cologne phonetics algorithm as `cologne(text)`, producing a phonetic code intended for approximate matching of German-language names.

```sql
CREATE EXTENSION cologne_pg;
SELECT cologne('MORELLO');

SELECT person_id, display_name
FROM people
WHERE cologne(display_name) = cologne('Müller');
```

Phonetic equality is a candidate-generation heuristic, not identity. Different names can share a code, spellings from other languages may behave poorly, and transliteration, punctuation, casing, Unicode normalization, and version changes can affect results. Always retain the original string and apply an additional ranking or human-review step.

If the function is used in an expression index or generated column, pin the exact implementation and rebuild dependent data after any behavior change. Test representative German names, diacritics, empty strings, `NULL`, very long input, and adversarial collision rates before using the result for search.
