PACKAGES:=$(patsubst unpacked/%,%,$(wildcard unpacked/*))

%: 
	backpack pack unpacked/$*

all: $(PACKAGES)
.PHONY: all

publish:
	mkdir -p ../pages/store/
	mv *.backpack ../pages/store/
.PHONY: publish