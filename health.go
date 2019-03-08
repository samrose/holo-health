package main
    import "bytes"
    import "fmt"
    import "io/ioutil"
    import "log"
    import "os"
    import "path/filepath"
    import "reflect"
    import "strconv"
    import "strings"
    import "github.com/jacobsa/go-serial/serial"
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
    b := []byte("A<")
    n, err := port.Write(b)
    if err != nil {
      log.Fatalf("port.Write: %v", err)
    }

    fmt.Println("Wrote", n, "bytes.")

}
func flash_yellow(){
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
func set_aurora(){
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
      log.Fatalf("[FATAL] serial.Open: %v", err)
    }

    // Make sure to close it later.
    defer port.Close()

    // Write 2 bytes to the port.
    b := []byte("A<")
    n, err := port.Write(b)
    if err != nil {
      log.Fatalf("[FATAL] port.Write: %v", err)
    }

    fmt.Println("Wrote", n, "bytes.")

}
func uuid() string{
    u, _ := ioutil.ReadFile("/proc/sys/kernel/random/uuid")
    //fmt.Println(string(u))
    return string(u)
}

func main(){
    fmt.Println(filepath.Match("/sys/bus/platform/devices/coretemp.0/hwmon/hwmon2/", "/sys/bus/platform/devices/coretemp.0/hwmon/hwmon2/temp*_input"))
    tempfiles := []string {
        "/sys/bus/platform/devices/coretemp.0/hwmon/hwmon2/temp2_input",
        "/sys/bus/platform/devices/coretemp.0/hwmon/hwmon2/temp3_input",
    }
    for _, each := range tempfiles {
        filerc, err := os.Open(each)
        if err != nil {
            log.Fatalf("[FATAL] read temp sensor failed with %v", err)
        }
        defer filerc.Close()
        buf := new(bytes.Buffer)
        buf.ReadFrom(filerc)

        contents := buf.String()
        fmt.Print(contents)
        contents = strings.TrimSuffix(contents, "\n")
        fmt.Println(reflect.TypeOf(contents))
        n, nerr := strconv.ParseInt(contents, 10, 64)
        if nerr == nil {
            fmt.Printf("%d of type %T", n, n)
        }
        if nerr != nil {
          log.Fatalf("[FATAL] error: %v", nerr)
        }
        uuid := uuid()
        fmt.Println(reflect.TypeOf(n))
        fmt.Println(n)
        if n > 79000 {
            flash_yellow()
            l := log.New(os.Stdout, "[WARNING] ", log.Ldate | log.Ltime)
            l.Printf("CPU temp is %s - %s", contents, uuid)
        }
        if n > 99000 {
            flash_red()
            l := log.New(os.Stdout, "[WARNING] ", log.Ldate | log.Ltime)
            l.Printf("CPU temp is %s - %s", contents, uuid)
        }
        if n < 79000 {
            set_aurora()
            l := log.New(os.Stdout, "[INFO] ", log.Ldate | log.Ltime)
            l.Printf("CPU temp is %s - %s", contents, uuid)
        }
    }
}

