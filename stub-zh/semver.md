## 用法

来源：

- [doc/semver.md](https://github.com/theory/pg-semver/blob/fdaf74754e704897b4632b211d9b7740930ebe69/doc/semver.md)
- [Changes](https://github.com/theory/pg-semver/blob/fdaf74754e704897b4632b211d9b7740930ebe69/Changes)
- [semver.control](https://github.com/theory/pg-semver/blob/fdaf74754e704897b4632b211d9b7740930ebe69/semver.control)
- [sql/semver--0.41.0--0.42.0.sql](https://github.com/theory/pg-semver/blob/fdaf74754e704897b4632b211d9b7740930ebe69/sql/semver--0.41.0--0.42.0.sql)

`semver` 提供符合语义版本 2.0.0 的数据类型、比较运算符及 `semverrange` 范围类型。0.42.0 修改了哈希算法，升级后必须重建所有基于 semver 值的哈希索引。

### 核心用法

```sql
CREATE EXTENSION semver;
SELECT '1.2.0'::semver > '1.2.0-b1'::semver;
SELECT '1.2.0'::semver = '1.2.0+build1'::semver;
SELECT to_semver('1.0'), is_semver('1.2.0');
SELECT get_semver_major('4.2.0'::semver);
SELECT '1.0.5'::semver <@ '[1.0.0,2.0.0)'::semverrange;
```

### 运行边界

`semver(text)` 及从 `text` 进行的类型转换要求合法的语义版本，`to_semver(text)` 提供较宽松的转换，`is_semver(text)` 检查格式有效性。`get_semver_major`、`get_semver_minor` 和 `get_semver_patch` 返回整数形式的版本部分，`get_semver_prerelease` 返回预发布文本。也支持从 `numeric`、`real`、`double precision`、`integer`、`bigint` 和 `smallint` 转换。0.42.0 会拒绝 Infinity 和 NaN，而不再产生格式错误的版本。

比较运算符 `=`、`<>`、`<`、`<=`、`>`、`>=` 以及 `MIN(semver)` 和 `MAX(semver)` 遵循语义版本顺序。预发布版排在对应正式版之前，构建元数据不影响优先级。支持 Btree 和哈希索引，`semverrange` 支持标准范围操作。

此原生扩展不受信任，需要管理员安装。安装 0.42.0 后，执行 `ALTER EXTENSION semver UPDATE`，并对所有包含 semver 键的已有哈希索引执行 REINDEX。仅更新扩展的 SQL 定义不会修复哈希变化带来的索引不兼容。本版还修正了以数字开头的字母数字预发布标识符排序，以及含多个加号的字符串校验。从较老版本升级时，还需检查上游兼容性说明，处理旧版曾接受的非法值。
