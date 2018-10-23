Sample GraphQL application using both https://github.com/graphql-go/graphql and https://github.com/graph-gophers/graphql-go

Download
========
go get github.com/svobodal/graphqlsample

Usage
=====

Option 1
--------

Build:
```
cd src/github.com/svobodal/graphqlsample/cmd/gqlgo
go build
```

Run server:
```
./gqlgo
```

or

Option 2
--------

Build:
```
cd src/github.com/svobodal/graphqlsample/cmd/gophers
go build
```

Run server:
```
./gophers
```

Test
----

```
curl -X POST http://localhost:7070/query -d '{"query": "{ video(id: 10){name} }"}'
```
