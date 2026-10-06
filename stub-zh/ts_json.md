## 用法

来源：

- [README.md](https://github.com/za-arthur/apod_fts/blob/635ddf0dae7dde6dc2e594ef1b2a7b4f6fd21d10/README.md)
- [modules/ts_json/ts_json--1.0.sql](https://github.com/za-arthur/apod_fts/blob/635ddf0dae7dde6dc2e594ef1b2a7b4f6fd21d10/modules/ts_json/ts_json--1.0.sql)
- [modules/ts_json/ts_json.c](https://github.com/za-arthur/apod_fts/blob/635ddf0dae7dde6dc2e594ef1b2a7b4f6fd21d10/modules/ts_json/ts_json.c)
- [modules/ts_json/ts_json.control](https://github.com/za-arthur/apod_fts/blob/635ddf0dae7dde6dc2e594ef1b2a7b4f6fd21d10/modules/ts_json/ts_json.control)

`ts_json` 是 APOD 全文检索演示中的小型辅助扩展，遍历 JSON/JSONB 标量值并以指定分隔符拼接为文本，供检索流水线使用。

### 核心用法

```sql
CREATE EXTENSION ts_json;
SELECT jsonb_values('{"title":"PostgreSQL","body":"database"}'::jsonb, ' ');
```

### 运行边界

`json_values(json, text)` 与 `jsonb_values(jsonb, text)` 返回文本，第二个参数是分隔符，而不是字段名；结果省略键名，包含字符串、数字、布尔及 null 值。扩展不会安装演示网站、RUM 索引或全文检索字典。

这是历史 C 扩展，未提供当前主版本兼容矩阵；安装要求超级用户，不要求预加载。在目标 PostgreSQL 版本上使用前应验证旧源码行为。
