import React, { useState } from "react";
import Form from "../../components/Form/Form";
import { PersonData } from "../../types";
import styles from "./Home.module.css";
import {
  GenerateExcel,
  OpenFile,
} from "../../../wailsjs/go/excel/ExcelService";

const Home: React.FC = () => {
  const [isLoading, setIsLoading] = useState(false);
  const [message, setMessage] = useState<{
    text: string;
    type: "success" | "error";
  } | null>(null);

  const handleSubmit = async (data: PersonData) => {
    setIsLoading(true);
    setMessage(null);

    try {
      const filePath = await GenerateExcel(data);

      setMessage({
        text: `Excel gerado com sucesso!`,
        type: "success",
      });

      // Abrir o arquivo
      await OpenFile(filePath);
    } catch (error) {
      console.error("Erro ao gerar Excel:", error);
      setMessage({
        text: `Erro ao gerar Excel: ${
          error instanceof Error ? error.message : "Erro desconhecido"
        }`,
        type: "error",
      });
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.card}>
        <h1 className={styles.title}>Gerador de Excel</h1>
        <p className={styles.subtitle}>
          Preencha os dados abaixo para gerar uma planilha Excel
        </p>

        {message && (
          <div className={`${styles.message} ${styles[message.type]}`}>
            {message.text}
          </div>
        )}

        <Form onSubmit={handleSubmit} isLoading={isLoading} />
      </div>
    </div>
  );
};

export default Home;
