## Usage

Sources:

- [Alibaba Cloud rds_encdb documentation](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/use-the-rds-encdb-extension-to-encrypt-sensitive-columns)
- [Alibaba Cloud RDS for PostgreSQL overview](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/what-is-apsaradb-rds-for-postgresql)

`rds_encdb` is an Alibaba Cloud RDS for PostgreSQL feature that encrypts selected columns as they leave a query result. It is provider-managed result-set masking, not portable at-rest column encryption, and Alibaba Cloud does not publish a community source package or control file for it.

### Enablement and Core Workflow

The official procedure applies only to RDS PostgreSQL 16 instances with minor engine version `20250228` or later. Direct self-installation is not generally supported: obtain provider authorization, enable the instance parameter, and use the privileged account designated by RDS.

```sql
CREATE EXTENSION rds_encdb;

CREATE TABLE test(a text, b text, c text);
INSERT INTO test VALUES ('foo', 'bar', 'hello world');

INSERT INTO rds_encdb.encryption_rule
VALUES
  (9, 'rule1', 'test', '1'),
  (10, 'rule1', 'test', '2');

SELECT * FROM rds_encdb.rules;
```

The instance parameter `rds_encdb.enable_encryption` must be set to `on` before installation. New rows in `rds_encdb.encryption_rule` take effect for existing sessions. The documented default algorithm is AES-256-GCM.

### Rules and Role Authorization

`rds_encdb.encryption_rule` identifies a protected table and column by `attrelid` and `attnum`, grouped under `rule_name`. The `rds_encdb.rules` view summarizes those entries by rule and table.

Accounts default to `RESTRICTED ACCESS`, which returns ciphertext for protected result columns. `FULL ACCESS` returns plaintext. A privileged account manages this state with `rds_encdb.setup_encryption_role(account, permission, expiration)` and `rds_encdb.remove_encryption_role(account)`; records are visible in `rds_encdb.encryption_role_auth`. The optional EncJDBC client can decrypt restricted results after establishing its key material.

### Limits and Security Boundary

The provider documentation excludes result sets returned by functions, non-`SELECT` operations such as cursors and prepare/execute, and queries using CTE or `UNION`. Test actual ORM and connection-pool behavior because prepared queries are a common client default.

Encryption is applied to returned values according to account policy; it does not prevent a sufficiently privileged account from reading plaintext, and it does not substitute for storage encryption, transport encryption, auditing, or least privilege. Availability, upgrades, key behavior, authorization, and support all remain within Alibaba Cloud's AliPG service boundary.
