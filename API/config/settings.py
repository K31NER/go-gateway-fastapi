from pydantic_settings import BaseSettings

class Settings(BaseSettings):
    CELERY_BROKER: str
    CELERY_BACKEND: str
    
    class Config:
        env_file = ".env"
        extra = "ignore"
        
settings = Settings()