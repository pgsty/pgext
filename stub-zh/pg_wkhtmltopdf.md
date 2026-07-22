## 用法

来源：

- [官方 README](https://github.com/RekGRpth/pg_wkhtmltopdf/blob/476e17cb2b4732f88b3c62b2c4de88948dff6a94/README.md)
- [官方扩展 SQL](https://github.com/RekGRpth/pg_wkhtmltopdf/blob/476e17cb2b4732f88b3c62b2c4de88948dff6a94/pg_wkhtmltopdf--1.0.sql)
- [官方 C 实现](https://github.com/RekGRpth/pg_wkhtmltopdf/blob/476e17cb2b4732f88b3c62b2c4de88948dff6a94/pg_wkhtmltopdf.c)

`pg_wkhtmltopdf` 在 PostgreSQL 后端内运行 libwkhtmltox/QtWebKit HTML-to-PDF 引擎，并以 `bytea` 返回 PDF。它只适合必须在数据库进程中渲染的场景；外部渲染服务通常能提供更安全的资源与网络边界。

### 核心流程

先设置输入页面，再在同一会话中渲染：

```sql
CREATE EXTENSION pg_wkhtmltopdf;

SELECT wkhtmltopdf_set_object_setting(
  'page',
  'https://example.com/invoices/123'
);

SELECT wkhtmltopdf_convert();
```

客户端必须把返回的 `bytea` 提取为原始字节。上游的 `COPY ... FORMAT binary` 示例写出的是 PostgreSQL 二进制 COPY 封装，而不是纯 PDF 文件；不能把该文件直接作为 PDF 提供。

### 函数与状态

- `wkhtmltopdf_set_object_setting(text,text)` 设置页面 URL 等对象选项。
- `wkhtmltopdf_set_global_setting(text,text)` 设置 libwkhtmltox 全局选项。
- `wkhtmltopdf_convert()` 同步执行转换并返回 `bytea`。

设置对象保存在后端静态状态中，扩展没有暴露 SQL 重置函数。应在受控的同一会话中完成设置与转换，不要假定事务回滚会恢复设置，并仔细测试重复调用。否则会话池复用可能把配置带给后续请求。

### 安全与运维

渲染会在数据库后端内执行网络获取、HTML 解析、JavaScript 和 PDF 生成。缓慢或恶意页面可能消耗后端 CPU/内存、阻塞工作进程、访问内部 HTTP 服务，或触发本地文件与旧浏览器引擎能力。应限制三个函数的 `EXECUTE` 权限，在 SQL 之外设置目标白名单，并在操作系统边界实施网络与本地文件控制。

扩展没有提供数据库级超时、取消策略、队列或渲染器隔离。调用应配合严格的语句超时和工作负载限制，但不能把这些措施等同于进程隔离。

构建链接 wkhtmltopdf 已归档的 QtWebKit C 库，因此库 ABI、安全修复、字体、证书与渲染行为都取决于主机。版本 `1.0` 没有文档化的 PostgreSQL 兼容矩阵。部署前应在确切服务端镜像上构建并运行有代表性及对抗性页面。
