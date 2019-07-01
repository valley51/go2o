// Autogenerated by Thrift Compiler (0.12.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"go2o/core/service/auto_gen/rpc/foundation_service"
	"go2o/core/service/auto_gen/rpc/ttype"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var _ = ttype.GoUnusedProtection__

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  string GetRegistry(string key)")
	fmt.Fprintln(os.Stderr, "   GetRegistries( keys)")
	fmt.Fprintln(os.Stderr, "   SearchRegistry(string key)")
	fmt.Fprintln(os.Stderr, "  Result CreateUserRegistry(string key, string defaultValue, string description)")
	fmt.Fprintln(os.Stderr, "  Result UpdateRegistry( registries)")
	fmt.Fprintln(os.Stderr, "  string ResourceUrl(string url)")
	fmt.Fprintln(os.Stderr, "  Result SetValue(string key, string value)")
	fmt.Fprintln(os.Stderr, "  Result DeleteValue(string key)")
	fmt.Fprintln(os.Stderr, "   GetRegistryV1( keys)")
	fmt.Fprintln(os.Stderr, "   GetValuesByPrefix(string prefix)")
	fmt.Fprintln(os.Stderr, "  string RegisterApp(SSsoApp app)")
	fmt.Fprintln(os.Stderr, "  SSsoApp GetApp(string name)")
	fmt.Fprintln(os.Stderr, "   GetAllSsoApp()")
	fmt.Fprintln(os.Stderr, "  bool SuperValidate(string user, string pwd)")
	fmt.Fprintln(os.Stderr, "  void FlushSuperPwd(string user, string pwd)")
	fmt.Fprintln(os.Stderr, "  string GetSyncLoginUrl(string returnUrl)")
	fmt.Fprintln(os.Stderr, "   GetAreaNames( codes)")
	fmt.Fprintln(os.Stderr, "   GetChildAreas(i32 code)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
	var m map[string]string = h
	return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
	parts := strings.Split(value, ": ")
	if len(parts) != 2 {
		return fmt.Errorf("header should be of format 'Key: Value'")
	}
	h[parts[0]] = parts[1]
	return nil
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	headers := make(httpHeaders)
	var parsedUrl *url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
	flag.Parse()

	if len(urlString) > 0 {
		var err error
		parsedUrl, err = url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
		if len(headers) > 0 {
			httptrans := trans.(*thrift.THttpClient)
			for key, value := range headers {
				httptrans.SetHeader(key, value)
			}
		}
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	iprot := protocolFactory.GetProtocol(trans)
	oprot := protocolFactory.GetProtocol(trans)
	client := foundation_service.NewFoundationServiceClient(thrift.NewTStandardClient(iprot, oprot))
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "GetRegistry":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetRegistry requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetRegistry(context.Background(), value0))
		fmt.Print("\n")
		break
	case "GetRegistries":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetRegistries requires 1 args")
			flag.Usage()
		}
		arg53 := flag.Arg(1)
		mbTrans54 := thrift.NewTMemoryBufferLen(len(arg53))
		defer mbTrans54.Close()
		_, err55 := mbTrans54.WriteString(arg53)
		if err55 != nil {
			Usage()
			return
		}
		factory56 := thrift.NewTJSONProtocolFactory()
		jsProt57 := factory56.GetProtocol(mbTrans54)
		containerStruct0 := foundation_service.NewFoundationServiceGetRegistriesArgs()
		err58 := containerStruct0.ReadField1(jsProt57)
		if err58 != nil {
			Usage()
			return
		}
		argvalue0 := containerStruct0.Keys
		value0 := argvalue0
		fmt.Print(client.GetRegistries(context.Background(), value0))
		fmt.Print("\n")
		break
	case "SearchRegistry":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "SearchRegistry requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.SearchRegistry(context.Background(), value0))
		fmt.Print("\n")
		break
	case "CreateUserRegistry":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "CreateUserRegistry requires 3 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		fmt.Print(client.CreateUserRegistry(context.Background(), value0, value1, value2))
		fmt.Print("\n")
		break
	case "UpdateRegistry":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "UpdateRegistry requires 1 args")
			flag.Usage()
		}
		arg63 := flag.Arg(1)
		mbTrans64 := thrift.NewTMemoryBufferLen(len(arg63))
		defer mbTrans64.Close()
		_, err65 := mbTrans64.WriteString(arg63)
		if err65 != nil {
			Usage()
			return
		}
		factory66 := thrift.NewTJSONProtocolFactory()
		jsProt67 := factory66.GetProtocol(mbTrans64)
		containerStruct0 := foundation_service.NewFoundationServiceUpdateRegistryArgs()
		err68 := containerStruct0.ReadField1(jsProt67)
		if err68 != nil {
			Usage()
			return
		}
		argvalue0 := containerStruct0.Registries
		value0 := argvalue0
		fmt.Print(client.UpdateRegistry(context.Background(), value0))
		fmt.Print("\n")
		break
	case "ResourceUrl":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "ResourceUrl requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.ResourceUrl(context.Background(), value0))
		fmt.Print("\n")
		break
	case "SetValue":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "SetValue requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.SetValue(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "DeleteValue":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "DeleteValue requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.DeleteValue(context.Background(), value0))
		fmt.Print("\n")
		break
	case "GetRegistryV1":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetRegistryV1 requires 1 args")
			flag.Usage()
		}
		arg73 := flag.Arg(1)
		mbTrans74 := thrift.NewTMemoryBufferLen(len(arg73))
		defer mbTrans74.Close()
		_, err75 := mbTrans74.WriteString(arg73)
		if err75 != nil {
			Usage()
			return
		}
		factory76 := thrift.NewTJSONProtocolFactory()
		jsProt77 := factory76.GetProtocol(mbTrans74)
		containerStruct0 := foundation_service.NewFoundationServiceGetRegistryV1Args()
		err78 := containerStruct0.ReadField1(jsProt77)
		if err78 != nil {
			Usage()
			return
		}
		argvalue0 := containerStruct0.Keys
		value0 := argvalue0
		fmt.Print(client.GetRegistryV1(context.Background(), value0))
		fmt.Print("\n")
		break
	case "GetValuesByPrefix":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetValuesByPrefix requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetValuesByPrefix(context.Background(), value0))
		fmt.Print("\n")
		break
	case "RegisterApp":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "RegisterApp requires 1 args")
			flag.Usage()
		}
		arg80 := flag.Arg(1)
		mbTrans81 := thrift.NewTMemoryBufferLen(len(arg80))
		defer mbTrans81.Close()
		_, err82 := mbTrans81.WriteString(arg80)
		if err82 != nil {
			Usage()
			return
		}
		factory83 := thrift.NewTJSONProtocolFactory()
		jsProt84 := factory83.GetProtocol(mbTrans81)
		argvalue0 := foundation_service.NewSSsoApp()
		err85 := argvalue0.Read(jsProt84)
		if err85 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.RegisterApp(context.Background(), value0))
		fmt.Print("\n")
		break
	case "GetApp":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetApp requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetApp(context.Background(), value0))
		fmt.Print("\n")
		break
	case "GetAllSsoApp":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetAllSsoApp requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetAllSsoApp(context.Background()))
		fmt.Print("\n")
		break
	case "SuperValidate":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "SuperValidate requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.SuperValidate(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "FlushSuperPwd":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "FlushSuperPwd requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.FlushSuperPwd(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "GetSyncLoginUrl":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetSyncLoginUrl requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetSyncLoginUrl(context.Background(), value0))
		fmt.Print("\n")
		break
	case "GetAreaNames":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetAreaNames requires 1 args")
			flag.Usage()
		}
		arg92 := flag.Arg(1)
		mbTrans93 := thrift.NewTMemoryBufferLen(len(arg92))
		defer mbTrans93.Close()
		_, err94 := mbTrans93.WriteString(arg92)
		if err94 != nil {
			Usage()
			return
		}
		factory95 := thrift.NewTJSONProtocolFactory()
		jsProt96 := factory95.GetProtocol(mbTrans93)
		containerStruct0 := foundation_service.NewFoundationServiceGetAreaNamesArgs()
		err97 := containerStruct0.ReadField1(jsProt96)
		if err97 != nil {
			Usage()
			return
		}
		argvalue0 := containerStruct0.Codes
		value0 := argvalue0
		fmt.Print(client.GetAreaNames(context.Background(), value0))
		fmt.Print("\n")
		break
	case "GetChildAreas":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetChildAreas requires 1 args")
			flag.Usage()
		}
		tmp0, err98 := (strconv.Atoi(flag.Arg(1)))
		if err98 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.GetChildAreas(context.Background(), value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}