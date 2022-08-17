package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fatih/color"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph"
	"github.com/jivison/gowon-indexer/lib/graph/generated"
	"github.com/jivison/gowon-indexer/lib/meta"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	startup()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	gqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	withMiddleware := meta.DoughnutMiddleware(meta.EnforcePassword(gqlHandler))

	srv := cors.AllowAll().Handler(withMiddleware)

	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	magenta := color.New(color.FgMagenta).SprintFunc()

	fmt.Printf("\nAPI running at %s\n", magenta(fmt.Sprintf("http://localhost:%s/graphql", port)))
	fmt.Printf("Playground running at %s\n", magenta(fmt.Sprintf("http://localhost:%s/playground", port)))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func startup() {
	godotenv.Load()

	const asciiArt = `888b     d888 d8b                                 888               888 888 
8888b   d8888 Y8P                                 888               888 888 
88888b.d88888                                     888               888 888 
888Y88888P888 888 888d888 888d888 .d88b.  888d888 88888b.   8888b.  888 888 
888 Y888P 888 888 888P"   888P"  d88""88b 888P"   888 "88b     "88b 888 888 
888  Y8P  888 888 888     888    888  888 888     888  888 .d888888 888 888 
888   "   888 888 888     888    Y88..88P 888     888 d88P 888  888 888 888 
888       888 888 888     888     "Y88P"  888     88888P"  "Y888888 888 888  미러볼`

	color.Cyan("\n\n" + asciiArt + "\n" + strings.Repeat("=", 83) + "\n\n")

	db.InitDB()
	fmt.Println("Connected to database")

}
