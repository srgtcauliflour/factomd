package engine

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	. "github.com/FactomProject/factomd/common/globals"
	"github.com/FactomProject/factomd/common/primitives"
	"github.com/FactomProject/factomd/elections"
)

func ParseCmdLine(args []string) *FactomParams {
	p := &Params // Global copy of decoded Params global.Params
	params := flag.NewFlagSet("factomd_params", flag.ExitOnError)

	ackBalanceHashPtr := params.Bool("balancehash", true, "If false, then don't pass around balance hashes")
	enablenetPtr := params.Bool("enablenet", true, "Enable or disable networking")
	waitEntriesPtr := params.Bool("waitentries", false, "Wait for Entries to be validated prior to execution of messages")
	listenToPtr := params.Int("node", 0, "Node Number the simulator will set as the focus")
	cntPtr := params.Int("count", 1, "The number of nodes to generate")
	netPtr := params.String("net", "tree", "The default algorithm to build the network connections")
	fnetPtr := params.String("fnet", "", "Read the given file to build the network connections")
	dropPtr := params.Int("drop", 0, "Number of messages to drop out of every thousand")
	journalPtr := params.String("journal", "", "Rerun a Journal of messages")
	journalingPtr := params.Bool("journaling", false, "Write a journal of all messages received. Default is off.")
	followerPtr := params.Bool("follower", false, "If true, force node to be a follower.  Only used when replaying a journal.")
	leaderPtr := params.Bool("leader", true, "If true, force node to be a leader.  Only used when replaying a journal.")
	dbPtr := params.String("db", "", "Override the Database in the Config file and use this Database implementation. Options Map, LDB, or Bolt")
	cloneDBPtr := params.String("clonedb", "", "Override the main node and use this database for the clones in a Network.")
	networkNamePtr := params.String("network", "", "Network to join: MAIN, TEST or LOCAL")
	peersPtr := params.String("peers", "", "Array of peer addresses. ")
	blkTimePtr := params.Int("blktime", 0, "Seconds per block.  Production is 600.")
	// TODO: Old fault mechanism -- remove
	//	faultTimeoutPtr := params.Int("faulttimeout", 99999, "Seconds before considering Federated servers at-fault. Default is 30.")
	runtimeLogPtr := params.Bool("runtimeLog", false, "If true, maintain runtime logs of messages passed.")
	exclusivePtr := params.Bool("exclusive", false, "If true, we only dial out to special/trusted peers.")
	PrefixNodePtr := params.String("Prefix", "", "Prefix the Factom Node Names with this value; used to create leaderless networks.")
	RotatePtr := params.Bool("Rotate", false, "If true, responsibility is owned by one leader, and Rotated over the leaders.")
	TimeOffsetPtr := params.Int("timedelta", 0, "Maximum timeDelta in milliseconds to offset each node.  Simulates deltas in system clocks over a network.")
	KeepMismatchPtr := params.Bool("keepmismatch", false, "If true, do not discard DBStates even when a majority of DBSignatures have a different hash")
	startDelayPtr := params.Int("startdelay", 10, "Delay to start processing messages, in seconds")
	DeadlinePtr := params.Int("Deadline", 1000, "Timeout Delay in milliseconds used on Reads and Writes to the network comm")
	CustomNetPtr := params.String("customnet", "", "This string specifies a custom blockchain network ID.")
	RpcUserflag := params.String("rpcuser", "", "Username to protect factomd local API with simple HTTP authentication")
	RpcPasswordflag := params.String("rpcpass", "", "Password to protect factomd local API. Ignored if rpcuser is blank")
	FactomdTLSflag := params.Bool("tls", false, "Set to true to require encrypted connections to factomd API and Control Panel") //to get tls, run as "factomd -tls=true"
	FactomdLocationsflag := params.String("selfaddr", "", "comma separated IPAddresses and DNS names of this factomd to use when creating a cert file")
	MemProfileRate := params.Int("mpr", 512*1024, "Set the Memory Profile Rate to update profiling per X bytes allocated. Default 512K, set to 1 to profile everything, 0 to disable.")
	exposeProfilePtr := params.Bool("exposeprofiler", false, "Setting this exposes the profiling port to outside localhost.")
	factomHomePtr := params.String("factomhome", "", "Set the Factom home directory. The .factom folder will be placed here if set, otherwise it will default to $HOME")

	logportPtr := params.String("logPort", "6060", "Port for pprof logging")
	portOverridePtr := params.Int("port", 0, "Port where we serve WSAPI;  default 8088")
	ControlPanelPortOverridePtr := params.Int("ControlPanelPort", 0, "Port for control panel webserver;  Default 8090")
	networkPortOverridePtr := params.Int("networkPort", 0, "Port for p2p network; default 8110")

	FastPtr := params.Bool("Fast", true, "If true, factomd will Fast-boot from a file.")
	FastLocationPtr := params.String("Fastlocation", "", "Directory to put the Fast-boot file in.")

	logLvlPtr := params.String("Loglvl", "none", "Set log level to either: none, debug, info, warning, error, fatal or panic")
	logJsonPtr := params.Bool("Logjson", false, "Use to set logging to use a json formatting")

	sim_stdinPtr := params.Bool("sim_stdin", true, "If true, sim control reads from stdin.")

	// Plugins
	PluginPath := params.String("plugin", "", "Input the path to any plugin binaries")

	// 	Torrent Plugin
	tormanager := params.Bool("tormanage", false, "Use torrent dbstate manager. Must have plugin binary installed and in $PATH")
	TorUploader := params.Bool("torupload", false, "Be a torrent uploader")

	// Logstash connection (if used)
	logstash := params.Bool("logstash", false, "If true, use Logstash")
	LogstashURL := params.String("logurl", "localhost:8345", "Endpoint URL for Logstash")

	sync2Ptr := params.Int("sync2", -1, "Set the initial blockheight for the second Sync pass. Used to force a total sync, or skip unnecessary syncing of entries.")

	params.StringVar(&p.DebugConsole, "debugconsole", "", "Enable DebugConsole on port. localhost:8093 open 8093 and spawns a telnet console, remotehost:8093 open 8093")
	params.StringVar(&p.StdoutLog, "stdoutlog", "", "Log stdout to a file")
	params.StringVar(&p.StderrLog, "stderrlog", "", "Log stderr to a file, optionally the same file as stdout")
	params.StringVar(&p.DebugLogRegEx, "debuglog", "off", "regex to pick which logs to save")
	params.IntVar(&elections.FaultTimeout, "faulttimeout", 30, "Seconds before considering Federated servers at-fault. Default is 30.")
	//params.CommandLine.Parse(args)

	// Good for unit testing
	disablePrometheus := params.Bool("disableprometheus", false, "Can be used to disable prometheus")

	params.Parse(args)

	p.AckbalanceHash = *ackBalanceHashPtr
	p.EnableNet = *enablenetPtr
	p.WaitEntries = *waitEntriesPtr
	p.ListenTo = *listenToPtr
	p.Cnt = *cntPtr
	p.Net = *netPtr
	p.Fnet = *fnetPtr
	p.DropRate = *dropPtr
	p.Journal = *journalPtr
	p.Journaling = *journalingPtr
	p.Follower = *followerPtr
	p.Leader = *leaderPtr
	p.Db = *dbPtr
	p.CloneDB = *cloneDBPtr
	p.PortOverride = *portOverridePtr
	p.Peers = *peersPtr
	p.NetworkName = *networkNamePtr
	p.NetworkPortOverride = *networkPortOverridePtr
	p.ControlPanelPortOverride = *ControlPanelPortOverridePtr
	p.LogPort = *logportPtr
	p.BlkTime = *blkTimePtr
	//	p.FaultTimeout = *faultTimeoutPtr
	p.RuntimeLog = *runtimeLogPtr
	p.Exclusive = *exclusivePtr
	p.Prefix = *PrefixNodePtr
	p.Rotate = *RotatePtr
	p.TimeOffset = *TimeOffsetPtr
	p.KeepMismatch = *KeepMismatchPtr
	p.StartDelay = int64(*startDelayPtr)
	p.Deadline = *DeadlinePtr
	p.CustomNet = primitives.Sha([]byte(*CustomNetPtr)).Bytes()[:4]
	p.RpcUser = *RpcUserflag
	p.RpcPassword = *RpcPasswordflag
	p.FactomdTLS = *FactomdTLSflag
	p.FactomdLocations = *FactomdLocationsflag
	p.MemProfileRate = *MemProfileRate
	p.Fast = *FastPtr
	p.FastLocation = *FastLocationPtr
	p.Loglvl = *logLvlPtr
	p.Logjson = *logJsonPtr
	p.Sim_Stdin = *sim_stdinPtr
	p.ExposeProfiling = *exposeProfilePtr

	p.PluginPath = *PluginPath
	p.TorManage = *tormanager
	p.TorUpload = *TorUploader

	p.UseLogstash = *logstash
	p.LogstashURL = *LogstashURL

	p.Sync2 = *sync2Ptr
	p.DisablePrometheus = *disablePrometheus

	if *factomHomePtr != "" {
		os.Setenv("FACTOM_HOME", *factomHomePtr)
	}

	// Handle the global (not factom server specific parameters
	if p.StdoutLog != "" || p.StderrLog != "" {
		handleLogfiles(p.StdoutLog, p.StderrLog)
	}

	fmt.Print("//////////////////////// Copyright 2017 Factom Foundation\n")
	fmt.Print("//////////////////////// Use of this source code is governed by the MIT\n")
	fmt.Print("//////////////////////// license that can be found in the LICENSE file.\n")

	if !isCompilerVersionOK() {
		fmt.Println("!!! !!! !!! ERROR: unsupported compiler version !!! !!! !!!")
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}

	// launch debug console if requested
	if p.DebugConsole != "" {
		launchDebugServer(p.DebugConsole)
	}

	return p
}

