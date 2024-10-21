package constants

import(
	"os"

	"github.com/joho/godotenv"
)
const (
	Pages = "privacy-policy|contact-us|terms-and-conditions|about-us|site-content|disclaimer"
	OpensearchURL = "https://vpc-phishbug-online-furc4jflhb4zfq2zxmacth4bdm.us-west-1.es.amazonaws.com"
	IndexOpenSearch = "phish-bug"
)

//Get Template Path
func GetTemplatePath() string{

	err := godotenv.Load("/home/ec2-user/.env")
	
	if err != nil {
		return ""
	}

	return os.Getenv("TEMPLATE_PATH")
}