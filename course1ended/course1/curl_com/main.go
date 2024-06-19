package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	var data = strings.NewReader(`{"audienceId":"85c59bd99a5e42b4811c198cd097d666","audienceName":"","slideId":96160636,"presentationId":4692922,"reactionType":"like"}`)
	req, err := http.NewRequest("POST", "https://audience.ahaslides.com/api/reaction/", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/116.0")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Origin", "https://audience.ahaslides.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://audience.ahaslides.com/knjb91sgis")
	req.Header.Set("Cookie", "mp_4c51563e79ee9579a57c1818c1368ec6_mixpanel=%7B%22distinct_id%22%3A%20%22189fa4023285d5-003786a17244e38-9762630-1fa400-189fa402329e8e%22%2C%22%24device_id%22%3A%20%22189fa4023285d5-003786a17244e38-9762630-1fa400-189fa402329e8e%22%2C%22%24initial_referrer%22%3A%20%22%24direct%22%2C%22%24initial_referring_domain%22%3A%20%22%24direct%22%7D; ahaFirstPage=https://audience.ahaslides.com/knjb91sgis; _gcl_au=1.1.682920079.1692120655; _ga=GA1.1.984959336.1692120655; _gid=GA1.2.903242149.1692120655; _ga_HJMZ53V9R3=GS1.1.1692120654.1.1.1692120654.60.0.0; _gat=1; _rdt_uuid=1692120654938.b0d05e41-b5fc-4bd4-a27d-b399a8e3650d; AWSALB=chUV/iPGeH0+o131QAxdhi7No2bo/qyFmo7wWIh/gIoYNAIrbDSRJqqazOe7uYcWpZLN8mMzz6rQzIGho8qlDA2fihGDtSAMzI4KjRclfwZxjyqAUZ5ZeNZZ9eXn; AWSALBCORS=chUV/iPGeH0+o131QAxdhi7No2bo/qyFmo7wWIh/gIoYNAIrbDSRJqqazOe7uYcWpZLN8mMzz6rQzIGho8qlDA2fihGDtSAMzI4KjRclfwZxjyqAUZ5ZeNZZ9eXn")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("TE", "trailers")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
