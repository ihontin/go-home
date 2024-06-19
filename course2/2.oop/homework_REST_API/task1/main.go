package main

//type MyClient struct {
//	client *http.Client
//}
//
//func NewMyClient(client *http.Client) *MyClient {
//	return &MyClient{
//		client: client,
//	}
//}
//
//func (c *MyClient) doRequest(url string) ([]byte, error) {
//	resp, err := c.client.Get(url)
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	return body, nil
//}
//
//func (c *MyClient) GetUser() (User, error) {
//	b, err := c.doRequest("https://someapi.com/users")
//	if err != nil {
//		return User{}, err
//	}
//	var user User
//	err = json.Unmarshal(b, &user)
//	if err != nil {
//		return User{}, err
//	}
//
//	return user, nil
//}
//
//type User struct {
//	Name string `json:"name"`
//	Age  int    `json:"age"`
//}

//---------------------------------------------------------------------

//type MyHTTPClient struct {
//	client *http.Client
//}
//
//func NewMyHTTPClient(client *http.Client) *MyHTTPClient {
//	return &MyHTTPClient{
//		client: client,
//	}
//}
//
//func (c *MyHTTPClient) DoRequest(url string) (*http.Response, error) {
//	response, err := c.client.Get(url)
//	if err != nil {
//		return nil, err
//	}
//	return response, nil
//}
//
//func main() {
//	client := &http.Client{}
//	myClient := NewMyHTTPClient(client)
//
//	url := "https://httpbin.org/get"
//	response, err := myClient.DoRequest(url)
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//	defer response.Body.Close()
//
//	fmt.Println("Response Status:", response.Status)
//
//	url2 := "https://www.google.com"
//	response2, err := myClient.DoRequest(url2)
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//	defer response2.Body.Close()
//
//	fmt.Println("Response Status:", response2.Status)
//}

//---------------------------------------------------------

//type MyHTTPClient struct {
//	client *http.Client
//}
//
//func NewMyHTTPClient(client *http.Client) *MyHTTPClient {
//	return &MyHTTPClient{
//		client: client,
//	}
//}
//
//func (c *MyHTTPClient) DoRequest(url string) (*http.Response, error) {
//	response, err := c.client.Get(url)
//	if err != nil {
//		return nil, err
//	}
//	return response, nil
//}
//
//func main() {
//	// Создаем фейковый сервер для тестирования
//	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintln(w, "Hello, World!")
//	}))
//	defer server.Close()
//
//	// Создаем мок-объект для http.Client
//	mockClient := &http.Client{
//		Transport: server.Transport,
//	}
//
//	myClient := NewMyHTTPClient(mockClient)
//
//	url := server.URL
//	response, err := myClient.DoRequest(url)
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//	defer response.Body.Close()
//
//	fmt.Println("Response Status:", response.Status)
//}
