package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/SVz777/cpu_emu/cpu"
)

func main() {
	c := cpu.NewTIS100("cpu0", cpu.MHz, 0)
	codes, err := os.ReadFile("./power.svz")
	if err != nil {
		log.Fatalf("read failed: %v", err)
	}
	prog, err := cpu.BuildProg(c, strings.Split(string(codes), "\n"))
	if err != nil {
		log.Fatalf("build prog failed: %v", err)
	}
	c.LoadProg(prog)
	base := cpu.NewDataBus(1)
	base.Connect(c.DataBus(cpu.LEFT))
	power := cpu.NewDataBus(1)
	power.Connect(c.DataBus(cpu.UP))
	out := cpu.NewDataBus(1)
	out.Connect(c.DataBus(cpu.RIGHT))
	go c.Run()
	base.Write(2)
	power.Write(5)
	fmt.Println(out.Read())
}
