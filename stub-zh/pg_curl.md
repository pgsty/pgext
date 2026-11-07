## 用法

来源：

- [Official source](https://api.pgxn.org/src/pg_curl/pg_curl-2.5.0/README.md)
- [Official source](https://api.pgxn.org/src/pg_curl/pg_curl-2.5.0/pg_curl.control)
- [Official source](https://api.pgxn.org/src/pg_curl/pg_curl-2.5.0/pg_curl--2.5.sql)
- [Official source](https://api.pgxn.org/src/pg_curl/pg_curl-2.5.0/pg_curl--2.4--2.5.sql)
- [Official source](https://api.pgxn.org/dist/pg_curl/2.5.0/META.json)

`pg_curl` 将 libcurl 传输操作暴露为 SQL 函数。上游发行版 2.5.0 使用 SQL 扩展版本 2.5。请求在数据库后端执行，使用服务器的网络与文件系统权限；安装扩展需要管理员。普通调用无需预加载。

### 在同一事务内执行请求

```sql
CREATE EXTENSION pg_curl;
BEGIN;
SELECT curl_easy_reset();
SELECT curl_easy_setopt_url('https://www.postgresql.org/');
SELECT curl_easy_perform();
SELECT convert_from(curl_easy_getinfo_data_in(), 'UTF8');
COMMIT;
```

`curl_easy_reset` 清理句柄，`curl_easy_setopt_url` 选择目标，`curl_easy_perform` 执行传输。`curl_easy_getinfo_data_in` 返回响应字节，`curl_easy_getinfo_header_in` 返回响应头。`pg_curl.transaction` 默认为 true，句柄会在事务结束时释放，因此多语句请求应放在同一个事务中。SQL 回滚不会撤销网络或文件操作。

### 访问权限与白名单

超级用户可为已有应用角色指定允许的前缀：

```sql
ALTER ROLE app_user SET pg_curl.whitelist = 'https://api.example.com/';
```

重新连接后角色默认值生效。未设置 `pg_curl.whitelist` 时，超级用户不受限制，其他角色默认拒绝访问；设置后也会限制超级用户。白名单生效时只允许匹配的 HTTP、HTTPS 和文件 URL，FTP、SMTP 等其他协议均被拒绝。HTTP 路径按路径段边界匹配，而非任意字符串前缀。

本地文件选项也会检查权限。路径解析后必须已存在，包括用于写出的 cookie-jar 文件。重定向目标会逐个检查，受限情况下跟随重定向要求 libcurl 7.80.0 及以上。受限时禁止显式设置代理和 DNS-over-HTTPS，但服务器环境中的代理仍可能生效。应谨慎限定允许的主机和路径，并保护存储的凭据及服务器本地文件。

### 其他操作与升级

`curl_postfield_append`、`curl_easy_setopt_postfields`、`curl_header_append` 用于构造表单或 JSON 请求。`curl_mime_data` 与 `curl_mime_file` 构造 MIME 部分，2.5 新增 `head text[]` 重载以传入多个头字段。头数组元素不得为 NULL，也不得包含回车或换行。MIME 数据重载接受文本或二进制数据。

升级前检查 `pg_extension.extversion`。本发行包包含 2.4 到 2.5 的升级脚本，没有提供 2.4.1 到 2.5 的路径，因此不能假定报告 2.4.1 的已有安装可自动升级。对于实际 SQL 版本为 2.4 的安装，先安装匹配文件，再执行：

```sql
ALTER EXTENSION pg_curl UPDATE TO '2.5';
```

向应用角色授予执行权限前，应先落实新的白名单规则。确认所用 libcurl 支持的协议，并在各数据库主机使用匹配的库文件。
