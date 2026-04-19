## 用法

来源：[README](https://github.com/bigsmoke/pg_text_semver/blob/master/README.md)，[META.json](https://github.com/bigsmoke/pg_text_semver/blob/master/META.json)，[Tag v1.2.1](https://github.com/bigsmoke/pg_text_semver/tree/v1.2.1)

`pg_text_semver` 在 PostgreSQL `text` 之上实现 Semantic Versioning 2.0.0，并使用 `semver` domain，而不是自定义 C type。

### 核心类型与函数

```sql
CREATE EXTENSION pg_text_semver;

SELECT '1.2.3'::semver < '2.0.0'::semver;
SELECT semver_cmp('1.2.3'::semver, '1.2.4'::semver);
SELECT semver_regexp(true);
SELECT '1.2.3-alpha.1+build5'::semver::semver_parsed;
```

- `semver`：基于 `text` 的 domain，并带有 SemVer 校验。
- `semver_parsed`：用于排序与索引的解析后 composite type。
- `semver_prerelease`：用于 prerelease 标识符的 domain。
- `semver_cmp(...)`：`semver` 与 `semver_parsed` 的比较函数。
- `semver_regexp(include_captures boolean)`：暴露校验 regex。

### 额外辅助函数

当前 README 还记录了 PGXN version range 辅助函数：

- `meta_pgxn_version_range(text)`
- `meta_pgxn_version_range_cmp(text, text)`
- `nonsemver_cmp(text, text, text)`

### 注意事项

- 该扩展更偏向规范导向、基于 `text` 的实现，而不是较低层的 C-based alternatives。
- 上游 README 仍然是权威用户参考；这次刷新主要是使内容与文档中的 1.2.1 helper set 对齐。
