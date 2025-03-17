import React, { useState } from "react";
import Button from "../Button/Button";
import styles from "./Form.module.css";
import { PersonData } from "../../types";

interface FormProps {
  onSubmit: (data: PersonData) => void;
  isLoading: boolean;
}

const Form: React.FC<FormProps> = ({ onSubmit, isLoading }) => {
  const [formData, setFormData] = useState<PersonData>({
    name: "",
    age: 0,
    height: 0,
    birthday: "",
  });

  const [errors, setErrors] = useState<Record<string, string>>({});

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value, type } = e.target;

    let parsedValue: string | number = value;
    if (type === "number") {
      parsedValue = value === "" ? 0 : parseFloat(value);
    }

    setFormData((prev) => ({
      ...prev,
      [name]: parsedValue,
    }));
  };

  const validateForm = (): boolean => {
    const newErrors: Record<string, string> = {};

    if (!formData.name.trim()) {
      newErrors.name = "Nome é obrigatório";
    }

    if (formData.age <= 0) {
      newErrors.age = "Idade deve ser maior que zero";
    }

    if (formData.height <= 0) {
      newErrors.height = "Altura deve ser maior que zero";
    }

    if (!formData.birthday) {
      newErrors.birthday = "Data de nascimento é obrigatória";
    }

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    if (validateForm()) {
      onSubmit(formData);
    }
  };

  return (
    <form className={styles.form} onSubmit={handleSubmit}>
      <div className={styles.formGroup}>
        <label htmlFor="name">Nome</label>
        <input
          type="text"
          id="name"
          name="name"
          value={formData.name}
          onChange={handleChange}
          className={errors.name ? styles.inputError : ""}
          placeholder="Digite seu nome"
        />
        {errors.name && <span className={styles.error}>{errors.name}</span>}
      </div>

      <div className={styles.formGroup}>
        <label htmlFor="age">Idade</label>
        <input
          type="number"
          id="age"
          name="age"
          value={formData.age || ""}
          onChange={handleChange}
          className={errors.age ? styles.inputError : ""}
          placeholder="Digite sua idade"
        />
        {errors.age && <span className={styles.error}>{errors.age}</span>}
      </div>

      <div className={styles.formGroup}>
        <label htmlFor="height">Altura (m)</label>
        <input
          type="number"
          id="height"
          name="height"
          step="0.01"
          value={formData.height || ""}
          onChange={handleChange}
          className={errors.height ? styles.inputError : ""}
          placeholder="Ex: 1.75"
        />
        {errors.height && <span className={styles.error}>{errors.height}</span>}
      </div>

      <div className={styles.formGroup}>
        <label htmlFor="birthday">Data de Nascimento</label>
        <input
          type="date"
          id="birthday"
          name="birthday"
          value={formData.birthday}
          onChange={handleChange}
          className={errors.birthday ? styles.inputError : ""}
        />
        {errors.birthday && (
          <span className={styles.error}>{errors.birthday}</span>
        )}
      </div>

      <Button type="submit" disabled={isLoading}>
        {isLoading ? "Gerando..." : "Gerar Excel"}
      </Button>
    </form>
  );
};

export default Form;
