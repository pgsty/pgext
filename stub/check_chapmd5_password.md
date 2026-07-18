## Usage

Sources:

- [Official extension control file](https://github.com/sshutdownow/pg_check-chap-md5/blob/5371628350775b7d51ac097bee24913705856fb1/check_chapmd5_password.control)

`check_chapmd5_password` — Function for verifying CHAP-MD5 authentication responses.

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "check_chapmd5_password";
SELECT extversion
FROM pg_extension
WHERE extname = 'check_chapmd5_password';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
