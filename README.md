# ShellScribeAI

ShellScribeAI is a powerful tool designed to dynamically generate and execute scripts based on user inputs through the OpenAI API. This tool supports both Linux and Windows environments, allowing it to interpret system information, generate necessary commands, and provide detailed explanations of command outputs. It is particularly useful for executing tasks across various operating systems using natural language commands.

## Features

- **Dynamic Script Generation**: Generates scripts based on user queries using OpenAI's API.
- **Cross-Platform Support**: Works seamlessly on both Linux and Windows operating systems.
- **Interactive Command Execution**: Prompts users for commands and executes them interactively.
- **Detailed Output Interpretation**: Provides natural language explanations of the command outputs.
- **Debug Mode**: Offers detailed logging for development and troubleshooting.

## Prerequisites

- **Go 1.16+**: This tool is developed in Go and requires Go version 1.16 or higher.
- **OpenAI API Key**: An OpenAI API key is needed for script generation and interpretation features.
- **Network Access**: The tool makes HTTP requests to the OpenAI API.

## Installation

1. **Clone the Repository**

   Clone the ShellScribeAI repository to your local machine.

   ```sh
   git clone https://github.com/NerdyNot/ShellScribeAI.git
   cd ShellScribeAI
   ```

2. **Initialize the Project**

   Initialize the Go module to manage project dependencies.

   ```sh
   go mod init github.com/YourUsername/ShellScribeAI
   ```

   If `go.mod` already exists, skip this step.

3. **Install Required Packages**

   Install the necessary packages. These packages are specified in the `go.mod` file and their versions are locked in `go.sum`.

   ```sh
   go get github.com/fatih/color
   go get github.com/manifoldco/promptui
   ```

   Alternatively, you can use `go mod tidy` to clean up and verify dependencies. This command will ensure that the `go.mod` and `go.sum` files are up to date and only include the necessary dependencies.

   ```sh
   go mod tidy
   ```

4. **Set Environment Variables**

   Set your OpenAI API key as an environment variable.

   ```sh
   export OPENAI_API_KEY="your_openai_api_key"
   ```

   For Windows PowerShell, set the environment variable as follows.

   ```ps
   $env:OPENAI_API_KEY="your_openai_api_key"
   ```

## Usage

### Running the Tool

To start ShellScribeAI, use the `go run` command in the terminal.

```sh
go run main.go
```

Alternatively, you can build the binary and execute it:

```sh
go build -o ShellScribeAI main.go
./ShellScribeAI
```

### Interactive Mode

When running the tool, it will prompt you to enter your commands interactively. Type your command and press Enter to execute it. The tool will analyze your query, generate the appropriate script, and execute it on your local system.

### Command-Line Arguments

ShellScribeAI supports a few command-line flags:

- `-d`: Enable debug mode to see detailed logs of the operations.

Example:

```sh
go run main.go -d
```

### Exit the Program

To exit the program, type `exit` or `quit` in the prompt.

## Configuration

ShellScribeAI does not require any configuration files. All configurations are managed through environment variables and command-line flags.

## Important Notices

1. **API Key Management**:
   - Ensure your OpenAI API key is secure and not exposed in your code or logs.
   - If the API key is not found in the environment variables, the tool will prompt you to enter it.

2. **Script Execution**:
   - The generated scripts can perform various system operations. Always review the generated scripts, especially when in debug mode, to ensure they are safe to execute.
   - Be cautious when running scripts with administrative or root privileges.

3. **User Responsibility**:
   - Users are responsible for the scripts executed on their systems. This tool can execute potentially dangerous commands if not used properly.
   - It is advisable to test scripts in a safe environment before running them on production systems.

## Contributing

Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Libraries and Licenses

### promptui

