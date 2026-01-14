from config.celery_config import celery_app
from celery.result import AsyncResult

def search_task_status(task_id: str):
    """ Varifica el estado de la tarea"""
    # Buscamos nuestra tarea
    task = AsyncResult(task_id,app=celery_app)
    
    # Validamos su estado
    if task.state == "PENDING":
        return {"status": "En cola..."}
    elif task.state == "RETRY":
        return {"status": "Reintentando..."}
    elif task.state == "STARTED":
        return {"status": "Procesando..."}
    elif task.state == "SUCCESS":
        return {"status": "Completado", "result": task.result}
    elif task.state == "FAILURE":
        return {"status": "Error", "details": str(task.info)}
    else:
        return {"status": task.state}