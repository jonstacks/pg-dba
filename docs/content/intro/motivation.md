---
title: "Motivation"
date: 2019-03-26T14:40:22-05:00
weight: 2
---

The motivation for pg-dba is to automate Postgres administation. We often found that our tables
were becoming bloated and the number of dead tuples was large despite our best efforts at using
the `auto_vacuum` and `auto_analyze` features of postgres.

Furthermore, we often found ourselves running a SQL script to find the tables that were bloated
and then selectively running `FULL VACUUM` on them. Some tables are just too large to run a
`FULL VACUUM` on since it would lock the table for too long and take down an app for too long.

With this in mind, we created **pg-dba**. While still in its infancy, it has helped to make the
task of keeping our tables performing much easier and anyone on a team can run it without having
to have much knowledge about what is happening under the hood. This program could even be run on
a cron during a scheduled maintenance window to help keep your DB performing at its best.
