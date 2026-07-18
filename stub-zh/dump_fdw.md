## 用法

来源：

- [固定提交的上游 README](https://github.com/MeetMe/dump_fdw/blob/1af2c525d1526442abca4722d3e9b3db55a6426f/README.md)
- [1.0 版安装 SQL](https://github.com/MeetMe/dump_fdw/blob/1af2c525d1526442abca4722d3e9b3db55a6426f/dump_fdw--1.0.sql)
- [固定提交的 C 实现](https://github.com/MeetMe/dump_fdw/blob/1af2c525d1526442abca4722d3e9b3db55a6426f/dump_fdw.c)

dump_fdw 是一个只读外部数据包装器，用来直接查询 PostgreSQL 自定义格式 dump 文件中的关系。它适合在受控环境中做备份数据的临时分析，或筛选少量行进行恢复。

### 配置与示例

安装扩展、创建 server 对象，再把外部表映射到服务器本地 dump 文件中的一个关系：

```sql
CREATE EXTENSION dump_fdw;

CREATE SERVER dump_server
FOREIGN DATA WRAPPER dump_fdw;

CREATE FOREIGN TABLE dumped_authors (
    id integer,
    last_name text,
    first_name text
)
SERVER dump_server
OPTIONS (
    file_name '/srv/backups/booktown.dump',
    schema_name 'public',
    relation_name 'authors'
);

SELECT * FROM dumped_authors ORDER BY id;
```

三个外部表选项都必须提供，列定义也必须与 dump 中的关系一致。只有超级用户才应创建或修改这些映射，因为文件名可以指向 PostgreSQL 服务进程可见的任意文件。

### 限制与安全

该包装器只实现扫描，没有插入、更新或删除回调。它解析旧版自定义 dump 的内部格式，并通过类似 CSV 的路径物化行；实际兼容性取决于 dump 格式、压缩方式、数据类型和 PostgreSQL 构建版本。

上游将 1.0 版称为原型，只声明测试过 PostgreSQL 9.1 至 9.4，并明确警告不要用于生产。源码自 2015 年后未再更新，但仓库并未归档。应在隔离的恢复主机上使用 dump 副本和低权限数据库，并以代表性数据验证结果后再依赖它。
