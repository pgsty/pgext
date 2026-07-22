## 用法

来源：

- [官方 README](https://github.com/RekGRpth/pg_htmldoc/blob/74259576cabbdaa2234d95a300a3c128637f665e/README.md)
- [扩展控制文件](https://github.com/RekGRpth/pg_htmldoc/blob/74259576cabbdaa2234d95a300a3c128637f665e/pg_htmldoc.control)
- [1.0 版扩展 SQL](https://github.com/RekGRpth/pg_htmldoc/blob/74259576cabbdaa2234d95a300a3c128637f665e/pg_htmldoc--1.0.sql)
- [C 实现](https://github.com/RekGRpth/pg_htmldoc/blob/74259576cabbdaa2234d95a300a3c128637f665e/pg_htmldoc.c)

`pg_htmldoc` 1.0 在 PostgreSQL 后端中嵌入 HTMLDOC。它把 HTML、服务器端文件或 URL 累积到后端本地文档中，再将文档渲染为 PDF 或 PostScript；扩展没有暴露 HTMLDOC 的 EPUB 或 HTML 输出模式。

### 核心流程

权限需求最低的输入路径是内存 HTML 字符串和 `bytea` 结果：

```sql
CREATE EXTENSION pg_htmldoc;

SELECT htmldoc_addhtml('<h1>Quarterly report</h1><p>Complete.</p>');
SELECT octet_length(convert2pdf());
```

转换前多次调用 `htmldoc_add*` 会把来源附加到同一文档。要取得渲染结果，可无参数调用 `convert2pdf()` 或 `convert2ps()`，并在客户端消费返回的 `bytea`。

### 函数

- `htmldoc_addhtml(html text) RETURNS boolean` 添加内存 HTML。
- `htmldoc_addfile(file text) RETURNS boolean` 读取数据库服务器操作系统用户可见的路径。
- `htmldoc_addurl(url text) RETURNS boolean` 从数据库服务器获取 URL。
- `convert2pdf()` 和 `convert2ps()` 以 `bytea` 返回渲染文档。
- `convert2pdf(file text)` 和 `convert2ps(file text)` 写入服务器端文件并返回 `boolean`。

### 安全与状态边界

文件重载不会访问 SQL 客户端的文件系统。输入和输出路径由 PostgreSQL 后端打开，URL 获取也使用后端的网络可达性。应撤销不受信任用户的执行权，以防服务器端文件泄露、在服务账号权限内创建任意文件，以及服务器端请求伪造：

```sql
REVOKE EXECUTE ON FUNCTION htmldoc_addfile(text) FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION htmldoc_addurl(text) FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION convert2pdf(text) FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION convert2ps(text) FROM PUBLIC;
```

C 实现将文档保存在当前 PostgreSQL 后端的进程级全局状态中，转换后不会立即显式清空。复用连接池会话因此可能携带或累积此前的文档状态。每个转换任务应使用受控的专用会话，并在依赖重复调用前验证隔离行为。

HTML、引用资源、URL 和输出大小都是不受信任的原生解析器输入，会消耗后端内存、CPU、网络和磁盘。应在扩展外配置语句超时、端点允许列表、输入大小限制和角色约束。
