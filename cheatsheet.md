# Svelte

npm create svelte@latest frontend
npm run dev
npm run dev -- --open
npm run build
npm run preview

> To deploy your app, you may need to install an [adapter](https://kit.svelte.dev/docs/adapters) for your target environment.

### Notes

- Assignment is required to trigger an update.
- Svelte's `<script>` blocks are run only when the component is created, so assignments within a `<script>` block are not automatically run again when a prop updates. You need to use `$:` which marks a statement as reactive.

# Tern

### To create a new migration:

tern new name_of_migration

### To migrate:

tern migrate

### To migrate up or down to a specific version:

tern migrate --destination 42

### To migrate up N versions:

tern migrate --destination +3

### To use a different migrations directory:

tern migrate --migrations path/to/migrations

### To use a different config file:

tern migrate --config path/to/tern.json

---

# SQL & Postgres

### Start:

sudo service postgresql start

### Login:

sudo -u postgres psql

### Change password:

\password <user>

### Use database:

\c <database_name>;

### List all databases:

\l

### List all tables:

\dt

### List all schemas:

\dn

### List all functions:

\df

### List all users:

\du

### List all views:

\dv

### List all sequences:

\ds

### List all indexes:

\d

### List all foreign keys:

\db

### List all rules:

\dr

### List all configurations:

\dy

### List all data types:

\dT

### List all extensions:

\dx

### List all collations:

\dC

### List all conversions:

\dF

### List all languages:

\dL

### List all servers:

\ds

### List all tablespaces:

\db

### List all roles:

\du

### List all settings:

\ds

### Create database:

CREATE DATABASE <database_name>;

### Create table:

CREATE TABLE <table_name> (
<column_name> <data_type> <constraint>,
<column_name> <data_type> <constraint>,
<column_name> <data_type> <constraint>
);

### Create column:

ALTER TABLE <table_name>
ADD COLUMN <column_name> <data_type> <constraint>;

### Create schema:

CREATE SCHEMA <schema_name>;

### Create function:

CREATE FUNCTION <function_name> (<arguments>)

### Create view:

CREATE VIEW <view_name> AS
<select_statement>;

### Create sequence:

CREATE SEQUENCE <sequence_name>
START WITH <start_value>
INCREMENT BY <increment_value>
MINVALUE <minimum_value>
MAXVALUE <maximum_value>
CACHE <cache_value>;

### Create index:

CREATE INDEX <index_name>
ON <table_name> (<column_name>);

### Create foreign key:

ALTER TABLE <table_name>
ADD CONSTRAINT <constraint_name>
FOREIGN KEY (<column_name>)
REFERENCES <table_name> (<column_name>);

### Create rule:

CREATE RULE <rule_name>
AS ON <event>
TO <table_name>
DO <action>;

### Create configuration:

CREATE CONFIGURATION <configuration_name>
AS <configuration_value>;

### Create data type:

CREATE TYPE <data_type_name>
AS (<column_name> <data_type>);

### Create extension:

CREATE EXTENSION <extension_name>
SCHEMA <schema_name>
VERSION <version>;

### Create collation:

CREATE COLLATION <collation_name>
(LC_COLLATE = <collate_value>,
LC_CTYPE = <ctype_value>);

### Create conversion:

CREATE CONVERSION <conversion_name>
FOR <source_encoding> TO <destination_encoding>
FROM <function_name>;

### Create language:

CREATE LANGUAGE <language_name>
HANDLER <function_name>
VALIDATOR <function_name>;

### Create server:

CREATE SERVER <server_name>
FOREIGN DATA WRAPPER <wrapper_name>
OPTIONS (<option_name> <option_value>);

### Create tablespace:

CREATE TABLESPACE <tablespace_name>
OWNER <role_name>
LOCATION <directory>;
