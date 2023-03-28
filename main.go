package main

// https://bartechcz.screenconnect.com/Bin/ConnectWiseControl.ClientSetup.deb?e=Access&y=Guest 
import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func main() {
	processK2("OVY122003214750")
	processK2("OVY122003214740")
	processK2("OVY122003214730")
	processK2("OVY118000256750")
	processK2("OVY118000256750")
	processK2("OVY118000257550")
	processK2("OVY118000761750")
	processK2("OVY118001700550")
	processK2("OVY118001700650")
	processK2("OVY118001700750")
	processK2("OVY118001700950")
	processK2("OVY118001701050")
	processK2("OVY118001701150")
	processK2("OVY118001701250")
	processK2("OVY118001701350")
	processK2("OVY118001701450")
	processK2("OVY118001701550")
	processK2("OVY118001701650")
	processK2("OVY118001701750")
	processK2("OVY118001701850")
	processK2("OVY118001702150")
	processK2("OVY118001702250")
	processK2("OVY118001702350")
	processK2("OVY118001702450")
	processK2("OVY118001703150")
	processK2("ONE118000012990")
	processK2("OVY118001706150")
	processK2("OVY118001706250")
	processK2("OVY118001706460")
	processK2("OVY118001708370")
	processK2("OVY118001709250")
	processK2("OVY118001709550")
	processK2("OVY118001709650")
	processK2("OVY118001709850")
	processK2("OVY118001709960")
	processK2("OVY118001710150")
	processK2("OVY118001710250")
	processK2("OVY118001710350")
	processK2("OVY118001710450")
	processK2("OVY118001710570")
	processK2("OVY118001710650")
	processK2("OVY118001710750")
	processK2("OVY118001710850")
	processK2("OVY118001711050")
	processK2("OVY118001711150")
	processK2("OVY118001711250")

}

func processK2(barcode string) {
	fmt.Println("Processing barcode: " + barcode)
	k2Url := "http://k2-web/k2sws/Formation/Special/KP_VYR_RaiseNotifyFromXMost/PAS?Barcode=" + barcode + "&x-auth="
	fmt.Println("Prehashed url:      " + k2Url)
	hash := getAuthHeader(k2Url, "xMost", "wGTmpI6VoCep241")
	fmt.Println("K2 created hash:    " + hash)
	updatedHash := strings.ReplaceAll(hash, "+", "%2b")
	fmt.Println("Updated hash:       " + hash)
	updatedK2Url := k2Url + updatedHash
	fmt.Println("Hashed url:         " + updatedK2Url)
	res, err := http.Get(updatedK2Url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Response:           " + res.Status)
	fmt.Println("Contains plus:      " + strconv.FormatBool(strings.Contains(updatedK2Url, "+")))
	c := res.Close
	fmt.Println("Open:               " + strconv.FormatBool(c))
	fmt.Println("")
}

func getAuthHeader(source_url string, username string, password string) string {
	hmacEncoder := hmac.New(md5.New, []byte(password))
	urlDecoded, _ := url.QueryUnescape(source_url)
	upperFullURL := strings.ToUpper(urlDecoded)
	hmacEncoder.Write([]byte(upperFullURL))
	hmacHash := base64.StdEncoding.EncodeToString(hmacEncoder.Sum(nil))
	return username + ":" + hmacHash
}
