package env

import (
	"fmt"
	"os"
)

var (
	LogLevel = Getter("LOG_LEVEL", "debug")
	GinMode  = Getter("GIN_MODE", "debug")

	// ------------------------ HTTP ----------------------- //
	HTTPPort           = fmt.Sprintf(":%s", Getter("HTTP_PORT", "6060"))
	HTTPApiRoot        = Getter("HTTP_API_ROOT", "api")
	HTTPApiVersion     = Getter("HTTP_API_VERSION", "v1")
	HTTPApiServiceName = Getter("HTTP_API_SERVICE_NAME", "stock")
	HTTPApiPkgPath     = Getter("HTTP_API_VERSION", "admin")

	// ------------------------ DB ------------------------- //
	DbDriver      = Getter("DB_DRIVER", "postgres")
	DbHost        = Getter("DB_HOST", "postgres_stock")
	DbPort        = Getter("DB_PORT", "5432")
	DbUser        = Getter("DB_USER", "stock")
	DbName        = Getter("DB_NAME", "stock")
	DbPassword    = Getter("DB_PASSWORD", "stock")
	DbSslMode     = Getter("DB_SSL_MODE", "disable")
	DbSslCertPath = Getter("DB_SSL_CERT_PATH", ".env/root.crt")

	// ------------------------ JAEGER ----------------------- //
	TraceHeader = Getter("TRACE_HEADER", "x-b3-traceid")

	//// ------------------------ FOCUS ----------------------- //
	//EditorJSMediaFolderName = Getter("EDITOR_JS_MEDIA_FOLDER_NAME", "editor_js")
	//ProductMediaFolderName  = Getter("PRODUCT_MEDIA_FOLDER_NAME", "product_media")
	//StoreMediaFolderName    = Getter("STORE_MEDIA_FOLDER_NAME", "store_media")
	//PromoMediaFolderName    = Getter("PROMO_MEDIA_FOLDER_NAME", "promo_media")
	//CdnUrl                  = Getter("CDN_URL", "http://cdn.dev.aeroidea.ru")

	// ----------------------- AWS S3 ---------------------- //
	AwsEndpoint        = Getter("AWS_ENDPOINT", "")
	AwsAccessKeyID     = Getter("AWS_ACCESS_KEY_ID", "")
	AwsSecretAccessKey = Getter("AWS_SECRET_ACCESS_KEY", "")
	AwsBucket          = Getter("AWS_BUCKET", "demo")
	AwsRegion          = Getter("AWS_REGION", "")
)

// Getter -
func Getter(key, defaultValue string) string {
	env, ok := os.LookupEnv(key)
	if ok {
		return env
	}
	return defaultValue
}
