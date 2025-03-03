package basecmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
	"github.com/yagyagoel1/quickdbClient/utils"
)

func sendPing() {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("err while establishing the connection", err)

	}
	defer conn.Close()
	data := "*1\r\n$4\r\nPING\r\n"
	_, err = conn.Write([]byte(data))
	if err != nil {
		fmt.Println("err while writing to the connection", err)
	}
	resp := utils.NewResp(conn)
	value, err := resp.Read()
	if err != nil {
		fmt.Println("error while reading the response", err)

	}
	utils.PrintOutput(value)

}

var PingCmd = &cobra.Command{
	Use:   "PING",
	Short: "it is used to check whether the server is active or no.",
	Run: func(cmd *cobra.Command, args []string) {
		sendPing()
	},
}
