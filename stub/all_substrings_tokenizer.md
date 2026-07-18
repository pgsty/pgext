## Usage

Sources:

- [Official extension control file](https://github.com/adjust/all_substrings_tokenizer/blob/f459a88981981659b7f21fab955ca1cbc802c219/all_substrings_tokenizer.control)
- [Official upstream documentation](https://github.com/adjust/all_substrings_tokenizer/blob/f459a88981981659b7f21fab955ca1cbc802c219/README.md)

`all_substrings_tokenizer` — Tokenizes text to all substrings bigger than a fixed character length.

The reviewed catalog snapshot records version `0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "all_substrings_tokenizer";
SELECT extversion
FROM pg_extension
WHERE extname = 'all_substrings_tokenizer';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
