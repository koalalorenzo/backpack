PACKAGES:=$(patsubst unpacked/%,%,$(wildcard unpacked/*))

%: 
	backpack pack unpacked/$*

all: $(PACKAGES)
.PHONY: all

publish:
	mkdir -p ../public/store/
	mv *.backpack ../public/store/
.PHONY: publish