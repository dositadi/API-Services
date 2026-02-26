install-migrate:
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz
	sudo mv migrate /usr/local/bin/migrate

install-make:
	sudo apt install make-guile

migrate-create:
	migrate create -ext sql -dir database/migrations -seq $name

migrate-up:
	migrate -path database/migrations -database "mysql://username:password@tcp(host)/db_name" up 

migrate-down: 
	migrate -path database/migrations -database "mysql://username:password@tcp(host)/db_name" down 1

seed-into-users:
	mysql -u username -pPASSWORD blog_db < database/seed/insert_users.sql

seed-into-blogs: 
	mysql -u username -pPASSWORD blog_db < database/seed/insert_blogs.sql

seed-into-comments:
	mysql -u username -pPASSWORD blog_db < database/seed/insert_comments.sql

seed-into-tags:
	mysql -u username -pPASSWORD blog_db < database/seed/insert_tags.sql

seed-into-blog_tags:
	mysql -u username -pPASSWORD blog_db < database/seed/insert_blog_tags.sql