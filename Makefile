post-compile: compile-proto
	$(GO) build -o $(BINDIR)/boilerplate $(HOMEDIR)/boilerplate/cmd/server/*.go
	$(GO) build -o $(BINDIR)/press $(HOMEDIR)/boilerplate/cmd/press/*.go

compile-proto: proto
	@for idl in \
		"protocols/arbiter" \
		"protocols/audio" \
		"protocols/event" \
		"protocols/image" \
		"protocols/text" \
	; do \
		GOPATH=$(GOPATH) && \
		PATH=$(GOPATH)/bin:$$PATH && \
		cd $$idl && \
		find . -maxdepth 10 ! -name "prediction.thrift" -type f -exec rm -rf {} \; && \
		kitex -module github.com/ishumei/krpc -service $$idl prediction.thrift && \
		cd -; \
	done

package-bin:
	mkdir -p $(OUTDIR)
	cp -rf $(HOMEDIR)/boilerplate/dist/* $(OUTDIR)

include Makefile.mk
