## 用法

来源：

- [1.0 版扩展 SQL](https://github.com/topretejal/lsm_idx/blob/8e71df09b49b4413dd18233705ccac223ccd75ce/lsm_idx--1.0.sql)
- [索引访问方法实现](https://github.com/topretejal/lsm_idx/blob/8e71df09b49b4413dd18233705ccac223ccd75ce/lsm_idx.c)
- [示例配置](https://github.com/topretejal/lsm_idx/blob/8e71df09b49b4413dd18233705ccac223ccd75ce/lsm_idx.conf)

`lsm_idx` 尽管名称指向 LSM，却只是未完成的 PostgreSQL 索引访问方法概念验证。1.0 版注册了访问方法和类似 B-tree 的操作符类，但其 C 处理器把所有扫描回调留空。它只能用于源码研究，不能用于应用索引。

### 仅用于一次性检查

在使用相同 PostgreSQL 头文件构建的隔离服务器中，可以检查扩展目录接口，而不要在重要数据上创建索引：

```sql
CREATE EXTENSION lsm_idx;

SELECT amname, amtype, amhandler::regproc
FROM pg_am
WHERE amname = 'lsm_idx';

SELECT opcname, opcintype::regtype
FROM pg_opclass
WHERE opcmethod = (SELECT oid FROM pg_am WHERE amname = 'lsm_idx')
ORDER BY 1, 2;
```

SQL 为常见整数、浮点、布尔、字符、二进制、日期时间、interval、网络相关、numeric、OID、text/name、money 与 UUID 类型定义了默认类，大多引用 PostgreSQL 内部 B-tree 支持函数。

表面上的创建语法如下：

```sql
CREATE INDEX experimental_lsm_idx
ON experimental_table USING lsm_idx (key_column);
```

不要在有价值的数据上执行该示例。处理器临时更改关系的访问方法 OID 后，把构建与插入委托给私有 B-tree 内部函数，并向 B-tree 元页面写入额外计数器，还会通过 `printf` 输出诊断信息。

### 缺失功能与风险

访问方法例程把 `ambeginscan`、`amrescan`、`amgettuple`、`amgetbitmap` 与 `amendscan` 设为空，也禁用了唯一索引、数组搜索和并行扫描。它没有实现 LSM 合并、层级管理、压实、恢复协议、代价模型或用户可见统计接口。

仓库包含 `shared_preload_libraries` 与 `lsm_idx.top_index_size` 示例，但 C 源码没有定义启动钩子或该名称的 GUC。这些配置行不能证明存在可用的预加载要求或可调选项。

1.0 版是 2020 年的单提交原型，没有 README、测试、升级路径或 PostgreSQL 兼容矩阵。控制文件允许重定位且未声明依赖。生产环境应使用 PostgreSQL 受支持的 B-tree 访问方法或维护中的替代品；继续开发该原型需要补全访问方法，并进行崩溃、WAL、vacuum、并发、规划器、升级与损坏测试。
