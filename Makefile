BINARY:=litra-ctl
MAIN:=.
SOURCES=$(shell find . -name \*.go)
PREFIX?=
SUDO?=sudo

PKGVERREL_git=$(shell git describe --long --dirty=+)
PKGVERREL=$(if $(PKGVERREL_git),$(patsubst v%,%,$(PKGVERREL_git)),$(error git describe failed))
PKGVER=$(firstword $(subst -, ,$(PKGVERREL)))
PKGREL=$(PKGVERREL:$(PKGVER)-%=%)

compile:
	go build -v ./...
.PHONY: compile

binary: $(BINARY)
.PHONY: binary
${BINARY}: Makefile go.mod go.sum $(SOURCES)
	go build -v -o $(BINARY) $(MAIN)

preinstall: binary
	rm -f packaging/completion/*
	mkdir -p packaging/completion
	set -eu ; for s in bash zsh fish ; do \
		./$(BINARY) completion $$s >packaging/completion/$(BINARY).$$s ; \
	done
.PHONY: preinstall
install: binary preinstall
	$(MAKE) install-only
.PHONY: install
install-only:
	$(SUDO) install -v $(BINARY) $(PREFIX)/usr/bin/
	$(SUDO) install -v -m 644 packaging/udev-rules/*.rules $(PREFIX)/lib/udev/rules.d/
	$(SUDO) install -v -m 644 packaging/completion/$(BINARY).bash /etc/bash_completion.d/$(BINARY)
	$(SUDO) install -v -m 644 packaging/completion/$(BINARY).zsh /usr/share/zsh/vendor-completions/_$(BINARY)
	$(SUDO) install -v -m 644 packaging/completion/$(BINARY).fish /usr/share/fish/vendor_completions.d/$(BINARY).fish
.PHONY: install-only
postinstall:
	$(SUDO) ./packaging/post-install.sh
.PHONY: postinstall

uninstall:
	$(SUDO) rm -vf "$(PREFIX)/usr/bin/$(BINARY)"
	cd packaging/udev-rules ; for f in *.rules ; do \
		$(SUDO) rm -vf "$(PREFIX)/lib/udev/rules.d/$$f" ; \
	done
.PHONY: uninstall

packages: package-deb
.PHONY: packages
package-deb: binary preinstall
# extra quoting on some args to work around checkinstall bugs:
# https://bugs.debian.org/cgi-bin/bugreport.cgi?bug=785441
	set -eu ; cd packaging/checkinstall && fakeroot checkinstall \
		--type=debian \
		--install=no \
		--fstrans=yes \
		--pkgarch=$(shell go env GOARCH) \
		--pkgname=litra-ctl \
		--pkgversion=$(PKGVER) \
		--pkgrelease=$(PKGREL) \
		--pkglicense=MIT \
		--pkggroup=utils \
		--pkgsource=https://github.com/fastcat/litra-ctl \
		--maintainer="'Matthew Gabeler-Lee <cheetah@fastcat.org>'" \
		--reset-uids=yes \
		--backup=no \
		$(MAKE) -C ../../ SUDO= install-only \
		</dev/null
.PHONY: package-deb

clean:
	rm -vf $(BINARY) packaging/checkinstall/*.deb
	rm -vf packaging/completion/$(BINARY).*
.PHONY: clean
