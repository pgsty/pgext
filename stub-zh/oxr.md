## 用法

来源：

- [官方文档](https://github.com/brunoenten/pg_oxr/blob/6df5bd5515c2e088a03d44a0357d3ae129dfe11e/README.md)
- [扩展控制文件](https://github.com/brunoenten/pg_oxr/blob/6df5bd5515c2e088a03d44a0357d3ae129dfe11e/oxr.control)
- [官方仓库](https://github.com/brunoenten/pg_oxr)

`oxr` 在 PostgreSQL 内调用 OpenExchangeRates API 并缓存汇率。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `oxr`：

```sql
CREATE EXTENSION oxr CASCADE;
```

经审查的控制文件或官方流程要求 `http`。只有服务器已安装这些扩展文件时，`CASCADE` 才会成功。

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
-- Store your API key (replace 'your_api_key' with your actual key)
SELECT oxr.set_config('api_key', 'your_api_key');
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `oxr.set_config` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `oxr.get_config` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `oxr.get_historical_rate` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `oxr.get_latest_rate` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `oxr.historical_rates` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `oxr.set_app_id` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 13, 14, 15, 16, 17, 18；不要推断未列出的主版本。
- 扩展会在 `oxr` 下固定或创建模式对象；权限与备份审查应包含这些对象。
- 目录生命周期为 abandoned；生产使用前应测试升级、备份恢复与服务器兼容性。
