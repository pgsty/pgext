## 用法

来源：

- [1.0 版扩展 SQL](https://github.com/RekGRpth/pg_wthtmltopdf/blob/dc68969a4821945d7743cd4ab4b93258dd365251/pg_wthtmltopdf--1.0.sql)
- [PostgreSQL C 包装层](https://github.com/RekGRpth/pg_wthtmltopdf/blob/dc68969a4821945d7743cd4ab4b93258dd365251/pg_wthtmltopdf.c)
- [Wt PDF 渲染器实现](https://github.com/RekGRpth/pg_wthtmltopdf/blob/dc68969a4821945d7743cd4ab4b93258dd365251/MyWPdfRenderer.cc)
- [扩展 control 文件](https://github.com/RekGRpth/pg_wthtmltopdf/blob/dc68969a4821945d7743cd4ab4b93258dd365251/pg_wthtmltopdf.control)

`pg_wthtmltopdf` 公开一个基于 Wt `WPdfRenderer` 与 libharu 的服务端 HTML 转 PDF 函数。该函数简单且固定：它只按 A4 纵向页面配置渲染，没有用于页面尺寸、方向、边距、字体、页眉或页脚的 SQL 选项。

### 核心流程

服务器共享库使用兼容的 Wt 与 libharu 版本构建后，创建扩展，并把 HTML 文本传给 `wthtmltopdf(text)`。

```sql
CREATE EXTENSION pg_wthtmltopdf;

SELECT wthtmltopdf(
  '<html><body><h1>Invoice</h1><p>Total: 42</p></body></html>'
);
```

渲染器设置 A4 纵向输出、压缩、UTF 编码、1.0 的边距值和 96 DPI。渲染在调用它的 PostgreSQL 后端中同步执行。

### 重要对象

- `wthtmltopdf(text) RETURNS text`：渲染所给 HTML 并返回生成的 PDF 字节。

### 二进制类型警告

虽然 SQL 签名声明为 `text`，C 函数却把原始 PDF 字节直接写进 PostgreSQL text datum。PDF 是二进制数据，可能包含数据库编码中的无效字节。因此普通文本客户端、类型转换、字符串函数、逻辑解码、转储或驱动可能拒绝、改写或损坏结果。这并不等价于安全的 `bytea` API。

采用前，必须以二进制结果模式测试实际客户端，并逐字节验证生成文件。客户端若不能安全取出原始负载，就不要使用该函数；应将扩展包装或修补为返回 `bytea`，或改在 PostgreSQL 外部渲染。

### 故障与安全说明

NULL 输入会报错。异常 HTML 或 Wt/libharu 故障可能引发 PostgreSQL 错误，某些渲染失败则返回 NULL。大型或恶意 HTML 会消耗数据库后端的 CPU 与内存，因此应限制 EXECUTE 权限、约束语句时长，并避免不可信输入。扩展的 SQL 接口没有 URL 获取、样式表、字体或沙箱策略；实际行为取决于链接的 Wt 软件栈。上游 README 为空，因此本次复核的 SQL 与源代码才是 `1.0` 版的权威接口。
