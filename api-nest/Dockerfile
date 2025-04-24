# Usa una imagen de Node.js
FROM node:20

# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos
COPY package*.json ./
RUN npm install

# Copia el resto del proyecto
COPY . .

# Compila el proyecto
RUN npm run build

# Expone el puerto
EXPOSE 3000

# Comando para iniciar el servidor
CMD ["npm", "run", "start:prod"]
