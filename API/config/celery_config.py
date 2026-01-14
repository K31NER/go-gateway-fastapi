from celery import Celery
from config.settings import settings

# Instanciamos celery
celery_app = Celery(
    "Ml_Tasks",
    broker=settings.CELERY_BROKER,
    backend=settings.CELERY_BACKEND,
    include=["workers.model_task"]
)

# Configuramos el cliente de celery
celery_app.conf.update(
    task_serializer="json",
    accept_content=["json"],
    result_serializer="json",
    timezone = "UTC"
)