from fastapi import APIRouter
from domain.data import Data
from use_case.predict_score import UseModel
from infrastructure.workers.predict_score_task import CeleryWorkerManager

router = APIRouter(prefix="/api/model",tags=["Model"])

@router.post("/predict")
async def predict_data(features: Data):
    celery = CeleryWorkerManager()
    use_case = UseModel(celery)
    
    # Ejecutamos la tarea y obtenemos directamente el ID (string)
    task_id = use_case.predict_task(features)
    
    return {
        "task_id": task_id,
        "status": "Prediccion en cola",
        "message": f"El modelo esta trabajando, consulte el estado en api/model/predict/{task_id}"
    }
    
@router.get("/predict/{task_id}")
async def get_prediction(task_id: str):
    celery = CeleryWorkerManager()
    use_case = UseModel(celery)
    
    # Buscamos nuestra tarea
    task_status = use_case.get_status(task_id)
    
    return task_status
    