package client

import (
	"context"

	pb "github.com/cloudingcity/grpc-chat/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Connect(addr string, username, password string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	log.Infof("UserName: %s", username)
	log.Infof("Password: %s", password)
	log.Infof("Connect to server: %s", addr)

	c := &Client{
		client: pb.NewChatClient(conn),
	}
	token, err := c.Connect(username, password)
	if err != nil {
		log.Fatalln(err)
	}
	log.Infof("Token: %s", token)
}

type Client struct {
	client pb.ChatClient
}

func (c *Client) Connect(username, password string) (string, error) {
	req := &pb.ConnectRequest{
		Username: username,
		Password: password,
	}
	resp, err := c.client.Connect(context.Background(), req)
	if err != nil {
		return "", err
	}
	return resp.Token, nil
}
