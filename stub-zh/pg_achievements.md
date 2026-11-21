## 用法

来源：

- [README.md](https://api.pgxn.org/src/pg_achievements/pg_achievements-1.0.0/README.md)
- [pg_achievements.control](https://api.pgxn.org/src/pg_achievements/pg_achievements-1.0.0/pg_achievements.control)
- [sql/pg_achievements--1.0.sql](https://api.pgxn.org/src/pg_achievements/pg_achievements-1.0.0/sql/pg_achievements--1.0.sql)
- [LICENSE](https://api.pgxn.org/src/pg_achievements/pg_achievements-1.0.0/LICENSE)

`pg_achievements` 根据数据库活动授予成就，并在发起操作的会话中显示通知。发行版 1.0.0 安装的 SQL 扩展版本为 1.0。它要求 Linux 上的 PostgreSQL 16–18、超级用户安装权限以及集群预加载；固定的 `achievements` 模式位于指定的一个数据库中。

### 启用与查看

将库加入现有预加载列表，选择保存数据的数据库，然后重启 PostgreSQL：

```ini
shared_preload_libraries = 'pg_achievements'
pg_achievements.database = 'postgres'
```

在该数据库中创建扩展。事件采集覆盖集群中的各个数据库。

```sql
CREATE EXTENSION pg_achievements;
SELECT tier, title, points, awarded_at FROM achievements.my;
SELECT code, current, threshold, unlocked FROM achievements.progress;
SELECT * FROM achievements.leaderboard;
SELECT achievements.counter('stmt_select');
SELECT achievements.sync();
```

`achievements.my` 和 `achievements.progress` 显示当前角色的成就与进度。`achievements.catalog_public` 按规则隐藏秘密成就的细节；排行榜公开未退出统计的角色分数。`achievements.sync(timeout_ms)` 等待后台进程落库，超时返回 false，默认超时为 5000 毫秒。

### 跟踪控制

```sql
ALTER ROLE robot SET pg_achievements.opt_out = on;
SET pg_achievements.notify = off;
SET pg_achievements.notice_style = ascii;
```

执行角色设置示例前，角色必须已经存在。退出统计会停止跟踪，关闭通知则只隐藏提示。`pg_achievements.track_nested` 由超级用户控制，默认关闭。后台刷写间隔默认为 1000 毫秒；修改存储数据库或可跟踪角色的容量需要重启。

`achievements.reset()` 清除调用者的进度，超级用户可以重置其他角色。`achievements.define()` 根据已有计数器定义自定义成就，`achievements.grant()` 手动授予成就；这些管理函数受到权限限制。删除角色会保留其数据，需由超级用户调用 `achievements.forget_role()` 清除。

### 运行边界

跟踪独立于用户事务，回滚不会撤销已经授予的成就。后台进程异步保存计数器，因此崩溃可能丢失最近尚未刷写的进度。扩展保存计数器与成就，不保存 SQL 正文、参数或对象名称。PUBLIC 不能直接读取底层表。

预备语句的解析期特征只在解析时计数，客户端元命令也无法被直接观察。向运行中的服务器恢复进度后，应重启 PostgreSQL 以刷新共享内存。调整预加载配置时应保留原有库。
