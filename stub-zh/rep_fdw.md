## 用法

来源：

- [官方 README](https://github.com/asotolongo/rep_fdw/blob/990fb7b36c5334db9dabf2ccfd1d8502bd057cad/README)
- [官方扩展 SQL](https://github.com/asotolongo/rep_fdw/blob/990fb7b36c5334db9dabf2ccfd1d8502bd057cad/rep_fdw--0.1.0.sql)
- [官方扩展控制文件](https://github.com/asotolongo/rep_fdw/blob/990fb7b36c5334db9dabf2ccfd1d8502bd057cad/rep_fdw.control)

`rep_fdw` 版本 `0.1.0` 是一个建立在 `postgres_fdw` 上的实验性触发行复制器。它将 INSERT、UPDATE 与 DELETE 操作同步镜像到生成的外部表，但没有持久队列、冲突协议、模式同步或可靠的 SQL 引用处理，不能作为生产复制系统。

### 核心流程

扩展在 `rep_fdw` 模式中创建对象，并依赖 `postgres_fdw`：

```sql
CREATE EXTENSION postgres_fdw;
CREATE EXTENSION rep_fdw;

SELECT rep_fdw.create_server('remote_db', '127.0.0.1', '5433', 'app');
SELECT rep_fdw.create_usermap('replica_user', 'secret', 'remote_db');
SELECT rep_fdw.create_f_table('public', 'orders', 'remote_db');
SELECT rep_fdw.generar_trigger('public.orders');
```

`create_f_table(text, text, text)` 检查本地表，并创建名称带 `f_` 前缀的外部表。`generar_trigger(text)` 安装 AFTER 行触发器，其 `tr_ftable()` 处理器会在本地 DML 期间将旧行或新行写入远端表。

### 正确性与安全边界

`create_usermap(text, text, text)` 为 PUBLIC 创建用户映射，并把所给密码存入 PostgreSQL 目录选项。这远比逐角色映射宽泛；应限制目录可见性、所有权与函数执行权限，并优先手工管理最小权限映射。

触发器通过字符串连接构造标识符、值和谓词。引号、NULL 值、特殊类型或恶意输入可能破坏语句或改变含义。UPDATE 与 DELETE 会用全部旧列值进行匹配，而不是使用声明的主键，因此重复行和并发变化存在歧义。

远端写入是同步的：FDW 错误可能中止本地 DML，而网络分区与事务恢复仍需专门测试。只能用可丢弃数据研究 `rep_fdw` 的方法。当正确性、重放、监控或恢复很重要时，应使用受维护的逻辑复制或基于队列的设计。
