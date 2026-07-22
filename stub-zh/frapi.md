## 用法

来源：

- [官方 README](https://github.com/adauhr/pg_frapi/blob/f6ac6decaf54558ed34beda4f35b9f027cb55835/README.md)
- [官方扩展 SQL](https://github.com/adauhr/pg_frapi/blob/f6ac6decaf54558ed34beda4f35b9f027cb55835/frapi--0.1.2.sql)
- [官方扩展控制文件](https://github.com/adauhr/pg_frapi/blob/f6ac6decaf54558ed34beda4f35b9f027cb55835/frapi.control)

`frapi` 是法国国家地址 API `api-adresse.data.gouv.fr` 的 alpha 阶段 SQL 包装器。版本 `0.1.2` 可将地址搜索与反向地理编码响应转换为带 PostGIS 点的行，但它通过不受信任的 `plsh` 代码执行网络请求，只适合隔离环境中的可信实验。

### 核心流程

扩展固定安装到 `frapi` 模式，并依赖 `plsh` 与 `postgis`：

```sql
CREATE EXTENSION plsh;
CREATE EXTENSION postgis;
CREATE EXTENSION frapi;

SELECT *
FROM frapi.adresse_search('8 boulevard du Port', NULL, NULL, NULL, 5);

SELECT *
FROM frapi.adresse_reverse(2.37, 48.85, 5);
```

`adresse_search_json(...)` 和 `adresse_reverse_json(...)` 返回原始 JSONB 响应。`adresse_search_format(jsonb)` 将响应映射为 `adresse_search` 复合类型，其中包含 `geometry(Point,4326)` 值。便捷函数 `adresse_search(...)` 与 `adresse_reverse(...)` 会组合这些步骤。

### 安全与服务边界

`get_url(text, numeric, numeric, integer)` 使用 `sleep` 与 `wget` 拼接 shell 命令。SQL 还会在没有完整 URL 编码的情况下连接查询参数。不得传入不受信任的字符串：shell 元字符或畸形参数可能改变命令行为，而且上游 README 明确说明 PL/sh 路径尚未接受安全审计。

结果依赖外部 DNS、TLS、API 可用性、速率限制和服务商响应模式。上游文档仅报告在 Ubuntu 14.04 与 PostgreSQL 9.5 上测试。任何非实验室用途之前，都应限制执行权限、施加网络控制，并验证当前平台兼容性与服务条款。
