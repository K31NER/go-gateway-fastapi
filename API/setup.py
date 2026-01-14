import time
import subprocess

def run_services():
    # Comandos para ejecutar
    # En Windows usamos 'solo' para Celery
    celery_cmd = [
        "celery", "-A", "config.celery_config", "worker", 
        "--loglevel=info", "-P", "solo"
    ]
    
    fastapi_cmd = [
        "uvicorn", "main:app", "--reload", "--port", "8000"
    ]

    print("--- Iniciando Worker de Celery ---")
    celery_process = subprocess.Popen(celery_cmd, shell=True)

    print("--- Iniciando Servidor FastAPI ---")
    fastapi_process = subprocess.Popen(fastapi_cmd, shell=True)

    try:
        # Mantener el script corriendo mientras los procesos estén vivos
        while True:
            time.sleep(1)
            # Si alguno muere, salimos
            if celery_process.poll() is not None:
                print("El proceso de Celery se ha detenido.")
                break
            if fastapi_process.poll() is not None:
                print("El proceso de FastAPI se ha detenido.")
                break
    except KeyboardInterrupt:
        print("\n--- Deteniendo servicios... ---")
    finally:
        celery_process.terminate()
        fastapi_process.terminate()
        print("Servicios detenidos.")

if __name__ == "__main__":
    run_services()
