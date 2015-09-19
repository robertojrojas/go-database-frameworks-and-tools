
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
       
       
CREATE TABLE public.todos
(
   id serial PRIMARY KEY,
   subject character varying(255), 
   description character varying(255), 
   completed boolean, 
   created_at date, 
   updated_at date
) 
WITH (
  OIDS = FALSE
)
;
ALTER TABLE public.todos
  OWNER TO go_user;   
   
ALTER TABLE public.todos ALTER COLUMN created_at
SET DEFAULT CURRENT_DATE;
   
ALTER TABLE public.todos ALTER COLUMN updated_at
SET DEFAULT CURRENT_DATE;  
 
INSERT INTO todos(subject, description, completed, created_at, updated_at)
    VALUES ('Learn Go', 'Go is really cool.', 'f', DEFAULT, DEFAULT);
    
INSERT INTO todos(subject, description, completed, created_at, updated_at)
    VALUES ('Buy Milk', null, 't', DEFAULT, DEFAULT);
    
INSERT INTO todos(subject, description, completed, created_at, updated_at)
    VALUES ('Shovel Snow', 'I need to move to Florida!', 'f', DEFAULT, DEFAULT);
    
INSERT INTO todos(subject, description, completed, created_at, updated_at)
    VALUES ('Watch Cybercrimes', 'Learning Interesting stuff.', 'f', DEFAULT, DEFAULT);
    
```  
    