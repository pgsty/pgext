

## 用法

> [pg_tiktoken](https://github.com/kelvich/pg_tiktoken)：用于 PostgreSQL 中 OpenAI 模型的 tiktoken 分词器。
> 来源：[README.md](https://raw.githubusercontent.com/kelvich/pg_tiktoken/main/README.md)

`pg_tiktoken` 是一个 PostgreSQL 扩展，使用 OpenAI 的 [tiktoken](https://github.com/openai/tiktoken) 库提供输入分词功能。它允许你直接在 SQL 中计数和编码 token，这在处理 OpenAI 模型的输入长度限制时非常有用。


--------

### 函数

#### tiktoken_count

计算给定编码或模型的 token 数量：

```sql
SELECT tiktoken_count('p50k_edit', 'A long time ago in a galaxy far, far away');
 tiktoken_count
----------------
             11
(1 row)
```

#### tiktoken_encode

获取给定编码或模型的 token ID：

```sql
SELECT tiktoken_encode('cl100k_base', 'A long time ago in a galaxy far, far away');
                  tiktoken_encode
----------------------------------------------------
 {32,1317,892,4227,304,264,34261,3117,11,3117,3201}
(1 row)
```

`tiktoken_count` 和 `tiktoken_encode` 的第一个参数都可以接受编码名称或 OpenAI 模型名称。


--------

### 支持的模型

| 编码名称 | OpenAI 模型 |
|----------|-------------|
| `cl100k_base` | ChatGPT 模型, `text-embedding-ada-002` |
| `p50k_base` | 代码模型, `text-davinci-002`, `text-davinci-003` |
| `p50k_edit` | 编辑模型如 `text-davinci-edit-001`, `code-davinci-edit-001` |
| `r50k_base`（或 `gpt2`） | GPT-3 模型如 `davinci` |
