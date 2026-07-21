## 用法

来源：

- [nominatim_fdw v2.0 README](https://github.com/jimjonesbr/nominatim_fdw/blob/v2.0/README.md)
- [nominatim_fdw v2.0 changelog](https://github.com/jimjonesbr/nominatim_fdw/blob/v2.0/CHANGELOG.md)
- [Extension control file](https://github.com/jimjonesbr/nominatim_fdw/blob/v2.0/nominatim_fdw.control)
- [官方Nominatim API概览](https://nominatim.org/release-docs/develop/api/Overview/)
- [OpenStreetMap Nominatim使用政策](https://operations.osmfoundation.org/policies/nominatim/)

`nominatim_fdw` 从PostgreSQL调用一个Nominatim地理编码服务。与传统的FDW不同，它暴露了用于搜索、逆向地理编码和OSM对象查找的函数；它不会创建可查询的外部表。

### 配置服务器

```sql
CREATE EXTENSION nominatim_fdw;

CREATE SERVER osm
  FOREIGN DATA WRAPPER nominatim_fdw
  OPTIONS (
    url 'https://nominatim.openstreetmap.org',
    connect_timeout '10',
    max_connect_retry '2',
    max_connect_redirect '1',
    accept_language 'en'
  );
```

公共OpenStreetMap端点有一个官方使用政策。对于持续或批量工作负载，请使用授权提供商或运行自己的Nominatim服务，根据要求标识应用程序，并遵守速率限制。

### 核心流程

自由文本搜索：

```sql
SELECT osm_id, display_name, lon, lat
FROM nominatim_search(
  server_name => 'osm',
  q => 'Neubrückenstraße 63, Münster, Germany'
);
```

逆向地理编码和对象查找：

```sql
SELECT osm_id, display_name, addressdetails
FROM nominatim_reverse(
  server_name => 'osm',
  lon => 7.6293,
  lat => 51.9648,
  addressdetails => true
);

SELECT osm_id, display_name
FROM nominatim_lookup(
  server_name => 'osm',
  osm_ids => 'W121736959'
);
```

### 重要对象

- `nominatim_search(...)` 实现了自由文本或结构化前向搜索。
- `nominatim_reverse(...)` 将经度和纬度解析为最近的合适OSM地址。
- `nominatim_lookup(...)` 获取节点、方式或关系标识符，如`N123`、`W456`或`R789`。
- `nominatim_fdw_version()` 报告扩展和主库版本。
- `nominatim_fdw_settings` 以行的形式暴露依赖和构建版本。
- 服务器选项包括`url`、代理配置、超时设置、重试和重定向限制，以及默认的`accept_language`。

所有端点函数都是`STRICT`：显式SQL中的`NULL`参数返回空结果而不发送请求。在2.0中它们正确声明为`VOLATILE`，因为响应是远程的且可以更改。

### 2.0 版本变更和注意事项

2.0版本验证逆向坐标、添加了`email`、`polygon_threshold`和`entrances`，暴露依赖设置，并修复了返回详细字段中的JSON转义。它还具有用户可见的变化：逆向输出使用`display_name`；`addressparts`变为`addressdetails`；地址细节默认为真用于逆向和查找；版本输出更短。在从1.3升级之前，请审查结果列的消费者。

每次调用都会在网络语句中执行网络I/O操作。请使用有限的超时设置，限制谁可以创建或修改服务器，并避免在一个大型查询中的每一行都调用公共服务。上游构建需要PostgreSQL 10或更高版本、libxml2 2.5或更高版本以及libcurl 7.74或更高版本。
