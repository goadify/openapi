generate_core:
	./bin/generate.sh generate_go core

generate_crud:
	./bin/generate.sh generate_go crud

generate_all:  generate_core generate_crud
