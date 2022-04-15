patch:
	./hack/make-rules/build.sh

.PHONY:clean
clean:patch
	go clean

.PHONY:deploy
deploy:patch
	make clean
	make patch
	scp patch root@c79:/root/

.PHONY:release
release:patch
	cp -f README.md patch ./patches/
	tar -czf patch.tgz ./patches