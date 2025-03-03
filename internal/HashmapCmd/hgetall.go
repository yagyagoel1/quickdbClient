package hashmapcmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
	"github.com/yagyagoel1/quickdbClient/utils"
)

func getAllValues(args []string) {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("err while connecting", err)

	}
	defer conn.Close()
	data := fmt.Sprintf("*2\r\n$7\r\nHGETALL\r\n$%d\r\n%s\r\n", len(args[0]), args[0])
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

var HgetAllCmd = &cobra.Command{
	Use:   "HGETALL",
	Short: "It is used to get a value in a hashmap",
	Run: func(cmd *cobra.Command, args []string) {
		getAllValues(args)

	},
}
