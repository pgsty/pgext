## 用法

来源：

- [官方扩展控制文件](https://github.com/mikeizbicki/pg_agg/blob/c64ecaee6441b73b7e2ed3bbb7c56b868a789e72/pg_agg.control)
- [官方扩展 SQL](https://github.com/mikeizbicki/pg_agg/blob/c64ecaee6441b73b7e2ed3bbb7c56b868a789e72/pg_agg--1.0.sql)

`pg_agg` 1.0 并没有实现控制文件注释所宣称的聚合查询索引支持。完整安装 SQL 只执行 `CREATE OR REPLACE LANGUAGE plpython3u`；它没有定义聚合、函数、类型、操作符、表、索引访问方法或配置参数。

### 核心流程

安装前应检查可用性，因为该扩展需要服务端的非受信任 PL/Python 3 模块：

```sql
SELECT name, default_version
FROM pg_available_extensions
WHERE name IN ('pg_agg', 'plpython3u');

CREATE EXTENSION pg_agg;

SELECT lanname, lanpltrusted
FROM pg_language
WHERE lanname = 'plpython3u';
```

`CREATE EXTENSION pg_agg` 唯一可观察的效果是创建或替换 `plpython3u` 语言。除此以外没有可记录的查询工作流。

### 重要对象

- `plpython3u` 是 1.0 版安装脚本中唯一出现的对象。
- `pg_agg` 是包装这一行安装动作的扩展记录；它自身不暴露任何用户 API。

### 运维说明

创建非受信任过程语言需要超级用户级权限，以及匹配 PostgreSQL 版本的服务端 PL/Python 软件包。非受信任 Python 函数以数据库服务器权限执行，因此即使本扩展自身没有创建函数，启用该语言仍是安全决策。除非后续权威上游版本提供真实实现，否则不要为聚合索引或查询加速选择 `pg_agg`；固定的 1.0 源码不具备这些能力。
