package APNs

import (
	"fmt"
	"github.com/sideshow/apns2/certificate"
	"github.com/sideshow/apns2/payload"
	"github.com/sideshow/apns2/token"
	"log"

	"github.com/sideshow/apns2"
)

func main() {

	processP8()

}

func processP12() {
	cert, err := certificate.FromP12File("../cert.p12", "")
	if err != nil {
		log.Fatal("Cert Error:", err)
	}

	notification := &apns2.Notification{}
	notification.DeviceToken = "11aa01229f15f0f0c52029d8cf8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7"
	notification.Topic = "com.sideshow.Apns2"
	notification.Payload = []byte(`{"aps":{"alert":"Hello!"}}`) // See Payload section below

	// If you want to test push notifications for builds running directly from XCode (Development), use
	// client := apns2.NewClient(cert).Development()
	// For apps published to the app store or installed as an ad-hoc distribution use Production()

	client := apns2.NewClient(cert).Production()
	res, err := client.Push(notification)

	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}

func processP8() {
	authKey, err := token.AuthKeyFromFile("../AuthKey_XXX.p8")
	if err != nil {
		log.Fatal("token error:", err)
	}

	token := &token.Token{
		AuthKey: authKey,
		// KeyID from developer account (Certificates, Identifiers & Profiles -> Keys)
		KeyID: "ABC123DEFG",
		// TeamID from developer account (View Account -> Membership)
		TeamID: "DEF123GHIJ",
	}

	notification := &apns2.Notification{}
	notification.DeviceToken = "11aa01229f15f0f0c52029d8cf8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7"
	notification.Topic = "com.sideshow.Apns2"
	notification.Payload = []byte(`{"aps":{"alert":"Hello!"}}`) // See Payload section below

	client := apns2.NewTokenClient(token)
	res, err := client.Push(notification)

	if err != nil {
		log.Fatal("Error:", err)
	}

	// {"aps":{"alert":"hello","badge":1},"key":"val"}

	payload := payload.NewPayload().Alert("hello").Badge(1).Custom("key", "val")

	notification.Payload = payload
	res, err = client.Push(notification)

	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}
