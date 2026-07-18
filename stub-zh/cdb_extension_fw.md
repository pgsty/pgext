## 用法

来源：

- [项目 README](https://github.com/rafatower/cdb_extension_fw/blob/a728599d718fce887ec53c788121bb27191cd7ef/README.md)
- [扩展 control 文件](https://github.com/rafatower/cdb_extension_fw/blob/a728599d718fce887ec53c788121bb27191cd7ef/cdb_extension_fw.control)
- [示例 SQL 源码](https://github.com/rafatower/cdb_extension_fw/tree/a728599d718fce887ec53c788121bb27191cd7ef/src/sql/src)

`cdb_extension_fw` 是一个实验性的 CartoDB 框架和演示扩展，而不是面向应用的功能。已审查的 `current` 构建会安装固定的 `cdb_extension_fw` schema，以及两个输出 NOTICE 的 PL/pgSQL 函数。

### 调用演示 API

```sql
CREATE EXTENSION cdb_extension_fw;

SELECT cdb_extension_fw.foo();
SELECT cdb_extension_fw.bar();
```

两个调用都返回 void，并输出包含各自函数名的 NOTICE。上游构建的安装检查也会在创建扩展后执行这两个调用。

### 范围与维护状态

该框架把 SQL、Python 源码与生成的扩展发布文件分离，但此仓库只包含这两个演示函数。它没有提供用于打包任意数据或配置的通用运行时 API。

control 文件使用非数字版本 `current`，把对象固定在 `cdb_extension_fw` schema 中，并将扩展标记为不可迁移。仓库没有发布 release；其限制说明也明确由扩展开发者自行保证升级和降级正确性。应将其视为历史构建系统源码，安装前检查生成的发布产物，不要让生产 schema 依赖这些演示函数。
