import os
import psycopg2

# Conexão com o banco de dados PostgreSQL
conn = psycopg2.connect(
    host="localhost", 
    database="Candidatos2024",
    user="postgres",
    password="gabson3xd"
)
cur = conn.cursor()

# Diretório onde estão as fotos
image_directory = r'YOUR_PATH\foto_cand2024_SP_div'

# Função para atualizar a foto no banco
def update_foto(sq_candidato, image_path):
    with open(image_path, 'rb') as file:
        binary_data = file.read()

# Atualiza a coluna 'foto' no candidato com o sq_candidato correspondente
    cur.execute("""
        UPDATE candidato
        SET foto = %s
        WHERE sq_candidato = %s
    """, (binary_data, sq_candidato))
    conn.commit()

# Percorre o diretório de fotos
for filename in os.listdir(image_directory):
    if filename.endswith(".jpg"):

# Extrai o ID do candidato a partir do nome do arquivo (removendo "FSP", "_div", e ".jpg")
        sq_candidato = filename.replace("FSP", "").replace("_div", "").replace(".jpg", "")
        
# Caminho completo da imagem
        image_path = os.path.join(image_directory, filename)
        
# Atualiza a foto para o candidato correspondente
        update_foto(sq_candidato, image_path)
        print(f"Foto de {filename} atualizada para o candidato {sq_candidato}")

# Fecha a conexão
cur.close()
conn.close()
