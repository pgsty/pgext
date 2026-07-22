## 用法

来源：

- [0.0.3 版官方 README](https://github.com/adjust/pg-geoip2lookup/blob/5c8144578f1c1461cc6ec0a12e18e9708de77220/Readme.md)
- [0.0.3 版扩展 SQL](https://github.com/adjust/pg-geoip2lookup/blob/5c8144578f1c1461cc6ec0a12e18e9708de77220/extension/geoip2lookup--0.0.3.sql)
- [PGXN 发行页面](https://pgxn.org/dist/geoip2lookup/)

`geoip2lookup` 是一个非受信 PL/Perl 扩展，在 PostgreSQL 服务器上读取 MaxMind MMDB 文件，并返回原始 JSONB 或有类型的地理位置记录。它适合偶发的地址补全，但每次调用都会打开数据库文件，上游实现也没有针对批量查询优化。

### 配置与查询

0.0.3 版安装到固定的 `geoip2lookup` 模式，依赖 `plperlu` 以及 Perl 的 `MaxMind::DB::Reader` 和 `JSON` 模块。配置服务器端目录与默认语言后，调用限定模式的包装函数：

```sql
CREATE EXTENSION geoip2lookup;

SET geoip2lookup.path = '/var/lib/GeoIP';
SET geoip2lookup.language = 'en';

SELECT *
FROM geoip2lookup.city('203.0.113.10'::inet);

SELECT geoip2lookup.raw_geoip2_json(
  '203.0.113.10'::inet,
  'City'
);
```

配置目录必须包含代码预期命名的文件，例如 `GeoIP2-City.mmdb`。底层重载接受显式路径，city/country 重载还接受语言。语言键区分大小写。

### 对象与 0.0.3 版实际接口

- `raw_geoip2_json` 以 JSONB 返回完整匹配 MMDB 记录。
- `city` 与 `country` 返回本地化名称、ISO/大洲标识和 GeoName 标识。
- `isp` 返回 ISP、组织与自治系统字段。
- `connection_type` 与 `anonymous_ip` 返回相应 MaxMind 数据库字段。
- 同名复合类型定义行布局。

发布的 0.0.3 版 SQL 存在必须如实面对的缺陷：单参数匿名 IP 包装器实际创建为 `anonymous_id`，ISP 行类型中一个字段拼写为 `autononous_system_number`。应在已安装构建上检查 `\df geoip2lookup.*` 与 `\dT+ geoip2lookup.*`，并使用这些实际标识符。

### 安全与运维

控制文件将扩展标记为仅限超级用户、不可重定位且依赖非受信 `plperlu`；没有预加载要求。函数以 PostgreSQL 服务器进程的文件系统权限读取路径。要限制执行权，不要把接受任意路径的重载暴露给不受信用户。

MMDB 数据是单独许可并定期更新的输入。应固定并审计数据库版本，原子刷新，并测试地址缺失行为。大批量补全更适合把结果载入维护表或使用面向批处理的服务，而不是在查询中反复打开文件。项目较旧且缺少自动回归夹具，因此必须在实际部署的 PostgreSQL、Perl、模块与 MMDB 版本组合上验证。
