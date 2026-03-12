

## 用法

> [pglogical_ticker: 准确查看 pglogical 复制延迟](https://github.com/enova/pglogical_ticker)

一个后台工作进程，定期更新 ticker 表以从提供者角度测量 pglogical 复制延迟。

### 启用

添加到 `postgresql.conf`：

```ini
shared_preload_libraries = 'pglogical,pglogical_ticker'
pglogical_ticker.database = 'mydb'
pglogical_ticker.naptime = 10      -- tick 间隔（秒，默认 10）
```

在提供者和所有订阅者上安装：

```sql
CREATE EXTENSION pglogical_ticker;
```

### 部署 Ticker 表

仅在**提供者**上运行（通过 pglogical 传播到订阅者）：

```sql
-- 部署 ticker 表（每个复制集一个）
SELECT pglogical_ticker.deploy_ticker_tables();

-- 将 ticker 表添加到复制
SELECT pglogical_ticker.add_ticker_tables_to_replication();
```

用于级联复制：

```sql
SELECT pglogical_ticker.deploy_ticker_tables('my_cascaded_set_name');
SELECT pglogical_ticker.add_ticker_tables_to_replication('my_cascaded_set_name');
```

### 手动 Tick

```sql
SELECT pglogical_ticker.tick();
```

### 启动 Ticker

如果在 `shared_preload_libraries` 中配置，ticker 在服务器启动时自动启动。否则：

```sql
SELECT pglogical_ticker.launch();

-- 或者，仅在复制集表存在时启动
SELECT pglogical_ticker.launch_if_repset_tables();
```

### 查看复制延迟

在**提供者**上：

```sql
SELECT * FROM pglogical_ticker.all_repset_tickers();
```

在**订阅者**上：

```sql
SELECT * FROM pglogical_ticker.all_subscription_tickers();
```

### 配置

- `pglogical_ticker.database` - 运行 ticker 的数据库
- `pglogical_ticker.naptime` - tick 间隔（秒，默认 10）
- `pglogical_ticker.restart_time` - 自动重启前等待秒数（默认 10，-1 禁用）
