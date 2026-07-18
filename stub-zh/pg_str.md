## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/publicmatt/pg_str/blob/5c49c2397a987feb6bcdc248a53abaf4297b7e36/README.md)
- [导出的 Rust 函数](https://github.com/publicmatt/pg_str/blob/5c49c2397a987feb6bcdc248a53abaf4297b7e36/src/lib.rs)
- [大小写转换模块](https://github.com/publicmatt/pg_str/blob/5c49c2397a987feb6bcdc248a53abaf4297b7e36/src/modules/case.rs)
- [Cargo 元数据](https://github.com/publicmatt/pg_str/blob/5c49c2397a987feb6bcdc248a53abaf4297b7e36/Cargo.toml)

`pg_str` 是一组由 Rust/pgrx 编写、受 Laravel 启发的字符串辅助函数。函数以 `str_` 前缀安装到 `public`，涵盖大小写转换、包含判断、切片、Markdown 渲染、随机字符串和 UUID 生成。

```sql
CREATE EXTENSION pg_str;
SELECT str_snake('Hello PostgreSQL');
SELECT str_contains_all('postgres extensions', ARRAY['postgres', 'ext']);
SELECT str_uuid();
```

README 检查表是最完整的 API 清单，但同时列出已实现与计划中的函数。依赖某个函数前，应在当前 Rust 模块中确认其存在。Markdown 渲染会生成 HTML，因此必须依据消费应用的信任模型对结果进行转义或净化。

0.3.1 使用较旧的 pgrx 版本，并声明 PostgreSQL 11–15 特性。GitHub 仓库是上游 Gitea 项目的镜像；在较新 PostgreSQL 上应实际验证构建，不能假定 ABI 兼容。