func isCompilerVersionOK() bool {
	goodenough := false

	if strings.Contains(runtime.Version(), "1.6") {
		goodenough = true
	}

	if strings.Contains(runtime.Version(), "1.7") {
		goodenough = true
	}

	if strings.Contains(runtime.Version(), "1.8") {
		goodenough = true
	}

	if strings.Contains(runtime.Version(), "1.9") {
		goodenough = true
	}

	if strings.Contains(runtime.Version(), "1.10") {
		goodenough = true
	}

	return goodenough
}

var handleLogfilesOnce sync.Once

func handleLogfiles(stdoutlog string, stderrlog string) {

	handleLogfilesOnce.Do(func() {

		var outfile *os.File
		var err error
		var wait sync.WaitGroup

		if stdoutlog != "" {
			// start a go routine to tee stdout to out.txt
			outfile, err = os.Create(stdoutlog)
			if err != nil {
				panic(err)
			}

			wait.Add(1)
			go func(outfile *os.File) {
				defer outfile.Close()
				defer os.Stdout.Close()                  // since I'm taking this away from  OS I need to close it when the time comes
				defer time.Sleep(100 * time.Millisecond) // Let the output all complete
				r, w, _ := os.Pipe()                     // Can't use the writer directly as os.Stdout so make a pipe
				oldStdout := os.Stdout
				os.Stdout = w
				wait.Done()
				// tee stdout to out.txt
				if _, err := io.Copy(io.MultiWriter(outfile, oldStdout), r); err != nil { // copy till EOF
					panic(err)
				}
			}(outfile) // stdout redirect func
		}

		if stderrlog != "" {
			if stderrlog != stdoutlog {
				outfile, err = os.Create(stderrlog)
				if err != nil {
					panic(err)
				}
			}

			wait.Add(1)
			go func(outfile *os.File) {
				if stderrlog != stdoutlog {
					defer outfile.Close()
				}
				defer os.Stderr.Close()                  // since I'm taking this away from  OS I need to close it when the time comes
				defer time.Sleep(100 * time.Millisecond) // Let the output all complete

				r, w, _ := os.Pipe() // Can't use the writer directly as os.Stdout so make a pipe
				oldStderr := os.Stderr
				os.Stderr = w
				wait.Done()
				if _, err := io.Copy(io.MultiWriter(outfile, oldStderr), r); err != nil { // copy till EOF
					panic(err)
				}
			}(outfile) // stderr redirect func
		}

		wait.Wait()                           // wait for the redirects to be active
		os.Stdout.WriteString("STDOUT Log\n") // Write any file header you want here e.g. node name and date and ...
		os.Stderr.WriteString("STDERR Log\n") // Write any file header you want here e.g. node name and date and ...
	})
}

