
## 用法

> [README](https://github.com/jimjonesbr/nominatim_fdw) | [Nominatim API](https://nominatim.org/)

`nominatim_fdw` 是一个 PostgreSQL 风格的 FDW 扩展，用于从 SQL 中调用 Nominatim 地理编码服务。该扩展围绕函数而不是外部表组织，并映射到 Nominatim 的 `search`、`reverse` 和 `lookup` 端点。

## 服务器配置

先创建扩展，再定义一个指向 Nominatim 端点的服务器：

```sql
CREATE EXTENSION nominatim_fdw;

CREATE SERVER osm
FOREIGN DATA WRAPPER nominatim_fdw
OPTIONS (url 'https://nominatim.openstreetmap.org');
```

README 记录的服务器选项包括：

- `url`，必填的端点 URL
- `http_proxy`
- `connect_timeout`，默认 `300`
- `max_connect_retry`，默认 `3`
- `max_connect_redirect`，其中 `0` 表示无限重定向

服务器选项可通过 `ALTER SERVER` 修改：

```sql
ALTER SERVER osm OPTIONS (ADD max_connect_retry '5');
ALTER SERVER osm OPTIONS (SET url 'https://a.new.url');
ALTER SERVER osm OPTIONS (DROP http_proxy);
```

代理凭据应放在用户映射中，而不是服务器定义中：

```sql
CREATE USER MAPPING FOR pguser
SERVER osm
OPTIONS (proxy_user 'myuser', proxy_password 'mysecret');
```

## 地理编码函数

### 搜索

`nominatim_search` 同时支持结构化查询和自由格式查询：

```sql
SELECT osm_id, ref, lon, lat, boundingbox
FROM nominatim_search(
  server_name => 'osm',
  street => 'Neubrueckenstrasse 63',
  city => 'Muenster',
  country => 'Germany'
);

SELECT osm_id, display_name, lon, lat
FROM nominatim_search(
  server_name => 'osm',
  q => '1600 Pennsylvania Avenue, Washington DC'
);
```

### 反向地理编码

```sql
SELECT osm_id, display_name, boundingbox
FROM nominatim_reverse(
  server_name => 'osm',
  lon => -77.0365,
  lat => 38.8977,
  zoom => 18,
  addressdetails => true
);
```

### OSM 对象查找

```sql
SELECT osm_id, display_name
FROM nominatim_lookup(
  server_name => 'osm',
  osm_ids => 'W121736959,R123456'
);
```

README 说明查找 ID 使用 OSM 类型前缀，例如 `N` 表示 node、`W` 表示 way、`R` 表示 relation。

## 说明

当前上游 README 列出的依赖要求包括：

- PostgreSQL 12 或更高版本
- `libxml2` 2.5.0 或更高版本
- `libcurl` 7.74.0 或更高版本

该扩展还暴露 `nominatim_fdw_version()` 用于版本检查，并支持通过 `ALTER EXTENSION nominatim_fdw UPDATE` 进行升级。
