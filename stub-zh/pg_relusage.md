


## 用法

> [pg_relusage: 记录 SQL 语句实际使用的关系](https://github.com/adept/pg_relusage)

pg_relusage 挂钩到查询执行器，记录每个 SQL 语句实际使用的关系（表、视图、索引等）。与语句日志不同，它在视图展开和未使用的连接消除之后报告关系。

### 工作原理

加载后，每个 SQL 语句会产生一条列出所有引用关系的日志消息：

```sql
SELECT * FROM pg_stats LIMIT 1;
```

产生日志输出：
```
relations used: pg_stats,pg_statistic,pg_class,pg_attribute,pg_namespace
```

### 加载

```sql
-- 按会话加载
LOAD 'pg_relusage';

-- 或在 postgresql.conf 中全局配置
shared_preload_libraries = 'pg_relusage'
```

### 配置

| 参数 | 默认值 | 描述 |
|-----------|---------|-------------|
| `pg_relusage.log_level` | `LOG` | 关系消息的日志级别 |
| `pg_relusage.rel_kinds` | `'riSvmfp'` | 要报告的关系类型（来自 `pg_class.relkind` 的单字母代码） |

关系类型代码：`r` = 表，`i` = 索引，`S` = 序列，`v` = 视图，`m` = 物化视图，`f` = 外部表，`p` = 分区表。

### 使用场景

该扩展适用于发现遗留数据库中未使用的对象。通过分析一段时间内记录的关系使用情况，可以识别应用程序实际访问了哪些表、视图和索引。
