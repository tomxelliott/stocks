# stocks
Stock Tracker written in Go.

For personal use but could easily be utilised by anyone who wants to track their own portfolio.
Simply amend the stocks list to suit.

I want to develop a basic front end for this and have the API be called every minute to ensure data is refreshed regularly.

To build image:

Note - I will streamline this image so that it uses less memory by using *FROM scratch* instead of the current method which bundles the Go runtime into the Docker image which isn't necessary for this application and leads to a bloated image.

```
docker build -t stocks .
```

Then to run:

```
docker run stocks
```
