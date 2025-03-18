# 🧮 Excel Generator

O **Excel Generator** é uma aplicação **PROTÓTIPO** desenvolvida em **Golang** com a ferramenta de desenvolvimento desktop **Wails** que permite ao usuário gerar planilhas Excel (.xlsx) a partir de entradas de dados. 
Se a planilha já existir, novos dados serão adicionados a ela, sem sobrescrever os dados anteriores.

O projeto foi pensando para ter uma boa organização e separação de responsabilidades entre as funções, além de uma interface agradável usando React e TypeScript.

### Funcionalidades:
- Interface gráfica simples e intuitiva.
- Validação de dados de entrada (nomes, idade, data de nascimento, etc.).
- Criação de novas planilhas a partir dos dados de input.

## 📋 Pré-requisitos

Certifique-se de ter instalado em sua máquina:

- [Golang](https://golang.org/dl/) (1.21 ou superior)
- [Wails](https://wails.io//) - Siga o tutorial no "Getting Started".

## 🚀 Instalação e Execução

1. **Clone o repositório:**

```bash
git clone https://github.com/SeuUsuario/excel-generator.git](https://github.com/Guizzs26/go-wails-lib-testing.git)
```

2. **Acesse o diretório do projeto:**

```bash
cd go-wails-lib-testing
```

3. **Instale as dependências do Wails:**

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

4. **Execute a aplicação:**

```bash
wails dev 
```

Será criado um localhost e aplicação desktop será executada e aberta no seu computador.

## 🤝 Contribuição

Sinta-se à vontade para enviar Pull Requests ou abrir Issues para relatar bugs ou sugerir melhorias.

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

