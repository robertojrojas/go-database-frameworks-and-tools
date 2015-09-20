
# Preparing DB for Samples

1- Install Postgresql

2- Start Server

3- Run the following SQL Script

```sql

CREATE ROLE "go_db" VALID UNTIL 'infinity';
   
CREATE ROLE "go_user" WITH PASSWORD 'password1'  LOGIN INHERIT  VALID UNTIL 'infinity';
GRANT "go_db" TO "go_user";

CREATE DATABASE "godos_development"
  WITH ENCODING='UTF8'
       OWNER="go_db"
       CONNECTION LIMIT=-1;
       
       
CREATE TABLE public.users
(
   id serial PRIMARY KEY,
   email character varying(255), 
   name character varying(255)
) 
WITH (
  OIDS = FALSE
)
;
ALTER TABLE public.users
  OWNER TO go_user;   
   
INSERT INTO users(email, name)
    VALUES ('roberto@email.com', 'Roberto');

    
``` 
 
# To Clean Set the Name
```sql

-- update users set name = NULL

-- update users set name = 'Roberto'

```
    