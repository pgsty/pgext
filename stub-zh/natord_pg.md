## 用法

来源：

- [官方 README](https://github.com/mpajkowski/natord_pg/blob/a1f2923934a1ae4179e559ea651641f1b975bef0/README.md)
- [官方 Rust 实现](https://github.com/mpajkowski/natord_pg/blob/a1f2923934a1ae4179e559ea651641f1b975bef0/src/lib.rs)
- [官方操作符 SQL](https://github.com/mpajkowski/natord_pg/blob/a1f2923934a1ae4179e559ea651641f1b975bef0/sql/operators.sql)

`natord_pg` 0.0.0 为 `text` 添加自然顺序比较，让连续数字按数值而不是纯字节顺序比较。它适合 `file2`、`file10` 这类标识符，并提供非默认的 B-tree 操作符类，用于对应的谓词和索引。

### 核心流程

```sql
CREATE EXTENSION natord_pg;

CREATE TABLE artifacts (name text PRIMARY KEY);
INSERT INTO artifacts VALUES ('file10'), ('file2'), ('file1');

SELECT name
FROM artifacts
ORDER BY name USING <~~;

CREATE INDEX artifacts_name_natord_idx
  ON artifacts USING btree (name natord_ops);

SELECT name
FROM artifacts
WHERE name ~>= 'file2' AND name <~~ 'file20'
ORDER BY name USING <~~;
```

上述自然顺序会把 `file2` 排在 `file10` 之前。由于 `natord_ops` 没有声明为 `text` 的默认操作符类，创建自然顺序索引时必须显式指定它。

### 重要对象

- `natord_gt(text, text)`、`natord_lt(text, text)`、`natord_ge(text, text)` 和 `natord_le(text, text)` 实现自然顺序比较。
- `natord_cmp(text, text)` 返回 -1、0 或 1，并作为 B-tree 比较支持函数。
- `<~~`、`~~>`、`<=~` 和 `~>=` 分别表示自然顺序的小于、大于、小于等于和大于等于。
- `natordfam` 和 `natord_ops` 定义 B-tree 操作符族与操作符类。

### 运维说明

固定提交中的 manifest 声明版本为 `0.0.0`，提供 PostgreSQL 13 至 17 的构建 feature，默认 feature 为 PostgreSQL 17。扩展不会替换数据库排序规则或普通 `text` 操作符；查询必须使用自然顺序操作符才能匹配 `natord_ops`。在依赖唯一性或范围语义前，应针对上游 `natord` crate 的行为测试标点、正负号、前导零、非 ASCII 文本和受排序规则影响的数据。
