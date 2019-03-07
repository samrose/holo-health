package main
    import "fmt"
    import "log"
    import "os/exec"
    //import "os"
    //import "time"
    import "github.com/jacobsa/go-serial/serial"
    //import "github.com/ssimunic/gosensors"
    //import "github.com/kolide/osquery-go"

func read_sensors(){
	out, err := exec.Command("cat /sys/bus/platform/devices/coretemp.0/hwmon/hwmon2/temp*_input").Output()
	if err != nil {
		log.Fatalf("read sensor failed with %v", err)
	}

	s := string(out)
    fmt.Println(s)
	//return s, nil
}

func flash_red(){
    // Set up options.
    options := serial.OpenOptions{
      PortName: "/dev/ttyUSB0",
      BaudRate: 19200,
      DataBits: 8,
      StopBits: 1,
      MinimumReadSize: 4,
    }

    // Open the port.
    port, err := serial.Open(options)
    if err != nil {
      log.Fatalf("serial.Open: %v", err)
    }

    // Make sure to close it later.
    defer port.Close()

    // Write 2 bytes to the port.
    b := []byte("P*")
    n, err := port.Write(b)
    if err != nil {
      log.Fatalf("port.Write: %v", err)
    }

    fmt.Println("Wrote", n, "bytes.")

}
func main(){
    read_sensors()
    flash_red()
}

