# stage 1: Definimos el constructor
FROM python:3.11-slim AS builder

# Definimos el directorio de trabajo
WORKDIR /install

# Evitamos basura innecesaria
ENV PIP_NO_CACHE_DIR=1 \
    PIP_DISABLE_PIP_VERSION_CHECK=1

# Copiamos los requerimientos
COPY requirements.txt .

# Realizamos la instalacion
RUN pip install --upgrade pip \
    && pip install --prefix=/install -r requirements.txt

# Stage 2: Corremos la apliacion

FROM python:3.11-slim

WORKDIR /app

# Copiamos las descargas del constructor (Builder)

# Definimos el directorio del builder con las dependencias necesarias
COPY --from=builder /install /usr/local 

# Copiamos los archivos
COPY . .

CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]
