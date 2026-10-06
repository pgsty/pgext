## 用法

来源：

- [packages/achievements/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/achievements/README.md)
- [packages/achievements/sql/pgpm-achievements--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/achievements/sql/pgpm-achievements--0.47.0.sql)
- [packages/achievements/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/achievements/Makefile)
- [packages/achievements/pgpm-achievements.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/achievements/pgpm-achievements.control)

`pgpm-achievements` 记录完成的步骤、成就及等级要求，由触发器在固定状态模式中更新成就状态。

### 核心用法

```sql
CREATE EXTENSION "pgpm-achievements" CASCADE;
SELECT * FROM status_public.levels;
SELECT * FROM status_public.level_requirements;
```

### 运行边界

`status_public.user_steps` 与 `status_public.user_achievements` 保存用户进度，`status_public.steps_required` 和 `status_public.user_achieved` 用于检查。接收用户写入前应审查 JWT 声明集成、行策略及 SECURITY DEFINER 辅助函数，并先配置等级要求。

0.47.0 是 SQL/PLpgSQL 扩展，没有自己的共享库，也不要求预加载。须先安装配套版本的依赖：`plpgsql`、`pgpm-jwt-claims`、`pgpm-verify`。 安装 SQL 要求平台角色 `anonymous`, `authenticated` 已存在，须先使用上游角色初始化流程，并审查授权。 控制文件允许非超级用户安装，但没有标记为 trusted；依赖、模式创建和角色授权权限仍须满足。上游未声明当前 PostgreSQL 主版本矩阵。