[promptui](https://github.com/manifoldco/promptui)
- **Description**: PromptUI is a library for creating command-line prompts with Go.
- **License**: BSD-3-Clause
- **Usage**: Used to create interactive prompts for user input.

### color

[color](https://github.com/fatih/color)
- **Description**: Color is a Go package for colorizing and styling console output.
- **License**: MIT
- **Usage**: Used to enhance the console output with colors.

### Go Standard Library

- The Go standard library provides numerous packages used extensively in this project for HTTP requests, JSON handling, and system interactions.

---

# ShellScribeAI

ShellScribeAI는 사용자의 입력을 기반으로 OpenAI API를 통해 동적으로 스크립트를 생성하고 실행하는 강력한 도구입니다. 이 도구는 Linux 및 Windows 환경을 모두 지원하며, 시스템 정보를 해석하고 필요한 명령을 생성하며 명령 출력에 대한 자세한 설명을 제공할 수 있습니다. 다양한 운영 체제에서 자연어를 통해 작업 명령을 수행할 때 유용합니다.

## 기능

- **동적 스크립트 생성**: 사용자의 질의에 따라 OpenAI의 API를 사용하여 스크립트를 생성합니다.
- **크로스 플랫폼 지원**: Linux 및 Windows 운영 체제에서 원활하게 작동합니다.
- **대화형 명령 실행**: 사용자의 명령을 대화형으로 요청하고 실행합니다.
- **상세한 출력 해석**: 명령 출력에 대한 자연어 설명을 제공합니다.
- **디버그 모드**: 개발 및 문제 해결을 위한 자세한 로그를 제공합니다.

## 사전 요구 사항

- **Go 1.16+**: 이 도구는 Go로 개발되었으며 Go 1.16 이상의 버전이 필요합니다.
- **OpenAI API 키**: 스크립트 생성 및 해석 기능을 사용하려면 OpenAI API 키가 필요합니다.
- **네트워크 접근**: 도구가 OpenAI API에 HTTP 요청을 보냅니다.

## 설치

1. **리포지토리 클론**

   ShellScribeAI 리포지토리를 로컬 머신에 클론합니다.

   ```sh
   git clone https://github.com/NerdyNot/ShellScribeAI.git
   cd ShellScribeAI
   ```

2. **프로젝트 초기화**

   Go 모듈을 초기화하여 프로젝트의 종속성을 관리합니다.

   ```sh
   go mod init github.com/YourUsername/ShellScribeAI
   ```

   `go.mod` 파일이 이미 존재하는 경우, 이 단계는 건너뛰십시오.

3. **필수 패키지 설치**

   필요한 패키지를 설치합니다. 이러한 패키지는 `go.mod` 파일에 명시되어 있으며, 해당 버전은 `go.sum` 파일에 고정되어 있습니다.

   ```sh
   go get github.com/fatih/color
   go get github.com/manifoldco/promptui
   ```

   또는, `go mod tidy` 명령을 사용하여 의존성을 정리하고 검증할 수 있습니다. 이 명령은 `go.mod`와 `go.sum` 파일이 최신 상태이며 필요한 의존성만 포함하도록 보장합니다.

   ```sh
   go mod tidy
   ```

4. **환경 변수 설정**

   OpenAI API 키를 환경 변수로 설정하십시오.

   ```sh
   export OPENAI_API_KEY="your_openai_api_key"
   ```

   Windows의 경우 PowerShell을 사용하여 환경 변수를 설정할 수 있습니다.

   ```ps
   $env:OPENAI_API_KEY="your_openai_api_key"
   ```

## 사용법

### 도구 실행

ShellScribeAI를 시작하려면 터미널에서 `go run` 명령을 사용하십시오.

```sh
go run main.go
```

또는, 바이너리를 빌드하고 실행할 수 있습니다:

```sh
go build -o ShellScribeAI main.go
./ShellScribeAI
```

### 대화형 모드

도구를 실행하면 명령을 대화형으로 입력하라는 메시지가 표시됩니다. 명령을 입력하고 Enter를 눌러 실행하십시오. 도구는 질의를 분석하고, 적절한 스크립트를 생성한 다음 로컬 시스템에서 실행합니다.

### 명령줄 인수

ShellScribeAI는 몇 가지 명령줄 플래그를 지원합니다:

- `-d`: 자세한 로그를 보려면 디버그 모드를 활성화하십시오.

예시:

```sh
go run main.go -d
```

### 프로그램 종료

프로그램을 종료하려면 프롬프트에 `exit` 또는 `quit`을 입력하십시오.

## 구성

ShellScribeAI는 구성 파일이

 필요하지 않습니다. 모든 구성은 환경 변수와 명령줄 플래그를 통해 관리됩니다.

## 중요 사항

1. **API 키 관리**:
   - OpenAI API 키를 안전하게 보관하고 코드나 로그에 노출되지 않도록 주의하십시오.
   - 환경 변수에서 API 키를 찾을 수 없는 경우, 도구가 API 키 입력을 요청합니다.

2. **스크립트 실행**:
   - 생성된 스크립트는 시스템 설정을 변경하거나 다양한 작업을 수행할 수 있습니다. 디버그 모드에서 생성된 스크립트를 검토하여 안전하게 실행할 수 있는지 확인하십시오.
   - 관리자 또는 루트 권한으로 스크립트를 실행할 때는 주의하십시오.

3. **사용자 책임**:
   - 사용자는 시스템에서 실행되는 스크립트에 대한 책임이 있습니다. 이 도구는 부적절하게 사용될 경우 잠재적으로 위험한 명령을 실행할 수 있습니다.
   - 프로덕션 시스템에서 실행하기 전에 안전한 환경에서 스크립트를 테스트하는 것이 좋습니다.

## 기여

개선 사항이나 버그 수정을 위한 이슈를 열거나 풀 리퀘스트를 제출해 주세요.

## 라이센스

이 프로젝트는 MIT 라이센스를 따릅니다. 자세한 내용은 [LICENSE](LICENSE) 파일을 참조하십시오.

## 라이브러리 및 라이센스

### promptui

[promptui](https://github.com/manifoldco/promptui)
- **설명**: promptui는 Go에서 명령줄 프롬프트를 생성하는 라이브러리입니다.
- **라이센스**: BSD-3-Clause
- **사용 용도**: 사용자 입력을 위한 대화형 프롬프트를 생성하는 데 사용됩니다.

### color

[color](https://github.com/fatih/color)
- **설명**: color는 콘솔 출력을 색상화하고 스타일링하기 위한 Go 패키지입니다.
- **라이센스**: MIT
- **사용 용도**: 콘솔 출력을 색상으로 꾸미기 위해 사용됩니다.

### Go 표준 라이브러리

- Go 표준 라이브러리는 이 프로젝트에서 HTTP 요청, JSON 처리 및 시스템 상호작용을 위해 광범위하게 사용됩니다.
