package main
    import "bytes"
    import "fmt"
    import "log"
    //import "os/exec"
    import "os"
    import "reflect"
    //import "strconv"
    //import "time"
    import "github.com/jacobsa/go-serial/serial"
    //import "github.com/ssimunic/gosensors"
    //import "github.com/kolide/osquery-go"


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
    b := []byte("Y*")
    n, err := port.Write(b)
    if err != nil {
      log.Fatalf("port.Write: %v", err)
    }

    fmt.Println("Wrote", n, "bytes.")

}
func main(){
	//out, err := exec.Command("cat /sys/bus/platform/devices/coretemp.0/hwmon/hwmon2/temp2_input").Output()
	filerc, err := os.Open("/sys/bus/platform/devices/coretemp.0/hwmon/hwmon2/temp2_input")
	if err != nil {
		log.Fatalf("read temp sensor failed with %v", err)
	}
    defer filerc.Close()
    buf := new(bytes.Buffer)
    buf.ReadFrom(filerc)

    contents := buf.String()

    fmt.Print(contents)
    fmt.Println(reflect.TypeOf(contents))
    if 25000 > 24000 {
        flash_red()
    }
}

