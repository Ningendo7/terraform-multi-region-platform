.PHONY: init fmt validate plan

init:
	terraform init

fmt:
	terraform fmt -check -recursive

validate:
	terraform validate

plan:
	terraform plan 