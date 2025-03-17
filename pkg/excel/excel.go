package excel

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/xuri/excelize/v2"
)

type PersonData struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Height   float64 `json:"height"`
	Birthday string  `json:"birthday"`
}

type ExcelService struct {
	ctx context.Context
}

func NewExcelService() *ExcelService {
	return &ExcelService{}
}

func (e *ExcelService) Startup(ctx context.Context) {
	e.ctx = ctx
}

// GenerateExcel cria um arquivo Excel com os dados fornecidos e retorna o caminho do arquivo salvo.
func (e *ExcelService) GenerateExcel(data PersonData) (string, error) {
	f := excelize.NewFile()

	// Definir cabeçalhos
	headers := []string{"Nome", "Idade", "Altura (m)", "Data de Nascimento"}
	for i, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		if err := f.SetCellValue("Sheet1", cell, header); err != nil {
			return "", fmt.Errorf("falha ao definir cabeçalho: %v", err)
		}
	}

	// Definir estilo para os cabeçalhos
	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:  true,
			Size:  12,
			Color: "#FFFFFF",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#4F81BD"},
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	})
	if err != nil {
		return "", fmt.Errorf("falha ao criar estilo: %v", err)
	}
	if err := f.SetCellStyle("Sheet1", "A1", "D1", style); err != nil {
		return "", fmt.Errorf("falha ao aplicar estilo: %v", err)
	}

	// Inserir dados
	if err := f.SetCellValue("Sheet1", "A2", data.Name); err != nil {
		return "", fmt.Errorf("falha ao inserir nome: %v", err)
	}
	if err := f.SetCellValue("Sheet1", "B2", data.Age); err != nil {
		return "", fmt.Errorf("falha ao inserir idade: %v", err)
	}
	if err := f.SetCellValue("Sheet1", "C2", data.Height); err != nil {
		return "", fmt.Errorf("falha ao inserir altura: %v", err)
	}
	if err := f.SetCellValue("Sheet1", "D2", data.Birthday); err != nil {
		return "", fmt.Errorf("falha ao inserir data de nascimento: %v", err)
	}

	// Ajustar largura das colunas
	if err := f.SetColWidth("Sheet1", "A", "D", 20); err != nil {
		return "", fmt.Errorf("falha ao ajustar largura das colunas: %v", err)
	}

	// Gerar nome do arquivo com timestamp
	timestamp := time.Now().Format("20060102_150405")
	fileName := fmt.Sprintf("dados_pessoa_%s.xlsx", timestamp)

	// Definir diretório de saída
	outputDir, err := getOutputDir()
	if err != nil {
		return "", fmt.Errorf("falha ao obter diretório de saída: %v", err)
	}

	// Criar diretório se não existir
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return "", fmt.Errorf("falha ao criar diretório: %v", err)
		}
	}

	// Salvar arquivo
	filePath := filepath.Join(outputDir, fileName)
	if err := f.SaveAs(filePath); err != nil {
		return "", fmt.Errorf("falha ao salvar arquivo: %v", err)
	}

	return filePath, nil
}

// getOutputDir retorna o diretório de saída apropriado para o sistema operacional.
func getOutputDir() (string, error) {
	// Usar diretório padrão do sistema
	var baseDir string
	switch runtime.GOOS {
	case "windows":
		baseDir = os.Getenv("USERPROFILE") // No Windows, o diretório do usuário é USERPROFILE
	case "darwin", "linux":
		baseDir = os.Getenv("HOME") // No macOS e Linux, o diretório do usuário é HOME
	default:
		return "", fmt.Errorf("sistema operacional não suportado")
	}

	// Diretório padrão para salvar os arquivos
	return filepath.Join(baseDir, "Documents", "GeradorExcel"), nil
}

// OpenFile abre o arquivo no programa padrão do sistema operacional.
func (e *ExcelService) OpenFile(path string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", "", path}
	case "darwin":
		cmd = "open"
		args = []string{path}
	case "linux":
		cmd = "xdg-open"
		args = []string{path}
	default:
		return fmt.Errorf("sistema operacional não suportado")
	}

	command := exec.Command(cmd, args...)
	if err := command.Start(); err != nil {
		return fmt.Errorf("falha ao abrir arquivo: %v", err)
	}

	return nil
}	