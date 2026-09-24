## 用法

来源：

- [README](https://api.pgxn.org/src/pg_htmldoc/pg_htmldoc-1.0.11/README.md)
- [Control file / 控制文件](https://api.pgxn.org/src/pg_htmldoc/pg_htmldoc-1.0.11/pg_htmldoc.control)
- [SQL](https://api.pgxn.org/src/pg_htmldoc/pg_htmldoc-1.0.11/pg_htmldoc--1.0.sql)

`pg_htmldoc` 内嵌 HTMLDOC，将排队的 HTML、文件或 URL 转换为 PDF 或 PostScript。发行版 1.0.11 保留控制版本 1.0，但权限和会话状态行为有重要变化。

### 核心工作流

下面的内存输入示例需要超级用户，因为标记内容仍可能引用服务端文件或 URL。无参数转换函数向客户端返回二进制数据。

```sql
CREATE EXTENSION pg_htmldoc;
SELECT htmldoc_addhtml('<h1>Quarterly report</h1><p>Complete.</p>');
SELECT octet_length(convert2pdf());
```

### 对象与状态

`htmldoc_addfile`、`htmldoc_addurl` 和 `htmldoc_addhtml` 向同一个后端内文档追加输入。`convert2pdf` 和 `convert2ps` 返回 `bytea`，传入文件名时则写入服务端路径。转换成功后会清空文档，再次转换前必须重新添加输入。空输入或无文档时调用转换会失败。

### 访问边界

`pg_htmldoc.whitelist` 是由超级用户控制的文件和 HTTP(S) 前缀列表。普通角色只能读取显式匹配的输入文件或 URL，空列表拒绝此类访问；对超级用户而言，非空列表缩小其可访问范围。白名单不能允许普通角色调用 `htmldoc_addhtml` 或文件输出重载，这些接口始终要求超级用户。返回渲染后的 `bytea` 本身不会写文件。

扩展可重定位，无需预加载，但需要 HTMLDOC 共享库。应限制原生解析器输入和外部资源访问，并注意连接池中多次添加输入的行为。SQL 函数授权与白名单用途不同，二者都需要正确配置。
