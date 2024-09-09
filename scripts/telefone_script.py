import csv
import random

# Número de pessoas no arquivo
num_pessoas = 77480

# Nome do arquivo CSV a ser gerado
output_file = "your_csv_name.csv"

# Gera uma lista de números de celular
celulares_por_pessoa = []
celular_count = 0

for i in range(1, num_pessoas + 1):

# Gera entre 1 e 3 números de celular para cada pessoa
    num_celulares = random.randint(1, 3)
    celulares = [f"(11) 9{random.randint(60000000, 99999999)}" for _ in range(num_celulares)]
    celulares_por_pessoa.append(celulares)

# Escreve o arquivo CSV
with open(output_file, mode="w", newline="") as file:
    writer = csv.writer(file)
    writer.writerow(["celulares"])  
    for celulares in celulares_por_pessoa:
        writer.writerow(celulares)

print(f"Arquivo {output_file} gerado com sucesso!")
