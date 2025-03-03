package hashmapcmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
	"github.com/yagyagoel1/quickdbClient/utils"
)

func setInMap(args []string) {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("err while connecting", err)

	}
	if len(args) < 3 {
		fmt.Println("The format is not correct expected THREE args Ex: HGET map key value")
		return
	}
	defer conn.Close()
	data := fmt.Sprintf("*4\r\n$4\r\nHSET\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n", len(args[0]), args[0], len(args[1]), args[1], len(args[2]), args[2])
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

var HsetCmd = &cobra.Command{
	Use:   "HSET",
	Short: "It is used to set a value in a hashmap",
	Run: func(cmd *cobra.Command, args []string) {
		setInMap(args)

	},
}
