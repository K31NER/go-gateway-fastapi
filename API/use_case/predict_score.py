from domain.data import Data
from repository.data_repository import DataRepository

class UseModel:
    
    def __init__(self, data_repo: DataRepository):
        self.data_repo = data_repo
        
    def predict_task(self, features: Data):
        """ Encola la tarea de prediccion"""
        
        task = self.data_repo.get_predict(features)
        
        return task
    
    def get_status(self, task_id: str):
        """ Obtine el el estado/resultado de la prediccion"""
        response = self.data_repo.get_status_task(task_id)
        
        return response