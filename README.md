# matching-timestamps

### Introduction
This is an implementation of a JSON/HTTP service, in golang, which returns the
matching timestamps of a periodic task.

A periodic task is described by the following properties:
- Period (every hour, every day, ...)
- Invocation point (where inside the period should be invoked)
- Timezone (days/months/years are timezone-depended)



### Endpoints


- A **GET** Method with endpoint **/ptlist?period={value}&tz={value}&t1={value}&t2={value}** , e.g., 
`http://localhost:8000/ptlist?period=1y&tz=Europe/Athens&t1=20180214T204603Z&t2=20211115T123456Z`.
  This endpoint returns all matching timestamps of a periodic task (ptlist) between 2 points in time (t1,
  t2). t1, t2 and the entries of ptlist are in UTC with seconds accuracy, in the following form:
  `20060102T150405Z`. The supported periods are: <number>h i.e., `1h`, <number>d i.e., `1d`, <number>mo i.e., `1mo`, <number>y i.e., `1y`.
  The invocation timestamp is at the start of the period (e.g. for 1h period a matching timestamp is considered the 20210729T010000Z).
  
  

### HowTo


#### Run matching-timestamps Service

`Ensure that you have Go installed on your system. `
- Navigate to the root directory of the project in your terminal.
- Run `go build ./cmd/main.go`  to build the binary executable for the service.
- Start the service by running the following command: ```./main```

###### Optional environment variables

```
SERVE_ON_PORT= Set service port, default value is 8000
```
