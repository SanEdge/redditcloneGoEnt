db_docs:
	dbdocs build doc/db.dbml

db_sqltodbml:
	sql2dbml --postgres doc/schema.sql -o doc/db.dbml

