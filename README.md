# One Purpose SpiceDB Plugin for Steampipe

Imagine you have workspaces, permissions, and users in your system, and you decide to use
[SpiceDB](https://authzed.com/spicedb)
as the underlying authorization system to control access to those resources. However, 
you prefer the SQL language. What can we do with this?

We can utilize **Foreign Data Wrapper** in PostgreSQL! It's not necessary to start from scratch with implementing an 
FDW extension, as we can use [Steampipe](https://steampipe.io).

This plugin leverages the spicedb schema from another [POC](https://github.com/merlante/inventory_access_poc/blob/main/schema/spicedb_bootstrap.yaml).

Let's pose an authorization question:

**Which workspaces have granted the 'inventory_all_read' permission to user 'userlarge'?**

I can ask with zed command this way:

```
zed permission lookup-resources workspace inventory_all_read user:userlarge
```

But with SQL and using this Steampipe plugin, I can query this way:

```sql
SELECT name FROM workspaces 
WHERE user_name='userlarge' AND permission='inventory_all_read'
```

# How to use this plugin ?


First, learn about Steampipe and install it from [Steampipe](https://steampipe.io/downloads)

# Prerequisite

You need to have SpiceDB installed and update the SpiceDB configuration in ./fdw-spicedb-poc/utils.go
I use mentioned [POC](https://github.com/merlante/inventory_access_poc/ for these reasons.

You also need to generate some relations; you can use the script `gen_relations.sh`:

For example - This generates 60000 workspaces for user60000 with permission inventory_all_read:
```
./gen_relations.sh 60000 user60000
```


Then, install Go, and follow these steps:

```asciidoc
$ git clone https://github.com/lpichler/fdw-spicedb-poc.git

$ cp ./config/fdw-spicedb-poc.spc ~/.steampipe/config

$ make

$ steampipe query

> SELECT name FROM workspaces WHERE user_name='userlarge' AND permission='inventory_all_read'

```

