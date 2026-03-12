

## Usage

> [pg_tiktoken](https://github.com/kelvich/pg_tiktoken): tiktoken tokenizer for use with OpenAI models in PostgreSQL.
> Source: [README.md](https://raw.githubusercontent.com/kelvich/pg_tiktoken/main/README.md)

`pg_tiktoken` is a PostgreSQL extension that provides input tokenization using OpenAI's [tiktoken](https://github.com/openai/tiktoken) library. It allows you to count and encode tokens directly in SQL, which is useful for managing input length limits when working with OpenAI models.


--------

### Functions

#### tiktoken_count

Count the number of tokens for a given encoding or model:

```sql
SELECT tiktoken_count('p50k_edit', 'A long time ago in a galaxy far, far away');
 tiktoken_count
----------------
             11
(1 row)
```

#### tiktoken_encode

Get the token IDs for a given encoding or model:

```sql
SELECT tiktoken_encode('cl100k_base', 'A long time ago in a galaxy far, far away');
                  tiktoken_encode
----------------------------------------------------
 {32,1317,892,4227,304,264,34261,3117,11,3117,3201}
(1 row)
```

Both `tiktoken_count` and `tiktoken_encode` accept either an encoding name or an OpenAI model name as the first argument.


--------

### Supported Models

| Encoding name | OpenAI models |
|---|---|
| `cl100k_base` | ChatGPT models, `text-embedding-ada-002` |
| `p50k_base` | Code models, `text-davinci-002`, `text-davinci-003` |
| `p50k_edit` | Edit models like `text-davinci-edit-001`, `code-davinci-edit-001` |
| `r50k_base` (or `gpt2`) | GPT-3 models like `davinci` |
