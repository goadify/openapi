generate_core:
	./bin/generate.sh generate_go core
	./bin/generate.sh generate_ts core

generate_crud:
	./bin/generate.sh generate_go crud
	./bin/generate.sh generate_ts crud

generate_navigation:
	./bin/generate.sh generate_go navigation
	./bin/generate.sh generate_ts navigation

generate_all:  generate_core generate_crud generate_navigation
