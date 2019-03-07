package main
    import "bytes"
    import "fmt"
    import "log"
    //import "os/exec"
    import "os"
    //import "time"
    import "github.com/jacobsa/go-serial/serial"
    //import "github.com/ssimunic/gosensors"
    //import "github.com/kolide/osquery-go"

func read_sensors() int{
	//out, err := exec.Command("cat /sys/bus/platform/devices/coretemp.0/hwmon/hwmon2/temp2_input").Output()
	filerc, err := os.Open("/sys/bus/platform/devices/coretemp.0/hwmon/hwmon2/temp2_input")
	if err != nil {
		log.Fatalf("read sensor failed with %v", err)
	}
    defer filerc.Close()
    buf := new(bytes.Buffer)
    contents := buf.ReadFrom(filerc)
    //contents := buf.ReadVarint()

    fmt.Print(contents)
    return contents

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
    b := []byte("R*")
    n, err := port.Write(b)
    if err != nil {
      log.Fatalf("port.Write: %v", err)
    }

    fmt.Println("Wrote", n, "bytes.")

}
func main(){
    read := read_sensors()
    if read > 24000 {
        flash_red()
    }
}

