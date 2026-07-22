## 用法

来源：

- [官方 README](https://github.com/clowder/pg_logfmt/blob/aaa2da2b71c6864264ad71b8d1d4a198d3fb5d9c/README.md)
- [SQL 函数实现](https://github.com/clowder/pg_logfmt/blob/aaa2da2b71c6864264ad71b8d1d4a198d3fb5d9c/src/lib.rs)
- [解析器实现](https://github.com/clowder/pg_logfmt/blob/aaa2da2b71c6864264ad71b8d1d4a198d3fb5d9c/src/parser.rs)

`pg_logfmt` 可解析文本中的 logfmt 键值对。它适合把应用日志行转换为 `jsonb`、逐行提取键，或在建立索引与分析前把键收集为文本数组。

### 核心流程

```sql
CREATE EXTENSION pg_logfmt;

SELECT logfmt_to_jsonb(
    'at=info method=POST path="/orders/42" status=204 bytes=0'
);

SELECT *
FROM logfmt_keys(
    'at=info method=POST path="/orders/42" status=204 bytes=0'
);

SELECT logfmt_keys_array(
    'at=info method=POST path="/orders/42" status=204 bytes=0'
);
```

`logfmt_to_jsonb(text)` 返回一个对象，其中值为字符串或 JSON null。`logfmt_keys(text)` 以集合形式返回键，`logfmt_keys_array(text)` 则把相同的键返回为数组。三个函数都声明为 immutable 且 parallel safe。

### 解析行为

解析器接受裸值与双引号值，后者可以包含空格。它可以跳过非 logfmt 前缀，从第一个可识别的 `key=value` 键值对开始解析，因此适用于带标签前缀的应用日志。未加引号的空值会变成 null；明确加引号的空字符串仍为空字符串。值中双重转义的引号会被 `logfmt_to_jsonb` 还原。

如果没有可识别的键值对，`logfmt_to_jsonb` 与 `logfmt_keys_array` 返回 NULL，而 `logfmt_keys` 不返回任何行。转换为 `jsonb` 时重复键只保留最后一个值，但返回键的函数会保留解析到的各次出现。

### 注意事项

已复核的上游版本为 `0.0.0`，其 Cargo 特性矩阵覆盖 PostgreSQL `14` 与 `15`。它是一个专用解析器，而不是自动推断 Schema 的系统：值不会转换为数字、布尔值或时间戳。依赖它完成采集前，应验证引号损坏、重复键、未知转义及自由文本混排等边界情况。不需要预加载或重启。
