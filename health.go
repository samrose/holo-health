package main
    import "bytes"
    import "fmt"
    import "io/ioutil"
    import "log"
    import "os"
    import "path/filepath"
    //import "reflect"
    import "strconv"
    import "strings"
    //import "time"
    import "github.com/jacobsa/go-serial/serial"
    //import "github.com/kolide/osquery-go"
    //import "github.com/coreos/go-systemd/dbus"

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
func flash_purple(){
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
    pattern := "/sys/bus/platform/devices/coretemp.0/hwmon/hwmon2/temp*_input"
    matches, merr := filepath.Glob(pattern)
    if merr != nil {
        log.Fatalf("[FATAL] read temp sensor failed with %v", merr)
    }

    //fmt.Println(matches)
    //tempfiles := []string { matches }
    for _, each := range matches {
        filerc, err := os.Open(each)
        if err != nil {
            log.Fatalf("[FATAL] read temp sensor failed with %v", err)
        }
        defer filerc.Close()
        buf := new(bytes.Buffer)
        buf.ReadFrom(filerc)

        contents := buf.String()
        //fmt.Print(contents)
        contents = strings.TrimSuffix(contents, "\n")
        //fmt.Println(reflect.TypeOf(contents))
        n, nerr := strconv.ParseInt(contents, 10, 64)
        if nerr == nil {
            fmt.Printf("%d of type %T", n, n)
        }
        if nerr != nil {
          log.Fatalf("[FATAL] error: %v", nerr)
        }
        uuid := uuid()
        //fmt.Println(reflect.TypeOf(n))
        //fmt.Println(n)
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
    enpattern := "/sys/class/net/en*/operstate"
    enmatches, enmerr := filepath.Glob(enpattern)
    if enmerr != nil {
        log.Fatalf("[FATAL] read operstate failed with %v", enmerr)
    }
    for _, each := range enmatches {
        enfilerc, ferr := os.Open(each)
        if ferr != nil {
            log.Fatalf("[FATAL] read operstate failed with %v", ferr)
        }
        defer enfilerc.Close()
        enbuf := new(bytes.Buffer)
        enbuf.ReadFrom(enfilerc)

        encontents := enbuf.String()
        enuuid := uuid()
        fmt.Print(encontents)
        if strings.TrimRight(encontents, "\n") == "down" {
            flash_purple()
            l := log.New(os.Stdout, "[WARNING] ", log.Ldate | log.Ltime)
            l.Printf("Network is %s - %s", encontents, enuuid)
        }
        if strings.TrimRight(encontents, "\n") == "up" {
            set_aurora()
            l := log.New(os.Stdout, "[INFO] ", log.Ldate | log.Ltime)
            l.Printf("Network is %s - %s", encontents, enuuid)
        }

    }
}

