from fastapi import APIRouter
from schemas.model import Data
from workers.model_task import predict_task
from workers.utils import search_task_status

router = APIRouter(prefix="/api/model",tags=["Model"])

@router.post("/predict")
async def predict_data(features: Data):
    
    task = predict_task.delay(features.model_dump())
    
    return {
        "task_id": task.id,
        "status":"Prediccion en cola",
        "message":f"El modelo esta trabajando, consulte el estado en api/model/predict/{task.id}"
    }
    
@router.get("/predict/{task_id}")
async def get_prediction(task_id: str):
    
    # Buscamos nuestra tarea
    task_status = search_task_status(task_id)
    
    return task_status
    