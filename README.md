To run:

```
docker-compose up -d
make build
./bin/server
./bin/client [cmd] -key=aString -value=aString


// example

./bin/client create -key=name -value=aid
``` 
 
Possible cmds:
- test
- create -key=[] -value=[]
- update -key=[] -value=[]
- get -key=[]  
- delete -key=[]  
- getHistory -key=[]  
- clear (clear tables)

- How would you support multiple users?
- - add user.id as foreign key
- How would you support answers with types other than string?
- - use interfaces and maps
- What are the main bottlenecks of your solution?
- - single connection to DB, without pool
- How would you scale the service to cope with thousands of requests?
- - multiple instances with load balancer or queues
