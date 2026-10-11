## 用法

来源：

- [README.md](https://github.com/JoshNewham1/pg_forecast/blob/0af23990c63d05b5b669d4af48689e47ffc6f307/README.md)
- [src/pg_forecast/Makefile](https://github.com/JoshNewham1/pg_forecast/blob/0af23990c63d05b5b669d4af48689e47ffc6f307/src/pg_forecast/Makefile)
- [src/pg_forecast/sql/functions/_common.sql](https://github.com/JoshNewham1/pg_forecast/blob/0af23990c63d05b5b669d4af48689e47ffc6f307/src/pg_forecast/sql/functions/_common.sql)
- [src/pg_forecast/pg_forecast.control](https://github.com/JoshNewham1/pg_forecast/blob/0af23990c63d05b5b669d4af48689e47ffc6f307/src/pg_forecast/pg_forecast.control)

`pg_forecast` 拟合 ARIMA 模型、预测时间序列数值，并通过触发器更新模型状态。此研究实现链接 libnlopt；Makefile 将源 SQL 片段拼接为 1.0 版安装脚本。

### 核心用法

```sql
CREATE EXTENSION pg_forecast;
CREATE TABLE time_series_data (timestamp timestamp, value double precision);
```

调用训练与预测函数之前，先向源表载入具有代表性的有序时间序列观测值。

```sql
SELECT create_forecast('autoarima', 'time_series_data', 'timestamp', 'value');
SELECT * FROM run_forecast('autoarima', 'time_series_data', 'timestamp', 'value', 10, '1 hour');
```

### 运行边界

`create_forecast()` 按源表、日期列与数值列训练具名模型。`run_forecast()` 根据预测步数与时间间隔返回日期和值，`remove_forecast()` 移除已注册模型。已实现的模型是 autoarima。模型表与插入触发器是持久数据库对象。多个例程使用 SECURITY DEFINER 和 public 模式，仅应部署到对象创建权限与函数访问受控的数据库。安装采用控制文件的默认超级用户要求。上游提供 PostgreSQL 18 安装示例，但未给出完整兼容矩阵或明确许可证声明。
