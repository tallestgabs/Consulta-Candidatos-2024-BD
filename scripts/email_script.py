import csv
import random

# Número de pessoas no arquivo
num_pessoas = 77480

# Nome do arquivo CSV a ser gerado
output_file = "your_csv_name.csv"

# Gera uma lista de emails
emails_por_pessoa = []
email_count = 0

for i in range(1, num_pessoas + 1):
    
# Gera entre 1 e 3 emails para cada pessoa
    num_emails = random.randint(1, 3)
    emails = [f"example@email{email_count + j}" for j in range(num_emails)]
    email_count += num_emails
    emails_por_pessoa.append(emails)

# Escreve o arquivo CSV
with open(output_file, mode="w", newline="") as file:
    writer = csv.writer(file)
    writer.writerow(["emails"])  # Cabeçalho da coluna
    for emails in emails_por_pessoa:
        writer.writerow(emails)

print(f"Arquivo {output_file} gerado com sucesso!")
