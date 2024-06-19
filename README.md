  # GO Driver

  Bem-vindo ao repositório do projeto de driver em Golang. Este documento fornece uma visão geral do projeto, instruções para instalação de dependências, execução de testes e outras informações úteis. Lembrando que esse projeto está sendo desenvolvido a partir da imersão do curso **AprendaGolang: https://aprendagolang.com.br/courses/imersao-backend/**.

  ## Visão Geral

  Este projeto tem como objetivo desenvolver um driver em Go para gerenciar operações de criação de pastas, uploads de pastas, arquivos e outras tarefas relacionadas. O driver será parte de um sistema que oferece funcionalidades robustas para manipulação de estruturas de arquivos e diretórios, com suporte a interações via CLI e acesso externo via API.

  ## Pré-requisitos

  Antes de começar, você precisará ter instalado em seu sistema:

  - [Go](https://golang.org/doc/install) (versão 1.16 ou superior)
  - Git

  ## Instalação

  1. **Clone o repositório**

     ```sh
     git clone https://github.com/JoaoPauloFontana/go-driver.git
     cd go-driver
     ```

  2. **Instale as dependências**

     Utilizamos o Go Modules para gerenciar as dependências. Para instalar as dependências do projeto, execute:

     ```sh
     go mod tidy
     ```

  ## Estrutura do Projeto

  Uma visão geral da estrutura de diretórios e arquivos do projeto:

  ```
  go-driver/
  ├── cmd/worker          # Comandos executáveis
  ├── docs/               # Documentos
  ├── internal/           # Pacotes internos
  ├── pkg/database        # Pacotes reutilizáveis
  ├── script/database     # Script para geração do banco
  ├── go.mod              # Arquivo de módulos do Go
  ├── go.sum              # Checksum das dependências
  └── README.md           # Este arquivo
  ```

  ## Execução

  Para executar a aplicação, utilize o comando:

  ```sh
  go run cmd/nome-do-comando/main.go
  ```

  ## Testes

  Para executar os testes, utilize o comando:

  ```sh
  go test ./internal/nome_da_pasta/... -v 
  ```

  Este comando irá rodar todos os testes presentes nos subdiretórios do projeto.

  ## Contribuição

  Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests. Para grandes mudanças, por favor, abra uma issue primeiro para discutir o que você gostaria de mudar.

  ## Contato

  Se você tiver alguma dúvida ou sugestão, sinta-se à vontade para entrar em contato.

  - João Paulo(mailto:jpfontana12@icloud.com)

  ---

  Agradecemos por usar e contribuir para o nosso projeto!
