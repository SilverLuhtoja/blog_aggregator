# blog_aggregator
to kill in use port : fuser -k 8080/tcp

# host = 127.0.0.1-user-pass

POSTGRES:
- start postgres server: sudo service postgresql start
- check postgres server status: sudo service postgresql status
- stop postgres server: sudo service postgresql stop

MIGRATION:
1. move to sql/schema directory
2. Migrate up : goose postgres postgres://user:pass@localhost:5432/blog_db  up  
Migrate down : goose postgres postgres://user:pass@localhost:5432/blog_db   down  
OR 
1. run bash migrate.sh to migrate

GENERATE SQL to GO (from root): sqlc generate