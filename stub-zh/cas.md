## 用法

来源：

- [扩展控制文件](https://github.com/philopon/cas/blob/e23971e2a65e5c9e50fa238cce9bb4209f041b36/cas.control)
- [0.1.0 版本扩展 SQL](https://github.com/philopon/cas/blob/e23971e2a65e5c9e50fa238cce9bb4209f041b36/cas--0.1.0.sql)
- [C 实现](https://github.com/philopon/cas/blob/e23971e2a65e5c9e50fa238cce9bb4209f041b36/cas.c)

`cas` 为形似美国化学文摘社登记号的值定义了一种定长 PostgreSQL 类型。它解析由连字符分隔的三个数字部分，允许读取各部分，并提供比较运算符与默认 B-tree 运算符类。它会存储输入的校验位，但不会验证 CAS 校验和算法。

### 核心流程

先确认软件包能在目标 PostgreSQL 构建上安装：

```sql
CREATE EXTENSION cas;

CREATE TABLE substance (
    registry_number cas PRIMARY KEY,
    name text NOT NULL
);

INSERT INTO substance VALUES ('57-09-0', 'example');

SELECT registry_number,
       cas_first(registry_number),
       cas_second(registry_number),
       cas_checkdigit(registry_number)
FROM substance;
```

值会依次按照第一部分、第二部分和存储的校验位进行数值排序。

### 对象索引

- `cas` 是带有 `first-second-check` 文本输入输出格式的 12 字节标量类型。
- `cas_first(cas)`、`cas_second(cas)` 和 `cas_checkdigit(cas)` 返回存储的各部分。
- 运算符 `=`、`!=`、`<`、`<=`、`>` 和 `>=` 比较各部分组成的元组。
- `cas_ops` 是默认 B-tree 运算符类。

### 运维说明

版本 `0.1.0` 可重定位，且未声明预加载要求。解析器只强制数字与连字符组成的格式，现实中校验位错误的值仍可保存。需要确保校验和正确时，应另外验证登记号。

固定版本的安装 SQL 在定义 PostgreSQL 基础类型通常所需的 shell 类型之前，就声明了引用 `cas` 的函数。部署前应在一次性数据库中测试 `CREATE EXTENSION cas`；全新安装的软件包可能需要修正扩展脚本。仓库没有提供 README 或发行文档，因此不能在已发布的控制文件、SQL、C 源码和回归文件之外推断类型转换、哈希索引或兼容性。
