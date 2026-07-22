## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/BrandwatchLtd/trunkfingerprint/blob/a4ed83cd533f233e70b23aac3ebb7d91248c7087/README.md)
- [1.3.0 版本 SQL 实现](https://github.com/BrandwatchLtd/trunkfingerprint/blob/a4ed83cd533f233e70b23aac3ebb7d91248c7087/trunkfingerprint--1.3.0.sql)
- [扩展控制文件](https://github.com/BrandwatchLtd/trunkfingerprint/blob/a4ed83cd533f233e70b23aac3ebb7d91248c7087/trunkfingerprint.control)

`trunkfingerprint` 为 PostgreSQL 目录结构以及可选的表数据计算可比较的 MD5 指纹。它适合检测由同一预期模式构建的数据库之间的漂移，而不是备份验证器或密码学完整性系统。

```sql
CREATE EXTENSION trunkfingerprint;
SELECT * FROM trunkfingerprint.get_db_fingerprint(
  p_level => 0,
  p_data => false,
  p_exclude_schemas => ARRAY['monitoring']::name[]
);
```

`p_level` 为 0 时返回一个指纹，为 1 时按被检查关系返回值，为 2 时输出对象级输入。`p_table` 限定扫描对象；`p_data` 选择仅结构、仅数据或两者；排除参数可从数据检查中移除模式、表或列。

实现会规范化许多 OID 引用并排除易变目录列，但指纹仍依赖 PostgreSQL 版本、扩展版本、对象定义、排序规则及工具自身的规范化规则。应只比较同类构建，并用级别 1 或 2 定位差异。

数据模式会扫描并序列化表内容，可能代价很高、让执行角色接触敏感值，并会随并发写入变化。应在一致快照或静止副本上运行，预算 I/O 和临时空间，并限制函数访问。MD5 抗碰撞性不足以提供对抗性证明；该用途应采用带认证的备份清单。
