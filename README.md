# stocks
Stock Tracker written in Go.

For personal use but could easily be utilised by anyone who wants to track their own portfolio.
Simply amend the stocks list to suit.

I want to develop a basic front end for this and have the API be called every minute to ensure data is refreshed regularly.

To build image:

```
$ CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
$ docker build -t stocks -f Dockerfile .
```

Then to run:

```
$ docker run -it stocks
```
