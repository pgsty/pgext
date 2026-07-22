## 用法

来源：

- [官方 README](https://github.com/RekGRpth/pg_mupdf/blob/a15db6ca07c5e95b8f2d78fe92774c659a3f3c54/README.md)
- [官方扩展 SQL](https://github.com/RekGRpth/pg_mupdf/blob/a15db6ca07c5e95b8f2d78fe92774c659a3f3c54/pg_mupdf--1.0.sql)
- [官方 C 实现](https://github.com/RekGRpth/pg_mupdf/blob/a15db6ca07c5e95b8f2d78fe92774c659a3f3c54/pg_mupdf.c)

`pg_mupdf` 在 PostgreSQL 后端内运行 MuPDF，把文档文本（尤其是 HTML）转换成 PDF 等二进制输出。它只提供一个同步转换函数；获取远程 HTML 和保存返回字节属于独立的应用职责。

### 核心流程

```sql
CREATE EXTENSION pg_mupdf;

WITH rendered AS (
  SELECT mupdf(
    '<html><body><h1>Quarterly report</h1></body></html>',
    input_type := 'html',
    output_type := 'pdf',
    options := 'compress',
    range := '1-N'
  ) AS pdf
)
SELECT octet_length(pdf), left(encode(pdf, 'hex'), 8) AS signature
FROM rendered;
```

应通过二进制安全的数据库驱动取回结果 `bytea`，再把这些字节写成输出文件。PostgreSQL `COPY ... FORMAT binary` 会添加 PostgreSQL 二进制 COPY 封装，本身不能生成原始 PDF 文件。

### 函数参考

`mupdf(data text, input_type text DEFAULT 'html', output_type text DEFAULT 'pdf', options text DEFAULT '', range text DEFAULT '1-N')` 通过 MuPDF 已注册的文档处理器打开输入，渲染指定页面范围，并把写入器输出作为 `bytea` 返回。可接受的类型名、写入选项和页面范围语法取决于所链接的 MuPDF 构建，因此应针对打包库进行验证。

### 运维边界

转换在调用方后端中运行，而且实现会创建无限容量的 MuPDF 存储。必须在函数外严格限制输入大小、页数、输出大小、超时与并发度；否则复杂或恶意内容可能耗尽数据库工作进程的 CPU 或内存。应持续修补 MuPDF，因为数据库进程会在自身地址空间内解析攻击者控制的文档格式。所有参数都应传入非 NULL 值，测试错误清理和实际链接的 MuPDF 版本；在审查文档信任、HTML 资源行为、字体、输出保真度与故障隔离前，不要广泛授予函数权限。
