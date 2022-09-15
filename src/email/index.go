package email

import (
	"fmt"
	"log"
	"os"

	//   "os"
	"github.com/joho/godotenv"
	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

type EmailConfigs struct {
	Key         string `mapstructure:"Key"`
	Secret      string `mapstructure:"Secret"`
	Customid    string `mapstructure:"Customid"`
	Owner       string `mapstructure:"Owner"`
	OwnerEmail  string `mapstructure:"OwnerEmail"`
	WebsiteLink string `mapstructure:"WebsiteLink"`
	Phone       string `mapstructure:"Phone"`
}

var apikey string = "e3666c5ac12e90ce7826cd2382f767f2"
var secretkey string = "479a0020bef2201427f63e146b3e33e8"

// func LoaddbConfig() (e EmailConfigs, err error) {
// 	viper.AddConfigPath(".")
// 	viper.SetConfigName("app")
// 	viper.SetConfigType("env")

// 	viper.AutomaticEnv()

// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		return
// 	}
// 	err = viper.Unmarshal(&e)
// 	return
// }

func Emailpay(name, customeremail, code string, amount string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file in routes", err)
	}
	// key := os.Getenv("EncryptionKey")
	// secret := os.Getenv("Secret")
	Customid := os.Getenv("Customid")
	Owner := os.Getenv("Owner")
	OwnerEmail := os.Getenv("OwnerEmail")
	WebsiteLink := os.Getenv("WebsiteLink")
	Phone := os.Getenv("Phone")
	mailjetClient := mailjet.NewMailjetClient(apikey, secretkey)
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: OwnerEmail,
				Name:  Owner,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: customeremail,
					Name:  name,
				},
			},
			Subject:  "Your Order at Nillavee was Successful",
			TextPart: "thank you forshoppign with us",
			HTMLPart: "<h3>Hello  " + name + "<br /> Thank you for shopping with us!</h3><br />Your Order of Ksh" + amount + " is being processed <br />We'll respond shortly.<br><br>Thanks!<br>" + Owner + "<br>Phone:" + Phone + "</br>Website: <a href='" + WebsiteLink + "'>" + WebsiteLink + "</a>",
			CustomID: Customid,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	//compile message to the shop owner
	// resa := []string{}
	// for _, m := range payment.Cart {
	// 	res := "<p>wedo: "+m.Name+" \n Specs"+m.Specs+"\n Decor"+m.Decor+"\n wedo Message"+m.Message+"\n Quantity"+fmt.Sprintf("%d", m.Quantity )+" \n Price"+fmt.Sprintf("%d", m.Price )+"</p> \n"
	// 	resa = append(resa, res)
	// }
	messagesInfo1 := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: OwnerEmail,
				Name:  Owner,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: OwnerEmail,
					Name:  Owner,
				},
			},
			Subject:  "Purchases made",
			TextPart: "An order has being placed",
			HTMLPart: "<h3>Hello " + Owner + "</h3><br /> You have an order of Ksh" + amount + " from " + name + "<br />Please handle it!<br />Ooh and have a lovely day! <a href=\"" + WebsiteLink + "/orders/show1/" + code + "\">This is the order content</a>",
			CustomID: Customid,
		},
	}
	messages1 := mailjet.MessagesV31{Info: messagesInfo1}
	// fmt.Println(payment)
	if customeremail != "" {
		_, err := mailjetClient.SendMailV31(&messages)
		if err != nil {
			log.Fatal(err)
		}
		_, err = mailjetClient.SendMailV31(&messages1)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("Data: %+v\n", res, res1)
		return
	}
	res1, err := mailjetClient.SendMailV31(&messages1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res1)
}
func Emailing() {
	// Key := os.Getenv("EncryptionKey")
	// Secret := os.Getenv("Secret")
	Customid := os.Getenv("Customid")
	Owner := os.Getenv("Owner")
	OwnerEmail := os.Getenv("OwnerEmail")
	// WebsiteLink := os.Getenv("WebsiteLink")
	// Phone := os.Getenv("Phone")
	mailjetClient := mailjet.NewMailjetClient(apikey, secretkey)
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: OwnerEmail,
				Name:  Owner,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: OwnerEmail,
					Name:  Owner,
				},
			},
			Subject:  "Hellos",
			TextPart: "thank you forshoppign with us",
			HTMLPart: "<h3>Hello Thank you fro shopping with us <a href='https://www.chantosweb.com/'>Wedding gowns</a>!</h3><br />Have a lovevely day!",
			CustomID: Customid,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}

func EmailingPassword(pass, email string) {
	// Key := os.Getenv("EncryptionKey")
	// Secret := os.Getenv("Secret")
	// Customid := os.Getenv("Customid")
	Owner := os.Getenv("Owner")
	OwnerEmail := os.Getenv("OwnerEmail")
	// WebsiteLink := os.Getenv("WebsiteLink")
	// Phone := os.Getenv("Phone")
	mailjetClient := mailjet.NewMailjetClient(apikey, secretkey)
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: OwnerEmail,
				Name:  Owner,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: email,
					Name:  "",
				},
			},
			Subject:  "Your password has being changed!",
			TextPart: "your password",
			HTMLPart: "<h3>your new password is " + pass + "</h3>!",
			CustomID: email,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}
