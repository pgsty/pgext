## 用法

来源：[README](https://github.com/jimjonesbr/nominatim_fdw/blob/master/README.md)，[Nominatim API](https://nominatim.org/)

`nominatim_fdw` 是一个面向 Nominatim 地理编码服务的 PostgreSQL foreign data wrapper。上游通过映射到 Nominatim `search`、`reverse` 和 `lookup` 端点的 SQL 函数暴露它，而不是通过 foreign tables。

### 创建服务器

```sql
CREATE EXTENSION nominatim_fdw;

CREATE SERVER osm
FOREIGN DATA WRAPPER nominatim_fdw
OPTIONS (url 'https://nominatim.openstreetmap.org');
```

上游记录的 server 选项包括：

- `url`（必填）
- `http_proxy`
- `connect_timeout`
- `max_connect_retry`
- `max_connect_redirect`

代理凭据应写在 user mapping 中：

```sql
CREATE USER MAPPING FOR pguser
SERVER osm
OPTIONS (proxy_user 'myuser', proxy_password 'mysecret');
```

### 地理编码函数

#### 结构化或自由格式搜索

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

#### 反向查询

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

#### OSM 对象查找

```sql
SELECT osm_id, display_name
FROM nominatim_lookup(
  server_name => 'osm',
  osm_ids => 'W121736959,R123456'
);
```

README 说明 OSM ID 需要带类型前缀，例如 `N` 表示 node、`W` 表示 way、`R` 表示 relation。

### 说明

- 上游要求 PostgreSQL 12+、`libxml2` 2.5.0+、`libcurl` 7.74.0+。
- 扩展暴露 `nominatim_fdw_version()` 用于查看版本。
- 当前 README 文档示例已使用 `CREATE EXTENSION ... WITH VERSION '1.3'` 与 `ALTER EXTENSION ... UPDATE TO '1.3'`，说明上游已超出这次 `1.2` 刷新目标。
