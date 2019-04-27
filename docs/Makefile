
doc-server:
	@docker build -t pg-dba-hugo .
	@docker run -d -p "1313:1313" -v $$(pwd):/doc --name pg-dba-docs pg-dba-hugo
	sleep 2 # Wait for doc server to come up
	open http://localhost:1313/pg-dba/

kill-doc-server:
	@docker rm -f pg-dba-docs
