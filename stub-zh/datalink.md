## 用法

来源：

- [上游 README 与概念验证警告](https://github.com/darold/datalink/blob/fe8853949aedc5ec7df0b0e6a1d8f49d9b57b4bd/README.md)
- [扩展 control 文件](https://github.com/darold/datalink/blob/fe8853949aedc5ec7df0b0e6a1d8f49d9b57b4bd/datalink.control)
- [SQL 安装脚本](https://github.com/darold/datalink/blob/fe8853949aedc5ec7df0b0e6a1d8f49d9b57b4bd/sql/datalink--0.5.0.sql)
- [后台工作进程实现](https://github.com/darold/datalink/blob/fe8853949aedc5ec7df0b0e6a1d8f49d9b57b4bd/datalink_bgw.c)

`datalink` `0.5.0` 版是 SQL/MED `DATALINK` 类型的概念验证。值可以引用已注册基目录下的 URL 或文件并携带注释；数据库控制模式尝试提供令牌访问、写时复制，以及外部文件的清理/恢复。

### 最小构造

模块依赖 `uuid-ossp` 和 `uri`，而且必须由超级用户安装。SQL 扩展库和工作进程是两个不同的二进制文件；应预加载 `datalink_bgw` 而不是 `datalink`，并在创建扩展前重启 PostgreSQL：

```conf
shared_preload_libraries = 'datalink_bgw'
```

```sql
CREATE EXTENSION "uuid-ossp";
CREATE EXTENSION uri;
CREATE EXTENSION datalink;
CREATE TABLE persons (id integer, picture datalink);
INSERT INTO persons VALUES (1, DLVALUE('img1.jpg', 'MyBaseDir', 'portrait'));
```

使用本地路径前，要先在 `pg_catalog.pg_datalink_bases` 注册基目录。工作进程定义了 `datalink.dl_naptime`、`datalink.dl_token_expiry`、`datalink.dl_base_directory`、`datalink.dl_token_path` 和 `datalink.dl_keep_max_copies`；启动前应复核其路径与保留值。PostgreSQL 操作系统账户需要相应文件系统权限。

上游明确声明它不能用于生产。它会在服务器端创建、复制、重命名、链接和删除文件，后台工作进程还会删除令牌文件及部分外部副本。默认值和语义仍处早期阶段，源码也没有当前兼容矩阵。只能在隔离目录树中用备份和低权限操作系统账户进行测试。
