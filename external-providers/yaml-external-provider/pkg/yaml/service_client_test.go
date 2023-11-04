package yaml

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/bombsimon/logrusr/v3"
	"github.com/konveyor/analyzer-lsp/jsonrpc2"
	"github.com/konveyor/analyzer-lsp/lsp/protocol"
	"github.com/konveyor/analyzer-lsp/provider"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// func TestRawRequests(t *testing.T) {
// 	// Create command struct for lsp
// 	lspCmd := exec.Command(
// 		// "gopls", "-vv", "-logfile", "debug-go.log", "-rpc.trace",
// 		// "pylsp", "--log-file", "debug-python.log",
// 		// "node", "server.js", "--stdio",
// 		"yaml-language-server", "--stdio",
// 	)

// 	lspStdin, err := lspCmd.StdinPipe()
// 	if err != nil {
// 		fmt.Println("Error creating pipe to lsp process:", err)
// 		return
// 	}

// 	lspStdout, err := lspCmd.StdoutPipe()
// 	if err != nil {
// 		fmt.Println("Error creating pipe to lsp process:", err)
// 		return
// 	}

// 	if err := lspCmd.Start(); err != nil {
// 		fmt.Println("Error starting lsp process:", err)
// 		return
// 	}

// 	// Close lsp process on program exit
// 	defer func() {
// 		if err := lspCmd.Wait(); err != nil {
// 			fmt.Println("Error waiting for lsp process to exit:", err)
// 		}
// 	}()

// 	send := func(json string) {
// 		r := fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(json), json)
// 		fmt.Printf("--- sending ---\n%s\n---------------\n", r)
// 		fmt.Fprintf(lspStdin, r)
// 	}

// 	recv := func() {
// 		reader := bufio.NewReader(lspStdout)

// 		for {
// 			// Read until the first occurrence of "\r\n\r\n"
// 			var result strings.Builder

// 			for {
// 				b, err := reader.ReadByte()
// 				if err != nil {
// 					fmt.Println("Error reading header:", err)
// 					return
// 				}

// 				result.WriteByte(b)

// 				if strings.HasSuffix(result.String(), "\r\n\r\n") {
// 					break
// 				}
// 			}

// 			header := result.String()

// 			// Convert content length to integer
// 			const prefix = "Content-Length: "
// 			start := strings.Index(header, prefix) + len(prefix)
// 			end := strings.Index(header[start:], "\r\n")

// 			contentLength, err := strconv.Atoi(header[start : start+end])
// 			if err != nil {
// 				fmt.Println("Error extracting content length:", err)
// 				return
// 			}

// 			// Read the specified number of bytes
// 			buffer := make([]byte, contentLength)
// 			_, err = io.ReadFull(reader, buffer)
// 			if err != nil {
// 				fmt.Println("Error reading content:", err)
// 				return
// 			}

// 			fmt.Printf("--- recving ---\n%s\n---------------\n", header+string(buffer))
// 		}
// 	}

// 	go recv()

// 	// send("{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"initialize\",\"params\":{\"processId\":0,\"rootUri\":\"file:///analyzer-lsp/examples/yaml\",\"capabilities\":{},\"extendedClientCapabilities\":{\"classFileContentsSupport\":true}}}")
// 	send("{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"initialize\",\"params\":{\"processId\":0,\"rootUri\":\"file:///home/parthiba/Documents/project/analyzer-lsp/examples/yaml\",\"capabilities\":{},\"extendedClientCapabilities\":{\"classFileContentsSupport\":true}}}")
// 	fmt.Scanln()
// 	// send("{\"jsonrpc\":\"2.0\",\"method\":\"initialized\",\"params\":{}}")
// 	fmt.Scanln()
// 	// send("{\"jsonrpc\":\"2.0\",\"id\":2,\"method\":\"textDocument/definition\",\"params\":{\"textDocument\":{\"uri\":\"file:///analyzer-lsp/examples/yaml/yaml.yaml\"},\"position\":{\"line\":0,\"character\":2}}}")
// 	// send("{\"jsonrpc\":\"2.0\",\"id\":2,\"method\":\"textDocument/definition\",\"params\":{\"textDocument\":{\"uri\":\"file:///home/parthiba/Documents/project/analyzer-lsp/examples/yaml/test.yaml\"},\"position\":{\"line\":0,\"character\":2}}}")
// 	send("{\"jsonrpc\":\"2.0\",\"id\":2,\"method\":\"yaml/get/all/jsonSchemas\",\"params\":{\"uri\":\"file:///home/parthiba/Documents/project/analyzer-lsp/schemas/custom-schema1.json\"}}")
// }

