# 175. Combine Two Tables

🟢 Difficulty: Easy

🔗 https://leetcode.com/problems/combine-two-tables/

---

## Approach

Start from `Person` so every person appears in the result, then use a `LEFT JOIN` on `personId` to attach an address when one exists. PostgreSQL returns `NULL` for the address columns when there is no matching row.

## PostgreSQL files

- [`solution.sql`](./postgresql/solution.sql) contains the query to submit.
- [`setup.sql`](./postgresql/setup.sql) creates temporary tables and loads the example data.
- [`run.sql`](./postgresql/run.sql) loads the setup and runs the solution in one `psql` session.

The portable PostgreSQL installation in this workspace is in `.tools/postgresql`, which is ignored by Git. Its database is stored in `.tools/postgresql/data`.

If the server is stopped (for example after restarting Windows), start it from the repository root:

```powershell
& .\.tools\postgresql\pgsql\bin\pg_ctl.exe -D .\.tools\postgresql\data -l .\.tools\postgresql\server.log -w -t 15 start
```

Then run the task:

```powershell
& .\.tools\postgresql\pgsql\bin\psql.exe -h localhost -U postgres -d postgres -f .\0000175_combineTwoTables\postgresql\run.sql
```

The output should contain Allen Wang with `NULL` city and state, and Bob Alice with New York City and New York. The row order may differ.

To stop the local server:

```powershell
& .\.tools\postgresql\pgsql\bin\pg_ctl.exe -D .\.tools\postgresql\data -m fast -w stop
```
