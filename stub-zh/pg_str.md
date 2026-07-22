## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/publicmatt/pg_str/blob/5c49c2397a987feb6bcdc248a53abaf4297b7e36/README.md)
- [导出的 Rust 模块](https://github.com/publicmatt/pg_str/blob/5c49c2397a987feb6bcdc248a53abaf4297b7e36/src/lib.rs)
- [大小写转换实现](https://github.com/publicmatt/pg_str/blob/5c49c2397a987feb6bcdc248a53abaf4297b7e36/src/modules/case.rs)
- [Markdown 实现](https://github.com/publicmatt/pg_str/blob/5c49c2397a987feb6bcdc248a53abaf4297b7e36/src/modules/markdown.rs)
- [Cargo 元数据](https://github.com/publicmatt/pg_str/blob/5c49c2397a987feb6bcdc248a53abaf4297b7e36/Cargo.toml)

`pg_str` 是用 Rust/pgrx 编写的一组 Laravel 风格字符串辅助函数。版本 `0.3.1` 把函数安装到 `public` schema 中，并且大多使用 `str_` 前缀，涵盖大小写转换、查找、切片、转写、词形变化、填充、Markdown 渲染、随机字符串和 UUID。它适合应用风格的便捷文本处理，但不能替代 PostgreSQL 能感知排序规则的操作符或全文搜索。

### 核心工作流

```sql
CREATE EXTENSION pg_str;

SELECT str_snake('Hello PostgreSQL');
SELECT str_contains_all(
  'postgres extensions',
  ARRAY['postgres', 'ext']
);
SELECT str_uuid();
```

### 函数组

- 位置与切片：`str_after`、`str_before`、`str_before_last`、`str_start`、`str_substr`、`str_split`、`str_length`、`str_replace` 和 `str_append`。
- 大小写与词形：`str_lower`、`str_upper`、`str_title`、`str_camel`、`str_kebab`、`str_snake`、`str_studly`、`str_scream`、`str_slug`、`str_singular` 和 `str_plural`。
- 测试与转换：`str_contains`、`str_contains_all`、`str_is_ascii` 和 `str_ascii`。
- 生成与展示：`str_random`、`str_uuid`、`str_markdown`、`pad_left`、`pad_right` 和 `pad_both`。

README 检查清单混合了已实现和规划中的名称。当前 Rust 模块才是权威依据；依赖一个函数前，应确认它是导出的 `#[pg_extern]` 函数并核对生成的 SQL 签名。

### 行为与兼容性说明

- `str_markdown` 把 Markdown 转换为 HTML 并启用删除线，但它不是 HTML 清理器。应在渲染边界把不可信输入视为活动内容，并按应用策略进行清理。
- 大小写转换、词形变化、转写和 slug 生成遵循 Rust 库行为，而不是数据库排序规则与区域设置。应测试多语言和领域专用词汇。
- `str_random` 会把有符号长度转换为分配大小；调用前必须拒绝负数或不合理的大值。未经所需安全属性评估，不要假定生成结果可作为密码或令牌。
- 某些填充路径混合使用字符计数和面向字节的字符串操作。应测试多字节填充字符串和长度；Unicode 行为关键时优先使用 PostgreSQL 核心函数。
- 版本 `0.3.1` 固定使用 pgrx `0.11.2`，并声明 PostgreSQL 11 到 15 的构建特性。为一个 PostgreSQL 大版本构建的二进制不能移植到另一个版本；尤其对更新版本，必须针对确切服务器大版本重新构建并测试。
