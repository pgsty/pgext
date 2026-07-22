## 用法

来源：

- [Official README](https://github.com/adjust/pg-ajversion/blob/502776cf0c17d7ad034a2b5bb36ba742f3e8334d/README.md)
- [Version 0.0.1 SQL API](https://github.com/adjust/pg-ajversion/blob/502776cf0c17d7ad034a2b5bb36ba742f3e8334d/ajversion--0.0.1.sql)
- [C input, output, and comparison implementation](https://github.com/adjust/pg-ajversion/blob/502776cf0c17d7ad034a2b5bb36ba742f3e8334d/src/ajversion.c)

`ajversion` 0.0.1 定义了一种紧凑、可排序的版本类型，用于保存最多三个数字分量的版本号。它适合存储并索引简单的 `major.minor.patch` 值，但不是完整的语义化版本实现：不支持预发布标识、构建元数据或严格输入校验。

### 核心用法

```sql
CREATE EXTENSION ajversion;

CREATE TABLE software_release (
  package text NOT NULL,
  version ajversion NOT NULL
);

INSERT INTO software_release VALUES
  ('demo', '6.7.4'),
  ('demo', '12.5.7');

SELECT version, major(version), minor(version), patch(version)
FROM software_release
WHERE version >= '7.0.0'::ajversion
ORDER BY version;
```

默认的 B-tree 与哈希操作符类可用于普通比较索引：

```sql
CREATE INDEX software_release_version_idx
ON software_release (version);
```

### 主要对象

- `ajversion`：四字节版本值。
- `major(ajversion)`、`minor(ajversion)`、`patch(ajversion)`：提取各分量。
- `=`、`<>`、`<`、`<=`、`>`、`>=`：按打包后的数字值比较。
- `btree_ajversion_ops`、`hash_ajversion_ops`：默认 B-tree 与哈希操作符类。
- 从 `integer`、`numeric`、`real`、`double precision` 的类型转换：把数值输入转换为 `ajversion`。

### 解析与取值边界

内部表示为主版本预留 12 位，为次版本和补丁版本各预留 10 位，因此有效范围分别是主版本 `0..4095`、次版本/补丁版本 `0..1023`；超出范围会报错。

输入解析有意采用宽松策略：它会在任意非数字字符之间扫描最多三段数字，缺失分量补零，并把输出规范化为三个分量。例如 `'15'` 会变成 `15.0.0`，完全没有数字的文本则变成 `0.0.0`。如果必须拒绝畸形文本，应在写入前另行校验；涉及预发布或构建字段的 SemVer 优先级也不应使用 `ajversion`。
