# opa-demo
Opa demo for self-learning
## Usage
Run the demo code
```Bash
$ go get -u github.com/gofiber/fiber/v2
$ go get -u github.com/open-policy-agent/opa/sdk
$ go build -o demo ./main
$ ./demo
 ┌───────────────────────────────────────────────────┐ 
 │                   Fiber v2.48.0                   │ 
 │               http://127.0.0.1:3000               │ 
 │       (bound on host 0.0.0.0 and port 3000)       │ 
 │                                                   │ 
 │ Handlers ............. 2  Processes ........... 1 │ 
 │ Prefork ....... Disabled  PID .............. 2552 │ 
 └───────────────────────────────────────────────────┘ 
```
Open another terminal
```Bash
$ curl -X POST -d "type=admin" http://127.0.0.1:3000/test_post
call post method!%
$ curl -X POST -d "type=normal_user" http://127.0.0.1:3000/test_post
only admin can access this method!% 
```
