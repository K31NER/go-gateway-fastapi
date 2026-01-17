from domain.data import Data
from abc import ABC, abstractmethod

class DataRepository(ABC):
    
    @abstractmethod
    def get_predict(self, data: Data) -> str | None:
        pass
    
    @abstractmethod
    def get_status_task(self, task_id: str) -> dict[str,float] | None:
        pass