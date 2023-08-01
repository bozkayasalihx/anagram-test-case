build: 
		@go build -o bin/anagramfinder
	
run: build
		@bin/anagramfinder anagrams.txt
test: 
		go clean -testcache && go test ./... -v -race -cover

