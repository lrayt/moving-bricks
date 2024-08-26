package transfer

import "testing"

func TestServer(t *testing.T) {
	svr := NewServer(":8080")
	if err1 := svr.Run(); err1 != nil {
		t.Errorf("run err:%s", err1.Error())
	}
}

func TestClient(t *testing.T) {
	client, err := NewClient("127.0.0.1:8080")
	if err != nil {
		t.Errorf("new client err:%s", err.Error())
		return
	}
	//res, err := client.Ping()
	//t.Log(res, err)
	err = client.Transfer()
	t.Log(err)
}

func TestClient2(t *testing.T) {
	client, err := NewClient("127.0.0.1:8080")
	if err != nil {
		t.Errorf("new client err:%s", err.Error())
		return
	}
	//res, err := client.Ping()
	//t.Log(res, err)
	err = client.Transfer()
	t.Log(err)
}
