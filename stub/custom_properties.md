## Usage

Sources:

- [Upstream README](https://github.com/GaryAustin1/custom-properties/blob/41f8434fab75171b0fe4080c0d150d85bcc16b18/README.md)
- [Extension control file](https://github.com/GaryAustin1/custom-properties/blob/41f8434fab75171b0fe4080c0d150d85bcc16b18/tle-custom-properties/custom_properties.control)
- [Trusted Language Extension SQL](https://github.com/GaryAustin1/custom-properties/blob/41f8434fab75171b0fe4080c0d150d85bcc16b18/tle-custom-properties/custom_properties--0.0.3.sql)
- [dbdev package page](https://database.dev/garyaustin/custom_properties)

`custom_properties` is a Supabase-oriented Trusted Language Extension that stores named properties for users and exposes helpers intended for row-level-security policies. The published `dbdev` package and available SQL are version `0.0.3`; the current control file says `0.0.4`, but no matching `0.0.4` install script is present at the reviewed commit.

### Package workflow

Generate a migration, review it, choose the destination schema, and apply it through the normal migration process:

```sh
dbdev add -o ./migrations -s extensions -v 0.0.3 package -n "garyaustin@custom_properties"
```

After installation, policy code can query the helpers (replace `user_roles` if a different schema was selected):

```sql
SELECT user_roles.user_has_property('Teacher');
SELECT user_roles.get_user_properties();
```

The SQL assumes Supabase objects such as `auth.users`, `auth.uid()`, `authenticated`, and `service_role`; it creates tables, RLS policies, grants, and `SECURITY DEFINER` functions. It also defines a disabled-by-default trigger that can rewrite `auth.users.raw_app_meta_data`. Audit the generated migration, function search paths, grants, and optional trigger before deployment. Do not assume the control-file version is installable until upstream supplies the corresponding SQL.
