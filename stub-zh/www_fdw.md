## 用法

来源：

- [官方 README](https://github.com/cyga/www_fdw/blob/2445772dddb44be8534527782deb8230f84015fe/README.md)
- [官方控制文件](https://github.com/cyga/www_fdw/blob/2445772dddb44be8534527782deb8230f84015fe/www_fdw.control)
- [官方安装 SQL](https://github.com/cyga/www_fdw/blob/2445772dddb44be8534527782deb8230f84015fe/sql/www_fdw.sql)
- [官方文档](https://github.com/cyga/www_fdw/wiki/Documentation)
- [官方示例](https://github.com/cyga/www_fdw/wiki/Examples)

`www_fdw` 是已废弃的外部数据包装器：它把 `SELECT` 转换成 HTTP 请求，再将 JSON、XML 或回调解码的响应转换成行。它面向早期 PostgreSQL FDW API，应视作历史集成代码，而不是当前服务器上仍受维护的 HTTP 客户端。

### 核心流程

创建扩展、包含服务基础 URI 的服务器、用户映射，以及一个外部表；外部表的列既可以作为请求参数，也可以作为响应字段。

```sql
CREATE EXTENSION www_fdw;

CREATE SERVER web_service
  FOREIGN DATA WRAPPER www_fdw
  OPTIONS (
    uri 'https://api.example.invalid/items',
    response_type 'json'
  );

CREATE USER MAPPING FOR current_user
  SERVER web_service;

CREATE FOREIGN TABLE web_items (
  category text,
  item_id text,
  title text
) SERVER web_service;

SELECT item_id, title
FROM web_items
WHERE category = 'database';
```

使用默认序列化器时，等值谓词会变成 URL 键值参数；其他运算符需要自定义请求回调。自动 JSON 或 XML 解码器会按广度优先顺序寻找第一个结构一致的数组，并把字段映射到外部表列。

### 服务器选项与回调

- `uri` 是基础端点；`uri_select` 与 `method_select` 自定义读取路径和 HTTP 方法。
- `request_user_agent` 与 `request_user_header` 添加请求头。
- `request_serialize_callback` 接收序列化的查询条件，可以构造自定义 URL 或请求体。
- `request_serialize_type` 支持 `log`、`null`、`json` 或 `xml` 表示。
- `response_type` 支持 `json`、`xml` 或 `other`。
- `response_deserialize_callback` 解析响应；`response_iterate_callback` 可以转换每个返回元组。
- `ssl_cert`、`ssl_key`、`cainfo` 与 `proxy` 把连接设置传给 libcurl。

虽然存在插入、更新和删除选项名，但官方文档明确说明未实现 FDW 写操作。

### 安全与兼容性

查询会让数据库服务器发出网络请求。应限制 FDW 与服务器的所有权，验证每个端点和代理，并防止不可信角色把数据库变成访问内部服务的 SSRF 通道。响应数据、请求头、证书路径和回调函数都应视作不可信集成输入。

`0.1.9` 发布版记录了 PostgreSQL 9.2 支持，并说明主分支的 PostgreSQL 9.5 支持未经测试。项目已不再维护，依赖旧式 libcurl/libjson/libxml 集成，规划能力也很有限。评估时应设置超时并实施外部隔离；在受支持的 PostgreSQL 版本上应优先采用仍受维护的服务摄取方案。
