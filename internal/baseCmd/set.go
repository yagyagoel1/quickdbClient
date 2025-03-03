package basecmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
	"github.com/yagyagoel1/quickdbClient/utils"
)

func Set(args []string) {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("error connecting to the server", err)
		return
	}
	if len(args) < 2 {
		fmt.Println("error the command at least require two arguments Ex: SET KEY 'VALUE' ")
		return
	}
	defer conn.Close()
	fmt.Println(args[1])
	data := []byte(fmt.Sprintf("*3\r\n$3\r\nSET\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n", len(args[0]), args[0], len(args[1]), args[1]))
	fmt.Println("da", data)
	log, err := conn.Write(data)
	fmt.Println("log", log)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	resp := utils.NewResp(conn)
	value, err := resp.Read()
	if err != nil {
		fmt.Println("error while reading the response", err)

	}
	utils.PrintOutput(value)
}

var SetCmd = &cobra.Command{
	Use:   "SET",
	Short: "Set the value against a key",
	Run: func(cmd *cobra.Command, args []string) {
		Set(args)
	},
}
