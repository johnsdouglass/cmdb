package main

import "github.com/spf13/pflag"

const (
	// configuration defaults support local development (i.e. "go run ...")
	// Server
	defaultServerAddress = "0.0.0.0"
	defaultServerPort    = "9090"
	defaultServerVersion = "/v1/"

	// Gateway
	defaultGatewayEnable      = true
	defaultGatewayAddress     = "0.0.0.0"
	defaultGatewayPort        = "8080"
	defaultGatewayURL         = "cmdb"
	defaultGatewaySwaggerFile = "pkg/pb/cmdb.swagger.json"
	defaultGatewaySwaggerUI   = "./apidoc-poc/swagger-ui-dist/"

	// Database
	defaultDatabaseEnable = true
	// DSN example: "postgres://postgres:postgres@postgres:5432/atlas_db?sslmode=disable"
	defaultDatabaseDSN      = ""
	defaultDatabaseType     = "postgres"
	defaultDatabaseHost     = "0.0.0.0"
	defaultDatabasePort     = "5432"
	defaultDatabaseName     = "cmdb"
	defaultDatabaseUser     = "postgres"
	defaultDatabasePassword = "postgres"
	defaultDatabaseSSL      = "disable"
	defaultDatabaseOption   = ""

	// PubSub
	defaultPubsubEnable       = false
	defaultPubsubAddress      = "pubsub.atlas"
	defaultPubsubPort         = "5555"
	defaultPubsubPublish      = "topic"
	defaultPubsubSubscribe    = "topic"
	defaultPubsubSubscriberID = "example_hello_subscriberid"

	// Authz
	defaultAuthzEnable  = false
	defaultAuthzAddress = "authz.atlas"
	defaultAuthzPort    = "5555"

	// Audit Logging
	defaultAuditEnable  = false
	defaultAuditAddress = "atlas.audit"
	defaultAuditPort    = "5555"

	// Tagging
	defaultTaggingEnable  = false
	defaultTaggingAddress = "atlas.tagging"
	defaultTaggingPort    = "5555"

	// Health
	defaultInternalEnable    = true
	defaultInternalAddress   = "0.0.0.0"
	defaultInternalPort      = "8081"
	defaultInternalHealth    = "/healthz"
	defaultInternalReadiness = "/ready"

	defaultConfigDirectory = "deploy/"
	defaultConfigFile      = ""
	defaultSecretFile      = ""
	defaultApplicationID   = "cmdb-app"

	// Logging
	defaultLoggingLevel = "debug"
)

var (
	// define flag overrides
	flagServerAddress = pflag.String("server.address", defaultServerAddress, "adress of gRPC server")
	flagServerPort    = pflag.String("server.port", defaultServerPort, "port of gRPC server")
	flagServerVersion = pflag.String("server.version", defaultServerVersion, "endpoint version of server")

	flagGatewayEnable      = pflag.Bool("gateway.enable", defaultGatewayEnable, "enable gatway")
	flagGatewayAddress     = pflag.String("gateway.address", defaultGatewayAddress, "address of gateway server")
	flagGatewayPort        = pflag.String("gateway.port", defaultGatewayPort, "port of gateway server")
	flagGatewayURL         = pflag.String("gateway.endpoint", defaultGatewayURL, "endpoint of gateway server")
	flagGatewaySwaggerFile = pflag.String("gateway.swaggerFile", defaultGatewaySwaggerFile, "directory of swagger.json file")
	flagGatewaySwaggerUI   = pflag.String("gateway.swaggerUI", defaultGatewaySwaggerUI, "directory of swagger-ui dist")

	flagDatabaseEnable   = pflag.Bool("database.enable", defaultDatabaseEnable, "enable database")
	flagDatabaseDSN      = pflag.String("database.dsn", defaultDatabaseDSN, "DSN of the database")
	flagDatabaseType     = pflag.String("database.type", defaultDatabaseType, "type of the database")
	flagDatabaseHost     = pflag.String("database.host", defaultDatabaseHost, "address of the database")
	flagDatabasePort     = pflag.String("database.port", defaultDatabasePort, "port of the database")
	flagDatabaseName     = pflag.String("database.name", defaultDatabaseName, "name of the database")
	flagDatabaseUser     = pflag.String("database.user", defaultDatabaseUser, "database username")
	flagDatabasePassword = pflag.String("database.password", defaultDatabasePassword, "database password")
	flagDatabaseSSL      = pflag.String("database.ssl", defaultDatabaseSSL, "database ssl mode")
	flagDatabaseOption   = pflag.String("database.option", defaultDatabaseOption, "define custom option to db driver")

	flagPubsubEnable       = pflag.Bool("atlas.pubsub.enable", defaultPubsubEnable, "enable application with pubsub")
	flagPubsubAddress      = pflag.String("atlas.pubsub.address", defaultPubsubAddress, "address or FQDN of the pubsub service")
	flagPubsubPort         = pflag.String("atlas.pubsub.port", defaultPubsubPort, "port of the pubsub service")
	flagPubsubPublish      = pflag.String("atlas.pubsub.publish", defaultPubsubPublish, "publisher topic")
	flagPubsubSubscribe    = pflag.String("atlas.pubsub.subscribe", defaultPubsubSubscribe, "subscriber topic")
	flagPubsubSubscriberID = pflag.String("atlas.pubsub.subscriber.id", defaultPubsubSubscriberID, "subscriber id")

	flagAuthzEnable  = pflag.Bool("atlas.authz.enable", defaultAuthzEnable, "enable application with authorization")
	flagAuthzAddress = pflag.String("atlas.authz.address", defaultAuthzAddress, "address or FQDN of the authorization service")
	flagAuthzPort    = pflag.String("atlas.authz.port", defaultAuthzPort, "port of the authorization service")

	flagAuditEnable  = pflag.Bool("atlas.audit.enable", defaultAuditEnable, "enable logging of gRPC requests on Atlas audit service")
	flagAuditAddress = pflag.String("atlas.audit.address", defaultAuditAddress, "address or FQDN of Atlas audit log service")
	flagAuditPort    = pflag.String("atlas.audit.port", defaultAuditPort, "port of Atlas audit log service")

	flagTaggingEnable  = pflag.Bool("atlas.tagging.enable", defaultTaggingEnable, "enable tagging")
	flagTaggingAddress = pflag.String("atlas.tagging.address", defaultTaggingAddress, "address or FQDN of Atlas tagging service")
	flagTaggingPort    = pflag.String("atlas.tagging.port", defaultTaggingPort, "port of Atlas tagging service")

	flagInternalEnable    = pflag.Bool("internal.enable", defaultInternalEnable, "enable internal http server")
	flagInternalAddress   = pflag.String("internal.address", defaultInternalAddress, "address of internal http server")
	flagInternalPort      = pflag.String("internal.port", defaultInternalPort, "port of internal http server")
	flagInternalHealth    = pflag.String("internal.health", defaultInternalHealth, "endpoint for health checks")
	flagInternalReadiness = pflag.String("internal.readiness", defaultInternalReadiness, "endpoint for readiness checks")

	flagConfigDirectory = pflag.String("config.source", defaultConfigDirectory, "directory of the configuration file")
	flagConfigFile      = pflag.String("config.file", defaultConfigFile, "directory of the configuration file")
	flagSecretFile      = pflag.String("config.secret.file", defaultSecretFile, "directory of the secrets configuration file")
	flagApplicationID   = pflag.String("app.id", defaultApplicationID, "identifier for the application")

	flagLoggingLevel = pflag.String("logging.level", defaultLoggingLevel, "log level of application")
)
