## 用法

来源：

- [项目 README](https://github.com/wulczer/pg-progress/blob/5d8343e933ff18396abc33e2ecabf4bb2c11a1d1/README.rst)
- [扩展 control 文件](https://github.com/wulczer/pg-progress/blob/5d8343e933ff18396abc33e2ecabf4bb2c11a1d1/progress.control)
- [0.0.1 版 SQL API](https://github.com/wulczer/pg-progress/blob/5d8343e933ff18396abc33e2ecabf4bb2c11a1d1/sql/progress--0.0.1.sql)
- [进度实现](https://github.com/wulczer/pg-progress/blob/5d8343e933ff18396abc33e2ecabf4bb2c11a1d1/src/progress.c)

`progress` 0.0.1 是 2013 年的原型，利用执行器插桩估算运行中查询的完成比例。它向独立监控会话提供 `pg_progress_update(integer)`、`pg_progress()` 和 `pg_progress_dot()`。

### 原型监控流程

服务器必须从上游项目修改过的 PostgreSQL 分支构建，并预加载该库：

```conf
shared_preload_libraries = 'progress'
```

重启并创建扩展后，监控会话向目标后端 PID 发信号并读取共享快照：

```sql
CREATE EXTENSION progress;

SELECT pg_progress_update(12345);
SELECT pg_progress();
SELECT pg_progress_dot();
```

`pg_progress_update` 要求目标后端计算快照，`pg_progress` 返回双精度估算值，`pg_progress_dot` 以 Graphviz DOT 形式返回插桩后的执行器树。仓库中的 `show-progress.py` 脚本负责协调查询与监控连接。

### 仅限原型的边界

该扩展无法针对原版 PostgreSQL 构建：它依赖修改过的 2013 年服务器源码中的私有执行器、插桩和进程信号 hook。它没有当前兼容路径，目录状态也标为 abandoned。不要把它加入受支持的生产 PostgreSQL 安装。

估算是启发式的，并受规划器基数误差影响。单个集群级共享快照会被监控者覆盖，因此并发观测并不按后端或数据库隔离。SQL 函数未声明为仅限超级用户，而向其他后端发信号并收集其执行器结构会增加工作量或泄露查询结构。若为了研究保留该原型，应在一次性数据环境中隔离运行，并显式限制函数执行权限。
