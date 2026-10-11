## 用法

来源：

- [docs/history.rst](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/docs/history.rst)
- [history/history--0.1.sql](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/history/history--0.1.sql)
- [docs/install.rst](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/docs/install.rst)
- [history/history.control](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/history/history.control)

`history` 是 Oliphant 的时态表扩展，可创建历史表并另行安装触发器，以选定时间精度保存基础表的版本。

### 核心用法

```sql
CREATE EXTENSION auth;
CREATE EXTENSION "merge";
CREATE EXTENSION history;
CREATE TABLE history_demo (id integer PRIMARY KEY, label text);
SELECT create_history_table('history_demo', 'day');
SELECT create_history_triggers('history_demo', 'day');
```

### 运行边界

这些是最初面向 PostgreSQL 9.x 编写文档的纯 SQL/PLpgSQL 工具，尚未验证当前主版本兼容性。安装沿用 control 默认权限，需要管理员；不含共享库，也不需要预加载。 需要先安装依赖 `auth` 和 `merge`。`create_history_table` 复制结构和当前数据，`create_history_triggers` 再连接后续写入；两个调用应使用相同的时间精度。按天保存的历史表示每天的状态，而非每次中间编辑。基础表需要合适的键。历史记录使用 `effective` 和 `expiry`，不会复制外键。基础表结构变更后需要维护生成的历史结构，删除扩展不能替代历史数据保留方案。
