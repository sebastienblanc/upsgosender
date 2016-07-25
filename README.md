#upsgosender [![GoDoc](https://godoc.org/github.com/sebastienblanc/upsgosender?status.png)](https://godoc.org/github.com/sebastienblanc/upsgosender)

This is a client sender in Go for the [AeroGear UnifiedPush Server](https://aerogear.org/push/) 

Download:
```shell
go get github.com/sebastienblanc/upsgosender
```

* * *
## Usage 
```go
criteria := &Criteria{Alias: []string{"seb", "bob"}}
message := &Message{Alert: "hello from #golang sender"}
unifiedMessage := &UnifiedMessage{Message: *message, Criteria: *criteria}
settings := &Settings{
	URL:           "https://mypushserver.com/ag-push",
	ApplicationID: "58f87fb7-829c-4c6f-a0eb-326d3017a94c",
	MasterSecret:  "3366736b-d52c-4115-87d3-c08095e87955"}
sender := NewSender(*settings)
sender.send(*unifiedMessage)
        
```


* * *
Automatically generated by [autoreadme](https://github.com/jimmyfrasche/autoreadme) on 2016.07.23