var launchDebugServerOnce sync.Once

func launchDebugServer(service string) {

	launchDebugServerOnce.Do(func() {

		// start a go routine to tee stderr to the debug console
		debugConsole_r, debugConsole_w, _ := os.Pipe() // Can't use the writer directly as os.Stdout so make a pipe
		var wait sync.WaitGroup
		wait.Add(1)
		go func() {

			r, w, _ := os.Pipe() // Can't use the writer directly as os.Stderr so make a pipe
			oldStderr := os.Stderr
			os.Stderr = w
			defer oldStderr.Close()                  // since I'm taking this away from  OS I need to close it when the time comes
			defer time.Sleep(100 * time.Millisecond) // let the output all complete
			wait.Done()
			if _, err := io.Copy(io.MultiWriter(oldStderr, debugConsole_w), r); err != nil { // copy till EOF
				panic(err)
			}
		}() // stderr redirect func

		//		wait.Add(1)
		//go func() {
		//
		//	r, w, _ := os.Pipe() // Can't use the writer directly as os.Stderr so make a pipe
		//	oldStdout := os.Stdout
		//	os.Stdout = w
		//	defer oldStdout.Close()                  // since I'm taking this away from  OS I need to close it when the time comes
		//	defer time.Sleep(100 * time.Millisecond) // let the output all complete
		//	wait.Done()
		//	if _, err := io.Copy(io.MultiWriter(oldStdout, debugConsole_w), r); err != nil { // copy till EOF
		//		panic(err)
		//	}
		//}() // stdout redirect func

		wait.Wait() // Let the redirection become active ...

		host, port := "localhost", "8093" // defaults
		if service != "" {
			parts := strings.Split(service, ":")
			if len(parts) == 1 { // No port
				parts = append(parts, port) // use default
			}
			if parts[0] == "" { //no
				parts[0] = host // use default
			}
			host, port = parts[0], parts[1]

			_, badPort := strconv.Atoi(port)
			if (host != "localhost" && host != "remotehost") || badPort != nil {
				panic("Malformed -debugconsole option. Should be localhost:[port] or remotehost:[port] where [port] is a port number")
			}
		}

		// Start a listener port to connect to the debug server
		ln, err := net.Listen("tcp", ":"+port)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Debug Server listening for %v on port %v\n", host, port)

		newStdInR, newStdInW, _ := os.Pipe() // Can't use the reader directly as os.Stdin so make a pipe

		// Accept connections (one at a time)
		go func() {
			for {
				fmt.Printf("Debug server waiting for connection.\n") // Does not accept a reconnect not sure why ... revist
				connection, err := ln.Accept()
				if err != nil {
					panic(err)
				}
				fmt.Printf("Debug server accepted a connection.\n")

				writer := bufio.NewWriter(connection) // if we want to send something back to the telnet
				reader := bufio.NewReader(connection)

				writer.WriteString("Hello from Factom Debug Console\n")
				writer.Flush()
				// copy stderr to debug console
				go func() {
					if _, err := io.Copy(writer, debugConsole_r); err != nil { // copy till EOF
						fmt.Printf("Error copying stderr to debug consol: %v\n", err)
					}
				}()

				// copy input from debug console to stdin
				if false { // not sure why this doesn't work -- revist down the road
					if _, err = io.Copy(newStdInW, reader); err != nil {
						panic(err)
					}
				} else {
					for { // copy input from debug console to stdin until eof
						writer.WriteString(">") // print a prompt
						writer.Flush()
						if buf, err := reader.ReadString('\n'); err != nil {
							if err == io.EOF {
								break
							} // This connection is closed
							if err != nil {
								panic(err)
							} // This listen has an error
						} else {
							newStdInW.WriteString(string(buf))
						}
					}
				}
				fmt.Printf("Client disconnected.\n")
			}
		}() // the accept routine

		if host == "localhost" {
			cmd := exec.Command("/usr/bin/gnome-terminal", "-x", "telnet", "localhost", port)
			err = cmd.Start()
			if err != nil {
				panic(err)
			}
			fmt.Printf("Debug terminal pid %v\n", cmd.Process.Pid)
		}
		os.Stdin = newStdInR               // start using the pipe as input
		time.Sleep(100 * time.Millisecond) // Let the redirection become active ...
	})
}
