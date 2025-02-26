<img src="Mirrorball Banner.png" alt="Mirrorball" width="500"/>

_모든 소린 나를 따라와_

This project is part of Gowon bot ([main repo](https://github.com/jivison/gowon))

## Running yourself

Ensure you have Go and Postgres properly installed. The indexer does not need to be in your gopath to run since it uses [Go modules](https://blog.golang.org/using-go-modules).

Install [sql-migrate](https://github.com/rubenv/sql-migrate), and then copy `dbconfig.example.yml` to `dbconfig.yml` and fill out your database credentials. Do the same with `.env.example`

_Note regarding 'WEBHOOK_URL', this is the url that the indexer posts to when it's done running tasks. If you have Gowon bot running, this defaults to localhost:3000. If you don't have/want Gowon bot running, I suggest using https://webhook.site for testing_

Ensure there exists a database called `gowon-indexer`, and then to run all migrations with `sql-migrate`, simply navigate the project root and run `sql-migrate up`.

To run the bot, simply run `go run server.go`. `go build` will generate an executable.

## Development

If you make any changes to the graphql schema file, you will need to rerun the code generator. You can do this by running `go run github.com/99designs/gqlgen generate`

## How to use

By default, the indexer runs on http://localhost:8080. Visiting that url will take you to a playground where you can view the schema and execute requests.

Requests should be made to http://localhost:8080/graphql

## General Structure

The indexer processes requests as follows:

1. `lib/graph/schema.resolvers.go` - This is the file that gqlgen generates that holds all the resolvers

2. `lib/controllers` - Every resolver hands off logic to a controller of the same name, accepting the same parameters and returning the same types as the resolver. Controllers typically handle high level logic.

3. `lib/services` - Similar functions grouped into services handle database interaction and most detailed logic

4. `lib/presenters` - Functions that "convert" database and other types into the generated "model" types

Some resolvers execute "tasks" (`lib/tasks`). Since these are asynchronous and run in a task server, they return a "task start response" attatched with a unique token. When the task is complete, the indexer posts to a url that it has completed the task.

## Any questions?

Somethings broken? Just curious how something works?

Feel free to shoot me a Discord dm at `@abbyfour`
or join the support server! https://discord.gg/9Vr7Df7TZf
