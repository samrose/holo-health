package main
    import "fmt"
    import "log"
    import "os"
    import "time"
    import "github.com/jacobsa/go-serial/serial"
    import "github.com/ssimunic/gosensors"
    import "github.com/kolide/osquery-go"

func read_sensors(){
	sensors, err := gosensors.NewFromSystem()
	// sensors, err := gosensors.NewFromFile("/path/to/log.txt")

	if err != nil {
		panic(err)
	}

	// Sensors implements Stringer interface,
	// so code below will print out JSON
	fmt.Println(sensors)

	// Also valid
	// fmt.Println("JSON:", sensors.JSON())

	// Iterate over chips
	for chip := range sensors.Chips {
		// Iterate over entries
		for key, value := range sensors.Chips[chip] {
			// If CPU or GPU, print out
			if key == "CPU" || key == "GPU" {
				fmt.Println(key, value)
			}
		}
	}

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

