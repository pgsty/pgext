## 用法

来源：

- [项目 README](https://github.com/ekayim/non_truncate/blob/036afbb78c4cbbbbb95a739fa420cb7b58b7d021/README.md)
- [扩展 control 文件](https://github.com/ekayim/non_truncate/blob/036afbb78c4cbbbbb95a739fa420cb7b58b7d021/non_truncate.control)
- [utility hook 实现](https://github.com/ekayim/non_truncate/blob/036afbb78c4cbbbbb95a739fa420cb7b58b7d021/non_truncate.c)

`non_truncate` 1.0 安装一个 PostgreSQL utility hook，使每条 `TRUNCATE` 语句都报错。它声明支持 PostgreSQL 13 及以上版本，并且必须在服务器启动时加载。

### 加载并登记扩展

把库加入 `shared_preload_libraries` 后重启 PostgreSQL：

```conf
shared_preload_libraries = 'non_truncate'
```

然后在需要通过扩展目录显示其存在的数据库中登记扩展：

```sql
CREATE EXTENSION non_truncate;

TRUNCATE sample_table;
-- ERROR: non_truncate extension is banning TRUNCATE statement.
```

### 集群级影响

拦截行为在共享库被预加载时开始，而不是在执行 `CREATE EXTENSION` 时开始。由于 `shared_preload_libraries` 是集群级设置，该 hook 也会影响没有扩展目录记录的数据库。此版本的安装 SQL 不创建任何可调用的 SQL 对象。

该 hook 会拒绝所有形式的 `TRUNCATE`，没有白名单、角色豁免、表豁免或会话开关。它无法阻止用 `DELETE`、替换表或其他 DDL 达到等价的数据清除效果，因此只是狭窄的防护栏，并非授权策略。部署前应测试它与其他所有 utility-hook 扩展的 hook 链接行为，并保留通过重启禁用它的操作流程。
