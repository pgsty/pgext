## 用法

来源：

- [packages/ltree-helpers/sql/pgpm-ltree-helpers--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/ltree-helpers/sql/pgpm-ltree-helpers--0.47.0.sql)
- [packages/ltree-helpers/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/ltree-helpers/Makefile)
- [packages/ltree-helpers/pgpm-ltree-helpers.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/ltree-helpers/pgpm-ltree-helpers.control)

`pgpm-ltree-helpers` 提供三个小型 SQL 转换函数：`ltree_helpers.to_path`、`ltree_helpers.to_slash` 和 `ltree_helpers.to_query`。

### 核心用法

```sql
CREATE EXTENSION "pgpm-ltree-helpers";
SELECT ltree_helpers.to_path('/team/docs');
SELECT ltree_helpers.to_slash('team.docs'::ltree);
SELECT ltree_helpers.to_query('/team/**');
```

### 运行边界

路径转换会移除开头的斜杠，再将分隔符替换为点。查询转换将单星号映射为一个标签，将双星号映射为 lquery 的多标签通配符。这只是字符串替换后进行类型转换，不是文件系统通配解析器，也不会转义标签。输入必须符合 ltree/lquery 语法，无效标签会触发转换错误；函数声明为严格且不可变。 依赖为 `plpgsql`, `ltree`, `pgpm-verify`。 安装脚本会向预先存在的角色 `authenticated`, `anonymous` 授权。 这是纯 SQL 模块，无需预加载动态库或重启；该修订未公布 PostgreSQL 主版本支持矩阵。模式布局固定，不可迁移。启用前须部署相匹配的完整依赖。
