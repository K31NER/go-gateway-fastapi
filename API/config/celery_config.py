from celery import Celery
from config.settings import settings

# Instanciamos celery
celery_app = Celery(
    "Ml_Tasks",
    broker=settings.CELERY_BROKER,
    backend=settings.CELERY_BACKEND,
    include=[
        "infrastructure.workers.predict_score_task"]
)

# Configuramos el cliente de celery
celery_app.conf.update(
    task_serializer="json",
    accept_content=["json"],
    result_serializer="json",
    timezone = "UTC",
    result_expires=3600 # 1 Hora
)
