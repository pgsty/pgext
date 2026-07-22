## 用法

来源：

- [官方 v1.13 README](https://github.com/jimjonesbr/oai_fdw/blob/v1.13/README.md)
- [官方 v1.13 控制文件](https://github.com/jimjonesbr/oai_fdw/blob/v1.13/oai_fdw.control)
- [官方 v1.13 更新日志](https://github.com/jimjonesbr/oai_fdw/blob/v1.13/CHANGELOG.md)
- [官方 OAI-PMH 规范](https://www.openarchives.org/OAI/openarchivesprotocol.html)

`oai_fdw` 1.13 是一个通过 HTTP 或 HTTPS 从 OAI-PMH 2.0 仓库采集元数据的外部数据包装器。它把 OAI 记录头和 XML 内容映射到外部表列，支持按仓库或 set 自动生成表，也可以把采集记录复制到本地表。

### 核心流程

先为仓库创建 server，再显式定义外部表或使用 `IMPORT FOREIGN SCHEMA`。每张表都需要 `oai_dc` 等元数据格式。

```sql
CREATE EXTENSION oai_fdw;
CREATE SCHEMA oai_data;

CREATE SERVER museum_oai
FOREIGN DATA WRAPPER oai_fdw
OPTIONS (url 'https://example.org/oai');

CREATE FOREIGN TABLE oai_data.records (
    id text OPTIONS (oai_node 'identifier'),
    content xml OPTIONS (oai_node 'content'),
    sets text[] OPTIONS (oai_node 'setspec'),
    datestamp timestamp OPTIONS (oai_node 'datestamp'),
    format text OPTIONS (oai_node 'metadataprefix'),
    status boolean OPTIONS (oai_node 'status')
)
SERVER museum_oai
OPTIONS (metadataprefix 'oai_dc');

SELECT id, datestamp
FROM oai_data.records
LIMIT 100;
```

如果没有 `setspec`，一次扫描可能采集仓库暴露的全部 set。可视情况使用 `from`、`until` 或 `setspec` 表选项和外层 `LIMIT`，但必须用远程服务验证实际请求行为。

### 导入与支持接口

- `IMPORT FOREIGN SCHEMA oai_repository` 创建一张覆盖整个仓库的表；`oai_sets` 为每个 OAI set 创建一张表。两者都要求 `metadataprefix` 导入选项，并支持 `LIMIT TO` 或 `EXCEPT`。
- 列选项 `oai_node` 把 `identifier`、`setspec`、`datestamp`、`content`、`metadataprefix` 或 `status` 映射到兼容的 PostgreSQL 类型。
- Server 选项包括 `url`、`http_proxy`、`connect_timeout`、`connect_retry`、`request_redirect` 和 `request_max_redirect`。
- `OAI_HarvestTable` 把外部表采集结果复制到本地表；OAI 支持函数暴露 set、格式等仓库元数据。

1.13 版把 `proxy_user` 和 `proxy_password` 移到了 user mapping。仓库 `user` 与 `password` 也应放在那里，并使用 HTTP Basic Authentication。

### 网络、凭据与一致性

查询会同步发出远程请求并解析不可信 XML。延迟、分页、服务限流、重定向、畸形响应及仓库变化都会影响结果。应把凭据放在最小权限 user mapping 中，限制创建或修改 server 的权限，并把可配置 URL/代理视为具有 SSRF 影响的出站网络访问。

OAI-PMH 是采集协议，不是事务数据库 API。重复扫描可能看到变化中的页面和删除记录；本地副本也会过期。应记录采集窗口与标识符，使刷新可重启，并显式协调删除/状态记录。1.13 版需要匹配的 libcurl/libxml2 构建依赖，还应针对每个目标仓库的元数据格式进行测试。
