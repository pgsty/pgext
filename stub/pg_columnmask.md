## Usage

Sources:

- [AWS dynamic masking overview](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Security.DynamicMasking.html)
- [AWS policy procedure reference](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Security.DynamicMasking.Procedures.html)
- [AWS predefined masking functions](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Security.DynamicMasking.PredefinedMaskingFunctions.html)
- [AWS security best practices](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Security.DynamicMasking.BestPractices.html)

`pg_columnmask` is an Amazon Aurora PostgreSQL extension for role-based dynamic column masking. It transforms values read by users covered by a masking policy while leaving stored data unchanged. This is an Aurora-managed feature, not a portable open-source extension.

### Core Workflow

Create the extension as `rds_superuser`, define a role, and add a policy as the table owner, `rds_superuser`, or configured policy administrator:

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

Reconnect sessions opened before extension installation, then test under the target role rather than only as an administrator:

```sql
BEGIN;
SET ROLE support_user;
SELECT email FROM public.customers;
ROLLBACK;
```

### Object Index

- `pgcolumnmask.create_masking_policy(name, regclass, jsonb, name[], int)` creates a weighted role policy.
- `alter_masking_policy`, `rename_masking_policy`, and `drop_masking_policy` manage policies.
- `mask_text` supports a mask character, visible prefix/suffix, and hash-mask mode.
- `mask_email` can mask the local part, domain, or both.
- `mask_timestamp` replaces all or one timestamp component.
- `pgcolumnmask.policy_admin_rolname` delegates policy administration through the Aurora cluster parameter group.

### Operational Notes

AWS currently documents availability on Aurora PostgreSQL `16.10` and later and `17.6` and later; confirm the target engine's extension version before deployment. Installation is per database, and sessions connected before installation must reconnect.

Masking is an additional display-control layer, not encryption or a substitute for privileges. Values read in `WHERE`, joins, subqueries, and `RETURNING` are masked, while constraints apply to real stored values. Table triggers can see unmasked transition rows and can leak them; review trigger code carefully. The independent GUCs `pgcolumnmask.restrict_dml_triggers_for_masked_users` and `pgcolumnmask.restrict_iot_triggers_for_masked_users` default to `false` and can prohibit relevant trigger execution when required. Test dumps, replication, exports, views, and all application DML as the actual masked roles.
