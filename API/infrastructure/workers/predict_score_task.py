from domain.data import Data
from celery.result import AsyncResult
from config.celery_config import celery_app
from repository.data_repository import DataRepository
from infrastructure.ml.predict_score import ScoreSalaryModel

# Instancia del modelo (Infrastructure)
model = ScoreSalaryModel()

@celery_app.task(
    bind=True,
    autoretry_for=(Exception,),
    retry_kwargs={"max_retries": 5, "countdown": 5}
)
def predict_task(self, features_dict: dict):
    features = Data(**features_dict)
    return model.get_predict(features)

class CeleryWorkerManager(DataRepository):
    """
    Clase de infraestructura para interactuar con Celery.
    Implementa la interfaz definida en la capa de Domain/Repository.
    """
    def get_predict(self, data: Data) -> str:
        """ Encola la tarea y devulve el id"""
        task = predict_task.delay(data.model_dump())
        return task.id
    
    def get_status_task(self, task_id: str):
        """Verifica el estado de la tarea"""
        response = AsyncResult(task_id, app=celery_app)
        
        if response.state == "PENDING":
            return {"status": "En cola..."}
        elif response.state == "RETRY":
            return {"status": "Reintentando..."}
        elif response.state == "STARTED":
            return {"status": "Procesando..."}
        elif response.state == "SUCCESS":
            return {"status": "Completado", "result": response.result}
        elif response.state == "FAILURE":
            return {"status": "Error", "details": str(response.info)}
        else:
            return {"status": response.state}
