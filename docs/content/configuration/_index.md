---
title: "Configuration"
date: 2019-03-26T11:07:21-05:00
chapter: true
weight: 3
pre: "<b>3. </b>"
---

### Chapter 3

# Configuration

Starting in version `0.4.0`, pg-dba will automatically read a `.env` file if one
is present in the current directory. Also starting in version `0.4.0`, pgdba respects
and prioritizes the following postgres environment variables over the old ones:

```
| Postgres Environment Variable | Old Environment Variable |
| ------------------------------|--------------------------|
| PGHOST                        | POSTGRES_HOST            |
| PGUSER                        | POSTGRES_USER            |
| PGDATABASE                    | POSTGRES_DB              |
| PGPASSWORD                    | POSTGRES_PASSWORD        |
| PGSSLMODE                     | SSL_MODE                 |
```

Currently, the execution of pg-dba is configured with environment variables.
Below is a list of the environment variables, their defaults, descriptions,
and allowed values:


| Environment Variable            | Default Value | Description                                       | Allowed Values                                                            |
|---------------------------------|---------------|---------------------------------------------------|---------------------------------------------------------------------------|
| ANALYZE_TIMEOUT_SECONDS         | 600           | The time in seconds before stopping an analyze    | Integer greater than 0                                                    |
| BLOAT_QUERY_TIMEOUT_SECONDS     | 30            | The time in seconds before stopping bloat query   | Integer greater than 0                                                    |
| FULL_VACUUM_TIMEOUT_SECONDS     | 600           | The time in seconds before stopping a full vacuum | Integer greater than 0                                                    |
| POSTGRES_DB                     | postgres      | The postgres DB                                   | `*`                                                                       |
| POSTGRES_HOST                   | localhost     | The postgres host                                 | `*`                                                                       |
| POSTGRES_PASSWORD               | ""            | The postgres password                             | `*`                                                                       |
| POSTGRES_USER                   | postgres      | The postgres user                                 | `*`                                                                       |
| LOG_LEVEL                       | info          | The level to output logs at                       | debug, info, warn, error                                                  |
| POST_ANALYZE                    | True          | Run an analyze after vacuum to update stats       | True, False                                                               |
| PRE_ANALYZE                     | True          | Run an analyze to update stats before vacuum      | True, False                                                               |
| SSL_MODE                        | require       | SSLMode for the connection.                       | See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters  |
| VERBOSE                         | False         | Run queries in VERBOSE mode                       | True, False                                                               |
| VACUUM_TIMEOUT_SECONDS          | 600           | The time in seconds before stopping a vacuum      | Integer greater than 0                                                    |
