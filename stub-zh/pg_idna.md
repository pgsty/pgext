## 用法

来源：

- [官方 README](https://github.com/owenthewizard/pg_idna/blob/cbb25acf09ba478c5ca828d722519bf5292334e1/README.md)
- [0.1.1 版变更记录](https://github.com/owenthewizard/pg_idna/blob/cbb25acf09ba478c5ca828d722519bf5292334e1/CHANGELOG.md)
- [SQL 可见的 Rust 实现](https://github.com/owenthewizard/pg_idna/blob/cbb25acf09ba478c5ca828d722519bf5292334e1/src/lib.rs)
- [可接受的转换选项](https://github.com/owenthewizard/pg_idna/blob/cbb25acf09ba478c5ca828d722519bf5292334e1/src/my_config.rs)

`pg_idna` 使用 Rust `idna` crate 的 UTS #46 实现，为 PostgreSQL 提供 Unicode IDNA 转换辅助函数。它可以在 Unicode 与 ASCII/Punycode 之间转换国际化域名标签，并提供轻量级分类函数。

### 核心流程

```sql
CREATE EXTENSION pg_idna;

SELECT idna_to_ascii('☕.us');
SELECT idna_to_unicode('xn--53h.us');
SELECT idna_is_ascii('example.org');
```

默认转换策略使用 ASCII 拒绝列表 `url`、连字符策略 `allow`，以及用于 ASCII 转换的 DNS 长度策略 `verify`。

### 函数与参数

- `idna_is_ascii(input text) RETURNS boolean` 判断每个字节是否都是 ASCII。
- `idna_is_punycode(input bytea) RETURNS boolean` 只是一个不区分大小写的 `xn--` 前缀提示；它不会校验标签其余部分。
- `idna_to_ascii(input, adl, h, dl)` 转换为 ASCII。`adl` 接受 `empty`、`std3` 或 `url`；`h` 接受 `allow`、`check_first_last` 或 `check`；`dl` 接受 `ignore`、`verify_allow_root_dot` 或 `verify`。
- `idna_to_unicode(input, adl, h)` 拒绝失败的转换。
- `idna_to_unicode_lossy(input, adl, h)` 在遇到无效 Punycode 时用 U+FFFD 替换字符返回输出。

可选参数都有默认值，因此普通 URL 场景只需输入值。验证 DNS 主机名而不是一般 URL 标签时，应显式指定策略。

### 注意事项

上游将项目标为 pre-alpha，明确要求不要用于生产，也不承诺稳定的版本方案。目录版本 `0.1.1-docs` 对应基于 0.1.1 发行版并附加文档修改的源码。

无效选项字符串、约束违规、畸形 Punycode 或转换失败会使 Rust 函数 panic，pgrx 会将其表现为 PostgreSQL 错误。绝不能使用 `idna_is_punycode` 作为校验。IDNA 转换也不能证明主机名存在、访问安全或视觉上无歧义；应用层仍需单独实施允许列表和易混淆名称策略。
