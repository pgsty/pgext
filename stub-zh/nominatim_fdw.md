

## 用法

> [nominatim_fdw: PostgreSQL 的 Nominatim 外部数据包装器](https://github.com/jimjonesbr/nominatim_fdw)

该 FDW 提供从 PostgreSQL 直接通过 SQL 函数访问 [Nominatim](https://nominatim.org/) 地理编码服务的能力，而非使用传统外部表。

### 创建服务器

```sql
CREATE EXTENSION nominatim_fdw;

CREATE SERVER osm FOREIGN DATA WRAPPER nominatim_fdw
  OPTIONS (url 'https://nominatim.openstreetmap.org');
```

**服务器选项：** `url`（必填，Nominatim 端点 URL）、`http_proxy`、`proxy_user`、`proxy_user_password`、`connect_timeout`（默认 300 秒）、`max_connect_retry`（默认 3）、`max_request_redirect`（0 = 无限制）。

### 地理编码（地址转坐标）

结构化搜索：

```sql
SELECT osm_id, ref, lon, lat, boundingbox
FROM nominatim_search(
  server_name => 'osm',
  street => 'Neubrueckenstrasse 63',
  city => 'Muenster',
  country => 'Germany'
);
```

自由格式搜索：

```sql
SELECT osm_id, display_name, lon, lat
FROM nominatim_search(
  server_name => 'osm',
  q => '1600 Pennsylvania Avenue, Washington DC'
);
```

**参数：** `q`（自由格式查询）、`street`、`city`、`county`、`state`、`country`、`postalcode`、`amenity`、`limit`（默认 10）、`addressdetails`、`extratags`、`namedetails`。

### 反向地理编码（坐标转地址）

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

**参数：** `lon`、`lat`（必填）、`zoom`（默认 18）、`addressdetails`、`extratags`、`namedetails`、`layer`。

### OSM 对象查找

```sql
SELECT osm_id, display_name
FROM nominatim_lookup(
  server_name => 'osm',
  osm_ids => 'W121736959,R123456'
);
```

**参数：** `osm_ids`（逗号分隔的 OSM ID 列表，带类型前缀：N=node，W=way，R=relation）、`addressdetails`、`extratags`、`namedetails`。

### 版本检查

```sql
SELECT nominatim_fdw_version();
```
