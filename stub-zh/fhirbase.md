## 用法

来源：

- [官方仓库 README](https://github.com/niquola/fhirbase/blob/a0a90ac9501e915daa22b72e76d97cec47e8be9d/README.md)
- [1.0 版扩展脚本](https://github.com/niquola/fhirbase/blob/a0a90ac9501e915daa22b72e76d97cec47e8be9d/fhirbase--1.0.sql)
- [原型开发 SQL](https://github.com/niquola/fhirbase/tree/a0a90ac9501e915daa22b72e76d97cec47e8be9d/dev)

`fhirbase` 是一个始于 2014 年、用于在 PostgreSQL 中表示和搜索 FHIR 资源的未完成原型。发布的 1.0 版扩展脚本不包含模式对象或可调用 SQL API；仓库的设计笔记与开发文件则探索了生成资源表和搜索索引。不能把这些实验当作已安装扩展的契约。

### 安装实际提供的内容

若 PostgreSQL 能找到这些历史文件，创建扩展会登记扩展，但版本化 SQL 脚本不会创建仓库中讨论的 patient 表、搜索函数或索引：

```sql
CREATE EXTENSION fhirbase;

SELECT extname, extversion
FROM pg_extension
WHERE extname = 'fhirbase';

SELECT classid::regclass, objid, deptype
FROM pg_depend
WHERE refobjid = (
  SELECT oid FROM pg_extension WHERE extname = 'fhirbase'
);
```

最后一个查询用于审计：上游 1.0 脚本没有可用的用户对象。控制文件指定了模块库，但仓库 Makefile 没有列出能够暴露 C API 的编译对象文件。

### 原型材料

`dev` 目录草拟了 FHIR 资源生成、JSON 支持的数据类型、插入/更新辅助函数以及字符串/令牌搜索索引。README 讨论不区分大小写和重音的 FHIR 搜索参数，但它是以另一个项目名称开头的设计文档，而且 Usage 段落为空。这些文件不会由 `fhirbase--1.0.sql` 安装。

真正需要 FHIR 存储的应用，应根据目标 FHIR 发行版推导当前模式，并独立验证资源版本、引用、术语、搜索参数、访问控制与迁移。把原型 SQL 复制到生产环境会形成不受支持的分支，而不是启用 1.0 版。

### 兼容性与维护

仓库最后一次提交在 2014 年，目录也没有记录活跃生命周期。已安装 SQL 没有权威兼容矩阵、升级路径、预加载要求或受支持的 FHIR 发行版。控制文件允许重定位，但这本身不能让原型投入运行。

本文只用于识别目录条目，以及上游缺少可用 SQL 接口这一事实。若另一个项目或打包产物也使用 `fhirbase` 名称，应检查其自己的控制文件和扩展 SQL，不要因名称相似而套用本仓库 API。
