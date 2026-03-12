

## 用法

> [geoip: 基于 IP 的 PostgreSQL 地理定位](https://github.com/tvondra/geoip)

该扩展提供基于 IP 的地理定位——你提供一个 IPv4 或 IPv6 地址，扩展会查找国家、城市、GPS 坐标、ASN 等信息。它需要 `ip4r` 扩展和来自 [MaxMind](https://www.maxmind.com) 的 GeoLite2 数据。

```sql
CREATE EXTENSION ip4r;
CREATE EXTENSION geoip;
```

### 函数

| 函数 | 说明 |
|------|------|
| `geoip_country_code(ip4\|ip6)` | 返回国家代码（2 字符） |
| `geoip_country(ip4\|ip6)` | 返回所有国家信息（代码、名称、网络） |
| `geoip_city_location(ip4\|ip6)` | 仅返回位置 ID（INT） |
| `geoip_city(ip4\|ip6)` | 返回所有城市信息（GPS、邮编等） |
| `geoip_asn(ip4\|ip6)` | 返回 ASN 名称和 IP 范围 |

### 示例

```sql
SELECT geoip_country_code('78.45.133.255'::ip4);
-- CZ

SELECT * FROM geoip.geoip_city('78.45.133.255'::ip4);
--  geoname_id | country_iso_code | city_name | postal_code | ...
-- ------------+------------------+-----------+-------------+----
--     3066399 | CZ               | Sardice   | 696 13      | ...

SELECT * FROM geoip.geoip_country('78.45.133.255'::ip4);
--     network     | country_iso_code | country_name
-- ----------------+------------------+--------------
--  78.45.128.0/17 | CZ               | Czechia

SELECT * FROM geoip.geoip_asn('78.45.133.255'::ip4);
--    network    | asn_number |      asn_name
-- --------------+------------+---------------------
--  78.44.0.0/15 |       6830 | Liberty Global B.V.
```


## 加载数据

该扩展需要来自 MaxMind 的 GeoLite2 CSV 数据。从 [MaxMind GeoLite2](https://dev.maxmind.com/geoip/geoip2/geolite2/) 下载 City、Country 和 ASN 数据集的 CSV 格式，然后加载：

```bash
cat GeoLite2-Country-Locations-en.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_country_locations FROM stdin WITH (FORMAT CSV, HEADER)'

cat GeoLite2-Country-Blocks-IPv4.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_country_blocks FROM stdin WITH (FORMAT CSV, HEADER)'

cat GeoLite2-Country-Blocks-IPv6.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_country_blocks FROM stdin WITH (FORMAT CSV, HEADER)'

cat GeoLite2-City-Locations-en.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_city_locations FROM stdin WITH (FORMAT CSV, HEADER)'

cat GeoLite2-City-Blocks-IPv4.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_city_blocks FROM stdin WITH (FORMAT CSV, HEADER)'

cat GeoLite2-City-Blocks-IPv6.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_city_blocks FROM stdin WITH (FORMAT CSV, HEADER)'

cat GeoLite2-ASN-Blocks-IPv4.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_asn_blocks FROM stdin WITH (FORMAT CSV, HEADER)'

cat GeoLite2-ASN-Blocks-IPv6.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_asn_blocks FROM stdin WITH (FORMAT CSV, HEADER)'
```

"locations" 文件有多种语言版本——选择适合你的即可。
