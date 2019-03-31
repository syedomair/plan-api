
run:    
	go run main.go router.go

rul:    
	psql -U khalid -d plan < database/table_def.sql; go run test/testURL/urlTest.go local

ruh:    
	go run test/testURL/urlTest.go 

tv:    
	go test ./... -v

ts:    
	go test ./... 
