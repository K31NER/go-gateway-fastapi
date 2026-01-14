import time, random
from schemas.model import Data

# Mock del modelo de machine learning
class ScoreSalaryModel:
    def predict(self, features: Data):
        
        # Simulamos tiempo
        time.sleep(random.randint(5,10))
        
        # Evitamos divison entre cero
        safe_experience = max(features.experience, 1)
        
        # Calculamos la puntacion de salario
        score = (features.salary / safe_experience) * 0.1
        
        return round(score,2)