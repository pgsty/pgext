## 用法

来源：

- [官方 README](https://github.com/lacanoid/pgwiki/blob/dbcc433acd6fb2dacc78d0621103a8e2d5c187ec/README.md)
- [官方扩展 SQL](https://github.com/lacanoid/pgwiki/blob/dbcc433acd6fb2dacc78d0621103a8e2d5c187ec/wiki.sql)
- [官方扩展控制文件](https://github.com/lacanoid/pgwiki/blob/dbcc433acd6fb2dacc78d0621103a8e2d5c187ec/wiki.control)

`wiki` 在 PostgreSQL 内实现带版本历史的 Wiki 数据模型与 SQL API，并可通过 PL/Perl 辅助函数访问外部 MediaWiki 实例。复核的 `0.1` 源码已经停止维护，且面向旧版 PostgreSQL；必须完成安全与兼容性审计后才能使用。

### 前置条件与设置

控制文件将扩展固定在 `wiki` 模式，依赖 `plperl` 与不受信任的 `plperlu`，并要求由超级用户安装。其 SQL 转储还会把对象所有者设为名为 `wiki` 的角色；安装前应检查脚本，并创建或映射预期的所有者。

```sql
CREATE EXTENSION plperl;
CREATE EXTENSION plperlu;
CREATE EXTENSION wiki;
```

由于 `plperlu` 可以访问宿主操作系统，只应由可信任的数据库管理员安装该扩展或授予其访问权限。

### 核心流程

使用 `wiki.page_set` 创建或替换页面，然后读取其当前正文或结构化元数据：

```sql
SELECT wiki.page_set('Welcome', 'Hello from PostgreSQL');

SELECT wiki.page_get('Welcome');
SELECT wiki.page_get_data('Welcome');
```

每次修改正文都会创建一个修订版本。SQL 还提供接收文本、XML 与 JSON 正文的重载，以及按节点或修订标识符读取页面的函数。

### 重要对象

- `wiki.page_set(...)` 写入页面正文并返回新的修订标识符。
- `wiki.page_get(...)` 返回当前或指定版本的页面正文。
- `wiki.page_get_data(...)` 以 JSON 返回页面元数据。
- `wiki.node`、`wiki.body`、`wiki.revision` 与 `wiki.links` 保存页面、内容修订和链接。
- `wiki.foreign_page_get(...)` 与 `wiki.foreign_pages_search(...)` 使用 `plperlu` 和 Perl `MediaWiki::API` 模块访问已配置的远端 Wiki。
- `wiki.html_tidy(...)` 依赖 Perl `HTML::Tidy` 模块。

### 运维注意事项

远端 Wiki 凭据保存在扩展的 `wiki.wiki` 表中，而不受信任的 Perl 函数可以发起网络与宿主级调用。需要检查表和函数授权、撤销过宽的默认权限、保护已存凭据，并限制出站网络访问。上游仓库最后更新于 2019 年，文档化兼容范围止于 PostgreSQL `12`，并未证明支持当前 PostgreSQL 版本。应固定精确源码、测试安装/升级/恢复流程；新部署更适合选择仍在维护的 Wiki 或应用层集成方案。
