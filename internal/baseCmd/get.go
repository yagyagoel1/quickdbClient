package basecmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
	"github.com/yagyagoel1/quickdbClient/utils"
)

func getValue(args []string) {

	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("err while connecting", err)
	}
	defer conn.Close()
	data := []byte(fmt.Sprintf("*2\r\n$3\r\nGET\r\n$%d\r\n%s\r\n", len(args[0]), args[0]))
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("error while writing ", err)

	}
	resp := utils.NewResp(conn)
	value, err := resp.Read()
	if err != nil {
		fmt.Println("error while reading the response", err)

	}
	utils.PrintOutput(value)

}

var GetCmd = &cobra.Command{
	Use:   "GET",
	Short: "it is used to get the value against the key",
	Run: func(cmd *cobra.Command, args []string) {
		getValue(args)
	},
}
