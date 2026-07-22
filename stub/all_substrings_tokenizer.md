## Usage

Sources:

- [Official README](https://github.com/adjust/all_substrings_tokenizer/blob/f459a88981981659b7f21fab955ca1cbc802c219/README.md)
- [Version 0.1 extension SQL](https://github.com/adjust/all_substrings_tokenizer/blob/f459a88981981659b7f21fab955ca1cbc802c219/sql/all_substrings_tokenizer--0.1.sql)
- [C implementation](https://github.com/adjust/all_substrings_tokenizer/blob/f459a88981981659b7f21fab955ca1cbc802c219/src/all_substrings_tokenizer.c)

`all_substrings_tokenizer` expands text into every contiguous substring containing at least three characters. It is a small multibyte-aware set-returning function that can feed custom search, matching, or demonstration pipelines; it is not a PostgreSQL text search parser or configurable tokenizer.

### Core Workflow

```sql
CREATE EXTENSION all_substrings_tokenizer;

SELECT token
FROM all_substrings_set('1二湖楽a');
```

The result contains `1二湖`, `1二湖楽`, `1二湖楽a`, `二湖楽`, `二湖楽a`, and `湖楽a`. Use the composite result's `token` field when joining it to other data:

```sql
SELECT d.id, s.token
FROM documents AS d
CROSS JOIN LATERAL all_substrings_set(d.title) AS s
WHERE char_length(s.token) = 3;
```

### Object Index

- `all_substrings_set(text) -> setof __substrings_token` emits the substrings.
- `__substrings_token` is a composite type with one `token text` attribute.

### Operational Notes

The minimum length is hard-coded as three characters; there is no SQL parameter or GUC to change it. Inputs shorter than three characters return no rows. The function is `IMMUTABLE` and `STRICT`, so SQL `NULL` also returns no rows.

Output grows quadratically with input length: a string of `n` characters produces `(n - 2) * (n - 1) / 2` rows. Bound input size before using the function on user-controlled text. Version `0.1` is relocatable and declares no preload requirement, but the upstream repository is archived and should be compatibility-tested on the target PostgreSQL release.
