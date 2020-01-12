service:
	docker-compose -f service-compose.yml up

infra:
	docker-compose -f terraform-compose.yml up
	if [ -f "./terraform/output/sqs_url.hcl" ];	then curl -X PUT -H "Content-Type: application/json" -d @./terraform/output/sqs_url.hcl http://localhost:8500/v1/kv/sqs/url; fi;
	if [ -f "./terraform/output/sqs_grp.hcl" ];	then curl -X PUT -H "Content-Type: application/json" -d @./terraform/output/sqs_grp.hcl http://localhost:8500/v1/kv/sqs/group; fi;

clean:
	docker-compose -f terraform-compose.yml down
	docker-compose -f service-compose.yml down