## 用法

- 来源: [GitHub 仓库](https://github.com/bigsmoke/pg_text_semver), [README](https://github.com/bigsmoke/pg_text_semver/blob/master/README.md)
- `pg_text_semver` 在 PostgreSQL `text` 之上实现了语义化版本规范 2.0.0，使用 `semver` 域以及相关辅助类型和函数。

```sql
CREATE EXTENSION pg_text_semver;
```

## 核心流程

README 强调了以下能力：

- 使用常规比较运算符比较 `semver` 值
- 直接调用 `semver_cmp(semver, semver)`
- 使用 `semver_regexp()` 校验并检查版本字符串
- 将解析后的值转换为 `semver_parsed` 以便排序和建索引
- 使用 `semver_prerelease` 进行预发布版本校验和比较

## 示例

上游 README 会引导读者查看 `test__pg_text_semver()` 过程，以获取类型、运算符和函数的具体示例。它还说明该扩展提供独立的 `semver_parsed` 类型，可序列化回 `semver` 或 `text`。

## 备注

README 将本项目与基于 C 的 `semver` 扩展作了对比：`pg_text_semver` 始终基于 `text` 域实现，重点是一个简洁、贴近规范的实现。
