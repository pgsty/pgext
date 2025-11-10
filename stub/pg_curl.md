
## Usage


```sql
CREATE EXTENSION pg_curl;
```

Perform HTTP Get:

```sql
-- wrap curl http get
CREATE OR REPLACE FUNCTION get(url TEXT) RETURNS TEXT LANGUAGE SQL AS $BODY$
WITH s AS (SELECT
               curl_easy_reset(),
               curl_easy_setopt_url(url),
               curl_easy_perform(),
               curl_easy_getinfo_data_in()
) SELECT convert_from(curl_easy_getinfo_data_in, 'utf-8') FROM s;
$BODY$;


SELECT get('https://www.postgresql.org/');
```


Perform Email SMTP:

```bash
CREATE OR REPLACE FUNCTION email(url TEXT, username TEXT, password TEXT, subject TEXT, sender TEXT, recipient TEXT, body TEXT, type TEXT) RETURNS TEXT LANGUAGE SQL AS $BODY$
    WITH s AS (SELECT
        curl_easy_reset(),
        curl_easy_setopt_mail_from(sender),
        curl_easy_setopt_password(password),
        curl_easy_setopt_url(url),
        curl_easy_setopt_username(username),
        curl_header_append('From', sender),
        curl_header_append('Subject', subject),
        curl_header_append('To', recipient),
        curl_mime_data(body, type:=type),
        curl_recipient_append(recipient),
        curl_easy_perform(),
        curl_easy_getinfo_header_in()
    ) SELECT curl_easy_getinfo_header_in FROM s;
$BODY$;
```

Perform FTP download:

```sql
CREATE OR REPLACE FUNCTION download(url TEXT, username TEXT, password TEXT) RETURNS BYTEA LANGUAGE SQL AS $BODY$
    WITH s AS (SELECT
        curl_easy_reset(),
        curl_easy_setopt_password(password),
        curl_easy_setopt_url(url),
        curl_easy_setopt_username(username),
        curl_easy_perform(),
        curl_easy_getinfo_data_in()
    ) SELECT curl_easy_getinfo_data_in FROM s;
$BODY$;
```
