## 用法

来源：

- [AWS 动态脱敏概述](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Security.DynamicMasking.html)
- [AWS 策略过程参考](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Security.DynamicMasking.Procedures.html)
- [AWS 预定义脱敏函数](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Security.DynamicMasking.PredefinedMaskingFunctions.html)
- [AWS 安全最佳实践](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Security.DynamicMasking.BestPractices.html)

`pg_columnmask` 是 Amazon Aurora PostgreSQL 的基于角色动态列脱敏扩展。它会转换受脱敏策略约束的用户读取到的值，但不改变存储数据。这是 Aurora 托管功能，不是可移植的开源扩展。

### 核心流程

以 `rds_superuser` 创建扩展，定义角色，再由表所有者、`rds_superuser` 或已配置的策略管理员添加策略：

```sql
CREATE EXTENSION pg_columnmask;

CREATE ROLE support_user;
GRANT SELECT ON public.customers TO support_user;

CALL pgcolumnmask.create_masking_policy(
    'customer_email_mask'::name,
    'public.customers'::regclass,
    jsonb_build_object(
        'email', 'pgcolumnmask.mask_email(email)'
    ),
    ARRAY['support_user']::name[],
    100
);
```

重新连接在扩展安装前已建立的会话，然后以目标角色测试，而不是只用管理员测试：

```sql
BEGIN;
SET ROLE support_user;
SELECT email FROM public.customers;
ROLLBACK;
```

### 对象索引

- `pgcolumnmask.create_masking_policy(name, regclass, jsonb, name[], int)` 创建带权重的角色策略。
- `alter_masking_policy`、`rename_masking_policy` 和 `drop_masking_policy` 管理策略。
- `mask_text` 支持掩码字符、可见前后缀和哈希掩码模式。
- `mask_email` 可以掩盖邮件本地部分、域名或两者。
- `mask_timestamp` 替换整个时间戳或一个时间戳分量。
- `pgcolumnmask.policy_admin_rolname` 通过 Aurora 集群参数组委派策略管理。

### 运维说明

AWS 当前文档说明该扩展可用于 Aurora PostgreSQL `16.10` 及以上版本和 `17.6` 及以上版本；部署前应确认目标引擎的扩展版本。扩展按数据库安装，安装前已连接的会话必须重新连接。

脱敏是额外的显示控制层，不是加密，也不能替代权限。`WHERE`、连接、子查询和 `RETURNING` 读取的值会被脱敏，而约束针对真实存储值执行。表触发器可以看到未脱敏的过渡行并可能泄露数据，应仔细审查触发器代码。独立的 GUC `pgcolumnmask.restrict_dml_triggers_for_masked_users` 和 `pgcolumnmask.restrict_iot_triggers_for_masked_users` 默认均为 `false`，需要时可禁止相关触发器执行。应以实际受限角色测试转储、复制、导出、视图以及全部应用 DML。
