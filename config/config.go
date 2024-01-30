package config

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func Init() {
	fmt.Println("config.init()")
}

func Loadconfig(path string) IConfig {
	envMap, err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("load env error: %v", err)
	}
	return &config{
		app: &app{
			host: envMap["APP_HOST"],
			port: func() int {
				port, err := strconv.Atoi(envMap["APP_PORT"])
				if err != nil {
					log.Fatalf("load port error: %v", err)
				}
				return port
			}(),
			name:    envMap["APP_NAME"],
			version: envMap["APP_VERSION"],
			readTimeout: func() time.Duration {
				r, err := strconv.Atoi(envMap["APP_READ_TIMEOUT"])
				if err != nil {
					log.Fatalf("load readTimeout error: %v", err)
				}
				return time.Duration(int64(r) * int64(math.Pow10(9)))
			}(),
			writeTimeout: func() time.Duration {
				w, err := strconv.Atoi(envMap["APP_WRITE_TIMEOUT"])
				if err != nil {
					log.Fatalf("load writeTimeout error: %v", err)
				}
				return time.Duration(int64(w) * int64(math.Pow10(9)))
			}(),
			bodyLimit: func() int {
				b, err := strconv.Atoi(envMap["APP_BODY_LIMIT"])
				if err != nil {
					log.Fatalf("load bodyLimit error: %v", err)
				}
				return b
			}(),
			fileLimit: func() int {
				f, err := strconv.Atoi(envMap["APP_FILE_LIMIT"])
				if err != nil {
					log.Fatalf("load fileLimit error: %v", err)
				}
				return f
			}(),
			gcpbucket: envMap["APP_GCP_BUCKET"],
		},
		db: &db{
			host: envMap["DB_HOST"],
			port: func() int {
				port, err := strconv.Atoi(envMap["DB_PORT"])
				if err != nil {
					log.Fatalf("load port error: %v", err)
				}
				return port
			}(),
			protocol: envMap["DB_PROTOCOL"],
			username: envMap["DB_USERNAME"],
			password: envMap["DB_PASSWORD"],
			database: envMap["DB_DATABASE"],
			sslMode:  envMap["DB_SSL_MODE"],
			maxConnections: func() int {
				mc, err := strconv.Atoi(envMap["DB_MAX_CONNECTIONS"])
				if err != nil {
					log.Fatalf("load maxConnections error: %v", err)
				}
				return mc
			}(),
		},
		jwt: &jwt{
			adminKey:  envMap["JWT_ADMIN_KEY"],
			secretKey: envMap["JWT_SECRET_KEY"],
			apiKey:    envMap["JWT_API_KEY"],
			accessExpiresAt: func() int {
				aea, err := strconv.Atoi(envMap["JWT_ACCESS_EXPIRES"])
				if err != nil {
					log.Fatalf("load accessExpiresAt error: %v", err)
				}
				return aea
			}(),
			refreshExpiresAt: func() int {
				rea, err := strconv.Atoi(envMap["JWT_REFRESH_EXPIRES"])
				if err != nil {
					log.Fatalf("load refreshExpiresAt error: %v", err)
				}
				return rea
			}(),
		},
	}
}

type IConfig interface {
	App() IAppConfig
	Db() IDbConfig
	Jwt() IJwtConfig
}

type config struct {
	app *app
	db  *db
	jwt *jwt
}

type IAppConfig interface {
	// host : port
	Url() string
	Name() string
	Version() string
	ReadTimeout() time.Duration
	WriteTimeout() time.Duration
	BodyLimit() int
	FileLimit() int
	GCPBucket() string
}

type app struct {
	host         string
	port         int
	name         string
	version      string
	readTimeout  time.Duration
	writeTimeout time.Duration
	bodyLimit    int
	fileLimit    int
	gcpbucket    string
}

func (c *config) App() IAppConfig {
	return c.app
}
func (a *app) Url() string                 { return fmt.Sprintf("%s:%d", a.host, a.port) }
func (a *app) Name() string                { return a.name }
func (a *app) Version() string             { return a.version }
func (a *app) ReadTimeout() time.Duration  { return a.readTimeout }
func (a *app) WriteTimeout() time.Duration { return a.writeTimeout }
func (a *app) BodyLimit() int              { return a.bodyLimit }
func (a *app) FileLimit() int              { return a.fileLimit }
func (a *app) GCPBucket() string           { return a.gcpbucket }

type IDbConfig interface {
	Url() string
	MaxOpenConns() int
}

type db struct {
	host           string
	port           int
	protocol       string
	username       string
	password       string
	database       string
	sslMode        string
	maxConnections int
}

func (c *config) Db() IDbConfig {
	return c.db
}

func (d *db) Url() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.host, d.port, d.username, d.password, d.database, d.sslMode)
}
func (d *db) MaxOpenConns() int { return d.maxConnections }

type IJwtConfig interface {
	AdminKey() []byte
	SecretKey() []byte
	ApiKey() []byte
	AccessExpiresAt() int
	RefreshExpiresAt() int
	SetJwtAccessExpires(t int)
	SetJwtRefreshExpires(t int)
}

type jwt struct {
	adminKey         string
	secretKey        string
	apiKey           string
	accessExpiresAt  int
	refreshExpiresAt int
}

func (c *config) Jwt() IJwtConfig {
	return c.jwt
}

func (j *jwt) AdminKey() []byte           { return []byte(j.adminKey) }
func (j *jwt) SecretKey() []byte          { return []byte(j.secretKey) }
func (j *jwt) ApiKey() []byte             { return []byte(j.apiKey) }
func (j *jwt) AccessExpiresAt() int       { return j.accessExpiresAt }
func (j *jwt) RefreshExpiresAt() int      { return j.refreshExpiresAt }
func (j *jwt) SetJwtAccessExpires(t int)  { j.accessExpiresAt = t }
func (j *jwt) SetJwtRefreshExpires(t int) { j.refreshExpiresAt = t }
