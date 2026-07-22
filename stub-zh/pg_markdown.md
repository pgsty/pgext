## 用法

来源：

- [官方 README](https://github.com/sycobuny/pg_markdown/blob/af587ec2e20f87642b6d1e96536036b39c0030da/README.md)
- [官方扩展控制文件](https://github.com/sycobuny/pg_markdown/blob/af587ec2e20f87642b6d1e96536036b39c0030da/pg_markdown.control)

`pg_markdown` 增加 `markdown` 数据类型，并在 PostgreSQL 内把已存储的 Markdown 转换成 HTML 或纯文本。它可以支持小型、以数据库为中心的发布流程，但当前审阅的代码和文档源自项目早期的 2011 年实现，使用前必须验证兼容性与安全性。

### 核心流程

以扩展类型保存 Markdown，并只在展示边界进行渲染：

```sql
CREATE EXTENSION pg_markdown;

CREATE TABLE blog_posts (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  body markdown NOT NULL
);

INSERT INTO blog_posts(body)
VALUES ('## Welcome

This is **Markdown** with a [link](https://www.postgresql.org/).');

SELECT to_html(body) FROM blog_posts WHERE id = 1;
SELECT to_text(body) FROM blog_posts WHERE id = 1;
```

### SQL 接口

- `markdown` 是扩展提供的存储类型。
- `to_html(markdown)` 以 `text` 返回渲染后的 HTML。
- `to_text(markdown)` 以 `text` 返回纯文本表示。

### 安全与兼容性

- 上游 README 只承诺 Markdown 转换，并未承诺 HTML 清理。应把 `to_html(markdown)` 输出视为不可信 HTML，并针对目标展示环境进行清理，尤其是作者可以提交原始 HTML、链接或图片时。
- 应固定并测试所接受的 Markdown 方言。这个旧版本没有说明现代 Markdown 变体、Unicode 边界情况、非法输入和不同渲染器之间的差异。
- 在大范围扫描中即时渲染会消耗数据库 CPU。可以把源 Markdown 作为权威数据，并在热点查询路径之外缓存已清理的输出。
- 升级前，应在目标 PostgreSQL 版本上测试类型二进制与文本 I/O、转储恢复、类型转换和已有数据；官方 README 没有声明支持当前 PostgreSQL 版本。
