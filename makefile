# Make will use bash instead of sh
SHELL := /usr/bin/env bash
#CC=clang # required by bazel

# GNU make man page
# http://www.gnu.org/software/make/manual/make.html

# For some strange reasons, intends & blanks shift in bash when calling 'make' so the formatting below should align intend at least on Bash on OSX.
.PHONY: help
help:
	@echo ' '
	@echo 'Setup: '
	@echo '    make check        	    	Checks whether all project requirements are present.'
	@echo '    make gen_keys        	Generates new API access keys and a master key for end to end encryption.'
	@echo '    make db_deploy        	Deploys & starts the DB container. Run just once to create. Then use docker container start/stop timescaledb.'
	@echo '    make db_configure   	Configures the initial DB. Run once to first initialize DB or run again for a hard DB reset which deletes all data!'
	@echo ' '
	@echo 'Dev: '
	@echo '    make build   		Builds the code base incrementally (fast). Use for coding.'
	@echo '    make rebuild   		Rebuilds all dependencies & the code base (slow). Use after go mod changes. '
	@echo '    make run   			Runs the default target defined in dev/run script. Use to run default binary.'
	@echo ' '
	@echo 'Docker: '
	@echo '    make build_docker   	Builds a docker image locally.'
	@echo '    make run_docker   		Run docker images locally.'
	@echo '    make publish_docker  	Publishes docker images to registry.'
	@echo '    make remove_docker	  	Removes OMX container & image. DB & data will remain intact and carry over.'
	@echo '    make replace_docker  	Replaces running OMX image with latest published image. DB & data will carry over to new version.'
	@echo '    make reset_docker    	(!) ALL DATA WILL BE LOST(!) Removes running OMX container, replaces it with latest local build, AND destroys & rebuilds DB.'
	@echo ' '

# Currently, no cloud deployment defined in script. Adapt & edit as needed.
#	@echo 'Deploy: '
#	@echo '    make cloud_setup   	Configures GCP and all cloud requirements for deployment.'
#	@echo '    make deploy   		Deploys image to standard target defined in deploy/deploy.sh script.'
#	@echo '    make redeploy   		Redeploys image to standard target defined in deploy/redeploy.sh script.'
#	@echo '    make teardown   		Removes entire deployment as defined in deploy/teardown.sh script.'
#   @echo ' '

# "---------------------------------------------------------"
# Setup
# "---------------------------------------------------------"
.PHONY: check
check:
	@source scripts/setup/check_requirements.sh

.PHONY: gen_keys
gen_keys:
	@source scripts/setup/gen_keys.sh

.PHONY: db_deploy
db_deploy:
	@source  scripts/db/db_setup.sh

.PHONY: db_configure
db_configure:
	@source  scripts/db/db_configure.sh

# "---------------------------------------------------------"
# Development
# "---------------------------------------------------------"
.PHONY: build
build:
	@source scripts/dev/build.sh

.PHONY: rebuild
rebuild:
	@source scripts/dev/rebuild.sh

.PHONY: run
run:
	@source scripts/dev/run.sh

# "---------------------------------------------------------"
# Docker image build, publish, run & replace
# "---------------------------------------------------------"
.PHONY: build_docker
build_docker :
	@source scripts/docker/image_build.sh

.PHONY: run_docker
run_docker :
	@source scripts/docker/image_run.sh

.PHONY: publish_docker
publish_docker :
	@source scripts/docker/image_push.sh

.PHONY: remove_docker
remove_docker :
	@source scripts/docker/image_remove.sh


.PHONY: replace_docker
replace_docker :
	@source scripts/docker/image_replace.sh

.PHONY: reset_docker
reset_docker :
	@source scripts/docker/image_hard_reset.sh

# "---------------------------------------------------------"
# Deployment
# "---------------------------------------------------------"
.PHONY: cloud_setup
cloud_setup:
	@source scripts/deploy/cloud_setup.sh

.PHONY: deploy
deploy:
	@source scripts/deploy/deploy.sh

.PHONY: redeploy
redeploy:
	@source scripts/deploy/redeploy.sh

.PHONY: teardown
teardown:
	@source scripts/deploy/teardown.sh