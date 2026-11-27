## 用法

来源：

- [扩展 control 文件](https://github.com/danolivo/execplan/blob/8591179dc1c95f93b2a4987bef648240dcf9b899/pg_execplan.control)
- [安装 SQL](https://github.com/danolivo/execplan/blob/8591179dc1c95f93b2a4987bef648240dcf9b899/init.sql)
- [安装脚本的生成规则](https://github.com/danolivo/execplan/blob/8591179dc1c95f93b2a4987bef648240dcf9b899/Makefile)
- [双节点集成测试](https://github.com/danolivo/execplan/blob/8591179dc1c95f93b2a4987bef648240dcf9b899/tests/rpl.sh)
- [项目说明与 PostgreSQL 内核补丁](https://github.com/danolivo/execplan/blob/8591179dc1c95f93b2a4987bef648240dcf9b899/README.md)

`pg_execplan` 是一个研究用扩展，用于在打过相同补丁的 PostgreSQL 12 开发版实例之间传递序列化的查询计划。`pg_repeater` 原型在接收端调用它的 SQL 接口。上游有多个修订都标为 `0.1`，但函数名有所不同。上述源码对应提供 `pg_exec_plan` 的修订。

### 在接收端数据库启用

先确认数据库运行的是匹配的内核和扩展库，再以超级用户安装。上游的集成测试在两个节点上都预加载了 `postgres_fdw` 和 `pg_execplan`。control 文件未声明 SQL 扩展依赖，安装时选定的模式以后不能更改。

```sql
CREATE EXTENSION pg_execplan;

SELECT p.oid::regprocedure
FROM pg_extension AS e
JOIN pg_depend AS d ON d.refclassid = 'pg_extension'::regclass
                   AND d.refobjid = e.oid AND d.deptype = 'e'
JOIN pg_proc AS p ON d.classid = 'pg_proc'::regclass AND p.oid = d.objid
WHERE e.extname = 'pg_execplan'
ORDER BY p.proname;
```

在这个修订中，`pg_store_query_plan` 将语句的查询计划写入服务端文件，`pg_exec_stored_plan` 读取已保存的计划，`pg_exec_plan` 接受查询文本及其序列化计划。应限制这些函数的调用权限，并仅使用可信输入。计划包含数据库内部数据结构，两端的源码修订和目录对象定义应当匹配。
