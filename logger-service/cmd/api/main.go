package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// defaultMaxUpload is the default max upload size (10 mb)
const defaultMaxUpload = 10485760

type Config struct {
	MaxJSONSize        int         // maximum size of JSON file we'll process
	MaxXMLSize         int         // maximum size of XML file we'll process
	MaxFileSize        int         // maximum size of uploaded files in bytes
	AllowedFileTypes   []string    // allowed file types for upload (e.g. image/jpeg)
	AllowUnknownFields bool        // if set to true, allow unknown fields in JSON
	ErrorLog           *log.Logger // the info log.
	InfoLog            *log.Logger // the error log.
	Models             data.Models
}

// New returns a new toolbox with sensible defaults.
func New() Config {
	return Config{
		MaxJSONSize: defaultMaxUpload,
		MaxXMLSize:  defaultMaxUpload,
		MaxFileSize: defaultMaxUpload,
		InfoLog:     log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog:    log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017"
	gRpcPort = "50001"
)

var client *mongo.Client

func main() {
	// connect to mongo
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}
	client = mongoClient

	// create a context in order to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// _ = client.Ping(ctx, readpref.Primary())

	// close connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	// start gRPC server
	go app.gRPCListen()

	// start RPC server
	go app.rpcListen()

	// register the RPC Server
	err = rpc.Register(new(RPCServer))

	// start web server
	app.serve()

}

func (app *Config) serve() {
	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	log.Printf("Starting logger service on port %s\n", webPort)
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func (app *Config) rpcListen() error {
	log.Printf("Starting RPC service on port %s\n", rpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))
	if err != nil {
		return err
	}
	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(rpcConn)
	}
}

func connectToMongo() (*mongo.Client, error) {
	// create connection options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// connect
	c, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Println("Error connecting:", err)
		return nil, err
	}

	return c, nil
}
