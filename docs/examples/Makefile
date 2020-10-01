PACKAGES:=$(patsubst unpacked/%,%,$(wildcard unpacked/*))

%: 
	backpack pack unpacked/$*

all: $(PACKAGES)
.PHONY: all

publish:
	mkdir -p ../../public/examples/
	mv *.backpack ../../public/examples/
.PHONY: publish