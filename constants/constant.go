package constants

import(
	"os"
	"strings"
	"webapp/types"

	"github.com/joho/godotenv"
)
const (
	Pages = "privacy-policy|contact-us|terms-and-conditions|about-us|site-content|disclaimer"
	OpensearchURL = "https://vpc-phishbug-online-furc4jflhb4zfq2zxmacth4bdm.us-west-1.es.amazonaws.com"
	IndexOpenSearch = "phish-bug"
	DATE_FORMAT = "January 2, 2006"
)

func GetAuthor() types.Author{
	return types.Author{Link: "Kunal"}
}

//Get Template Path
func GetTemplatePath() string{

	err := godotenv.Load("/home/ec2-user/.env")
	
	if err != nil {
		return ""
	}

	return os.Getenv("TEMPLATE_PATH")
}


//Get Template Path
func GetOpenSearchAddress() []string{

	err := godotenv.Load("/home/ec2-user/.env")
	
	if err != nil {
		return []string{}
	}

	urls := os.Getenv("OPEN_SEARCH_ADDRESS")

	return strings.Split(urls, ",")
}

//Get Template Path
func GetENVKey(key string) string{

	err := godotenv.Load("/home/ec2-user/.env")
	
	if err != nil {
		return ""
	}

	return os.Getenv(key)
}