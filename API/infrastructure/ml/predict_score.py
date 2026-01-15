import time, random
from domain.data import Data
from repository.data_repository import DataRepository

# Mock del modelo de machine learning
class ScoreSalaryModel(DataRepository):
    
    def get_predict(self, features: Data) -> float | None:
        
        # Simulamos tiempo
        time.sleep(random.randint(5,10))
        
        # Evitamos divison entre cero
        safe_experience = max(features.experience, 1)
        
        try:
            # Calculamos la puntacion de salario
            score = (features.salary / safe_experience) * 0.1
        except Exception:
            return None
        
        return round(score,2)
    
    def get_status_task(self, task):
        return super().get_status_task(task)