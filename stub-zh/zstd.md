
## 用法

| 函数                                                                                 | 返回类型    |
|--------------------------------------------------------------------------------------|-------------|
| <code>zstd_compress(*data* bytea [, *dictionary* bytea [, *level* integer ]])</code> | `bytea`     |
| <code>zstd_decompress(*data* bytea [, *dictionary* bytea ])</code>                   | `bytea`     |
| <code>zstd_length(*data* bytea)</code>                                               | `integer`   |

`zstd_compress` 对提供的 `data` 进行压缩，返回一个 Zstandard 帧。可以选择提供预设的 `dictionary`（字典）。也可以覆盖默认的压缩 `level`（级别），有效值范围为 `1`（最快速度）到 `22`（最高压缩率），默认级别为 `3`。

如果需要在不使用字典的情况下覆盖压缩级别，请将 `dictionary` 设置为 `NULL`。

`zstd_decompress` 对 `data` 中提供的 Zstandard 帧进行解压缩，返回解压后的数据。可以选择提供与压缩时相同的预设 `dictionary`（字典）。

`zstd_length` 返回给定 Zstandard 帧的解压后长度。如果 `ZSTD_getFrameContentSize()` 可用，当长度未知时返回 `NULL`。如果该函数不可用，则无法区分错误、未知解压长度和零解压长度这几种情况。


### 示例

基本的压缩/解压缩示例：

```sql
CREATE EXTENSION zstd;

SELECT zstd_compress('hello world');
-- zstd_compress
-- --------------------------------------------
-- \x28b52ffd200b59000068656c6c6f20776f726c64

SELECT convert_from(zstd_decompress('\x28b52ffd200b59000068656c6c6f20776f726c64'), 'utf-8');
-- convert_from
-- --------------
--  hello world
```

指定压缩级别（`1` 为最快速度，`22` 为最高压缩率，默认为 `3`）

```sql
CREATE EXTENSION zstd;

SELECT zstd_compress('hello world',  NULL, 10);
-- zstd_compress
-- --------------------------------------------
-- \x28b52ffd200b59000068656c6c6f20776f726c64

SELECT convert_from(zstd_decompress('\x28b52ffd200b59000068656c6c6f20776f726c64'), 'utf-8');
-- convert_from
-- --------------
--  hello world
```