// For debugging
type EvaluateCall struct {
	ServiceClient provider.ServiceClient
	Cap           string
	ConditionInfo []byte
}

// For debugging
func TestStuff(t *testing.T) {
	logrusLog := logrus.New()
	logrusLog.SetOutput(os.Stdout)
	logrusLog.SetFormatter(&logrus.TextFormatter{})
	// need to do research on mapping in logrusr to level here TODO
	logrusLog.SetLevel(logrus.Level(5))

	log := logrusr.New(logrusLog)

	ctx := context.TODO()
	client := NewGenericProvider()

	// golangSC, err := client.Init(ctx, log, provider.InitConfig{
	// 	Location: "/home/jonah/Projects/analyzer-lsp/examples/golang",
	// 	// WorkspaceFolders: []string{
	// 	// 	"/home/jonah/Projects/analyzer-lsp/examples/golang",
	// 	// },
	// 	AnalysisMode: "full",
	// 	ProviderSpecificConfig: map[string]interface{}{
	// 		"name":          "go",
	// 		"lspServerPath": "gopls",
	// 		"lspArgs":       []interface{}{"-vv", "-logfile", "debug-go.log", "-rpc.trace"},
	// 	},
	// })
	// // golangSC.(*genericServiceClient).rpc.AddHandler(FileHandler{os.Stdout})
	// _ = golangSC

	// pythonSC, err := client.Init(ctx, log, provider.InitConfig{
	// 	Location: "/home/jonah/Projects/analyzer-lsp/examples/python",
	// 	// WorkspaceFolders: []string{
	// 	// 	"/home/jonah/Projects/analyzer-lsp/examples/python",
	// 	// },
	// 	AnalysisMode: "full",
	// 	ProviderSpecificConfig: map[string]interface{}{
	// 		"name":          "python",
	// 		"lspServerPath": "pylsp",
	// 		"lspArgs":       []interface{}{"--log-file", "debug-python.log"},
	// 		// Would like to use WorkspaceFolders and DependencyFolders instead of this
	// 		"referencedOutputIgnoreContains": []interface{}{
	// 			"file:///home/jonah/Projects/analyzer-lsp/examples/python/__pycache__",
	// 			"file:///home/jonah/Projects/analyzer-lsp/examples/python/.venv",
	// 		},
	// 	},
	// })
	// // pythonSC.(*genericServiceClient).rpc.AddHandler(FileHandler{os.Stdout})
	// _ = pythonSC

	yamlSC, err := client.Init(ctx, log, provider.InitConfig{
		Location: "/home/parthiba/Documents/yaml-server-test/analyzer-lsp/examples/yaml",
		// WorkspaceFolders: []string{
		// 	"/home/jonah/Projects/analyzer-lsp/examples/golang",
		// },
		AnalysisMode: "full",
		ProviderSpecificConfig: map[string]interface{}{
			"name":          "yaml",
			"lspServerPath": "/usr/local/bin/yaml-language-server",
			"lspArgs":       []interface{}{"--stdio"},
		},
	})
	// yamlSC.(*genericServiceClient).rpc.AddHandler(FileHandler{os.Stdout})
	_ = yamlSC
	yamlSC.(*genericServiceClient).rpc.AddHandler(jsonrpc2.FileHandler{os.Stdout})
	var result []byte

	// s := protocol.

	// sUri := map[string]any{
	// 	"uri": "file:///home/parthiba/Documents/project/analyzer-lsp/schemas/custom-schema1.json",
	// }

	// m := "file:///home/parthiba/Documents/project/analyzer-lsp/schemas/custom-schema1.json"
	// m := "file:///home/parthiba/Documents/project/analyzer-lsp/examples/yaml/test.yaml"

	yUri := map[string]interface{}{
		"uri": "file:///home/parthiba/Documents/project/analyzer-lsp/examples/yaml/test.yaml",
	}

	// m := os.DevNull

	err = yamlSC.(*genericServiceClient).rpc.Notify(ctx, "yaml/registerCustomSchemaRequest", &protocol.InitializedParams{})
	if err != nil {
		fmt.Printf("notify error - %v", err)
	}
	yamlSC.(*genericServiceClient).rpc.Call(ctx, "custom/schema/request", yUri, &result)
	return

	// if err != nil {
	// 	panic(err)
	// }

	var res []byte
	_ = res

	// params := protocol.TextDocumentPositionParams{
	// 	TextDocument: protocol.TextDocumentIdentifier{
	// 		URI: "file:///home/jonah/Projects/analyzer-lsp/examples/yaml/yaml.yaml",
	// 	},
	// 	Position: protocol.Position{
	// 		Line:      0,
	// 		Character: 2,
	// 	},
	// }

	// yamlSC.(*genericServiceClient).rpc.Call(ctx, "textDocument/definition", params, &res)

	// params := protocol.DocumentSymbolParams{
	// 	TextDocument: protocol.TextDocumentIdentifier{
	// 		URI: "/home/jonah/Projects/analyzer-lsp/examples/python/file_a.py",
	// 	},
	// }

	// pythonSC.(*genericServiceClient).rpc.Call(ctx, "textDocument/documentSymbol", params, &res)

	// p0 := protocol.WorkspaceSymbolParams{
	// 	Query: "dummy.HelloWorld",
	// }

	// golangSC.(*genericServiceClient).rpc.Call(ctx, "workspace/symbol", p0, &res)

	// p1 := protocol.DocumentSymbolParams{
	// 	TextDocument: protocol.TextDocumentIdentifier{
	// 		URI: "/home/jonah/Projects/analyzer-lsp/examples/python/file_b.py",
	// 	},
	// }

	// pythonSC.(*genericServiceClient).rpc.Call(ctx, "textDocument/documentSymbol", p1, &res)

	// return

	// Everything under here doesn't matter for now

	var calls []EvaluateCall
	calls = append(
		calls,
		// EvaluateCall{pythonSC, "referenced", []byte(`{"referenced":{pattern: "def hello_world"}}`)},
		// EvaluateCall{pythonSC, "referenced", []byte(`{"referenced":{pattern: "speak"}}`)},
		// EvaluateCall{pythonSC, "referenced", []byte(`{"referenced":{pattern: "create_custom_resource_definition"}}`)},
		// EvaluateCall{yamlSC, "referenced", []byte(`{"referenced":{pattern: "blah"}}`)},
		// EvaluateCall{golangSC, "referenced", []byte(`{"referenced":{pattern: "v1beta1.CustomResourceDefinition"}}`)},
		// EvaluateCall{golangSC, "referenced", []byte(`{"referenced":{pattern: "HelloWorld"}}`)},
	)

	var responses []provider.ProviderEvaluateResponse

	for _, call := range calls {
		response, err := call.ServiceClient.Evaluate(ctx, call.Cap, call.ConditionInfo)
		if err != nil {
			panic(err)
		}

		responses = append(responses, response)

		fmt.Printf("Service Client: %s\n", call.ServiceClient.(*genericServiceClient).config.ProviderSpecificConfig["name"])
		fmt.Printf("Evaluated: %s, %s\n", call.Cap, string(call.ConditionInfo))
		fmt.Printf("Incidents:\n")
		b, _ := yaml.Marshal(response.Incidents)
		s := string(b)
		fmt.Printf("%s\n", s)

		fmt.Printf("Matched: %v\n", response.Matched)
		fmt.Printf("TemplateContext: %v\n", response.TemplateContext)
	}
}
