package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
)

type Data struct {
	Code   string
	Output string
}

func main() {
	address := "localhost:8080"
	index := template.Must(template.ParseFiles("Index.html"))

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "text/html")

		if req.Method == http.MethodGet {
			data := Data{
				Code:   "package main\r\n\nimport \"fmt\"\r\n\nfunc main() {\n  fmt.Println(\"Hello, World!\")\n}",
				Output: "Output will appear here after you send!",
			}

			err := index.Execute(rw, data)
			if err != nil {
				log.Fatalf("template execution: %s", err)
			}

		} else if req.Method == http.MethodPost {
			code := req.PostFormValue("code")

			goCmd := fmt.Sprintf(
				"touch main.go && echo -e %q > main.go && go fmt main.go && go run main.go", code)
			dockerArgs := []string{"run", "--rm", "golang", "bash", "-c", goCmd}
			cmd := exec.Command("docker", dockerArgs...)

			var output string
			stdoutStderr, err := cmd.CombinedOutput()
			if err != nil {
				output = fmt.Sprintf("Output:\n%s\nError: %s\n", stdoutStderr, err)
			} else {
				output = fmt.Sprintf("%s", stdoutStderr)
			}

			data := Data{
				Code:   code,
				Output: output,
			}

			err = index.Execute(rw, data)
			if err != nil {
				log.Fatalf("template execution: %s", err)
			}
		}
	})

	log.Printf("Listening on %s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
