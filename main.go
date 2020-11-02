/*
* Author: Shakeeb Shams
* Version: 1.2
*/
package main
import ("net"; "net/url"; "strconv"; "strings"; "time"; "crypto/tls"; "fmt"; "io/ioutil"; "math"; "os"; "sort")

/*
* The main() function is the driver function of the program
*/
func main() {
	args := os.Args[1:]
	if (args[0] == "--profile" && len(args) == 3) {
		u := urlChecker(args[1])
		count, err := strconv.Atoi(args[2])
		if (err != nil || count <= 0) {invalid("")}
		getProfiling(u, count)
	} else if (len(args) == 0 || args[0] == "--help") {
		helpReq()
	} else if (args[0] == "--url") {
		u := urlChecker(args[1])
		resp, _ := getResponse(u)
		fmt.Println(resp)
	} else {invalid("")}
}

/*
* Helper function to notify user of invalid input or errors within program
*/
func invalid(passedError string) {
	if passedError != "" {println(passedError)} else {println("Invalid inputs, please try again.")}
	os.Exit(0)
}

/*
* Function to prompt user of various usages of the program
*/
func helpReq() {
	println("Usage:\n")
	println("\tgo run . --url <URL>\t\t\t\tMake an HTTP request to the specified URL")
	println("\tgo run . --profile <URL> <Number of requests>\tProfile specified number of requests to the specified URL")
	println("\tgo run . --help\t\t\t\t\tPrints this Usage page")
}

/*
* URL sanity checker
* Return: sanitized URL
*/
func urlChecker(passedurl string) *url.URL {
	if passedurl[:4] != "http" {passedurl = "https://" + passedurl}
	u, err := url.Parse(passedurl)
	if err != nil {invalid("")}
	if u.Path == "" {u.Path = "/"}
	return u
}

/*
* The getResponse() function creates a socket and connects to the address 
* Return: response and status code
*/
func getResponse(u *url.URL) (string, int) {
	timeout, _ := time.ParseDuration("10s")
	conn, err := tls.DialWithDialer(&(net.Dialer{Timeout: timeout}), "tcp", u.Hostname()+":https", nil)
	if err != nil {invalid("TCP Dialup failed - " + err.Error())}
	rt := fmt.Sprintf("GET %v HTTP/1.0\r\n", u.Path)
	rt += fmt.Sprintf("Host: %v\r\n", u.Hostname())
	rt += fmt.Sprintf("\r\n")

	defer conn.Close()
	_, err = conn.Write([]byte(rt))
	if err != nil {invalid("Connection failed - " + err.Error())}
	resp, err := ioutil.ReadAll(conn)
	if err != nil {invalid("Failed to get response - " + err.Error())}

	status, _ := strconv.Atoi(strings.Split(string(resp), " ")[1])
	return string(resp), status
}

/*
* Function to calculate profiling stats of a website
*/
func getProfiling(u *url.URL, count int) {
	var totalTimeArr []int
	var errors []int
	maxSize := float64(0)
	minSize := math.MaxFloat64
	sumTotalTime := 0
	for i := 1; i <= count; i++ {
		startTime := time.Now()
		resp, status := getResponse(u)
		totalTime := int(time.Since(startTime).Milliseconds())
		sumTotalTime += totalTime
		totalTimeArr = append(totalTimeArr, totalTime)
		if status != 200 {errors = append(errors, status)}
		maxSize = math.Max(maxSize, float64(len(resp)))
		minSize = math.Min(minSize, float64(len(resp)))
	}
	sort.Ints(totalTimeArr)
	println("PROFILE OF ", strings.ToLower(u.Hostname()), "\n")
	println("NUMBER OF REQUESTS: ", count)
	println("FASTEST TIME: ", totalTimeArr[count-1], "ms")
	println("SLOWEST TIME: ", totalTimeArr[0], "ms")
	println("MEAN TIME: ", int(float64(sumTotalTime)/float64(count)), "ms")
	println("MEDIAN TIME: ", totalTimeArr[count/2], "ms")
	println("PERCENTAGE OF SUCCESSFULL REQUESTS: ", int(((float64(count)-float64(len(errors)))*100.0)/float64(count)), "%")
	fmt.Printf("ERROR CODES: ")
	for _, value := range errors {fmt.Printf("%d ", value)}
	println("")
	println("SIZE OF SMALLEST RESPONSE (bytes): ", int(minSize))
	println("SIZE OF LARGEST RESPONSE (bytes): ", int(maxSize))
}