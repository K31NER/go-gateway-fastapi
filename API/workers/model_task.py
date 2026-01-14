from schemas.model import Data
from celery.result import AsyncResult
from model.model import ScoreSalaryModel
from config.celery_config import celery_app

model = ScoreSalaryModel()

@celery_app.task(
    bind=True, # Enlazamos la funcion con celery
    autoretry_for=(Exception,), # Reintento por errores
    retry_kwargs={  # Argumentos de reintentos
        "max_retries":5,
        "countdown":5}
    )

def predict_task(self, features: dict):
    """ Predice la puntuacion salariar de usuario"""
    try:
        # Convertimos el diccionario de nuevo a un objeto Data
        data_obj = Data(**features)
        result = model.predict(data_obj)
    except Exception as e:
        raise Exception(f"Error al calcular resultado: {str(e)}")
    return {"prediction":result}

