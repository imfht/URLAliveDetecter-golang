package URLAliveDetecter

import "testing"

func Test_doRequest(t *testing.T) {
	doRequest("http://www.baidu.com")
}

func TestWorker(t *testing.T) {
	var urls = []string{
	}
	Worker(urls)
}

func TestWorkFile(t *testing.T) {
	WorkFile("urls.txt")
}
